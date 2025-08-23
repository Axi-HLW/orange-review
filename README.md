# 项目简介

本项目是一个使用Go语言编写的电商评价服务。

---

业务上包含C端（用户端）、B端（商家端）以及O端（平台端），其中：
- C端：创建评价、查看评价
- B端：查看评价、投诉评价
- O端：查看评价、处理投诉、删除评价

---

技术上考虑到电商的评论系统属于「读多写少」的场景，采用CQRS架构：
- **写链路**：MySQL
- **读链路**：ElasticSearch + Redis
- **数据同步**：Canal + Kafka + 消费者脚本

---

本仓库由2个服务+1个API网关组成，均采用字节的CloudWeGo框架实现，分别是：
- **API网关（Hertz）**：暴露HTTP接口给C端、B端、O端用户。
- **评论服务（Kitex）**：暴露gRPC接口，提供核心评论读写能力。
- **唯一ID生成服务（Kitex）**：生成分布式趋势递增唯一ID。

# 架构图

![电商评论服务架构图](./picture/Review-Service-Arch.png)

# API概览


# 数据库表设计

# 参考

参考了七米老师的项目思路，但是完全使用字节开源的：Kitex + Hertz框架实现（而不是教程中的B站开源的：Kratos框架）。

- 七米老师原项目地址：[七米老师github](https://github.com/Q1mi/go-micro-service-and-cloud-native-course/tree/master/lesson115/review-service)
- 字节CloudWeGo官网：[字节CloudWeGo](https://cloudwego.cn/zh/)
- 字节官网项目地址：[字节微服务项目Demo](https://github.com/cloudwego/biz-demo)

## ES

- 查找流程 & 插入流程：[ES的查找与插入笔记](./blog/es.md)