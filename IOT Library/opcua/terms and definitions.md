- AddressSpace

  是 OPC UA 服务器对其客户暴露的信息集合

- Aggregate

  是一个函数用于从原始数据计算导出数据

- Alarm

  是一个跟状态有关的，需要被关注的事件

- Attribute

  所有的属性被 OPC UA 定义，而不是由客户端或者服务器。属性(attributes)是 AddressSpace 中唯一允许有数据值的元素。

- Client

  是给符合规范的 OPC UA 服务器发送消息的软件应用

- Communication Stack
    是软件模块处于应用跟硬件之间的一层，它提供了众多函数用于编码跟解码，生成发送的消息，并且解码、解密、解包收到的消息。

- Discovery

    OPC UA客户端获取关于OPC UA服务器信息的过程，包括了端点地址以及安全信息。

- EventNotifier

    节点的特殊属性，表示一个客户端有可能订阅到了这个节点来接收事件发生的通知。

- Node
    地址空间(AddressSpace)的根本组成

-     