# 更新日志

本文件记录项目的重要变更。

## [0.1.0] - 2026-03-25

项目首次发布。

### 新增

- 新增 `agent-observability` 服务，用于从 OpenSearch 查询智能体 Trace 数据。
- 新增 Trace 查询接口，支持原始 DSL 检索和按 `conversation_id` 查询，并生成 Swagger 文档。
- 新增 `agent-observability` 的 Docker、Helm 和 GitHub Actions 构建发布流程。
- 新增 `otelcol-contribute-chart` Helm Chart，用于在 Kubernetes 中部署 OpenTelemetry Collector Contrib。
- 新增 Collector 默认 OTLP 接入与 OpenSearch 导出配置，支持 Trace 和 Log 链路处理。
- 新增仓库级中英文 README，说明 Tracing AI 的定位、架构、核心能力与快速开始方式。

### 文档

- 新增 `agent-observability/docs/` 下的 Agent 链路系统文档，包括 PRD、设计文档、API 描述与 Swagger 产物。
