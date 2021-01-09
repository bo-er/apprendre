## 项目创建

### 什么是 spring mvc

![spring mvc](../pictures/springmvc.png)

首先用户的请求会到达 Servlet(Controller 层)，然后根据请求调用相应的 Java Bean，并把所有的显示结果交给 JSP 去完成，这样的模式我们就称为 MVC 模式。

传统的模型层被拆分为了业务层(Service)和数据访问层（DAO,Data Access Object）。 在 Service 下可以通过 Spring 的声明式事务操作数据访问层。而 Controller 是业务层开始执行的地方。

- **M 代表 模型（Model）**
  模型是什么呢？ 模型就是数据，就是 dao,bean。上面的 Servcie 跟 Repository 就是夹带事务的模型层。
- **V 代表 视图（View）**
  视图是什么呢？ 就是网页, JSP，用来展示模型中的数据
- **C 代表 控制器（controller)**
  控制器是什么？ 控制器的作用就是把不同的数据(Model)，显示在不同的视图(View)上，Servlet 扮演的就是这样的角色。Spring 的 Controller 配合 Servlet 起到了控制器的作用。

## 相关依赖包介绍，Web 容器的选择和替换

## RESTFUL 服务介绍

## Swagger-ui 配置和使用介绍

## 服务器端口修改方法介绍

## 独立打包和启动时修改启动参数方法介绍

