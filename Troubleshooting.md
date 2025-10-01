# 部署排查记录

## 运行环境与依赖
- **kafka-ui 端口冲突**：`docker/kafka/docker-compose.yaml` 默认映射 `8888:8080`，服务器已有进程占用 8888，导致容器启动失败。处理方式：释放端口或修改映射，例如改为 `18088:8080` 后重新 `docker compose up -d kafka-ui`。
- **持久化目录**：为防系统盘写满，需要把 MySQL、Elasticsearch 等容器的数据卷挂载到自定义路径（如 `/data/orange-review/mysql/data`、`/data/orange-review/es/data`），并保证目录对容器 UID/GID 可写。

## 构建与进程管理
- **Go 构建 OOM**：编译 `github.com/elastic/go-elasticsearch/v8/typedapi` 时多次出现 `signal: killed`。根因是机器内存不足，被 OOM Killer 终止。解决方案：增加物理内存或 swap（例如新增 8G 交换文件）、设置 `GOMAXPROCS=1` 或在资源更充足的环境编译再上传二进制。
- **review_job 无 build.sh**：`app/review_job` 需手动 `go build -o output/bin/review_job main.go`。
- **保持进程存活**：直接在 Xshell 前台运行会在断线时退出。推荐使用 `tmux new -s orange` 运行，或编写 systemd unit，使服务随系统启动并在异常退出时自动拉起。

## Kafka / Canal 排障
1. **配置文件挂载**：Canal 镜像从 `/home/admin/canal-server/conf` 读取配置，需要将宿主机目录（如 `/opt/canal/conf`）挂载至该路径。
2. **缺少 canal.mq.servers**：默认模板未包含，导致 Canal 输出 Kafka 时报错。必须单独一行写入 `canal.mq.servers = <broker>`，且不要在同一行添加 `#` 注释。
3. **端口与网络**：最初使用 `29092` 连接外部 Kafka 导致 `TimeoutException`。最终方案：要么在 compose 中设置 `DOCKER_HOST_IP=<宿主机IP>` 并使用 `EXTERNAL://<IP>:9092`，要么将 Canal 加入 Kafka 的 docker 网络并连接 `kafka1:19092`。
4. **安全组与防火墙**：对外访问 Kafka 需放行 9092（云安全组 + `firewall-cmd --permanent --add-port=9092/tcp`）。
5. **Topic 不一致**：Canal 输出 `topicreview`，而 `review_job` 初始订阅 `topic1`，导致消费不到。统一 topic 后问题解决。
6. **调试手段**：使用 `docker logs canal-server` 检查 Canal 报错；使用 `docker exec kafka1 kafka-console-consumer --bootstrap-server kafka1:19092 --topic <topic> --from-beginning` 验证 Kafka 中的消息。

## review_job 配置注意点
- `conf.GetEnv()` 固定返回 "./"，因此实际加载的是 `app/review_job/conf/conf.yaml`。部署时必须在服务器上编辑该文件（而非本地副本）。
- 曾多次看到日志 `Kafka:{Addrs:[localhost:9092]}`，说明配置尚未更新。改成实际地址 `139.224.15.94:9092` 并重启即可。
- 之后又由于订阅的 topic 与 Canal 不符导致日志无输出。修改 `kafka.topic` 与 Canal 保持一致（`topicreview`），消费日志即恢复，并在写入 ES 时打印 `result:"created"`。

## Elasticsearch 相关问题
- **容器无法启动**：日志报 `failed to obtain node locks ... /usr/share/elasticsearch/data`。原因是数据目录不可写或缺失。需要在宿主机创建目录并赋予 UID/GID 1000 权限，调整 compose 中的数据卷映射，并设置 `vm.max_map_count=262144`。
- **资源不足**：原服务器可用内存不足 1GB，ES 频繁重启。最终将 Elasticsearch + Kibana 部署到独立服务器，只运行这两个服务。
- **配置同步**：迁移后需更新：
  - `app/review/conf/<env>/conf.yaml` 中的 `elasticsearch.addr`
  - `app/review_job/conf/conf.yaml` 中的 `elasticsearch.addr`
  服务重启后通过 `curl http://<ES>/review/_search` 或 Kibana 验证数据是否写入。

## 验证链路的流程
1. 通过网关创建评价：
   ```bash
   curl -X POST http://<服务器IP>:8080/v1/review \
     -d "user_id=1001&order_id=5001&store_id=1&score=5&service_score=5&express_score=5" \
     -d "content=great+product&pic_info=[]&video_info=&anonymous=false&sku_id=2001&spu_id=3001"
   ```
2. 确认 MySQL `review_info` 表新增记录。
3. 使用 Kafka CLI 或 Kafka UI 查看目标 topic 是否收到 Canal 推送的消息。
4. 观察 `review_job` 日志：若消费成功，会打印 `message at topic/...` 以及 `result:"created"`。
5. Elasticsearch `review` 索引能查询到新文档，Kibana Discover 亦可看到。
6. 调用读接口 `GET /v1/store/{store_id}/reviews?page=1&size=10`，验证 Redis/ES 读链是否通畅。

## 常用调试命令
- Kafka 生产消息：`docker exec -it kafka1 kafka-console-producer --broker-list kafka1:19092 --topic <topic>`。
- Kafka 消费消息：`docker exec -it kafka1 kafka-console-consumer --bootstrap-server kafka1:19092 --topic <topic> --from-beginning --timeout-ms 1000`。
- Elasticsearch 查询：`curl http://<ES>/review/_search?pretty`。
- tmux 断线重连：`tmux attach -t orange`。

> 以上问题与处理方案为本次部署排查过程的完整记录，后续复现或排障时可按此清单逐项核对。
