# Nicehouse

#### 介绍
Nicehouse 是一个开源房产管理服务平台，开放聚合线上线下房产资源，打造一个全方位的房产服务生态市场，为消费者提供优质房产服务资源。

#### 软件架构
本项目结合 Go Workspace、Go Module 以及 Makefile 共同组成 Monorepo 项目结构，所有服务 / 模块共享一个 go.mod。每个服务参考 kratos-layout 均采用整洁架构分为 3 层：service (DDD application)、biz (DDD domain)、data (DDD repository)，其中 biz 与 data 运用了依赖倒置原则，从而消除了 repo 与 DO (Domain Object) 的耦合，同时也避免了出现 Go package 循环引用。

| 现阶段采用的主要开源技术产品：                               |
| ------------------------------------------------------------ |
| 微服务框架：Dubbo-Go (Triple协议)                                  配置中心 / 注册中心：Nacos |
| SQL 数据库：MariaDB                                                          ORM   框架：GORM |
| WEB Server 框架：Gin                                                          Swagger 文档生成工具：gin-swagger |
| protobuf 参数校验插件：protoc-gen-validate                  错误处理：go-kratos/errors |
| 依赖注入：wire                                                                      ULID 生成工具：oklog/ulid |


#### 安装教程

1.  安装最新 Go v1.19.4（本项目使用了 Go 泛型）；
2.  安装部署 MariaDB v10.6.11、Nacos v2.1.2；
3.  安装 Goland 2022.3 或者安装最新 VSCode-Go 扩展和 VSCode-proto 扩展。

#### 使用说明

本项目同时兼容 Windows、Linux、macOS，更推荐在**龙芯**机器上使用。其中，Windows 上推荐 wsl2 + Cygwin64 配合使用。

1.  Clone 本项目到本地；
2.  创建数据库`nicehouse`;
3.  配置环境变量`export DUBBO_GO_CONFIG_PATH=../conf/direct.yaml`;
4.  在 `nicehouse` 根目录下执行 `make init` 以安装项目所需的 cli 工具；
5.  使用 Goland 或 VSCode 打开项目后，找到`nicehouse/app/xxx/cmd/app.go` 然后启动 Debugger。

注意：

1. Goland 的 Package path 需要配置到 `github.com/yrcs/nicehouse/app/xxx/cmd` 

2. Goland 的 Working directory 需要配置到 `.../nicehouse/app/xxx/cmd`

3. Goland 的 Output directory 需要 配置到 `.../nicehouse/app/xxx/build`

   **以上 3 点都需配置，不然 Debugger 的启动速度会很慢！其中的 xxx 表示具体的模块目录**

最后，即使你上面都配置了。无论 Goland，还是 VSCode 在启动 Debugger 时仍然总是会出现时快时慢的情况。究其原因，是这两款 IDE 对于 Monorepo 这类工程结构的适配性不够好。

**小技巧**

1. 在项目根目录执行的 `make xxx` 命令针对所有模块，而在 `app/xxx` 目录下执行的 `make xxx` 命令只针对该 `xxx` 模块。可以执行 `make help` 列出所有命令。
2. 执行 `make init` 后就可以使用 `imports-formatter` 格式化 Go 源码中的 imports 块了。只需在项目根目录下执行 `imports-formatter ./...` 执行完后再看看源码变化 ^_^

#### **参与贡献**

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
