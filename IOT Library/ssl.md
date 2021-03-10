### Why SSL exists

- 加密

  加密的作用是隐藏从一台计算机发送到另外一台计算机的数据。

- 身份认证

  确保计算机所连接的对象是可以信任的。

### Encrption-how?

1. 计算机之间商量好如何加密

   - client 给 server 发送 hello 消息,消息包括了如下内容:
     - key exchange method
       - RSA
       - Diffie-Hellman
       - DSA
     - Cipher
       - RC4
       - Triple DES
       - AES
     - Hash: 用于生成 `message authentication code`,跟上面提到的 message 一起发送，用于确保消息的完整性。
       - HMAC-MD5
       - HMAC-SHA
     - Version
     - Random number: 用于生成 session key
   - server 给 client 发送 hello 消息,告诉客户端它选择哪一个 Key,哪一个 Cipher,哪一个 Hash 方式。

2. 服务端给客户端发送一个证书(certificate)
3. 开始加密数据
   - 客户端跟服务端同时开始计算一个 `master secret code`
   - 客户端要求服务端加密
