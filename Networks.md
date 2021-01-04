| 序号 | 网络名称    | 协议              | 协议数据单位 | 定位        |
| ---- | ----------- | ----------------- | ------------ | ----------- |
| 5    | Application | HTTP,SMTP,etc...  | Messages     | n/a         |
| 4    | Transport   | TCP/UDP           | Segment      | Port #'s    |
| 3    | Network     | IP                | Datagram     | IP address  |
| 2    | Data Link   | Ethernet, Wi-Fi   | Frames       | MAC address |
| 1    | Physical    | 10 Base T, 802.11 | Bits         | n/a         |

1. Physical layer

   用来表示物理连接着的计算机

2. Data Link

   负责定义一种通用的方式来解析物理层的信号，这样网络设备之间才能交流。数据链路层使用Ethernet协议。这个协议的作用是通过网络或者链路给某一个结点node发送数据

3. Network Layer

   允许不同的网络间进行交流，需要通过一种叫做路由器的设备。Internetwork就是连接在一起的网络的集合，最著名的是Internet。网络层所使用的最多的协议是IP协议。(Internet Protocol) ,它是互联网以及小型网络的基础。基于网络的程序一般分类为服务端跟客户端，一台电脑上可能会有多个应用，多个应用间的数据传输不会相互干扰，这归功于Transport Layer。

4. Transport Layer

   传输层需要找到接收数据的程序。每当提到TCP（Transmition Control Protocol),你可能会想到TCP/IP.之所以它们被放到一起是因为TCP主要使用了IP协议。另外一种常用的协议也是基于IP协议，称为UDP(User Data Protocol)

   TCP跟UDP的区别是TCP有确保数据可靠的传输的机制而UDP没有。

5. Application

   应用层

小结：

5层协议可以这样理解：

网络层类似送快递的货车，数据链路层相当于道路，网络层相当于高德导航，传输层相当于快递员把货物搬上楼，应用层就是货物本身。



------

**Cabels**

就是网线，有copper铜线跟fiber光纤两种

**Crosstalk**

当一根线的电流脉冲被另外一个线检测到，高层的协议可以识别到数据的异常并且要求重新传输。

光纤的每根管子大概跟人的头发丝一样细。

------

**Hub**

它是一个物理层的设备，允许同时连接多台计算机，连接到一台HUB的计算机，发送的数据会传播给连接到HUB的全部计算机。

**Collision domain**

它是一个网络时间片段，在这个片段内只有一台设备可以获取数据。如果多个设备同时发送数据，通过网线发送的电流脉冲将彼此干扰。这也是为什么现在HUB都不用了。

------

Switch

跟Hub类似，区别在于HUB是一个物理层的设备，而Switch是数据链路层的设备。

也就是说Switch知道不同的数据应该被送往哪个系统。

Hubs跟Switches主要用于单个网络的连接，这种网络一般称为LAN(Local Area Netwrok)

------

Router

A device that knows how to forward data between independent networks

既然路由器可以在不同的网络间转发数据，那么说明它是网络层的设备。

Border Gateway Protocal (BGP)

路由器们通过BGP协议来彼此交流，BGP使得路由器可以从最优的路线传输数据

------

Twisted pair cable (双绞线)

双绞线（twisted pair，TP）是一种[综合布线](https://baike.baidu.com/item/综合布线/4282)工程中最常用的传输介质，是由两根具有绝缘保护层的铜[导线](https://baike.baidu.com/item/导线/1413914)组成的。把两根绝缘的铜导线按一定密度互相绞在一起，每一根导线在传输中[辐射](https://baike.baidu.com/item/辐射/5676)出来的电波会被另一根线上发出的电波抵消，有效降低信号干扰的程度。

双绞线一般由两根22～26号绝缘铜导线相互缠绕而成，“双绞线”的名字也是由此而来。实际使用时，双绞线是由多对双绞线一起包在一个绝缘电缆套管里的。如果把一对或多对双绞线放在一个绝缘套管中便成了[双绞线电缆](https://baike.baidu.com/item/双绞线电缆/3393109) [1] ，但**日常生活中一般把“双绞线电缆”直接称为“双绞线”**。

与其他传输介质相比，双绞线在传输距离，[信道宽度](https://baike.baidu.com/item/信道宽度/1208670)和数据传输速度等方面均受到一定限制，但价格较为低廉。

普通的网线就是双绞线。通常的网线有两个灯，黄色的是Link LED,绿色的是Activity LED

------

MAC address

是一个全局唯一的标识符依附在一个独立的网络接口上，它由6组（一组两个）16进制数表示，因此是48（6X2X4)位

Ethernet使用Mac地址来明确发送数据的设备跟接收数据的设备的地址。

------

Unicast(单播)

In computer networking, **unicast** is a one-to-one transmission from one point in the network to another  point; that is, one sender and one receiver, each identified by a  network address. **Unicast** is in contrast to multicast and broadcast which are one-to-many transmissions.

Mac地址的第一个16进制数的最后一位如果是0则表示ethernet frame只针对目标地址即单播，也就是说数据会被发送到冲突域中的全部机器，但是只有mac地址是目标地址的设备可以接收到这些数据。

Mac地址的第一个16进制数的最后一位如果是1则表示ethernet处于**multicast** 多播模式

ethernet的最后一种模式是广播模式，广播模式通过一种特殊的地址实现： FF:FF:FF:FF:FF:FF

------

Data Package

数据包是一个概括的(all-encompassing)名词用来表示任何单一的通过网络链路传输的二进制数据

Ethernet Frame

