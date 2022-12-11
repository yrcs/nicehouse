# Nicehouse

#### 介绍
Nicehouse 是一个开源房产管理服务平台，开放聚合线上线下房产资源，打造一个全方位的房产服务生态市场，为消费者提供优质房产服务资源。

#### 软件架构
本项目结合 Go Workspace、Go Module 以及 Makefile 共同组成 Monorepo 项目结构，所有服务 / 模块共享一个 go.mod。每个服务参考 [kratos-layout](https://go-kratos.gitee.io/blog/go-project-layout) 均采用整洁架构分为 3 层：service (DDD application)、biz (DDD domain)、data (DDD repository)，各层之间采用 [wire 依赖注入](https://go-kratos.gitee.io/blog/go-project-wire)。其中 biz 与 data 之间运用了依赖倒置原则，这样不仅消除了 repo 与 DO (Domain Object) 的耦合，同时也避免了出现 Go package 循环引用。

| 现阶段采用的主要开源技术产品                                 |
| ------------------------------------------------------------ |
| 微服务框架：Dubbo-Go (Triple协议)　　　　　　　　　配置中心 / 注册中心：Nacos |
| SQL 数据库：MariaDB　　　　　　　　　　　　　　　ORM   框架：GORM |
| WEB Server 框架：Gin　　　　　　　　　　　　　　　文档生成工具：gin-swagger |
| protobuf 参数校验插件：protoc-gen-validate　　　　　错误处理：go-kratos/errors |
| 依赖注入：wire　　　　　　　　　　　　　　　　　　ULID 生成工具：oklog/ulid |


#### 安装教程

1.  安装最新 Go v1.19.4（本项目使用了 Go 泛型）；
2.  安装部署 MariaDB v10.6.11、Nacos v2.1.2；
3.  下载对应操作系统的 [protoc](https://github.com/protocolbuffers/protobuf/releases)（如：protoc-21.11-linux-x86_64.zip、protoc-21.11-win64.zip），然后解压到 `$GOPATH`；
4.  安装 Goland 2022.3 或者安装最新 VSCode-Go 扩展和 VSCode-proto 扩展。

#### 使用说明

本项目同时兼容 Windows、Linux、macOS，更推荐在**龙芯**机器上使用。其中，Windows 上推荐 wsl2 + Cygwin64 配合使用。

1.  Clone 本项目到本地；
2.  创建数据库`nicehouse`;
3.  配置环境变量`export DUBBO_GO_CONFIG_PATH=../conf/direct.yaml`;
4.  在 `nicehouse` 根目录下执行 `make init` 以安装项目所需的 cli 工具；
5.  使用 Goland 或 VSCode 打开项目后，找到`nicehouse/app/xxx/cmd/app.go` 然后启动 Debugger。

Goland 还需添加额外配置项：

一、添加本项目中的 `third_party` 文件夹到 Goland Protocol Buffers 的 Import Paths 中。

   依次点击菜单 `File` > `Settings...` > `Languages & Frameworks` > `Protocol Buffers` 最后点击 `+` 进行添加。

二、添加 Debug / Run 配置项：

​    依次点击菜单 `Run` > `Edit Configurations...` 然后点击 `+` 进行添加，最后按下面 3 点对应填写。

1. Goland 的 Package path 需要配置到 `github.com/yrcs/nicehouse/app/xxx/cmd` 
2. Goland 的 Working directory 需要配置到 `.../nicehouse/app/xxx/cmd`
3. Goland 的 Output directory 需要 配置到 `.../nicehouse/app/xxx/build`

**以上 3 点都需配置，不然 Debugger 的启动速度会很慢！其中的 xxx 表示具体的模块目录**

先启动 Provider 端（没有以 BFF 开头的模块）后启动 Consumer 端（以 BFF 开头的模块）。接着进行 http api 接口测试：http://localhost:8080/swagger/index.html

最后，即使你上面都配置了。Goland 在启动 Debugger 时仍然总是会出现时快时慢的情况（VSCode 也有此情况）。究其原因，是这两款 IDE 对于 Monorepo 这类工程结构的适配性不够好。

#### 开发说明

本项目整体分两种服务：Provider 服务和 Consumer 服务（即以 BFF - "Backend For Frontend" 开头的服务）。

每个服务的全局变量（如以 Err 开头的自定义的哨兵错误） / 函数都应放到 biz 层。因为 service 层与 data 层都依赖 biz 层，方便各层调用。

开发流程：api -> biz -> service -> data

api 接口服务层（用于定义 DTO 对象和 rpc 接口）：

	1. 定义 proto 文件并执行 `make proto-gen` 生成 *.pb.go 文件；
	1. 定义 xxx_http.go 路由文件并在里面定义各 http 请求方法路由（没有 http 服务需求则不需要定义该文件）。

biz 业务逻辑层：

​	定义 DO (Domain Object) 对象和 data 层所依赖的 repo 接口。

service 服务层：

​	实现了 api 的 server interface，主要完成 DTO -> DO 的转换以及一些简单的业务逻辑。

data 存储层：

​	定义 PO (Persistant Object) 对象，完成 DO -> PO 的转换以及返回 DO。

**小技巧**

1. 在项目根目录执行的 `make xxx` 命令针对所有模块，而在 `app/xxx` 目录下执行的 `make xxx` 命令只针对该 `xxx` 模块。可以执行 `make help` 列出所有命令。
2. 执行 `make init` 后就可以使用 `imports-formatter` 格式化 Go 源码中的 imports 块了。只需在项目根目录下执行 `imports-formatter ./...` 还可以将该工具集成到 Goland 中，参见 [Goland 使用 imports-formatter 工具快速速格式化 import 代码块](https://xie.infoq.cn/article/e9e229f7c468026e9ce17af25)。执行完后再看看源码变化 ^_^

#### **参与贡献**

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
