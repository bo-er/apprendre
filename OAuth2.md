## 目标

- 了解 OAuth 是什么
- 工作过程
- 使用 Keycloak 来保护前端应用

## 什么是 OAuth2.0

OAuth2.0 允许第三方应用获得有限制的权限来访问 HTTP 服务，要么是通过用户授权来实现这一点，要么第三方应用代表自己实现(纯后端应用)。
OAuth 不是 API 或者某种服务，它是用于授权的开放标准，任何人都可以实现它。OAuth 有 2.0 跟 1.0 两个版本，2.0 不向后兼容，现在广泛使用的是 2.0 版本。后面提到的 OAuth 指 OAuth2.0

之所以有 OAuth 是因为传统的基于 HTTP BASIC 的账号密码登录具有很多不便利以及风险点，为了设计一个更好的 web 系统，实现单点登录，统一的身份认证成为了一种规范。统一的身份认证引入了授权系统，由于提供 API 的后端系统信任授权系统，因此只要客户端拥有来自授权系统所提供的安全令牌，就能访问后端系统提供的 API。这种无状态的不基于 session 的安全策略，也让后端系统更方便的实现了可扩展性。

OAuth 的框架由授权中心来实现，Keycloak 就是实现了该框架的一种授权中心。

## OAuth 的主要组成部分

### Scopes

Scopes 包括 Scopes to grant 跟 Scopes to deny,也就是应用获取你的权限的时候可以访问、不可以访问哪些内容。

### 组成

OAuth 框架有下面这些较为重要的名词:

- Authentication : 验证用户的身份
- Authorization : 根据用户的身份来决定用户可以访问哪些资源
- 资源拥有者(Resource Owner): 用户
- 资源服务器(Resource Server): 一般是提供 API 接口的后端应用
- 客户(Client): 想要访问用户数据的应用
- 授权服务器(Authorization Server): OAuth 的核心

### OAuth Tokens

- Access Token : 客户使用 Access Token 来访问 API 资源,通常存活期很短,往往是几分钟几小时,而不是几天或者几周。
- Refresh Token : 它的生存期更长，几天几周或者几月，它的作用是获取新的 access token。
- JWT : 包括 header(头信息)、payload(载体)以及 signature(签名)三部分。实际上 OAuth 没有规定 token 应该是什么样的格式。通常你会使用`JSON WEB TOKEN`(JWT).JWT 允许你使用签名来存储信息，并且使用密钥解密信息。
  - header
    通常包含了 token 类型的信息以及用于生产签名的算法。例如:
    ```
    {typ:'jwt',alg'HS256'}
    ```
  - payload
    JSON 格式的数据，可以自定义存放的数据。
  - signature
    签名是头信息、数据载体、签名的哈希加密值。

## OAuth Flows(授权类型)

- Implicit Flow

  之所以称其为隐式流是因为所有的通信都在浏览器上发生。
  这种方式只要用户登录成功立即返回访问令牌以及身份令牌给客户端，去掉了服务器返回 authorization_code 给客户端然后客户端使用 authorization_code、client_id 以及 client_secret 验证客户端身份的过程，只适合纯前端应用。由于授权服务器只能通过验证请求地址、验证请求授权权限列表然后将访问令牌放入响应体的 cookie 中（明确 domain 是重定向地址)发送重定向响应，这个重定向地址一定要很具体，而且必须使用 https 协议来传输。

- Authorization Code Flow

  验证的流程是:

  - 客户端获取一个授权码
  - 客户端使用授权码获取 access token(refresh token 是可选的)

  这种授权模式假设资源拥有者跟客户端所在的设备不同。
  ![Authorization Code Flow](https://i2.wp.com/blogs.innovationm.com/wp-content/uploads/2019/07/blog-open1.png?w=1141)

- Client Credential Grant Flow

  相比前两种需要用户输入登录账号以及密码来使用浏览器作为客户端发送请求，Client Credential Grant 没有图形化的登录界面，也就是说这种方式用于服务端的访问授权。发送授权请求的应用会发送带有 client_id、client_secret、grant_type 信息的请求给授权应用，授权应用检查无误后立即返回访问令牌。如果把 client_id 理解为用户名，client_secret 理解为登录密码，也可以说这种方式是属于后端应用的 Implicit Flow。公司目前使用的通用权限管理 CP 系统就属于这类认证方式。

- Resource Owner Password Flow.

  这种方式用户直接在客户端上输入账号与密码，而不是授权服务上，使用 Postman 获取 Keycloak 的 admin-cli token 就是采取这种方式。客户端会将用户账号密码、client_id、grant_type 等信息发送给授权服务以验证身份，验证通过后授权服务直接返回访问令牌。

## Keycloak

### Keycloak 的三种认证方式

上面介绍的标准 Oauth2 有四种认证流，Keycloak 只有三种的原因是将 Authorization Code Flow 跟 Client Credential Flow 合并为了一种。Keycloak 上的认证流分别是：

- confidential

  这种方式支持通过 Authorization code flow 以及 Client credentials grant 来获取 token。在微服务中调用其他微服务的一方应该使用 confidential 模式。采用这种方式连接的客户端需要存储 keycloak 上的应用密钥 secret。

- public

  对于需要浏览器登录、而且无法保证安全存储 keycloak 应用密钥 secret 的服务器端应用可以采取这种方式。由于没有密钥，public 模式必须要严格保证重定向 URL 的准确性以及使用 HTTPS 协议来传输文本。

- bearer-only

  采取这种方式意味着 Keycloak 应用只接受合法的 Bearer token 所发送的请求，如果采用这种方式 keycloak 不支持浏览器登录,因为完全是客户端跟 keylcoak 之间的鉴权所以意味着这种方式只支持纯后端应用。比如后端应用 caller 需要调用后端应用 callee 时应用 callee 可以采取这种方式。(Bearer token 是完全具有访问权限的 token)
