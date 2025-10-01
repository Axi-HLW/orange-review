# 启动流程手册

## 1. 基础环境要求
- 操作系统：Linux 服务器（建议≥4C/8G，Elasticsearch 单独部署更佳）。
- 已安装：Docker、Docker Compose、Go 1.20+、tmux（或 systemd）、curl、git。
- 确保云安全组/防火墙开放以下端口：
  - Consul `8500`
  - MySQL `3306`
  - Redis `6379`
  - Kafka `9092`（EXTERNAL）、`19092`（内部）
  - Canal `11111`（如需调试）
  - Elasticsearch `9200`（HTTP）、`5601`（Kibana）
  - Hertz 网关 `8080`

## 2. 获取代码与准备目录
```bash
cd /opt
git clone <仓库地址> orange-review
cd orange-review
```

### 2.1 数据与配置目录
- 创建持久化目录（示例）：
  ```bash
  sudo mkdir -p /data/orange-review/mysql/data
  sudo mkdir -p /data/orange-review/es/data
  sudo chown -R 1000:1000 /data/orange-review/es
  ```
- 如果使用 Canal，自定义配置放在 `/opt/canal/conf/`（稍后挂载）。
- 准备 Go 运行环境：`export GO_ENV=test`（开发环境默认读取 `conf/test`）。

## 3. 启动基础容器
### 3.1 Consul + MySQL + Redis
```bash
cd docker
docker compose up -d        # 启动 consul、mysql、redis
```
如需修改数据卷映射，在 `docker/docker-compose.yaml` 中调整。

### 3.2 Kafka 套件
```bash
cd kafka
export DOCKER_HOST_IP=<宿主机IP>
docker compose up -d
```
- `kafka-ui` 默认占用 8888，如冲突需在 compose 中调整。
- 验证：`docker exec kafka1 kafka-topics --bootstrap-server kafka1:19092 --list`

### 3.3 Elasticsearch & Kibana
若 ES 部署在独立服务器：
- 在目标机器运行 `docker/es/docker-compose.yaml`。
- 确保 `vm.max_map_count=262144`，数据目录对 UID/GID 1000 可写。
- 启动后确认 `curl http://<ES_IP>:9200` 正常。

### 3.4 Canal Server（可选，用于 MySQL→Kafka 同步）
- 准备配置 `/opt/canal/conf/canal.properties`、`/opt/canal/conf/example/instance.properties`：
  ```
  canal.serverMode = kafka
  canal.mq.servers = <Kafka 外部地址>
  canal.instance.master.address = <mysql_host>:3306
  canal.instance.dbUsername = canal
  canal.instance.dbPassword = canalpass
  canal.mq.topic = topicreview
  ```
- 启动容器：
  ```bash
  docker run -d --name canal-server \
    -p 11111:11111 \
    --network kafka_default \
    -v /opt/canal/conf:/home/admin/canal-server/conf \
    canal/canal-server:v1.1.6
  ```
- 确认日志 `docker logs -f canal-server` 无报错。

## 4. 配置应用
### 4.1 网关 & 服务配置
根据实际环境修改以下文件：
- `app/gateway/conf/<env>/conf.yaml`
- `app/review/conf/<env>/conf.yaml`
- `app/generator/conf/<env>/conf.yaml`
设置服务地址、MySQL/Redis/Consul 信息。

### 4.2 review_job 配置
编辑 `app/review_job/conf/conf.yaml`：
```yaml
elasticsearch:
  addr: "http://<ES_IP>:9200"
  index: "review"
kafka:
  addrs:
    - "<宿主机IP>:9092"    # 与 Canal 输出一致
  group_id: "review-job"
  topic: "topicreview"
```

### 4.3 Canal	réview_job Topic 对齐
- 确保 `canal.instance.filter.regex` 覆盖目标表。
- `canal.mq.topic` 与 `review_job` 的 `kafka.topic` 相同。

## 5. 构建项目
在每个子服务目录执行：
```bash
cd app/generator && go mod download && sh build.sh
cd app/review && go mod download && sh build.sh
cd app/gateway && go mod download && sh build.sh
cd app/review_job && go mod download && go build -o output/bin/review_job
```
如机器内存紧张，可先增加 swap 并设置 `export GOMAXPROCS=1`。

## 6. 启动服务顺序
1. **ID 生成服务（generator）**
   ```bash
   cd app/generator
   GO_ENV=test ./output/bin/generator
   ```
2. **评论服务（review）**
   ```bash
   cd app/review
   GO_ENV=test ./output/bin/review
   ```
3. **HTTP 网关（gateway）**
   ```bash
   cd app/gateway
   GO_ENV=test ./output/bin/gateway
   ```
4. **Kafka → ES 同步任务（review_job）**
   ```bash
   cd app/review_job
   GO_ENV=test ./output/bin/review_job
   ```
> 建议使用 tmux 或 systemd 管理进程。示例 systemd unit：`ExecStart=/opt/orange-review/app/<svc>/output/bin/<svc>`，`Environment=GO_ENV=test`。

## 7. 联调与验证
1. 创建评价（网关接口）：见步骤 4 中 `curl` 示例。
2. MySQL 校验：
   ```bash
   docker exec -it mysql mysql -uroot -p1234 -D gomalldb -e "SELECT review_id FROM review_info ORDER BY id DESC LIMIT 1;"
   ```
3. Kafka 校验：
   ```bash
   docker exec -it kafka1 kafka-console-consumer --bootstrap-server kafka1:19092 --topic topicreview --from-beginning --timeout-ms 1000
   ```
4. review_job 日志应出现 `message at topic/...`、`result:"created"`。
5. Elasticsearch：
   ```bash
   curl -XPOST "http://<ES_IP>:9200/review/_search?pretty" \
     -H 'Content-Type: application/json' \
     -d '{"query":{"term":{"review_id":{"value":"<新ID>"}}}}'
   ```
6. 读接口验证：
   ```bash
   curl "http://<服务器IP>:8080/v1/store/1/reviews?page=1&size=10"
   ```

## 8. 常见问题提示
- ES 启动失败多为数据目录权限或 `vm.max_map_count` 未设置。
- Canal 报 `TimeoutException` 往往是 `canal.mq.servers` 错误或 9092 未开放。
- review_job 日志长期无输出，检查 `kafka.topic` 是否与 Canal 一致。
- 若某服务意外退出，在 `log/` 目录或 tmux 窗口查日志。

## 9. 日常运维
- 使用 `tmux attach -t orange` 或 `systemctl status <svc>` 查看状态。
- 定期备份 MySQL 数据卷与 Elasticsearch 索引（可用快照）。
- 更新流程：`git pull` → 重新 `go build` → 滚动重启各服务。
- 查看 Kafka/ES 的监控指标，以便提前发现积压或资源紧张。

> 按此流程逐步执行，可在新环境快速搭建并验证完整的电商评价链路。
