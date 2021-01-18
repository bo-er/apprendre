# 后端架构发展历程

## 单机架构

- 应用数量与用户数都较少，可以把应用和数据库部署在同一台服务器上
- 随着用户数的增长，应用和数据库之间竞争资源，单机性能不足以支撑业务

## 第一次演进：数据库与应用分开部署

- 应用和数据库分别独占服务器资源，显著提高两者各自性能。
- 随着用户数的增长，并发读写数据库成为瓶颈

## 第二次演进：引入本地缓存和分布式缓存

- 在应用服务器上增加本地缓存，并在外部增加分布式缓存，缓存热门访问信息或页面等。通过缓存能把绝大多数请求在读写数据库前拦截掉，大大降低数据库压力
- 缓存抗住了大部分的访问请求，随着用户数的增长，并发压力主要落在单机的应用上，响应逐渐变慢

## 第三次演进：引入反向代理实现负载均衡

- 在多台服务器上分别部署应用，使用反向代理软件（Nginx）把请求均匀分发到每个应用中。此处假设应用最多支持 100 个并发，Nginx 最多支持 50000 个并发，那么理论上 Nginx 把请求分发到 500 个应用上，就能抗住 50000 个并发。其中涉及的技术包括：Nginx、HAProxy，两者都是工作在网络第七层的反向代理软件，主要支持 http 协议，还会涉及 session 共享、文件上传下载的问题。
- 反向代理使应用服务器可支持的并发量大大增加，但并发量的增长也意味着更多请求穿透到数据库，单机的数据库最终成为瓶颈

## 第四次演进：数据库读写分离

- 把数据库划分为读库和写库，读库可以有多个，通过同步机制把写库的数据同步到读库，对于需要查询最新写入数据场景，可通过在缓存中多写一份，通过缓存获得最新数据。其中涉及的技术包括：Mycat，它是数据库中间件，可通过它来组织数据库的分离读写和分库分表，客户端通过它来访问下层数据库，还会涉及数据同步，数据一致性的问题。
- 业务逐渐变多，不同业务之间的访问量差距较大，不同业务直接竞争数据库，相互影响性能

## 第五次演进：数据库按业务分库

- 把订单、用户等等数据存放到不同的数据库
- 把不同业务的数据保存到不同的数据库中，使业务之间的资源竞争降低，对于访问量大的业务，可以部署更多的服务器来支撑。
- 随着用户数的增长，单机的写库会逐渐会达到性能瓶颈

## 第六次演进：把大表拆分为小表

- 比如针对评论数据，可按照商品 ID 进行 hash，路由到对应的表中存储；针对支付记录，可按照小时创建表，每个小时表继续拆分为小表，使用用户 ID 或记录编号来路由数据。只要实时操作的表数据量足够小，请求能够足够均匀的分发到多台服务器上的小表，那数据库就能通过水平扩展的方式来提高性能。其中前面提到的 Mycat 也支持在大表拆分为小表情况下的访问控制。
- 这种做法显著的增加了数据库运维的难度，对 DBA 的要求较高。数据库设计到这种结构时，已经可以称为分布式数据库，但是这只是一个逻辑的数据库整体，数据库里不同的组成部分是由不同的组件单独来实现的，如分库分表的管理和请求分发，由 Mycat 实现，SQL 的解析由单机的数据库实现，读写分离可能由网关和消息队列来实现，查询结果的汇总可能由数据库接口层来实现等等，这种架构其实是 MPP（大规模并行处理）架构的一类实现。
- 数据库和应用都能够水平扩展，可支撑的并发大幅提高，随着用户数的增长，最终单机的 Nginx 会成为瓶颈

## 第七次演进：使用 LVS 或 F5 来使多个 Nginx 负载均衡

- 由于瓶颈在 Nginx，因此无法通过两层的 Nginx 来实现多个 Nginx 的负载均衡。图中的 LVS 和 F5 是工作在网络第四层的负载均衡解决方案，其中 LVS 是软件，运行在操作系统内核态，可对 TCP 请求或更高层级的网络协议进行转发，因此支持的协议更丰富，并且性能也远高于 Nginx，可假设单机的 LVS 可支持几十万个并发的请求转发；F5 是一种负载均衡硬件，与 LVS 提供的能力类似，性能比 LVS 更高，但价格昂贵。由于 LVS 是单机版的软件，若 LVS 所在服务器宕机则会导致整个后端系统都无法访问，因此需要有备用节点。可使用 keepalived 软件模拟出虚拟 IP，然后把虚拟 IP 绑定到多台 LVS 服务器上，浏览器访问虚拟 IP 时，会被路由器重定向到真实的 LVS 服务器，当主 LVS 服务器宕机时，keepalived 软件会自动更新路由器中的路由表，把虚拟 IP 重定向到另外一台正常的 LVS 服务器，从而达到 LVS 服务器高可用的效果。
- 由于 LVS 也是单机的，随着并发数增长到几十万时，LVS 服务器最终会达到瓶颈，此时用户数达到千万甚至上亿级别，用户分布在不同的地区，与服务器机房距离不同，导致了访问的延迟会明显不同

## 第八次演进：通过 DNS 轮询实现机房间的负载均衡

- 在 DNS 服务器中可配置一个域名对应多个 IP 地址，每个 IP 地址对应到不同的机房里的虚拟 IP。当用户访问http://cn.chinasws.com/时，DNS服务器会使用轮询策略或其他策略，来选择某个IP供用户访问。此方式能实现机房间的负载均衡，至此，系统可做到机房级别的水平扩展，千万级到亿级的并发量都可通过增加机房来解决，系统入口处的请求并发量不再是问题。
- 随着数据的丰富程度和业务的发展，检索、分析等需求越来越丰富，单单依靠数据库无法解决如此丰富的需求

## 第九次演进：引入 NoSQL 数据库和搜索引擎等技术

- 当数据库中的数据多到一定规模时，数据库就不适用于复杂的查询了，往往只能满足普通查询的场景。对于统计报表场景，在数据量大时不一定能跑出结果，而且在跑复杂查询时会导致其他查询变慢，对于全文检索、可变数据结构等场景，数据库天生不适用。因此需要针对特定的场景，引入合适的解决方案。如对于海量文件存储，可通过分布式文件系统 HDFS 解决，对于 key value 类型的数据，可通过 HBase 和 Redis 等方案解决，对于全文检索场景，可通过搜索引擎如 ElasticSearch 解决，对于多维分析场景，可通过 Kylin 或 Druid 等方案解决。
- 当然，引入更多组件同时会提高系统的复杂度，不同的组件保存的数据需要同步，需要考虑一致性的问题，需要有更多的运维手段来管理这些组件等。
- 引入更多组件解决了丰富的需求，业务维度能够极大扩充，随之而来的是一个应用中包含了太多的业务代码，业务的升级迭代变得困难

## 第十次演进：大应用拆分为小应用

- 按照业务板块来划分应用代码，使单个应用的职责更清晰，相互之间可以做到独立升级迭代。这时候应用之间可能会涉及到一些公共配置，可以通过分布式配置中心 Zookeeper 来解决。
- 不同应用之间存在共用的模块，由应用单独管理会导致相同代码存在多份，导致公共功能升级时全部应用代码都要跟着升级

## 第十一次演进：复用的功能抽离成微服务

- 如用户管理、订单、支付、鉴权等功能在多个应用中都存在，那么可以把这些功能的代码单独抽取出来形成一个单独的服务来管理，这样的服务就是所谓的微服务，应用和服务之间通过 HTTP、TCP 或 RPC 请求等多种方式来访问公共服务，每个单独的服务都可以由单独的团队来管理。此外，可以通过 Dubbo、SpringCloud 等框架实现服务治理、限流、熔断、降级等功能，提高服务的稳定性和可用性。
- 不同服务的接口访问方式不同，应用代码需要适配多种访问方式才能使用服务，此外，应用访问服务，服务之间也可能相互访问，调用链将会变得非常复杂，逻辑变得混乱

## 第十二次演进：引入企业服务总线 ESB 屏蔽服务接口的访问差异

- 通过 ESB 统一进行访问协议转换，应用统一通过 ESB 来访问后端服务，服务与服务之间也通过 ESB 来相互调用，以此降低系统的耦合程度。这种单个应用拆分为多个应用，公共服务单独抽取出来来管理，并使用企业消息总线来解除服务之间耦合问题的架构，就是所谓的 SOA（面向服务）架构，这种架构与微服务架构容易混淆，因为表现形式十分相似。个人理解，微服务架构更多是指把系统里的公共服务抽取出来单独运维管理的思想，而 SOA 架构则是指一种拆分服务并使服务接口访问变得统一的架构思想，SOA 架构中包含了微服务的思想。
- 业务不断发展，应用和服务都会不断变多，应用和服务的部署变得复杂，同一台服务器上部署多个服务还要解决运行环境冲突的问题，此外，对于如大促这类需要动态扩缩容的场景，需要水平扩展服务的性能，就需要在新增的服务上准备运行环境，部署服务等，运维将变得十分困难

## 第十三次演进：引入容器化技术实现运行环境隔离与动态服务管理

- 目前最流行的容器化技术是 Docker，最流行的容器管理服务是 Kubernetes(K8S)，应用/服务可以打包为 Docker 镜像，通过 K8S 来动态分发和部署镜像。Docker 镜像可理解为一个能运行你的应用/服务的最小的操作系统，里面放着应用/服务的运行代码，运行环境根据实际的需要设置好。把整个“操作系统”打包为一个镜像后，就可以分发到需要部署相关服务的机器上，直接启动 Docker 镜像就可以把服务起起来，使服务的部署和运维变得简单。
- 使用容器化技术后服务动态扩缩容问题得以解决，但是机器还是需要公司自身来管理，在非大促的时候，还是需要闲置着大量的机器资源来应对大促，机器自身成本和运维成本都极高，资源利用率低

## 第十四次演进：以云平台承载系统

- 系统可部署到公有云上，利用公有云的海量机器资源，解决动态硬件资源的问题，在大促的时间段里，在云平台中临时申请更多的资源，结合 Docker 和 K8S 来快速部署服务，在大促结束后释放资源，真正做到按需付费，资源利用率大大提高，同时大大降低了运维成本。
- 所谓的云平台，就是把海量机器资源，通过统一的资源管理，抽象为一个资源整体，在之上可按需动态申请硬件资源（如 CPU、内存、网络等），并且之上提供通用的操作系统，提供常用的技术组件（如 Hadoop 技术栈，MPP 数据库等）供用户使用，甚至提供开发好的应用，用户不需要关系应用内部使用了什么技术，就能够解决需求（如音视频转码服务、邮件服务、个人博客等）。在云平台中会涉及如下几个概念：
- IaaS：基础设施即服务。对应于上面所说的机器资源统一为资源整体，可动态申请硬件资源的层面；
- PaaS：平台即服务。对应于上面所说的提供常用的技术组件方便系统的开发和维护；
- SaaS：软件即服务。对应于上面所说的提供开发好的应用或服务，按功能或性能要求付费。

## 架构设计遵循的原则

- N+1 设计。系统中的每个组件都应做到没有单点故障；
- 回滚设计。确保系统可以向前兼容，在系统升级时应能有办法回滚版本；
- 禁用设计。应该提供控制具体功能是否可用的配置，在系统出现故障时能够快速下线功能；
- 监控设计。在设计阶段就要考虑监控的手段；(API 层级)
- 多活数据中心设计。若系统需要极高的高可用，应考虑在多地实施数据中心进行多活，至少在一个机房断电的情况下系统依然可用；
- 采用成熟的技术。刚开发的或开源的技术往往存在很多隐藏的 bug，出了问题没有商业支持可能会是一个灾难；
- 资源隔离设计。应避免单一业务占用全部资源；
- 架构应能水平扩展。系统只有做到能水平扩展，才能有效避免瓶颈问题；
- 非核心则购买。非核心功能若需要占用大量的研发资源才能解决，则考虑购买成熟的产品；
- 使用商用硬件。商用硬件能有效降低硬件故障的机率；
- 快速迭代。系统应该快速开发小功能模块，尽快上线进行验证，早日发现问题大大降低系统交付的风险；
- 无状态设计。服务接口应该做成无状态的，当前接口的访问不依赖于接口上次访问的状态。

# 微服务简介

## 官方定义

The microservice architectural style is an approach to developing a single application as a suite of small services, each running in its own process and communicating with lightweight mechanisms, often an HTTP resource API. These services are built around business capabilities and independently deployable by fully automated deployment machinery. There is a bare minimum of centralized management of these services , which may be written in different programming languages and use different data storage technologies.– James Lewis and Martin Fowler

微服务架构风格是一种将单个应用程序开发为一套小型服务的方法，每个服务都在自己的进程中运行，并通过轻量级机制（通常是 HTTP 资源 API）进行通信。这些服务围绕业务能力构建，并可通过全自动部署机制独立部署。对这些服务进行最低限度的集中管理，这些服务可能用不同的编程语言编写，使用不同的数据存储技术。

- 一些列的独立的服务共同组成系统
- 单独部署，跑在自己的进程里
- 每个服务为独立的业务开发
- 分布式的管理
- 自动化运维（DevOps）
- 容错
- 快速演化

引入技术：`Docker` `container` `Kubernetes` `Helm` `Harbor` `elasticsearch(日志相关)` `filebeat(日志相关)` `Kibana(日志相关)` `Drone` `Git`

## 优点

- 开发效率高
- 代码维护容易
- 部署灵活
- 稳定性高
- 扩展性高

## 缺点

- 开发复杂，分布式管理
- 会有重复开发
- 分布式的管理开销
- 分布式的调用开销

## 怎么具体实践微服务

### 客户端如何访问这些服务？

后台有 N 个服务，前台就需要记住管理 N 个服务，一个服务下线/更新/升级，前台就要重新部署，这明显不服务我们拆分的理念，特别当前台是移动应用的时候，通常业务变化的节奏更快,所以，一般在后台 N 个服务和 UI 之间一般会一个代理或者叫 API Gateway，他的作用包括

- 提供统一服务入口，让微服务对前台透明
- 聚合后台的服务，节省流量，提升性能
- 提供安全，过滤，流控等 API 管理功能

我的理解其实这个 API Gateway 可以有很多广义的实现办法，可以是一个软硬一体的盒子，也可以是一个简单的 MVC 框架，甚至是一个 Node.js 的服务端。他们最重要的作用是为前台（通常是移动应用）提供后台服务的聚合，提供一个统一的服务出口，解除他们之间的耦合，不过 API Gateway 也有可能成为单点故障点或者性能的瓶颈。

引入技术：`Ingress` `Traefik` `kong`

### 服务之间如何通信？

RPC 与异步消息的方式在分布式系统中有特别广泛的应用，他既能减低调用服务之间的耦合，又能成为调用之间的缓冲，确保消息积压不会冲垮被调用方，同时能保证调用方的服务体验，继续干自己该干的活，不至于被后台性能拖慢。不过需要付出的代价是一致性的减弱，需要接受数据最终一致性；还有就是后台服务一般要实现幂等性，因为消息发送出于性能的考虑一般会有重复（保证消息的被收到且仅收到一次对性能是很大的考验)

引入技术：`gRPC` `Restful` `kafka` `emqx` `nsq`

### 这么多服务，怎么找?

在微服务架构中，一般每一个服务都是有多个拷贝，来做负载均衡。一个服务随时可能下线，也可能应对临时访问压力增加新的服务节点。服务之间如何相互感知？服务如何管理？这就是服务发现的问题了。一般有两类做法，也各有优缺点。基本都是通过 zookeeper 等类似技术做服务注册信息的分布式管理。当服务上线时，服务提供者将自己的服务信息注册到 ZK（或类似框架），并通过心跳维持长链接，实时更新链接信息。服务调用者通过 ZK 寻址，根据可定制算法，找到一个服务，还可以将服务信息缓存在本地以提高性能。当服务下线时，ZK 会发通知给服务客户端。

引入技术：`Consul(目前使用的)` `etcd` `Istio(下一代技术)`

### 服务挂了怎么办？

分布式最大的特性就是网络是不可靠的。通过微服务拆分能降低这个风险，不过如果没有特别的保障，结局肯定是噩梦。我们刚遇到一个线上故障就是一个很不起眼的 SQL 计数功能，在访问量上升时，导致数据库 load 彪高，影响了所在应用的性能，从而影响所有调用这个应用服务的前台应用。所以当我们的系统是由一系列的服务调用链组成的时候，我们必须确保任一环节出问题都不至于影响整体链路。相应的手段有很多：

- 重试机制
- 限流
- 熔断机制
- 负载均衡
- 降级（本地缓存）

引入技术： `Prometheus` `Grafana` `Skywalking` `Hystrix` `Nginx` `kenyata`

# gRPC 简介

## 什么是 RPC

远程过程调用，是分布式系统中不同节点间流行的通信方式。

## 为什么需要 RPC

- 解决分布式系统中，服务之间的调用问题。
- 远程调用时，要能够像本地调用一样方便，让调用者感知不到远程调用的逻辑。
- 节省传输流量，提高传输效率，用二进制传输。

## RPC 与 Restful 区别

其实这两者并不是一个维度的概念，总得来说 RPC 涉及的维度更广。 如果硬要比较，那么可以从 RPC 风格的 url 和 Restful 风格的 url 上进行比较。 比如你提供一个查询订单的接口，用 RPC 风格，你可能会这样写：

```
/queryOrder?orderId=123
```

用 Restful 风格

```
Get
/order?orderId=123
```

RPC 是面向过程，Restful 是面向资源，并且使用了 Http 动词。从这个维度上看，Restful 风格的 url 在表述的精简性、可读性上都要更好。

## RPC 与 gRPC

要实现一个 RPC 不算难，难的是实现一个高性能高可靠的 RPC 框架。 gRPC 是一个高性能、开源和通用的 RPC 框架，面向移动和 HTTP/2 设计。 使用 gRPC 能让我们更容易编写跨语言的分布式代码。

## 安装 protobuf 工具链

grpc 使用 protobuf 作为 IDL(interface description language)

### 安装 protobuf

[下载地址](https://github.com/protocolbuffers/protobuf/releases)

```
# 安装
mv protoc /usr/local/protoc
# 查看 protoc
protoc --version
```

### 安装 protoc-gen-go

protoc-gen-go 是 Go 的 protoc 编译插件，protobuf 内置了许多高级语言的编译器，但没有 Go 的。

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

运行 protoc -h 命令可以发现内置的只支持以下语言

```
protoc -h
...
--cpp_out=OUT_DIR           Generate C++ header and source.
--csharp_out=OUT_DIR        Generate C# source file.
--java_out=OUT_DIR          Generate Java source file.
--js_out=OUT_DIR            Generate JavaScript source.
--objc_out=OUT_DIR          Generate Objective C header and source.
--php_out=OUT_DIR           Generate PHP source file.
--python_out=OUT_DIR        Generate Python source file.
--ruby_out=OUT_DIR          Generate Ruby source file.
...
```

所以我们使用 protoc 编译生成 Go 版的 grpc 时，需要先安装此插件。

### 安装 grpc-go 库

grpc-go 包含了 Go 的 grpc 库。

```
go get google.golang.org/grpc
```

可能会被墙掉了，使用如下方式手动安装。

```
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto

cd $GOPATH/src/
go install google.golang.org/grpc
```

# proto3 语言指南

## protobuf 语法

https://developers.google.com/protocol-buffers/docs/overview https://developers.google.com/protocol-buffers/docs/proto3

### 定义消息类型（Defining A Message Type）

First let’s look at a very simple example. Let’s say you want to define a search request message format, where each search request has a query string, the particular page of results you are interested in, and a number of results per page. Here’s the .proto file you use to define the message type.

让我们先看一个 proto3 的查找请求参数的消息格式的例子，这个请求参数例子模仿分页查找请求，他有一个请求参数字符串，有一个当前页的参数还有一个每页返回数据大小的参数，proto 文件内容如下：

不要覆盖或者修改，而是往下加序号，否则如果其他程序也使用了`SearchRequest`那么覆盖修改后会导致错误。

```java
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```

- The first line of the file specifies that you’re using proto3 syntax: if you don’t do this the protocol buffer compiler will assume you are using proto2. This must be the first non-empty, non-comment line of the file.
- 第一行的含义是限定该文件使用的是`proto3`的语法， syntax = “proto3”;如果你不这样做，协议缓冲区编译器将假定你正在使用`proto2`的语法，文件的第一个必须是非空注释行。
- The SearchRequest message definition specifies three fields (name/value pairs), one for each piece of data that you want to include in this type of message. Each field has a name and a type.
- SearchRequest 定义有三个承载消息的属性，每一个被定义在 SearchRequest 消息体中的字段，都是由数据类型和属性名称组成。

#### 指定字段类型（Specifying Field Types）

In the above example, all the fields are scalar types: two integers (page_number and result_per_page) and a string (query). However, you can also specify composite types for your fields, including enumerations and other message types.

在上面的例子中，所有的属性都是标量，两个整型(page_number、result_per_page)和一个字符串(query)，你还可以在指定复合类型，包括枚举类型或者其他的消息类型。

#### 分配标量（Assigning Field Numbers)

As you can see, each field in the message definition has a unique number. These field numbers are used to identify your fields in the message binary format, and should not be changed once your message type is in use. Note that field numbers in the range 1 through 15 take one byte to encode, including the field number and the field’s type (you can find out more about this in Protocol Buffer Encoding). Field numbers in the range 16 through 2047 take two bytes. So you should reserve the numbers 1 through 15 for very frequently occurring message elements. Remember to leave some room for frequently occurring elements that might be added in the future.

就像所看见的一样，每一个被定义在消息中的字段都会被分配给一个唯一的标量，这些标量用于标识你定义在二进制消息格式中的属性，标量一旦被定义就不允许在使用过程中再次被改变。标量的值在 1 ～ 15 的这个范围里占一个字节编码(详情请参看 谷歌的 Protocol Buffer Encoding )。https://developers.google.com/protocol-buffers/docs/encoding。范围16到2047的字段编号采用两个字节。因此，应该为经常出现的消息元素保留数字1到15。记住为将来可能添加的频繁出现的元素留出一些空间。

The smallest field number you can specify is 1, and the largest is 229 - 1, or 536,870,911. You also cannot use the numbers 19000 through 19999 (FieldDescriptor::kFirstReservedNumber through FieldDescriptor::kLastReservedNumber), as they are reserved for the Protocol Buffers implementation - the protocol buffer compiler will complain if you use one of these reserved numbers in your .proto. Similarly, you cannot use any previously reserved field numbers.

您可以指定的最小字段数是 1，最大的字段数是 229-1，即 536,870,911。你也不能使用 19000 到 19999 这些数字(FieldDescriptor: : kFirstReservedNumber through FieldDescriptor: kLastReservedNumber) ，因为它们是为协议缓冲实现保留的，同样，也不能使用任何以前使用过的字段编号。

#### 指定属性规则（Specifying Field Rules）

Message fields can be one of the following:

消息字段可以是下列字段之一:

- singular: a well-formed message can have zero or one of this field (but not more than one). And this is the default field rule for proto3 syntax.
- 单数: 一个正确的消息可以有零个或者一个这样的消息属性(但是不要超过一个)，这是 proto3 语法的默认字段规则
- repeated: this field can be repeated any number of times (including zero) in a well-formed message. The order of the repeated values will be preserved.
- 重复: 该字段可以在格式良好的消息中重复任意次数(包括零次)。重复值的顺序将被保留

In proto3, repeated fields of scalar numeric types use packed encoding by default.

在 proto3 中，标量数字类型的重复字段默认使用压缩编码

You can find out more about packed encoding in Protocol Buffer Encoding.

您可以在 Protocol Buffer Encoding (https://developers.google.com/protocol-buffers/docs/encoding)中找到关于打包编码的更多信息。

#### 添加更多的消息类型(Adding More Message Types)

Multiple message types can be defined in a single .proto file. This is useful if you are defining multiple related messages – so, for example, if you wanted to define the reply message format that corresponds to your SearchResponse message type, you could add it to the same .proto:

在一个 proto 文件中可以定义多个消息类型，你可以在一个文件中定义一些相关的消息类型，上面的例子 proto 文件中只有一个请求查找的消息类型，现在可以为他多添加一个响应的消息类型，具体如下：

```
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

message SearchResponse {
    ....
}
```

#### 添加注释（Adding Comments）

To add comments to your .proto files, use C/C++-style // and /_ … _/ syntax.

proto 文件中的注释使用的是 c/c++中的单行注释 // 和多行注释 /_ … _/ 语法风格，

```
/* SearchRequest represents a search query, with pagination options to
 * indicate which results to include in the response. */

message SearchRequest {
  string query = 1;
  int32 page_number = 2;  // Which page number do we want?
  int32 result_per_page = 3;  // Number of results to return per page.
}
```

#### 保留属性（Reserved Fields）

If you update a message type by entirely removing a field, or commenting it out, future users can reuse the field number when making their own updates to the type. This can cause severe issues if they later load old versions of the same .proto, including data corruption, privacy bugs, and so on. One way to make sure this doesn’t happen is to specify that the field numbers (and/or names, which can also cause issues for JSON serialization) of your deleted fields are reserved. The protocol buffer compiler will complain if any future users try to use these field identifiers.

为了避免在加载相同的.proto 的旧版本，包括数据损坏，隐含的错误等，这可能会导致严重的问题的方法是指定删除的字段的字段标签（和/或名称，也可能导致 JSON 序列化的问题）被保留。 如果将来的用户尝试使用这些字段标识符，协议缓冲区编译器将会报错。

```
message Foo {
  reserved 2;
  reserved "foo", "bar";
}
```

上述例子定义保留属性为”foo”, “bar”，定义保留属性位置为 2，即在 2 这个位置上不可以定义属性，如:string name=2;是不允许的，编译器在编译 proto 文件的时候如果发现，2 这个位置上有属性被定义则会报错。

Note that you can’t mix field names and field numbers in the same reserved statement.

注意，不能在同一个保留语句中混合字段名和字段编号。

#### 你的原型产生了什么？ (What’s Generated From Your .proto?)

When you run the protocol buffer compiler on a .proto, the compiler generates the code in your chosen language you’ll need to work with the message types you’ve described in the file, including getting and setting field values, serializing your messages to an output stream, and parsing your messages from an input stream.

编译器根据`.proto`生成你选择语言的代码，包括工作的消息类型，文件中描述，字段值，序列化你的消息到一个输出流，解析你的消息从输入流。

… - For Go, the compiler generates a .pb.go file with a type for each message type in your file. - 对于 Go 语言，编译器生成一个后缀为 `.pb.go` 的文件 …

### 数据类型（Scalar Value Types）

A scalar message field can have one of the following types – the table shows the type specified in the .proto file, and the corresponding type in the automatically generated class:

一个信息标量具有如下表格所示的数据类型，下表主要是对.proto 文件的值类型和 go 的值类型的对照表

| .proto Type | Notes                                                                                                                                           | Go Type |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------- |
| double      |                                                                                                                                                 | float64 |
| float       |                                                                                                                                                 | float32 |
| int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32   |
| int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64   |
| uint32      | Uses variable-length encoding.                                                                                                                  | uint32  |
| uint64      | Uses variable-length encoding.                                                                                                                  | uint64  |
| sint32      | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int32   |
| sint64      | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | int64   |
| fixed32     | Always four bytes. More efficient than uint32 if values are often greater than 228.                                                             | uint32  |
| fixed64     | Always eight bytes. More efficient than uint64 if values are often greater than 256.                                                            | uint64  |
| sfixed32    | Always four bytes.                                                                                                                              | int32   |
| sfixed64    | Always eight bytes.                                                                                                                             | int64   |
| bool        |                                                                                                                                                 | bool    |
| string      | A string must always contain UTF-8 encoded or 7-bit ASCII text, and cannot be longer than 232                                                   | String  |
| bytes       | May contain any arbitrary sequence of bytes no longer than 232.                                                                                 | []byte  |

详情参看官方文档

### 默认值（Default Values）

When a message is parsed, if the encoded message does not contain a particular singular element, the corresponding field in the parsed object is set to the default value for that field. These defaults are type-specific:

当解析消息时，如果编码消息不包含特定的值，解析对象中的相应字段将设置为该字段的默认值。这些缺省值对于不同类型是不同的:

- For strings, the default value is the empty string.默认值是空字符串,注意不是 null
- For bytes, the default value is empty bytes.默认值是空 bytes
- For bools, the default value is false.默认值是 false
- For numeric types, the default value is zero.默认值是 0
- For enums, the default value is the first defined enum value, which must be 0.默认值是第一个枚举值,即 0
- For message fields, the field is not set. Its exact value is language-dependent. See the generated code guide for details.

The default value for repeated fields is empty (generally an empty list in the appropriate language).

重复字段的默认值为空（在相对应的编程语言中通常是一个空的 list）.

Note that for scalar message fields, once a message is parsed there’s no way of telling whether a field was explicitly set to the default value (for example whether a boolean was set to false) or just not set at all: you should bear this in mind when defining your message types. For example, don’t have a boolean that switches on some behaviour when set to false if you don’t want that behaviour to also happen by default. Also note that if a scalar message field is set to its default, the value will not be serialized on the wire.

请注意，对于标量消息字段，一旦消息被解析，就无法判断字段是显式设置为默认值(例如，是否将布尔值设置为 false)还是根本没有设置: 在定义消息类型时应该牢记这一点。例如，如果您不希望某些行为在默认情况下也发生，那么就不要设置一个布尔值，该布尔值在设置为 false 时可以开启某些行为。还要注意，如果将标量消息字段设置为默认值，则该值将不会在连接上序列化。

See the generated code guide for your chosen language for more details about how defaults work in generated code.

有关生成的代码的默认工作方式的更多详细信息，请参阅所选语言的生成代码指南。（https://developers.google.com/protocol-buffers/docs/reference/go-generated）

### 枚举（Enumerations）

When you’re defining a message type, you might want one of its fields to only have one of a pre-defined list of values. For example, let’s say you want to add a corpus field for each SearchRequest, where the corpus can be UNIVERSAL, WEB, IMAGES, LOCAL, NEWS, PRODUCTS or VIDEO. You can do this very simply by adding an enum to your message definition with a constant for each possible value.

在定义消息类型时，您可能希望其中一个字段只具有一个预定义的值列表。例如，假设您想为每个 SearchRequest 添加一个语料库字段，其中语料库可以是 `UNIVERSAL`、 `WEB`、 `IMAGES`、 `LOCAL`、`NEWS`、 `PRODUCTS` 或 `VIDEO`。您可以非常简单地做到这一点，方法是在消息定义中为每个可能的值添加一个带常量的枚举。

In the following example we’ve added an enum called Corpus with all the possible values, and a field of type Corpus:

在下面的例子中，我们添加了一个名为 Corpus 的枚举，包含所有可能的值，以及一个类型为 Corpus 的字段:

```java
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 4;
}
```

As you can see, the Corpus enum’s first constant maps to zero: every enum definition must contain a constant that maps to zero as its first element. This is because:

如上例中所示，Corpus 枚举类型的第一个枚举值是 0，每一个枚举值定义都会与一个常量映射，而这些常量的第一个常量值必须为 0，原因如下：

- There must be a zero value, so that we can use 0 as a numeric default value.
- 必须有一个 0 作为值，以至于我们可是使用 0 作为默认值
- The zero value needs to be the first element, for compatibility with the proto2 semantics where the first enum value is always the default.
- 第一个元素的值取 0，用于与第一个元素枚举值作为默认值的 proto2 语义兼容

You can define aliases by assigning the same value to different enum constants. To do this you need to set the allow_alias option to true, otherwise the protocol compiler will generate an error message when aliases are found.

枚举类型允许你定义别名，别名的作用是分配不中的标量，使用相同的常量值，使用别名只需要在定义枚举类型的第一行中添加 allow_alias 选项，并将值设置为 true 即可，如果没有设置该值就是用别名，在编译的时候会报错。

```
message MyMessage1 {
  enum EnumAllowingAlias {
    option allow_alias = true;
    UNKNOWN = 0;
    STARTED = 1;
    RUNNING = 1;
  }
}
message MyMessage2 {
  enum EnumNotAllowingAlias {
    UNKNOWN = 0;
    STARTED = 1;
    // RUNNING = 1;  // Uncommenting this line will cause a compile error inside Google and a warning message outside.
  }
}
```

Enumerator constants must be in the range of a 32-bit integer. Since enum values use varint encoding on the wire, negative values are inefficient and thus not recommended. You can define enums within a message definition, as in the above example, or outside – these enums can be reused in any message definition in your .proto file. You can also use an enum type declared in one message as the type of a field in a different message, using the syntax _MessageType_._EnumType_.

枚举器常量必须在 32 位整数的范围内。由于枚举值在线上使用 varint 编码，负值的效率很低，因此不推荐使用。你可以在消息定义中定义枚举，就像上面的例子一样，也可以在外面定义–这些枚举可以在你的.proto 文件中的任何消息定义中重复使用。你也可以使用在一个消息中声明的枚举类型作为不同消息中的字段类型，使用的语法是*MessageType*._EnumType_.

When you run the protocol buffer compiler on a .proto that uses an enum, the generated code will have a corresponding enum for Java or C++, a special EnumDescriptor class for Python that’s used to create a set of symbolic constants with integer values in the runtime-generated class.

当你在使用枚举的.proto 上运行协议缓冲编译器时，生成的代码会有一个相应的枚举，用于 Java 或 C++，一个特殊的 EnumDescriptor 类，用于在运行时生成的类中创建一组带有整数值的符号常量。

> :warning: **Caution:** the generated code may be subject to language-specific limitations on the number of enumerators (low thousands for one language). Please review the limitations for the languages you plan to use. :warning: **警告:** 生成的代码可能受到特定语言的枚举数限制(单种语言的数量低于千)。请检查您计划使用的语言的限制

During deserialization, unrecognized enum values will be preserved in the message, though how this is represented when the message is deserialized is language-dependent. In languages that support open enum types with values outside the range of specified symbols, such as C++ and Go, the unknown enum value is simply stored as its underlying integer representation. In languages with closed enum types such as Java, a case in the enum is used to represent an unrecognized value, and the underlying integer can be accessed with special accessors. In either case, if the message is serialized the unrecognized value will still be serialized with the message.

在反序列化过程中，未被识别的枚举值将被保存在消息中，不过在消息被反序列化时，如何表示这些值取决于语言。在支持开放枚举类型的语言中，其值在指定的符号范围之外，如 C++和 Go，未知的枚举值将被简单地存储为其底层的整数表示。在具有封闭的枚举类型的语言中，如 Java，枚举中的一个 case 被用来表示一个未识别的值，底层的整数可以用特殊的访问器来访问。无论哪种情况，如果消息被序列化，未被识别的值仍然会和消息一起序列化。

For more information about how to work with message enums in your applications, see the generated code guide for your chosen language.

关于如何在应用程序中使用消息枚举的更多信息，请参见所选语言的生成代码指南。

#### Reserved Values

If you update an enum type by entirely removing an enum entry, or commenting it out, future users can reuse the numeric value when making their own updates to the type. This can cause severe issues if they later load old versions of the same .proto, including data corruption, privacy bugs, and so on. One way to make sure this doesn’t happen is to specify that the numeric values (and/or names, which can also cause issues for JSON serialization) of your deleted entries are reserved. The protocol buffer compiler will complain if any future users try to use these identifiers. You can specify that your reserved numeric value range goes up to the maximum possible value using the max keyword.

如果你更新一个枚举类型，完全删除一个枚举条目，或者将其注释出来，那么未来的用户在对该类型进行更新时可以重新使用数值。如果他们以后加载同一.proto 的旧版本，这可能会导致严重的问题，包括数据损坏、隐私错误等。确保这种情况不会发生的一种方法是指定保留你删除条目的数值（和/或名称，这也会导致 JSON 序列化的问题）。如果未来的用户试图使用这些标识符，协议缓冲区编译器会抱怨。您可以使用 max 关键字指定您保留的数值范围上升到最大可能的值。

```
enum Foo {
  reserved 2, 15, 9 to 11, 40 to max;
  reserved "FOO", "BAR";
}
```

Note that you can’t mix field names and numeric values in the same reserved statement.

请注意，不能在同一个保留语句中混合使用字段名和数值。

### 引用其他的消息类型（Using Other Message Types）

You can use other message types as field types. For example, let’s say you wanted to include Result messages in each SearchResponse message – to do this, you can define a Result message type in the same .proto and then specify a field of type Result in SearchResponse:

你可以使用其他消息类型作为字段类型。例如，假设你想在每个 SearchResponse 消息中包含 Result 消息–要做到这一点，你可以在同一个.proto 中定义一个 Result 消息类型，然后在 SearchResponse 中指定一个 Result 类型的字段。

```
message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```

#### 导入其他 proto 中定义的消息 （Importing Definitions）

Note that this feature is not available in Java.

注意，这个特性在 Java 中是不可用的。

In the above example, the Result message type is defined in the same file as SearchResponse – what if the message type you want to use as a field type is already defined in another .proto file?

在上面的例子中，Result 和 SearchResponse 消息类型被定义在同一个.proto 文件中，如果把他们分成两个文件定义，应该如何引用呢？

You can use definitions from other .proto files by importing them. To import another .proto’s definitions, you add an import statement to the top of your file:

proto 中为我们提供了 import 关键字用于引入不同.proto 文件中的消息类型,你可以在你的.proto 文件的顶部加入如下语句因为其他.proto 文件的消息类型：

```
import "myproject/other_protos.proto";
```

By default you can only use definitions from directly imported .proto files. However, sometimes you may need to move a .proto file to a new location. Instead of moving the .proto file directly and updating all the call sites in a single change, now you can put a dummy .proto file in the old location to forward all the imports to the new location using the import public notion. import public dependencies can be transitively relied upon by anyone importing the proto containing the import public statement. For example:

默认情况下，您只能使用直接导入的.proto 文件中的定义。然而，有时您可能需要将.proto 文件移动到一个新的位置。与其直接移动.proto 文件并在一次更改中更新所有的调用站点，现在您可以在旧位置放置一个虚拟的.proto 文件，以使用 import public 概念将所有的导入转发到新位置。 import public 依赖可以被任何导入包含 import public 声明的 proto 的人中转依赖。例如

```
// new.proto
// All definitions are moved here
// old.proto
// This is the proto that all clients are importing.
import public "new.proto";
import "other.proto";
// client.proto
import "old.proto";
// You use definitions from old.proto and new.proto, but not other.proto
```

The protocol compiler searches for imported files in a set of directories specified on the protocol compiler command line using the -I/–proto_path flag. If no flag was given, it looks in the directory in which the compiler was invoked. In general you should set the –proto_path flag to the root of your project and use fully qualified names for all imports.

协议编译器使用-I/–proto_path 标志在协议编译器命令行指定的一组目录中搜索导入的文件。如果没有给出标志，则在编译器被调用的目录中查找。一般来说，你应该将–proto_path 标志设置为项目的根目录，并且对所有的导入文件使用完全限定的名称。

#### Using proto2 Message Types

It’s possible to import proto2 message types and use them in your proto3 messages, and vice versa. However, proto2 enums cannot be used directly in proto3 syntax (it’s okay if an imported proto2 message uses them).

可以导入 proto2 消息类型并在 proto3 消息中使用它们，反之亦然。但是，proto2 枚举不能直接在 proto3 语法中使用（如果导入的 proto2 消息使用了它们，也是可以的）。

### 内嵌类型(Nested Types)

You can define and use message types inside other message types, as in the following example – here the Result message is defined inside the SearchResponse message:

你可以在其他消息类型中定义和使用消息类型，在下面的例子中在 SearchResponse 消息体中定义了一个 Result 消息并使用。

```
message SearchResponse {
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```

If you want to reuse this message type outside its parent message type, you refer to it as _Parent_._Type_:

如果想在其他的消息体引用 Result 这个消息，可以 Parent.Type 这样引用，例子：

```go
message SomeOtherMessage {
  SearchResponse.Result result = 1;
}
```

You can nest messages as deeply as you like:

消息还可以深层的嵌套定义，如下例子：

```java
message Outer {                  // Level 0
  message MiddleAA {  // Level 1
    message Inner {   // Level 2
      int64 ival = 1;
      bool  booly = 2;
    }
  }
  message MiddleBB {  // Level 1
    message Inner {   // Level 2
      int32 ival = 1;
      bool  booly = 2;
    }
  }
}
```

### 更新消息类型（Updating A Message Type）

If an existing message type no longer meets all your needs – for example, you’d like the message format to have an extra field – but you’d still like to use code created with the old format, don’t worry! It’s very simple to update message types without breaking any of your existing code. Just remember the following rules:

如果现有的消息类型不再满足您的所有需要——例如，您希望消息格式有一个额外的字段——但是您仍然希望使用用旧格式创建的代码，不要担心！在不破坏任何现有代码的情况下更新消息类型非常简单。只要记住以下规则:

- Don’t change the field numbers for any existing fields.
- 不要改变任何现有字段的字段号。
- If you add new fields, any messages serialized by code using your “old” message format can still be parsed by your new generated code. You should keep in mind the default values for these elements so that new code can properly interact with messages generated by old code. Similarly, messages created by your new code can be parsed by your old code: old binaries simply ignore the new field when parsing. See the Unknown Fields section for details.
- 如果你添加了新的字段，任何使用 “旧 “消息格式的代码序列化的消息仍然可以被你新生成的代码解析。你应该记住这些元素的默认值，这样新代码就可以与旧代码生成的消息正确地交互。同样，你的新代码创建的消息也可以被你的旧代码解析：旧的二进制文件在解析时只是忽略新字段。详情请参阅未知字段部分。
- Fields can be removed, as long as the field number is not used again in your updated message type. You may want to rename the field instead, perhaps adding the prefix “OBSOLETE\_”, or make the field number reserved, so that future users of your .proto can’t accidentally reuse the number.
- 字段可以被删除，只要字段号不在你更新的消息类型中再次使用。你可能想重新命名这个字段，也许加上前缀 “OBSOLETE\_“，或者保留这个字段号，这样你的.proto 的未来用户就不会意外地重复使用这个号码。
- int32, uint32, int64, uint64, and bool are all compatible – this means you can change a field from one of these types to another without breaking forwards- or backwards-compatibility. If a number is parsed from the wire which doesn’t fit in the corresponding type, you will get the same effect as if you had cast the number to that type in C++ (e.g. if a 64-bit number is read as an int32, it will be truncated to 32 bits).
- `int32`、`uint32`、`int64`、`uint64`和`bool`都是兼容的–这意味着你可以将一个字段从这些类型中的一个改成另一个，而不会破坏向前或向后的兼容性。如果从线上解析出一个不适合相应类型的数字，你将得到与在 C++中把数字投到该类型中一样的效果（例如，如果一个 64 位的数字被读作 int32，它将被截断为 32 位）
- sint32 and sint64 are compatible with each other but are not compatible with the other integer types.
- sint32 和 sint64 相互兼容，但与其他整数类型不兼容。
- string and bytes are compatible as long as the bytes are valid UTF-8.
- 字符串和字节是兼容的，只要字节是有效的 UTF-8。
- Embedded messages are compatible with bytes if the bytes contain an encoded version of the message.
- 如果字节包含消息的编码版本，则嵌入式消息与字节兼容。
- fixed32 is compatible with sfixed32, and fixed64 with sfixed64.
- fixed32 与 sfixed32 兼容，fixed64 与 sfixed64 兼容。
- For string, bytes, and message fields, optional is compatible with repeated. Given serialized data of a repeated field as input, clients that expect this field to be optional will take the last input value if it’s a primitive type field or merge all input elements if it’s a message type field. Note that this is not generally safe for numeric types, including bools and enums. Repeated fields of numeric types can be serialized in the packed format, which will not be parsed correctly when an optional field is expected.
- 对于字符串、字节和消息字段，`option`与`repeat`兼容。给定重复字段的序列化数据作为输入，如果是基元类型字段，期望该字段是可选的客户端将取最后一个输入值，如果是消息类型字段，则合并所有输入元素。请注意，这对于数值类型，包括`bools`和`enums`，一般来说并不安全。数值类型的重复字段可以以打包格式进行序列化，当期望使用可选字段时，它将不会被正确解析。
- enum is compatible with int32, uint32, int64, and uint64 in terms of wire format (note that values will be truncated if they don’t fit). However be aware that client code may treat them differently when the message is deserialized: for example, unrecognized proto3 enum types will be preserved in the message, but how this is represented when the message is deserialized is language-dependent. Int fields always just preserve their value.
- `enum`与`int32`、`uint32`、`int64`和`uint64`在线格式上是兼容的（注意，如果不适合，值会被截断）。然而要注意，当消息被反序列化时，客户端代码可能会以不同的方式处理它们：例如，未被识别的 proto3 枚举类型将在消息中被保留，但当消息被反序列化时，如何表示这一点取决于语言。Int 字段总是只保留其值。
- Changing a single value into a member of a new `oneof` is safe and binary compatible. Moving multiple fields into a new `oneof` may be safe if you are sure that no code sets more than one at a time. Moving any fields into an existing `oneof` is not safe.
- 将一个值改成一个新的值的成员是安全的，二进制兼容的。将多个字段移动到一个新的值中可能是安全的，如果你确定没有代码集一次超过一个。将任何字段移动到现有的值中是不安全的。

### 未知字段（Unknown Fields）

Unknown fields are well-formed protocol buffer serialized data representing fields that the parser does not recognize. For example, when an old binary parses data sent by a new binary with new fields, those new fields become unknown fields in the old binary.

未知字段是格式良好的协议缓冲区序列化数据，代表解析器不认识的字段。例如，当一个旧的二进制文件解析一个带有新字段的新二进制文件发送的数据时，这些新字段就成为旧二进制文件中的未知字段。

Originally, proto3 messages always discarded unknown fields during parsing, but in version 3.5 we reintroduced the preservation of unknown fields to match the proto2 behavior. In versions 3.5 and later, unknown fields are retained during parsing and included in the serialized output.

最初，proto3 消息在解析过程中总是丢弃未知字段，但在 3.5 版本中，我们重新引入了未知字段的保留，以匹配 proto2 的行为。在 3.5 及以后的版本中，未知字段在解析过程中被保留，并包含在序列化输出中。

### 任意消息类型（Any）

> :warning: **目前不要使用**

The Any message type lets you use messages as embedded types without having their .proto definition. An Any contains an arbitrary serialized message as bytes, along with a URL that acts as a globally unique identifier for and resolves to that message’s type. To use the Any type, you need to import google/protobuf/any.proto.

Any 消息类型允许您将消息作为嵌入类型使用，而不需要。 原始定义。 任意包含一个任意序列化的字节消息，以及一个作为全局唯一标识符的 URL，并解析为该消息的类型。 要使用 Any 类型，你需要导入 google / protobuf / Any。 原始的。

```go
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

The default type URL for a given message type is type.googleapis.com/_packagename_._messagename_.

给定消息类型的默认类型 URL 是 type.googleapis.com/_packagename_._messagename_。

Different language implementations will support runtime library helpers to pack and unpack Any values in a typesafe manner – for example, in Java, the Any type will have special pack() and unpack() accessors, while in C++ there are PackFrom() and UnpackTo() methods:

不同的语言实现将支持运行时库 helpers 以类型安全的方式打包和解包任何值，例如，在 Java 中，Any 类型将有特殊的 pack()和 unpack()访问器，而在 C++中则有 PackFrom()和 UnpackTo()方法。

```C++
// Storing an arbitrary message type in Any.
NetworkErrorDetails details = ...;
ErrorStatus status;
status.add_details()->PackFrom(details);

// Reading an arbitrary message from Any.
ErrorStatus status = ...;
for (const Any& detail : status.details()) {
  if (detail.Is<NetworkErrorDetails>()) {
    NetworkErrorDetails network_error;
    detail.UnpackTo(&network_error);
    ... processing network_error ...
  }
}
```

Currently the runtime libraries for working with Any types are under development.

目前正在开发用于处理任何类型的运行时库。

If you are already familiar with proto2 syntax, the Any can hold arbitrary proto3 messages, similar to proto2 messages which can allow extensions.

如果你已经熟悉 proto2 语法，Any 可以保存任意的 proto3 消息，类似于 proto2 消息，可以允许扩展。

### Oneof

> :warning: **对性能没有要求，不要使用**

If you have a message with many fields and where at most one field will be set at the same time, you can enforce this behavior and save memory by using the oneof feature.

如果您有一条包含许多字段的消息，并且最多同时设置一个字段，那么您可以通过使用其中一个特性来强制执行此行为并节省内存。

Oneof fields are like regular fields except all the fields in a oneof share memory, and at most one field can be set at the same time. Setting any member of the oneof automatically clears all the other members. You can check which value in a oneof is set (if any) using a special case() or WhichOneof() method, depending on your chosen language.

其中一个字段类似于常规字段，只不过共享内存中的一个字段中的所有字段都是常规字段，而且最多可以同时设置一个字段。设置其中的任何成员都会自动清除所有其他成员。根据所选择的语言，可以使用特殊 case ()或 WhichOneof ()方法检查 one of 中的哪个值被设置(如果有的话)。

#### Oneof 使用（Using Oneof）

To define a oneof in your .proto you use the oneof keyword followed by your oneof name, in this case test_oneof:

定义一个你的。你可以使用 one of keyword 后面跟着你的 one of name，在这个例子中 test one of:

```
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
```

You then add your oneof fields to the oneof definition. You can add fields of any type, except map fields and repeated fields.

然后将其中一个字段添加到该字段的定义中。您可以添加任何类型的字段，除了映射字段和重复字段。

In your generated code, oneof fields have the same getters and setters as regular fields. You also get a special method for checking which value (if any) in the oneof is set. You can find out more about the oneof API for your chosen language in the relevant API reference.

在生成的代码中，其中一个字段具有与常规字段相同的 getter 和 setter。 您还可以获得一个特殊的方法来检查其中一个设置了哪个值(如果有的话)。 您可以在相关的 API 参考文献中找到更多关于所选语言的 API。

#### Oneof 特性（Oneof Features）

- Setting a oneof field will automatically clear all other members of the oneof. So if you set several oneof fields, only the last field you set will still have a value.
- 设置一个字段将自动清除该字段的所有其他成员。因此，如果您设置了多个字段之一，那么只有最后设置的字段仍然具有值。

```
SampleMessage message;
message.set_name("name");
CHECK(message.has_name());
message.mutable_sub_message();   // Will clear name field.
CHECK(!message.has_name());
```

- If the parser encounters multiple members of the same oneof on the wire, only the last member seen is used in the parsed message.
- 如果解析器在连接中遇到同一个成员的多个成员，则只有最后看到的成员用于解析消息。
- A oneof cannot be repeated.
- oneof 是不能重复的。
- Reflection APIs work for oneof fields.
- 反射 APIs 适用于一个字段。
- If you set a oneof field to the default value (such as setting an int32 oneof field to 0), the “case” of that oneof field will be set, and the value will be serialized on the wire.
- 如果将 oneof 字段设置为默认值(例如将 int32 oneof 字段设置为 0)，则将设置该字段的 “case” ，并在连接上序列化该值。
- If you’re using C++, make sure your code doesn’t cause memory crashes. The following sample code will crash because sub_message was already deleted by calling the set_name() method.
- 如果你使用 c + + ，确保你的代码不会导致内存崩溃。下面的示例代码将崩溃，因为通过调用 set*name ()方法已经删除了 sub * message。

```
SampleMessage message;
SubMessage* sub_message = message.mutable_sub_message();
message.set_name("name");      // Will delete sub_message
sub_message->set_...            // Crashes here
```

- Again in C++, if you Swap() two messages with oneofs, each message will end up with the other’s oneof case: in the example below, msg1 will have a sub_message and msg2 will have a name.
- 同样在 c + + 中，如果您使用 Swap ()两条消息中的一条，每条消息将以另一条消息中的一条结束: 在下面的例子中，msg1 将有一条子消息，msg2 将有一个名称。

```
SampleMessage msg1;
msg1.set_name("name");
SampleMessage msg2;
msg2.mutable_sub_message();
msg1.swap(&msg2);
CHECK(msg1.has_sub_message());
CHECK(msg2.has_name());
```

#### 向后兼容性问题（Backwards-compatibility issues）

Be careful when adding or removing oneof fields. If checking the value of a oneof returns None/NOT_SET, it could mean that the oneof has not been set or it has been set to a field in a different version of the oneof. There is no way to tell the difference, since there’s no way to know if an unknown field on the wire is a member of the oneof.

添加或删除一个字段时要小心。如果检查 oneof 的值返回 None/NOT_SET，这可能意味着 oneof 没有被设置，或者它已经被设置为 oneof 的不同版本中的一个字段。没有办法区分，因为没有办法知道电线上的未知字段是否是。

##### 标签重用问题（Tag Reuse Issues）

- Move fields into or out of a oneof: You may lose some of your information (some fields will be cleared) after the message is serialized and parsed. However, you can safely move a single field into a new oneof and may be able to move multiple fields if it is known that only one is ever set.
- 将字段移入或移出`oneof`。在电文被序列化和解析后，你可能会丢失一些信息（一些字段将被清除）。但是，您可以安全地将一个字段移入一个新的`oneof`中，如果知道只有一个字段被设置，您也可以将多个字段移入。
- Delete a oneof field and add it back: This may clear your currently set oneof field after the message is serialized and parsed.
- 删除一个`oneof`字段，然后再添加回来。这可能会在电文序列化和解析后清除当前设置的`oneof`字段。
- Split or merge oneof: This has similar issues to moving regular fields.
- 拆分或合并`oneof`。这与移动常规领域有类似的问题。

### Maps

If you want to create an associative map as part of your data definition, protocol buffers provides a handy shortcut syntax: 如果你想创建一个关联映射作为你数据定义的一部分，`protocol buffers`提供了一个方便的快捷语法:

```
map<key_type, value_type> map_field = N;
```

…where the key_type can be any integral or string type (so, any scalar type except for floating point types and bytes). Note that enum is not a valid key_type. The value_type can be any type except another map.

其中 key_type 可以是任何积分或字符串类型(所以，除了浮点类型和字节以外，任何标量类型)。注意，enum 不是有效的 key_type。value_type 可以是任何类型，除了另一个映射。

So, for example, if you wanted to create a map of projects where each Project message is associated with a string key, you could define it like this:

例如，如果你想创建一个项目映射，其中每个项目消息都与一个字符串键相关联，你可以这样定义:

proto 支持 map 属性类型的定义，语法如下： map map_field = N; key_type 可以是任何整数或字符串类型（除浮点类型和字节之外的任何标量类型,枚举类型也是不合法的 key 类型），value_type 可以是任何类型的数据。

```
map<string, Project> projects = 3;
```

- Map fields cannot be repeated.
- Map 字段不能重复
- Wire format ordering and map iteration ordering of map values is undefined, so you cannot rely on your map items being in a particular order.
- When generating text format for a .proto, maps are sorted by key. Numeric keys are sorted numerically.
- When parsing from the wire or when merging, if there are duplicate map keys the last key seen is used. When parsing a map from text format, parsing may fail if there are duplicate keys.
- If you provide a key but no value for a map field, the behavior when the field is serialized is language-dependent. In C++, Java, and Python the default value for the type is serialized, while in other languages nothing is serialized.

The generated map API is currently available for all proto3 supported languages. You can find out more about the map API for your chosen language in the relevant API reference.

生成的映射 API 目前可用于所有支持 proto3 的语言。map 更具体的使用方式参看 API。

#### 向后兼容性（Backwards compatibility）

The map syntax is equivalent to the following on the wire, so protocol buffers implementations that do not support maps can still handle your data:

```
message MapFieldEntry {
  key_type key = 1;
  value_type value = 2;
}
```

repeated MapFieldEntry map_field = N; Any protocol buffers implementation that supports maps must both produce and accept data that can be accepted by the above definition.

### 包（Packages）

You can add an optional package specifier to a .proto file to prevent name clashes between protocol message types.

可以为 proto 文件指定包名，防止消息命名冲突。

```
package foo.bar;
message Open { ... }
```

You can then use the package specifier when defining fields of your message type:

当你在为消息类型定义属性的时候，你可以通过命名.类型的形式来使用已经定义好的消息类型，如下：

```
message Foo {
  ...
  foo.bar.Open open = 1;
  ...
}
```

The way a package specifier affects the generated code depends on your chosen language:

包说明符影响生成代码的方式取决于您选择的语言:

… - In Go, the package is used as the Go package name, unless you explicitly provide an option go_package in your .proto file. - 在 Go 语言中，使用 Go 包名 …

#### 包和名称解析（Packages and Name Resolution）

Type name resolution in the protocol buffer language works like C++: first the innermost scope is searched, then the next-innermost, and so on, with each package considered to be “inner” to its parent package. A leading ‘.’ (for example, .foo.bar.Baz) means to start from the outermost scope instead.

在 protocol buffer 语言中，类型名称解析的工作原理类似于 c + + : 首先搜索最内层的作用域，然后搜索下一个最内层的作用域，依此类推，每个包都被认为是其父包的“ inner”。一个引导(例如:。Foo.bar.意思是从最外层开始。

The protocol buffer compiler resolves all type names by parsing the imported .proto files. The code generator for each language knows how to refer to each type in that language, even if it has different scoping rules.

协议缓冲区编译器通过解析导入的原始文件。每种语言的代码生成器都知道如何引用该语言中的每种类型，即使它有不同的作用域规则。

### 服务定义 (Defining Services)

If you want to use your message types with an RPC (Remote Procedure Call) system, you can define an RPC service interface in a .proto file and the protocol buffer compiler will generate service interface code and stubs in your chosen language. So, for example, if you want to define an RPC service with a method that takes your SearchRequest and returns a SearchResponse, you can define it in your .proto file as follows:

如果你想在 RPC(远程过程调用)中使用已经定义好的消息类型，你可以在.proto 文件中定一个消息服务接口,protocol buffer 编译器会生成对应语言的接口代码。
参数跟返回值都是前面定义的消息类型。

```
service SearchService {
    //  方法名  方法参数                 返回值
    rpc Search(SearchRequest) returns (SearchResponse);
}
```

The most straightforward RPC system to use with protocol buffers is gRPC: a language- and platform-neutral open source RPC system developed at Google. gRPC works particularly well with protocol buffers and lets you generate the relevant RPC code directly from your .proto files using a special protocol buffer compiler plugin.

最简单易用的 RPC 系统就是 gRPC: Google 开发的一个语言和平台无关的开源 RPC 系统。这个系统可以用于协议缓冲。gRPC 特别适用于协议缓冲，它可以让你直接从你的。原型文件使用特殊的协议缓冲编译器插件。

If you don’t want to use gRPC, it’s also possible to use protocol buffers with your own RPC implementation. You can find out more about this in the Proto2 Language Guide.

如果你不想使用 gRPC，你也可以在你自己的 RPC 实现中使用协议缓冲。你可以在《 Proto2 语言指南》中找到更多相关信息。

There are also a number of ongoing third-party projects to develop RPC implementations for Protocol Buffers. For a list of links to projects we know about, see the third-party add-ons wiki page.

还有一些正在进行的第三方项目正在开发 RPC 的实施协议缓冲。有关我们所知道的项目的链接列表，请参阅第三方添加项 wiki 页面。

### JSON Mapping

Proto3 supports a canonical encoding in JSON, making it easier to share data between systems. The encoding is described on a type-by-type basis in the table below.

If a value is missing in the JSON-encoded data or if its value is null, it will be interpreted as the appropriate default value when parsed into a protocol buffer. If a field has the default value in the protocol buffer, it will be omitted in the JSON-encoded data by default to save space. An implementation may provide options to emit fields with default values in the JSON-encoded output.

| proto3                 | JSON          | JSON example                              | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ---------------------- | ------------- | ----------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| message                | object        | `{"fooBar": v, "g": null, …}`             | Generates JSON objects. Message field names are mapped to lowerCamelCase and become JSON object keys. If the `json_name` field option is specified, the specified value will be used as the key instead. Parsers accept both the lowerCamelCase name (or the one specified by the `json_name` option) and the original proto field name. `null` is an accepted value for all field types and treated as the default value of the corresponding field type. |
| enum                   | string        | `"FOO_BAR"`                               | The name of the enum value as specified in proto is used. Parsers accept both enum names and integer values.                                                                                                                                                                                                                                                                                                                                               |
| map                    | object        | `{"k": v, …}`                             | All keys are converted to strings.                                                                                                                                                                                                                                                                                                                                                                                                                         |
| repeated V             | array         | `[v, …]`                                  | `null` is accepted as the empty list `[]`.                                                                                                                                                                                                                                                                                                                                                                                                                 |
| bool                   | true, false   | `true, false`                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| string                 | string        | `"Hello World!"`                          |                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| bytes                  | base64 string | `"YWJjMTIzIT8kKiYoKSctPUB+"`              | JSON value will be the data encoded as a string using standard base64 encoding with paddings. Either standard or URL-safe base64 encoding with/without paddings are accepted.                                                                                                                                                                                                                                                                              |
| int32, fixed32, uint32 | number        | `1, -10, 0`                               | JSON value will be a decimal number. Either numbers or strings are accepted.                                                                                                                                                                                                                                                                                                                                                                               |
| int64, fixed64, uint64 | string        | `"1", "-10"`                              | JSON value will be a decimal string. Either numbers or strings are accepted.                                                                                                                                                                                                                                                                                                                                                                               |
| float, double          | number        | `1.1, -10.0, 0, "NaN", "Infinity"`        | JSON value will be a number or one of the special string values “NaN”, “Infinity”, and “-Infinity”. Either numbers or strings are accepted. Exponent notation is also accepted. -0 is considered equivalent to 0.                                                                                                                                                                                                                                          |
| Any                    | `object`      | `{"@type": "url", "f": v, … }`            | If the Any contains a value that has a special JSON mapping, it will be converted as follows: `{"@type": xxx, "value": yyy}`. Otherwise, the value will be converted into a JSON object, and the `"@type"` field will be inserted to indicate the actual data type.                                                                                                                                                                                        |
| Timestamp              | string        | `"1972-01-01T10:00:20.021Z"`              | Uses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9 fractional digits. Offsets other than “Z” are also accepted.                                                                                                                                                                                                                                                                                                       |
| Duration               | string        | `"1.000340012s", "1s"`                    | Generated output always contains 0, 3, 6, or 9 fractional digits, depending on required precision, followed by the suffix “s”. Accepted are any fractional digits (also none) as long as they fit into nano-seconds precision and the suffix “s” is required.                                                                                                                                                                                              |
| Struct                 | `object`      | `{ … }`                                   | Any JSON object. See `struct.proto`.                                                                                                                                                                                                                                                                                                                                                                                                                       |
| Wrapper types          | various types | `2, "2", "foo", true, "true", null, 0, …` | Wrappers use the same representation in JSON as the wrapped primitive type, except that `null` is allowed and preserved during data conversion and transfer.                                                                                                                                                                                                                                                                                               |
| FieldMask              | string        | `"f.fooBar,h"`                            | See `field_mask.proto`.                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ListValue              | array         | `[foo, bar, …]`                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| Value                  | value         |                                           | Any JSON value. Check [google.protobuf.Value](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Value) for details.                                                                                                                                                                                                                                                                                            |
| NullValue              | null          |                                           | JSON null                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| Empty                  | object        | `{}`                                      | An empty JSON object                                                                                                                                                                                                                                                                                                                                                                                                                                       |

#### JSON options

A proto3 JSON implementation may provide the following options:

- Emit fields with default values: Fields with default values are omitted by default in proto3 JSON output. An implementation may provide an option to override this behavior and output fields with their default values.
- Ignore unknown fields: Proto3 JSON parser should reject unknown fields by default but may provide an option to ignore unknown fields in parsing.
- Use proto field name instead of lowerCamelCase name: By default proto3 JSON printer should convert the field name to lowerCamelCase and use that as the JSON name. An implementation may provide an option to use proto field name as the JSON name instead. Proto3 JSON parsers are required to accept both the converted lowerCamelCase name and the proto field name.
- Emit enum values as integers instead of strings: The name of an enum value is used by default in JSON output. An option may be provided to use the numeric value of the enum value instead.

### 选项（Options）

Individual declarations in a .proto file can be annotated with a number of options. Options do not change the overall meaning of a declaration, but may affect the way it is handled in a particular context. The complete list of available options is defined in google/protobuf/descriptor.proto.

Some options are file-level options, meaning they should be written at the top-level scope, not inside any message, enum, or service definition. Some options are message-level options, meaning they should be written inside message definitions. Some options are field-level options, meaning they should be written inside field definitions. Options can also be written on enum types, enum values, oneof fields, service types, and service methods; however, no useful options currently exist for any of these.

Here are a few of the most commonly used options:

- java_package (file option): The package you want to use for your generated Java classes. If no explicit java_package option is given in the .proto file, then by default the proto package (specified using the “package” keyword in the .proto file) will be used. However, proto packages generally do not make good Java packages since proto packages are not expected to start with reverse domain names. If not generating Java code, this option has no effect.

```
option java_package = "com.example.foo";
```

- java_outer_classname (file option): The class name for the outermost Java class (and hence the file name) you want to generate. If no explicit java_outer_classname is specified in the .proto file, the class name will be constructed by converting the .proto file name to camel-case (so foo_bar.proto becomes FooBar.java). If not generating Java code, this option has no effect.

```
option java_outer_classname = "Ponycopter";
```

- java_multiple_files (file option): If false, only a single .java file will be generated for this .proto file, and all the Java classes/enums/etc. generated for the top-level messages, services, and enumerations will be nested inside of an outer class (see java_outer_classname). If true, separate .java files will be generated for each of the Java classes/enums/etc. generated for the top-level messages, services, and enumerations, and the Java “outer class” generated for this .proto file won’t contain any nested classes/enums/etc. This is a Boolean option which defaults to false. If not generating Java code, this option has no effect.

```
option java_multiple_files = true;
```

- optimize_for (file option): Can be set to SPEED, CODE_SIZE, or LITE_RUNTIME. This affects the C++ and Java code generators (and possibly third-party generators) in the following ways:
  - SPEED (default): The protocol buffer compiler will generate code for serializing, parsing, and performing other common operations on your message types. This code is highly optimized.
  - CODE_SIZE: The protocol buffer compiler will generate minimal classes and will rely on shared, reflection-based code to implement serialialization, parsing, and various other operations. The generated code will thus be much smaller than with SPEED, but operations will be slower. Classes will still implement exactly the same public API as they do in SPEED mode. This mode is most useful in apps that contain a very large number .proto files and do not need all of them to be blindingly fast.
  - LITE_RUNTIME: The protocol buffer compiler will generate classes that depend only on the “lite” runtime library (libprotobuf-lite instead of libprotobuf). The lite runtime is much smaller than the full library (around an order of magnitude smaller) but omits certain features like descriptors and reflection. This is particularly useful for apps running on constrained platforms like mobile phones. The compiler will still generate fast implementations of all methods as it does in SPEED mode. Generated classes will only implement the MessageLite interface in each language, which provides only a subset of the methods of the full Message interface. `option optimize_for = CODE_SIZE;`
- cc_enable_arenas (file option): Enables arena allocation for C++ generated code.
- objc_class_prefix (file option): Sets the Objective-C class prefix which is prepended to all Objective-C generated classes and enums from this .proto. There is no default. You should use prefixes that are between 3-5 uppercase characters as recommended by Apple. Note that all 2 letter prefixes are reserved by Apple.
- deprecated (field option): If set to true, indicates that the field is deprecated and should not be used by new code. In most languages this has no actual effect. In Java, this becomes a @Deprecated annotation. In the future, other language-specific code generators may generate deprecation annotations on the field’s accessors, which will in turn cause a warning to be emitted when compiling code which attempts to use the field. If the field is not used by anyone and you want to prevent new users from using it, consider replacing the field declaration with a reserved statement.

```
int32 old_field = 6 [deprecated = true];
```

#### 自定义选项 (Custom Options)

Protocol Buffers also allows you to define and use your own options. This is an advanced feature which most people don’t need. If you do think you need to create your own options, see the Proto2 Language Guide for details. Note that creating custom options uses extensions, which are permitted only for custom options in proto3.

### 生成你的类 (Generating Your Classes)

To generate the Java, Python, C++, Go, Ruby, Objective-C, or C# code you need to work with the message types defined in a .proto file, you need to run the protocol buffer compiler protoc on the .proto. If you haven’t installed the compiler, download the package and follow the instructions in the README. For Go, you also need to install a special code generator plugin for the compiler: you can find this and installation instructions in the golang/protobuf repository on GitHub.

要生成 Java、 Python、 c + + 、 Go、 Ruby、 Objective-C 或 c # 代码，您需要使用。原型文件，你需要运行协议缓冲编译器协议。原始的。如果尚未安装编译器，请下载该包并按照 README 中的说明操作。对于 Go，您还需要为编译器安装一个特殊的代码生成器插件: 您可以在 GitHub 上的 golang/protobuf 存储库中找到这个插件和安装指令。

The Protocol Compiler is invoked as follows: 协议编译器的调用方式如下:

```
protoc --proto_path=IMPORT_PATH --cpp_out=DST_DIR --java_out=DST_DIR --python_out=DST_DIR --go_out=DST_DIR --ruby_out=DST_DIR --objc_out=DST_DIR --csharp_out=DST_DIR path/to/file.proto
```

- IMPORT_PATH specifies a directory in which to look for .proto files when resolving import directives. If omitted, the current directory is used. Multiple import directories can be specified by passing the –proto_path option multiple times; they will be searched in order.-I=\_IMPORT*PATH* can be used as a short form of –proto_path.
- 指定要在其中查找的目录
- You can provide one or more output directives: …
  - –go_out generates Go code in DST_DIR. See the Go generated code reference for more.
  - 生成 Go 代码 …
- You must provide one or more .proto files as input. Multiple .proto files can be specified at once. Although the files are named relative to the current directory, each file must reside in one of the IMPORT_PATHs so that the compiler can determine its canonical name.
- 你必须提供一个或多个`.proto`文件作为输入

​ 5 Go Protocol 基础

​

​

​

# Protocol Buffer Basics: Go

This tutorial provides a basic Go programmer’s introduction to working with protocol buffers, using the [proto3](https://developers.google.com/protocol-buffers/docs/proto3) version of the protocol buffers language. By walking through creating a simple example application, it shows you how to

本教程提供了一个基本的 Go 程序员使用协议缓冲区语言的[proto3](https://developers.google.com/protocol-buffers/docs/proto3)版本来处理协议缓冲区的介绍。通过创建一个简单的示例应用程序，它向您展示了如何进行以下操作

- Define message formats in a `.proto` file.
- Use the protocol buffer compiler.
- Use the Go protocol buffer API to write and read messages.

This isn’t a comprehensive guide to using protocol buffers in Go. For more detailed reference information, see the [Protocol Buffer Language Guide](https://developers.google.com/protocol-buffers/docs/proto3), the [Go API Reference](https://pkg.go.dev/google.golang.org/protobuf/proto), the [Go Generated Code Guide](https://developers.google.com/protocol-buffers/docs/reference/go-generated), and the [Encoding Reference](https://developers.google.com/protocol-buffers/docs/encoding).

这并不是一份在 Go 中使用协议缓冲区的全面指南。更详细的参考信息，请参见 [协议缓冲区语言指南](https://developers.google.com/protocol-buffers/docs/proto3)、[Go API 参考](https://pkg.go.dev/google.golang.org/protobuf/proto)、[Go 生成代码指南](https://developers.google.com/protocol-buffers/docs/reference/go-generated)和 [编码参考](https://developers.google.com/protocol-buffers/docs/encoding)。

## Why use protocol buffers?

The example we’re going to use is a very simple “address book” application that can read and write people’s contact details to and from a file. Each person in the address book has a name, an ID, an email address, and a contact phone number.

我们要使用的例子是一个非常简单的 “地址簿 “应用程序，它可以从一个文件中读取和写入人们的联系信息。地址簿中的每个人都有一个名字、一个 ID、一个电子邮件地址和一个联系电话。

How do you serialize and retrieve structured data like this? There are a few ways to solve this problem:

如何对这样的结构化数据进行序列化和检索？有几种方法可以解决这个问题。

- Use [gobs](https://golang.org/pkg/encoding/gob/) to serialize Go data structures. This is a good solution in a Go-specific environment, but it doesn’t work well if you need to share data with applications written for other platforms.
- 使用[gobs](https://golang.org/pkg/encoding/gob/)来序列化 Go 数据结构。在 Go 专用的环境中，这是一个很好的解决方案，但如果你需要与为其他平台编写的应用程序共享数据，它就不好用了。
- You can invent an ad-hoc way to encode the data items into a single string – such as encoding 4 ints as “12:3:-23:67”. This is a simple and flexible approach, although it does require writing one-off encoding and parsing code, and the parsing imposes a small run-time cost. This works best for encoding very simple data.
- 你可以发明一种特别的方式将数据项编码成一个单一的字符串–比如将 4 个 ints 编码为 “12:3:-23:67”。这是一种简单而灵活的方法，尽管它确实需要编写一次性的编码和解析代码，而且解析会带来少量的运行时成本。这对于编码非常简单的数据最有效。
- Serialize the data to XML. This approach can be very attractive since XML is (sort of) human readable and there are binding libraries for lots of languages. This can be a good choice if you want to share data with other applications/projects. However, XML is notoriously space intensive, and encoding/decoding it can impose a huge performance penalty on applications. Also, navigating an XML DOM tree is considerably more complicated than navigating simple fields in a class normally would be.
- 将数据序列化为 XML。这种方法非常有吸引力，因为 XML 是（某种程度上）人类可读的，而且有很多语言的绑定库。如果你想与其他应用程序/项目共享数据，这可能是一个不错的选择。然而，XML 是出了名的空间密集型，编码/解码会对应用程序造成巨大的性能损失。此外，浏览一个 XML DOM 树比浏览一个类中的简单字段要复杂得多。

Protocol buffers are the flexible, efficient, automated solution to solve exactly this problem. With protocol buffers, you write a `.proto` description of the data structure you wish to store. From that, the protocol buffer compiler creates a class that implements automatic encoding and parsing of the protocol buffer data with an efficient binary format. The generated class provides getters and setters for the fields that make up a protocol buffer and takes care of the details of reading and writing the protocol buffer as a unit. Importantly, the protocol buffer format supports the idea of extending the format over time in such a way that the code can still read data encoded with the old format.

协议缓冲区正是解决这个问题的灵活、高效、自动化的解决方案。有了协议缓冲区，你写一个`.proto`描述你想存储的数据结构。由此，协议缓冲编译器创建一个类，该类以高效的二进制格式实现协议缓冲数据的自动编码和解析。生成的类为构成协议缓冲区的字段提供了 getter 和 setter，并将协议缓冲区作为一个单元进行读写的细节处理。重要的是，协议缓冲区格式支持随着时间的推移而扩展格式的想法，这样代码仍然可以读取用旧格式编码的数据。

## Where to find the example code

Our example is a set of command-line applications for managing an address book data file, encoded using protocol buffers. The command `add_person_go` adds a new entry to the data file. The command `list_people_go` parses the data file and prints the data to the console.

我们的例子是一套管理通讯录数据文件的命令行应用程序，使用协议缓冲区进行编码。命令 “add_person_go “在数据文件中添加一个新条目。命令`list_people_go`解析数据文件并将数据打印到控制台。

You can find the complete example in the [examples directory](https://github.com/protocolbuffers/protobuf/tree/master/examples) of the GitHub repository.

你可以在 GitHub 仓库的[examples 目录](https://github.com/protocolbuffers/protobuf/tree/master/examples)中找到完整的例子。

## Defining your protocol format

To create your address book application, you’ll need to start with a `.proto` file. The definitions in a `.proto` file are simple: you add a _message_ for each data structure you want to serialize, then specify a name and a type for each field in the message. In our example, the `.proto` file that defines the messages is [`addressbook.proto`](https://github.com/protocolbuffers/protobuf/blob/master/examples/addressbook.proto).

要创建你的地址簿应用程序，你需要从一个`.proto`文件开始。`.proto`文件中的定义很简单：你为你要序列化的每个数据结构添加一个*消息*，然后为消息中的每个字段指定一个名称和一个类型。在我们的例子中，定义消息的`.proto`文件是[`addressbook.proto`](https://github.com/protocolbuffers/protobuf/blob/master/examples/addressbook.proto)。

The `.proto` file starts with a package declaration, which helps to prevent naming conflicts between different projects.

`.proto`文件以包声明开始，这有助于防止不同项目之间的命名冲突。

```
syntax = "proto3";
package tutorial;

import "google/protobuf/timestamp.proto";
```

The `go_package` option defines the import path of the package which will contain all the generated code for this file. The Go package name will be the last path component of the import path. For example, our example will use a package name of “tutorialpb”.

`go_package`选项定义了包的导入路径，它将包含该文件的所有生成代码。Go 包名将是导入路径的最后一个路径组件。例如，我们的例子将使用 “tutorialpb “的包名。

```
option go_package = "github.com/protocolbuffers/protobuf/examples/go/tutorialpb";
```

Next, you have your message definitions. A message is just an aggregate containing a set of typed fields. Many standard simple data types are available as field types, including `bool`, `int32`, `float`, `double`, and `string`. You can also add further structure to your messages by using other message types as field types.

接下来，你有你的消息定义。一个消息只是一个包含一组类型字段的集合。许多标准的简单数据类型都可以作为字段类型，包括 `bool`, `int32`, `float`, `double`, 和 `string`。你也可以通过使用其他消息类型作为字段类型来为你的消息添加进一步的结构。

```
message Person {
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;

  google.protobuf.Timestamp last_updated = 5;
}

// Our address book file is just one of these.
message AddressBook {
  repeated Person people = 1;
}
```

In the above example, the `Person` message contains `PhoneNumber` messages, while the `AddressBook` message contains `Person` messages. You can even define message types nested inside other messages – as you can see, the `PhoneNumber` type is defined inside `Person`. You can also define `enum` types if you want one of your fields to have one of a predefined list of values – here you want to specify that a phone number can be one of `MOBILE`, `HOME`, or `WORK`.

在上面的例子中，`Person` 消息包含 `PhoneNumber` 消息，而 `AddressBook` 消息包含 `Person` 消息。你甚至可以定义嵌套在其他消息中的消息类型–如你所见，`PhoneNumber`类型被定义在`Person`中。你也可以定义`enum`类型，如果你想让你的一个字段有一个预定义的值列表–这里你想指定一个电话号码可以 `MOBILE` 、 `HOME` 或 `WORK` 中的一个。

The ” = 1”, “ = 2” markers on each element identify the unique “tag” that field uses in the binary encoding. Tag numbers 1-15 require one less byte to encode than higher numbers, so as an optimization you can decide to use those tags for the commonly used or repeated elements, leaving tags 16 and higher for less-commonly used optional elements. Each element in a repeated field requires re-encoding the tag number, so repeated fields are particularly good candidates for this optimization.

每个元素上的”=1”、”=2 “标记标识了该字段在二进制编码中使用的唯一 “标签”。标签号 1-15 比更高的数字需要少一个字节的编码，所以作为优化，你可以决定将这些标签用于常用的或重复的元素，而将标签 16 及以上的留给不常用的可选元素。重复字段中的每个元素都需要重新编码标签号，所以重复字段特别适合做这个优化。

If a field value isn’t set, a [default value](https://developers.google.com/protocol-buffers/docs/proto3#default) is used: zero for numeric types, the empty string for strings, false for bools. For embedded messages, the default value is always the “default instance” or “prototype” of the message, which has none of its fields set. Calling the accessor to get the value of a field which has not been explicitly set always returns that field’s default value.

如果没有设置字段值，则使用[默认值](https://developers.google.com/protocol-buffers/docs/proto3#default)：数字类型为 0，字符串为空字符串，ools 为 false。对于嵌入式消息，默认值总是消息的 “默认实例 “或 “原型”，它没有设置任何字段。调用访问器来获取一个没有被显式设置的字段的值，总是返回该字段的默认值。

If a field is `repeated`, the field may be repeated any number of times (including zero). The order of the repeated values will be preserved in the protocol buffer. Think of repeated fields as dynamically sized arrays.

如果一个字段是 “repeated”，那么这个字段可以被重复任何次数（包括零）。重复值的顺序将被保留在协议缓冲区中。把重复字段看作是动态大小的数组。

You’ll find a complete guide to writing `.proto` files – including all the possible field types – in the [Protocol Buffer Language Guide](https://developers.google.com/protocol-buffers/docs/proto3). Don’t go looking for facilities similar to class inheritance, though – protocol buffers don’t do that.

你可以在[协议缓冲区语言指南](https://developers.google.com/protocol-buffers/docs/proto3)中找到编写`.proto`文件的完整指南–包括所有可能的字段类型。不要去寻找类似于类继承的功能，尽管–协议缓冲区不做这些。

## Compiling your protocol buffers

Now that you have a `.proto`, the next thing you need to do is generate the classes you’ll need to read and write `AddressBook` (and hence `Person` and `PhoneNumber`) messages. To do this, you need to run the protocol buffer compiler `protoc` on your `.proto`:

现在你已经有了一个`.proto`，接下来你需要做的是生成你需要的类，以便读写`AddressBook`（以及因此 `Person` 和 `PhoneNumber` ）消息。要做到这一点，你需要在你的`.proto`上运行协议缓冲编译器`protoc`。

1. If you haven’t installed the compiler, [download the package](https://developers.google.com/protocol-buffers/docs/downloads) and follow the instructions in the README.
2. 如果你还没有安装编译器，[下载软件包](https://developers.google.com/protocol-buffers/docs/downloads)，并按照 README 中的说明进行操作。
3. Run the following command to install the Go protocol buffers plugin:
4. 运行以下命令安装 Go 协议缓冲区插件。

```
   go install google.golang.org/protobuf/cmd/protoc-gen-go
```

The compiler plugin

```
   protoc-gen-go
```

will be installed in

```
   $GOPATH/bin
   protoc
```

1. Now run the compiler, specifying the source directory (where your application’s source code lives – the current directory is used if you don’t provide a value), the destination directory (where you want the generated code to go; often the same as
2. 现在运行编译器，指定源目录（你的应用程序的源代码所在的地方–如果你没有提供一个值，则使用当前目录），目标目录（你想让生成的代码去的地方；通常与你的应用程序的源代码所在的地方相同。

```
   protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```

This generates `github.com/protocolbuffers/protobuf/examples/go/tutorialpb/addressbook.pb.go` in your specified destination directory.

这将在你指定的目标目录下生成`github.com/protocolbuffers/protobuf/examples/go/tutorialpb/addressbook.pb.go`。

## The Protocol Buffer API

Generating `addressbook.pb.go` gives you the following useful types:

生成`addressbook.pb.go`后，你可以得到以下有用的类型。

- An `AddressBook` structure with a `People` field.
- 一个带有 `People` 字段的 `AddressBook` 结构。
- A `Person` structure with fields for `Name`, `Id`, `Email` and `Phones`.
- 一个 `Person` 结构，有 `Name`, `Id`, `Email` 和 `Phones` 字段。
- A `Person_PhoneNumber` structure, with fields for `Number` and `Type`.
- The type `Person_PhoneType` and a value defined for each value in the `Person.PhoneType` enum.
- 类型`Person_PhoneType`和为`Person.PhoneType`枚举中的每个值定义的值。

You can read more about the details of exactly what’s generated in the [Go Generated Code guide](https://developers.google.com/protocol-buffers/docs/reference/go-generated), but for the most part you can treat these as perfectly ordinary Go types.

你可以在[Go Generated Code guide](https://developers.google.com/protocol-buffers/docs/reference/go-generated)中阅读更多关于具体生成的细节，但在大多数情况下，你可以把这些类型当作完全普通的 Go 类型。

Here’s an example from the [`list_people` command’s unit tests](https://github.com/protocolbuffers/protobuf/blob/master/examples/list_people_test.go) of how you might create an instance of Person:

下面是[`list_people` command’s unit tests](https://github.com/protocolbuffers/protobuf/blob/master/examples/list_people_test.go)中的一个例子，说明你如何创建 Person 的实例。

```
p := pb.Person{
        Id:    1234,
        Name:  "John Doe",
        Email: "jdoe@example.com",
        Phones: []*pb.Person_PhoneNumber{
                {Number: "555-4321", Type: pb.Person_HOME},
        },
}
```

## Writing a Message

The whole purpose of using protocol buffers is to serialize your data so that it can be parsed elsewhere. In Go, you use the `proto` library’s [Marshal](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#Marshal) function to serialize your protocol buffer data. A pointer to a protocol buffer message’s `struct` implements the `proto.Message` interface. Calling `proto.Marshal` returns the protocol buffer, encoded in its wire format. For example, we use this function in the [`add_person` command](https://github.com/protocolbuffers/protobuf/blob/master/examples/add_person.go):

使用协议缓冲区的整个目的是将你的数据序列化，以便它可以在其他地方被解析。在 Go 中，你使用`proto`库的[Marshal](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#Marshal)函数来序列化你的协议缓冲区数据。一个指向协议缓冲区消息的`struct`指针实现了`proto.Message`接口。调用`proto.Marshal`返回协议缓冲区，以其线格式编码。例如，我们在[`add_person` command](https://github.com/protocolbuffers/protobuf/blob/master/examples/add_person.go)中使用这个函数。

```
book := &pb.AddressBook{}
// ...

// Write the new address book back to disk.
out, err := proto.Marshal(book)
if err != nil {
        log.Fatalln("Failed to encode address book:", err)
}
if err := ioutil.WriteFile(fname, out, 0644); err != nil {
        log.Fatalln("Failed to write address book:", err)
}
```

## Reading a Message

To parse an encoded message, you use the `proto` library’s [Unmarshal](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#Unmarshal) function. Calling this parses the data in `buf` as a protocol buffer and places the result in `pb`. So to parse the file in the [`list_people` command](https://github.com/protocolbuffers/protobuf/blob/master/examples/list_people.go), we use:

要解析一个编码的消息，你可以使用`proto`库的[Unmarshal](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#Unmarshal)函数。调用这个函数会将`buf`中的数据解析为一个协议缓冲区，并将结果放在`pb`中。因此，为了解析[`list_people`命令](https://github.com/protocolbuffers/protobuf/blob/master/examples/list_people.go)中的文件，我们使用。

```
// Read the existing address book.
in, err := ioutil.ReadFile(fname)
if err != nil {
        log.Fatalln("Error reading file:", err)
}
book := &pb.AddressBook{}
if err := proto.Unmarshal(in, book); err != nil {
        log.Fatalln("Failed to parse address book:", err)
}
```

## Extending a Protocol Buffer

Sooner or later after you release the code that uses your protocol buffer, you will undoubtedly want to “improve” the protocol buffer’s definition. If you want your new buffers to be backwards-compatible, and your old buffers to be forward-compatible – and you almost certainly do want this – then there are some rules you need to follow. In the new version of the protocol buffer:

在你发布使用你的协议缓冲区的代码之后，你迟早会毫无疑问地想要 “改进 “协议缓冲区的定义。如果你想让你的新缓冲区向后兼容，而你的旧缓冲区向前兼容–你几乎肯定是想这样做的–那么你需要遵循一些规则。在新版本的协议缓冲区中。

- you _must not_ change the tag numbers of any existing fields.
- 你不能改变任何现有字段的标签号。
- you _may_ delete fields.
- 你 **可以** 删除字段。
- you _may_ add new fields but you must use fresh tag numbers (i.e. tag numbers that were never used in this protocol buffer, not even by deleted fields).
- 你 **可以** 添加新的字段，但必须使用新的标签号（即在这个协议缓冲区中从未使用过的标签号，甚至被删除的字段也不例外）。

(There are [some exceptions](https://developers.google.com/protocol-buffers/docs/proto3#updating) to these rules, but they are rarely used.)

(这些规则有[一些例外](https://developers.google.com/protocol-buffers/docs/proto3#updating)，但很少使用。)

If you follow these rules, old code will happily read new messages and simply ignore any new fields. To the old code, singular fields that were deleted will simply have their default value, and deleted repeated fields will be empty. New code will also transparently read old messages.

如果你遵循这些规则，老代码会很高兴地读取新消息，并简单地忽略任何新字段。对旧代码来说，被删除的单数字段将简单地具有其默认值，而被删除的重复字段将是空的。新代码也会透明地读取旧消息。

However, keep in mind that new fields will not be present in old messages, so you will need to do something reasonable with the default value. A type-specific [default value](https://developers.google.com/protocol-buffers/docs/proto3#default) is used: for strings, the default value is the empty string. For booleans, the default value is false. For numeric types, the default value is zero.

但是，请记住，新字段不会出现在旧消息中，所以你需要对默认值做一些合理的处理。使用特定类型的[默认值](https://developers.google.com/protocol-buffers/docs/proto3#default)：对于字符串，默认值是空字符串。对于 booleans，默认值是 false。对于数值类型，默认值为零。

### Reserved Values

If you [update](https://developers.google.com/protocol-buffers/docs/proto3#updating) an enum type by entirely removing an enum entry, or commenting it out, future users can reuse the numeric value when making their own updates to the type. This can cause severe issues if they later load old versions of the same `.proto`, including data corruption, privacy bugs, and so on. One way to make sure this doesn’t happen is to specify that the numeric values (and/or names, which can also cause issues for JSON serialization) of your deleted entries are `reserved`. The protocol buffer compiler will complain if any future users try to use these identifiers. You can specify that your reserved numeric value range goes up to the maximum possible value using the `max` keyword.

如果你[update](https://developers.google.com/protocol-buffers/docs/proto3#updating)通过完全删除一个枚举条目来删除一个枚举类型，或者将其注释出来，那么未来的用户在对该类型进行更新时可以重新使用数值。如果他们以后加载同一`.proto`的旧版本，这可能会导致严重的问题，包括数据损坏、隐私错误等。确保这种情况不会发生的一种方法是指定你删除的条目的数值（和/或名称，这也会导致 JSON 序列化的问题）是 “保留 “的。如果未来的用户试图使用这些标识符，协议缓冲区编译器会抱怨。您可以使用`max`关键字指定您保留的数值范围到最大可能的值。

# Go Protocol 代码生成

# Go Generated Code

This page describes exactly what Go code the protocol buffer compiler generates for any given protocol definition. Any differences between proto2 and proto3 generated code are highlighted - note that these differences are in the generated code as described in this document, not the base API, which are the same in both versions. You should read the [proto2 language guide](https://developers.google.com/protocol-buffers/docs/proto) and/or the [proto3 language guide](https://developers.google.com/protocol-buffers/docs/proto3) before reading this document.

本页确切地描述了协议缓冲编译器为任何给定协议定义生成的 Go 代码。proto2 和 proto3 生成的代码之间的任何差异都会被高亮显示–请注意，这些差异是在本文档中描述的生成的代码中，而不是在基本 API 中，这两个版本中的 API 是相同的。在阅读本文档之前，你应该阅读[proto2 语言指南](https://developers.google.com/protocol-buffers/docs/proto)和/或[proto3 语言指南](https://developers.google.com/protocol-buffers/docs/proto3)。

## 编译器调用(Compiler Invocation)

The protocol buffer compiler requires a plugin to generate Go code. Install it with:

协议缓冲编译器需要一个插件来生成 Go 代码。安装它的方法是

```
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

This will install a `protoc-gen-go` binary in `$GOBIN`. Set the `$GOBIN` environment variable to change the installation location. It must be in your `$PATH` for the protocol buffer compiler to find it. The protocol buffer compiler produces Go output when invoked with the `--go_out` flag. The parameter to the `--go_out` flag is the directory where you want the compiler to write your Go output. The compiler creates a single source file for each `.proto` file input. The name of the output file is created by replacing the `.proto` extension with `.pb.go`. The `.proto` file should contain a `go_package` option specifying the full import path of the Go package that contains the generated code.

这将在`$GOBIN`中安装`protoc-gen-go`二进制文件。设置`$GOBIN`环境变量来改变安装位置。它必须在你的`$PATH`中，协议缓冲区编译器才能找到它。协议缓冲区编译器在调用`--go_out`标志时，会产生 Go 输出。`--go_out`标志的参数是你想让编译器写 Go 输出的目录。编译器为每个输入的`.proto`文件创建一个源文件。输出文件的名称是通过将`.proto`扩展名替换为`.pb.go`来创建的。`.proto`文件应该包含一个`go_package`选项，指定包含生成代码的 Go 包的完整导入路径。

```
option go_package = "example.com/foo/bar";
```

The subdirectory of the output directory the output file is placed in depends on the `go_package` option and the compiler flags:

输出文件被放置在输出目录的子目录中，取决于`go_package`选项和编译器的标志。

- By default, the output file is placed in a directory named after the Go package’s import path. For example, a file `protos/foo.proto` with the above `go_package` option results in a file named `example.com/foo/bar/foo.pb.go`.
- 默认情况下，输出文件被放置在一个以 Go 包的导入路径命名的目录中。例如，使用上述`go_package`选项的文件`protos/foo.proto`的结果是一个名为`example.com/foo/bar/foo.pb.go`的文件。
- If the `--go_opt=module=$PREFIX` flag is given to `protoc`, the specified directory prefix is removed from the output filename. For example, a file `protos/foo.proto` with the above `go_package` option and the flag `--go_opt=module=example.com/foo` results in a file named `bar/foo.pb.go`.
- 如果`--go_opt=module=$PREFIX`标志给了`protoc`，那么指定的目录前缀会从输出文件名中删除。例如，如果一个文件`protos/foo.proto`带有上述`go_package`选项和`--go_opt=module=example.com/foo`标志，则会产生一个名为`bar/foo.pb.go`的文件。
- If the `--go_opt=paths=source_relative` flag is given to `protoc`, the output file is placed in the same relative directory as the input file. For example, the file `protos/foo.proto` results in a file named `protos/foo.pb.go`.
- 如果给`protoc`加上`--go_opt=paths=source_relative`标志，则输出文件会被放置在与输入文件相同的相对目录中。例如，文件`protos/foo.proto`的结果是一个名为`protos/foo.pb.go`的文件。

When you run the proto compiler like this:

当你这样运行 proto 编译器时。

```
protoc --proto_path=src --go_out=build/gen --go_opt=paths=source_relative src/foo.proto src/bar/baz.proto
```

the compiler will read the files `src/foo.proto` and `src/bar/baz.proto`. It produces two output files: `build/gen/foo.pb.go` and `build/gen/bar/baz.pb.go`.

编译器将读取文件`src/foo.proto`和`src/bar/baz.proto`。它产生两个输出文件。`build/gen/foo.pb.go`和`build/gen/bar/baz.pb.go`。

The compiler automatically creates the directory `build/gen/bar` if necessary, but it will _not_ create `build` or `build/gen`; they must already exist.

如果有必要，编译器会自动创建`build/gen/bar`目录，但它不会创建`build`或`build/gen`；它们必须已经存在。

## 包管理（Packages）

Source `.proto` files should contain a `go_package` option specifying the full Go import path for the package containing the file. If there is no `go_package` option, the compiler will try to guess at one. A future release of the compiler will make the `go_package` option a requirement.

源文件`.proto`应包含一个`go_package`选项，指定包含该文件的包的完整 Go 导入路径。如果没有`go_package`选项，编译器将尝试猜测一个。未来的编译器版本将把`go_package`选项作为一项要求。

The Go package name of generated code will be the last path component of the `go_package` option. For example:

生成的代码的 Go 包名将是 `go_package` 选项的最后一个路径成分。例如

```
// The Go package name is "timestamppb".
option go_package = "google.golang.org/protobuf/types/known/timestamppb";
```

The import path is used to determine which import statements must be generated when one `.proto` file imports another `.proto` file. For example, if `a.proto` imports `b.proto`, the generated `a.pb.go` file needs to import the Go package which contains the generated `b.pb.go` file (unless both files are in the same package).

导入路径用于确定当一个`.proto`文件导入另一个`.proto`文件时必须生成哪些导入语句。例如，如果`a.proto`导入`b.proto`，则生成的`a.pb.go`文件需要导入包含生成的`b.pb.go`文件的 Go 包（除非两个文件都在同一个包中）。

The import path is also used to construct output filenames. See the “Compiler Invocation” section above for details.

导入路径也被用来构造输出文件名。详见上面 “编译器调用 “一节。

The `go_package` option may also include an explicit package name separated from the import path by a semicolon. For example: “example.com/foo;package_name”. This usage is discouraged, since it is almost always clearer for the package name to correspond to the import path (the default). As an alternative to the `go_package` option, the Go import path for a `.proto` file may be specified on the command line with the `--go_opt=M=$FILENAME=$IMPORT_PATH` flag to `protoc`.

`go_package`选项也可以包含一个明确的包名，用分号与导入路径分开。例如：”example.com/foo;“。”example.com/foo;package_name”。这种用法是不推荐的，因为包名与导入路径（默认）相对应会更清楚。作为 “go_package “选项的替代方案，可以在命令行中用”–go_opt=M=$FILENAME=$IMPORT_PATH “标记来指定`.protoc`文件的导入路径。

## 消息（Messages）

Given a simple message declaration:

一个简单的消息声明：

```
message Foo {}
```

the protocol buffer compiler generates a struct called `Foo`. A `*Foo` implements the [`proto.Message`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#Message) interface.

协议缓冲区编译器生成一个名为`Foo`的结构。`*Foo`实现了[`proto.Message`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#Message)接口

The [`proto` package](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc) provides functions which operate on messages, including conversion to and from binary format.

[`proto` package](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc)提供了对消息进行操作的函数，包括二进制格式的转换和转换。

The `proto.Message` interface defines a `ProtoReflect` method. This method returns a [`protoreflect.Message`](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect?tab=doc#Message) which provides a reflection-based view of the message.

`proto.Message` 接口定义了一个 `ProtoReflect` 方法。该方法返回一个[`protoreflect.Message`](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect?tab=doc#Message)，提供基于反射的消息视图。

The `optimize_for` option does not affect the output of the Go code generator.

`optimize_for`选项不影响 Go 代码生成器的输出。

### 嵌套/内嵌类型(Nested Types)

A message can be declared inside another message. For example:

一个消息可以在另一个消息里面声明。例如：

```
message Foo {
  message Bar {
  }
}
```

In this case, the compiler generates two structs: `Foo` and `Foo_Bar`.

在这种情况下，编译器会生成两个结构。 `Foo` 和 `Foo_Bar`。

## 字段(Fields)

The protocol buffer compiler generates a struct field for each field defined within a message. The exact nature of this field depends on its type and whether it is a singular, repeated, map, or oneof field.

协议缓冲区编译器为消息中定义的每个字段生成一个结构字段。这个字段的确切性质取决于它的类型，以及它是单字段、重复字段、映射字段还是 `oneof` 字段

Note that the generated Go field names always use camel-case naming, even if the field name in the `.proto` file uses lower-case with underscores ([as it should](https://developers.google.com/protocol-buffers/docs/style)). The case-conversion works as follows:

请注意，生成的 Go 字段名总是使用驼峰大写命名，即使`.proto`文件中的字段名使用带下划线的小写([应该如此](https://developers.google.com/protocol-buffers/docs/style))。大小写转换的工作原理如下。

1. The first letter is capitalized for export. If the first character is an underscore, it is removed and a capital X is prepended.
2. 输出时第一个字母大写。如果第一个字符是下划线，则去掉下划线，前面加一个大写的 X
3. If an interior underscore is followed by a lower-case letter, the underscore is removed, and the following letter is capitalized.
4. 如果一个内部下划线后面是一个小写字母，则去掉下划线，并将后面的字母大写。

Thus, the proto field `foo_bar_baz` becomes `FooBarBaz` in Go, and `_my_field_name_2` becomes `XMyFieldName_2`.

因此，原字段`foo_bar_baz`在 Go 中变成了`FooBarBaz`，`_my_field_name_2`变成了`XMyFieldName_2`。

### Singular Scalar Fields (proto2)

For either of these field definitions:

```
optional int32 foo = 1;
required int32 foo = 1;
```

the compiler generates a struct with an `*int32` field named `Foo` and an accessor method `GetFoo()` which returns the `int32` value in `Foo` or the default value if the field is unset. If the default is not explicitly set, the [zero value](https://golang.org/ref/spec#The_zero_value) of that type is used instead (`0` for numbers, the empty string for strings).

编译器生成一个结构，其中有一个名为 “Foo “的 “\*int32 “字段和一个访问器方法 “GetFoo()“，该方法返回 “Foo “中的 “int32 “值，如果该字段未设置，则返回默认值。如果没有明确设置默认值，则使用该类型的[零值](https://golang.org/ref/spec#The_zero_value)代替(数字为`0`，字符串为空字符串)。

For other scalar field types (including `bool`, `bytes`, and `string`), `*int32` is replaced with the corresponding Go type according to the [scalar value types table](https://developers.google.com/protocol-buffers/docs/proto#scalar).

对于其他标量字段类型(包括`bool`、`bytes`和`string`)，`*int32`将根据[标量值类型表](https://developers.google.com/protocol-buffers/docs/proto#scalar)用相应的 Go 类型代替。

### Singular Scalar Fields (proto3)

For this field definition:

```
int32 foo = 1;
```

The compiler will generate a struct with an `int32` field named `Foo` and an accessor method `GetFoo()` which returns the `int32` value in `Foo` or the [zero value](https://golang.org/ref/spec#The_zero_value) of that type if the field is unset (`0` for numbers, the empty string for strings).

编译器将生成一个结构，其中有一个名为`Foo`的`int32`字段和一个访问器方法`GetFoo()`，该方法将返回`Foo`中的`int32`值，如果字段未设置，则返回该类型的[零值](https://golang.org/ref/spec#The_zero_value)（数字为`0`，字符串为空字符串）。

For other scalar field types (including `bool`, `bytes`, and `string`), `int32` is replaced with the corresponding Go type according to the [scalar value types table](https://developers.google.com/protocol-buffers/docs/proto3#scalar). Unset values in the proto will be represented as the [zero value](https://golang.org/ref/spec#The_zero_value) of that type (`0` for numbers, the empty string for strings).

对于其他标量字段类型(包括`bool`、`bytes`和`string`)，`int32`将根据[标量值类型表](https://developers.google.com/protocol-buffers/docs/proto3#scalar)用相应的 Go 类型代替。原子中未设置的值将用该类型的[零值](https://golang.org/ref/spec#The_zero_value)来表示（数字为`0`，字符串为空字符串）。

### Singular Message Fields

Given the message type:

```
message Bar {}
```

For a message with a `Bar` field:

```
// proto2
message Baz {
  optional Bar foo = 1;
  // The generated code is the same result if required instead of optional.
  // 如果是必填项而不是可选项，生成的代码结果是一样的。
}

// proto3
message Baz {
  Bar foo = 1;
}
```

The compiler will generate a Go struct

```
type Baz struct {
        Foo *Bar
}
```

Message fields can be set to `nil`, which means that the field is unset, effectively clearing the field. This is not equivalent to setting the value to an “empty” instance of the message struct.

消息字段可以设置为 “nil”，这意味着该字段被取消设置，有效地清除了该字段。这不等于将值设置为消息结构的 “空 “实例。

The compiler also generates a `func (m *Baz) GetFoo() *Bar` helper function. This function returns a `nil` `*Bar` if `m` is nil or `foo` is unset. This makes it possible to chain get calls without intermediate `nil` checks.

编译器还生成了一个`func (m *Baz) GetFoo() *Bar`帮助函数。如果`m`为 nil 或`foo`未设置，该函数返回一个` nil``*Bar `。这使得它可以在没有中间`nil`检查的情况下进行链式获取调用。

### 重复字段（Repeated Fields）

Each repeated field generates a slice of `T` field in the struct in Go, where `T` is the field’s element type. For this message with a repeated field:

每一个重复的字段都会在 Go 结构中生成一个`T`字段的片断，其中`T`是字段的元素类型。对于这个带有重复字段的消息。

```
message Baz {
  repeated Bar foo = 1;
}
```

the compiler generates the Go struct:

```
type Baz struct {
        Foo  []*Bar
}
```

Likewise, for the field definition `repeated bytes foo = 1;` the compiler will generate a Go struct with a `[][]byte` field named `Foo`. For a repeated [enumeration](https://developers.google.com/protocol-buffers/docs/reference/go-generated#enum) `repeated MyEnum bar = 2;`, the compiler generates a struct with a `[]MyEnum` field called `Bar`.

同样，对于字段定义`repeated bytes foo = 1;`，编译器将生成一个带有`[][]byte`字段的 Go 结构，名为`Foo`。对于重复[枚举](https://developers.google.com/protocol-buffers/docs/reference/go-generated#enum)`repeated MyEnum bar = 2;`，编译器会生成一个带有`[]MyEnum`字段的结构，名为`Bar`。

The following example shows how to set the field:

```
baz := &Baz{
  Foo: []*Bar{
    {}, // First element.
    {}, // Second element.
  },
}
```

To access the field, you can do the following:

```
foo := baz.GetFoo() // foo type is []*Bar.
b1 := foo[0] // b1 type is *Bar, the first element in foo.
```

### Map Fields

Each map field generates a field in the struct of type `map[TKey]TValue` where `TKey` is the field’s key type and `TValue` is the field’s value type. For this message with a map field:

每一个 map 字段都会在`map[TKey]TValue`类型的结构中生成一个字段，其中`TKey`是字段的键类型，`TValue`是字段的值类型。对于这个带有 map 字段的消息。

```
message Bar {}

message Baz {
  map<string, Bar> foo = 1;
}
```

the compiler generates the Go struct:

```
type Baz struct {
        Foo map[string]*Bar
}
```

### Oneof Fields

For a oneof field, the protobuf compiler generates a single field with an interface type `isMessageName_MyField`. It also generates a struct for each of the [singular fields](https://developers.google.com/protocol-buffers/docs/reference/go-generated#singular-scalar) within the oneof. These all implement this `isMessageName_MyField` interface.

对于一个 oneof 字段，protobuf 编译器会生成一个接口类型为 “isMessageNme_MyField “的单字段。它还为 oneof 中的每个[单字段](https://developers.google.com/protocol-buffers/docs/reference/go-generated#singular-scalar)生成一个结构。这些都实现了这个`isMessageName_MyField`接口。

For this message with a oneof field:

```
package account;
message Profile {
  oneof avatar {
    string image_url = 1;
    bytes image_data = 2;
  }
}
```

the compiler generates the structs:

```
type Profile struct {
        // Types that are valid to be assigned to Avatar:
        //      *Profile_ImageUrl
        //      *Profile_ImageData
        Avatar isProfile_Avatar `protobuf_oneof:"avatar"`
}

type Profile_ImageUrl struct {
        ImageUrl string
}
type Profile_ImageData struct {
        ImageData []byte
}
```

Both `*Profile_ImageUrl` and `*Profile_ImageData` implement `isProfile_Avatar` by providing an empty `isProfile_Avatar()` method.

`*Profile_ImageUrl`和`*Profile_ImageData`都通过提供一个空的`isProfile_Avatar()`方法来实现`isProfile_Avatar`。

The following example shows how to set the field:

下面的例子展示了如何设置该字段。

```
p1 := &account.Profile{
  Avatar: &account.Profile_ImageUrl{"http://example.com/image.png"},
}

// imageData is []byte
imageData := getImageData()
p2 := &account.Profile{
  Avatar: &account.Profile_ImageData{imageData},
}
```

To access the field, you can use a type switch on the value to handle the different message types.

要访问该字段，可以使用值的类型开关来处理不同的消息类型。

```
switch x := m.Avatar.(type) {
case *account.Profile_ImageUrl:
        // Load profile image based on URL
        // using x.ImageUrl
case *account.Profile_ImageData:
        // Load profile image based on bytes
        // using x.ImageData
case nil:
        // The field is not set.
default:
        return fmt.Errorf("Profile.Avatar has unexpected type %T", x)
}
```

The compiler also generates get methods `func (m *Profile) GetImageUrl() string` and `func (m *Profile) GetImageData() []byte`. Each get function returns the value for that field or the zero value if it is not set.

编译器还生成了 get 方法 `func (m *Profile) GetImageUrl() string` 和 `func (m *Profile) GetImageData() []byte` 。每个 get 函数都返回该字段的值，如果没有设置，则返回零值。

## 枚举（Enumerations）

Given an enumeration like:

```
message SearchRequest {
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 1;
  ...
}
```

the protocol buffer compiler generates a type and a series of constants with that type.

协议缓冲区编译器会生成一个类型和一系列的常量。

For enums within a message (like the one above), the type name begins with the message name:

对于消息中的 enums（如上图），类型名称以消息名称开头。

```
type SearchRequest_Corpus int32
```

For a package-level enum:

对于包级枚举：

```
enum Foo {
  DEFAULT_BAR = 0;
  BAR_BELLS = 1;
  BAR_B_CUE = 2;
}
```

the Go type name is unmodified from the proto enum name:

```
type Foo int32
```

This type has a `String()` method that returns the name of a given value.

该类型有一个`String()`方法，返回给定值的名称。

The `Enum()` method initializes freshly allocated memory with a given value and returns the corresponding pointer:

`Enum()`方法用一个给定的值初始化新分配的内存，并返回相应的指针。

```
func (Foo) Enum() *Foo
```

The protocol buffer compiler generates a constant for each value in the enum. For enums within a message, the constants begin with the enclosing message’s name:

协议缓冲区编译器为枚举中的每个值生成一个常量。对于消息中的枚举，常量以包围消息的名称开始。

```
const (
        SearchRequest_UNIVERSAL SearchRequest_Corpus = 0
        SearchRequest_WEB       SearchRequest_Corpus = 1
        SearchRequest_IMAGES    SearchRequest_Corpus = 2
        SearchRequest_LOCAL     SearchRequest_Corpus = 3
        SearchRequest_NEWS      SearchRequest_Corpus = 4
        SearchRequest_PRODUCTS  SearchRequest_Corpus = 5
        SearchRequest_VIDEO     SearchRequest_Corpus = 6
)
```

For a package-level enum, the constants begin with the enum name instead:

对于一个包级的枚举，常量以枚举名开头。

```
const (
        Foo_DEFAULT_BAR Foo = 0
        Foo_BAR_BELLS   Foo = 1
        Foo_BAR_B_CUE   Foo = 2
)
```

The protobuf compiler also generates a map from integer values to the string names and a map from the names to the values:

protobuf 编译器还生成了一个从整数值到字符串名称的映射，以及一个从名称到值的映射。

```
var Foo_name = map[int32]string{
        0: "DEFAULT_BAR",
        1: "BAR_BELLS",
        2: "BAR_B_CUE",
}
var Foo_value = map[string]int32{
        "DEFAULT_BAR": 0,
        "BAR_BELLS":   1,
        "BAR_B_CUE":   2,
}
```

Note that the `.proto` language allows multiple enum symbols to have the same numeric value. Symbols with the same numeric value are synonyms. These are represented in Go in exactly the same way, with multiple names corresponding to the same numeric value. The reverse mapping contains a single entry for the numeric value to the name which appears first in the .proto file.

请注意，`.proto`语言允许多个枚举符号具有相同的数值。具有相同数值的符号是同义词。在 Go 中，这些符号以完全相同的方式表示，同一个数值对应多个名称。反向映射包含一个数值到 `.proto` 文件中最先出现的名称的单一条目。

## Extensions (proto2)

Given an extension definition:

```
extend Foo {
  optional int32 bar = 123;
}
```

The protocol buffer compiler will generate an [`protoreflect.ExtensionType`](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect?tab=doc#ExtensionType) value named `E_Bar`. This value may be used with the [`proto.GetExtension`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#GetExtension), [`proto.SetExtension`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#SetExtension), [`proto.HasExtension`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#HasExtension), and [`proto.ClearExtension`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#ClearExtension) functions to access an extension in a message. The `GetExtension` function and `SetExtension` functions respectively accept and return an `interface{}` value containing the extension value type.

协议缓冲区编译器将生成一个名为 `E_Bar` 的[`protoreflect.ExtensionType`](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect?tab=doc#ExtensionType)值。这个值可以与[`proto.GetExtension`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#GetExtension)、[`proto.SetExtension`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#SetExtension)、[`proto.HasExtension`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#HasExtension)和[`proto.ClearExtension`](https://pkg.go.dev/google.golang.org/protobuf/proto?tab=doc#ClearExtension)函数一起使用，以访问消息中的扩展。`GetExtension`函数和`SetExtension`函数分别接受并返回一个包含扩展值类型的`interface{}`值。

For singular scalar extension fields, the extension value type is the corresponding Go type from the [scalar value types table](https://developers.google.com/protocol-buffers/docs/proto3#scalar).

对于单数标量扩展字段，扩展值类型是[标量值类型表](https://developers.google.com/protocol-buffers/docs/proto3#scalar)中对应的 Go 类型。

For singular embedded message extension fields, the extension value type is `*M`, where `M` is the field message type.

对于单数内嵌消息扩展字段，扩展值类型为`*M`，其中`M`为字段消息类型。

For repeated extension fields, the extension value type is a slice of the singular type.

对于重复的扩展字段，扩展值类型是单数类型的一个片断。

For example, given the following definition:

例如，给定以下定义。

```
extend Foo {
  optional int32 singular_int32 = 1;
  repeated bytes repeated_string = 2;
  optional Bar repeated_message = 3;
}
```

Extension values may be accessed as:

扩展值的访问方式为：

```
m := &somepb.Foo{}
proto.SetExtension(m, extpb.E_SingularInt32, int32(1))
proto.SetExtension(m, extpb.E_RepeatedString, []string{"a", "b", "c"})
proto.SetExtension(m, extpb.E_RepeatedMessage, &extpb.Bar{})

v1 := proto.GetExtension(m, extpb.E_SingularInt32).(int32)
v2 := proto.GetExtension(m, extpb.E_RepeatedString).([][]byte)
v3 := proto.GetExtension(m, extpb.E_RepeatedMessage).(*extpb.Bar)
```

Extensions can be declared nested inside of another type. For example, a common pattern is to do something like this:

扩展可以在另一个类型中声明嵌套。例如，一个常见的模式是这样做的。

```
message Baz {
  extend Foo {
    optional Baz foo_ext = 124;
  }
}
```

In this case, the `ExtensionType` value is named `E_Baz_Foo`.

在这种情况下，`ExtensionType`值被命名为`E_Baz_Foo`。

## 服务 (services)

The Go code generator does not produce output for services by default. If you enable the [gRPC](https://www.grpc.io/) plugin (see the [gRPC Go Quickstart guide](https://github.com/grpc/grpc-go/tree/master/examples)) then code will be generated to support gRPC.

Go 代码生成器默认不生成服务的输出。如果您启用[gRPC](https://www.grpc.io/)插件(参见[gRPC Go 快速入门指南](https://github.com/grpc/grpc-go/tree/master/examples))，那么将生成支持 gRPC 的代码。
