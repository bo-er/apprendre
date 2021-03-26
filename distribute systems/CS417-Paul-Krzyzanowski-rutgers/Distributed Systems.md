# Distributed Systems

## [Paul Krzyzanowski](https://www.cs.rutgers.edu/~pxk/417/notes)

## Networking

Without shared memory, we need a way for collections of systems (computers or other endpoint devices) to communicate. To do so, they use a **communication network**. If every communicating pair of hosts would have a dedicated physical connection between them, then there would be no sharing of the network infrastructure and we have a true **physical circuit**. This is not practical since it limits the ability for arbitrary computers to talk with each other concurrently. It is also incredibly wasteful in terms of resource use: the circuit would be present even if no data is flowing.

What is needed is a way to **share** the network infrastructure among all connected systems. The challenge in doing so is to allow these systems to talk but avoid **collisions**, the case when two nodes on a network transmit at the same time, on the same channel, and on the same connection. Both signals then get damaged and data does is not transmitted, or is garbled. This is the **multiple access problem**: _how do you coordinate multiple transmitters on a network to ensure that each of them can send their messages?_

There are three broad approaches that enable us to do this:

1. **Channel partitioning** divides the communication channel into “slots”. If the network is divided into short, fixed-length time slots, we have **Time Division Multiplexing**, or **TDM**. Each host must communicate only during its assigned time slot. Routers may connect multiple networks together. When two hosts need to communicate, they establish an end-to-end route called a **virtual circuit**. It is called a _virtual_ circuit because the setup of this route configures routers and assigns communication slots. This provides the _illusion_ of a true circuit switched network in that all messages are guaranteed to arrive in order with constant latency and a guaranteed bandwidth. The switched telephone network is an example of virtual circuit switching, providing a maximum delay of 150 milliseconds and digitizing voice to a constant 64 kbps data rate.

   If the network is partitioned into frequency bands, or channels, then we have **Frequency Division Multiplexing**, or **FDM**. This defines a **broadband** network. Cable TV is an example of a broadband network, transmitting many channels simultaneously, each in using a well-defined frequency band.

   The problem with a channel partitioning approach is that it is wasteful. Network bandwidth is allocated even if there is nothing to transmit.

2. **Taking turns** requires that we create some means of granting permission for a system to transmit. A polling protocol uses a master node to poll other nodes in sequence, offering each a chance to transmit their data. A token passing protocol circulates a special message, called a _token_, from machine to machine. When a node has a token, it is allowed to transmit and must then pass the token to its neighbor.

   The problem with the taking turns approach is that a dead master or lost token can bring the network to a halt. Handling failure cases is complex. This method was used by networks such as IBM’s Token Ring Network but is largely dead now.

3. A **random access** protocol does not use scheduled time slots and allows nodes to transmit at arbitrary times in variable size time slots. This technique is known as **packet switching**. Network access is accomplished via **statistical multiplexing**. A data stream is segmented into multiple variable-size packets. Since these packets will be intermixed with others, each packet must be identified and addressed. Packet switched networks generally cannot provide guaranteed bandwidth or constant latency. Ethernet is an example of a packet-switched network.

Packet switching is the dominant means of data communication today. The packets in a packet-switched network are called **datagrams** and are characterized by unreliable delivery with no guarantees on arrival time or arrival order. Each datagram is fully self-contained with no reliance on previous or future datagrams. This form of communication is also known as **connectionless service**. There is no need to set up a communication session and hence no concept of a connection. Neither routers nor endpoints need to maintain any state as they have to do with a virtual circuit; there is no concept of where a system is in its conversation.

## OSI reference model

Data networking is generally implemented as a layered stack of several **protocols** — each responsible for a specific aspect of networking. The **OSI reference model** defines seven layers of network protocols. Some of the more interesting ones are: the network, transport, and presentation layers.

- 1.  Physical

Deals with hardware, connectors, voltage levels, frequencies, etc.

- 2.  Data link

  Sends and receives packets on the physical network. Ethernet packet transmission is an example of this layer. Connectivity at the link layer defines the **local area network** (**LAN**).

- 3.  Network

Relays and routes data to its destination. This is where networking gets interesting because we are no longer confined to a single physical network but can route traffic between networks. IP, the Internet Protocol, is an example of this layer.

- 4.  Transport

Provides a software endpoint for networking. Now we can talk application-to-application instead of machine-to-machine. TCP/IP and UDP/IP are examples of this layer.

- 5.  Session

Manages multiple logical connections over a single communication link. Examples are SSL (Secure Sockets Layer) tunnels, remote procedure call connection management, and HTTP 1.1.

- 6.  Presentation

Converts data between machine representations. Examples are data representation formats such as XML, JSON, XDR (for ONC remote procedure calls), NDR (for Microsoft COM+ remote procedure calls), and ASN.1 (used for encoding cryptographic keys and digital certificates).

- 7.  Application

  This is a catch-all layer that includes every application-specific communication protocol. For example, SMTP (sending email), IMAP (receiving email), FTP (file transfer), HTTP (getting web pages).

## Data link layer

Ethernet and Wi-Fi (the 802.11 family of protocols) are the most widely used link-layer technologies for local area networks. Both Wi-Fi and Ethernet use the same addressing format and were designed to freely interoperate at the link layer.

Ethernet provides packet-based, in-order, unreliable, connectionless communications. It occupies layers one and two of the OSI model. There is no acknowledgement of packet delivery. This means that a packet may be lost or mangled and the sender will not know. Communication is **connectionless**, which means that there is no need to set up a path between the sender and receiver and no need for either the sender or receiver to maintain any information about the state of communications; packets can be sent and received spontaneously. Messages are delivered in the order they were sent. Unlike IP-based wide-area networking, there are no multiple paths that may cause a scrambling of the sequence of messages.

Interfaces communicating at the link layer must use link-layer addressing. A **MAC address** (for example, an Ethernet address) is different from, and unrelated to, an IP address. An Ethernet MAC address is globally unique to a device and there is no expected grouping of such addresses within a local area network. IP addresses on a LAN, on the other hand, will share a common network prefix.

## Network layer: IP Networking

The **Internet Protocol** (IP) is a network layer protocol that handles the interconnection of multiple local and wide-area networks and the routing logic between the source and destination. It is a logical network whose data is transported by physical networks (such as Ethernet, for example). The IP layer provides unreliable, connectionless datagram delivery of packets between nodes (e.g., computers).

The key principles that drove the design of the Internet are:

1. Support the **interconnection of networks**. The Internet is a _logical_ network that spans multiple physical networks, each of which may have different characteristics. IP demands nothing of these underlying networks except an ability to try to deliver packets.
2. IP assumes **unreliable** communication. That does _not_ mean that most packets will get lost! It means that delivery is not guaranteed. If reliable delivery is needed, software on the receiver will have to detect lost data and ask the sender to retransmit it. Think of mail delivery: most mail gets to its destination but once in a while, a letter gets lost or takes a really long time to arrive.
3. **Routers** connect networks together. A router is essentially a dedicated computer with multiple network links. It receives packets from one network and decides which outgoing link to send the packet.
4. **No central control** of the network. The precursor of the Internet was the ARPAnet, built to connect companies and universities working on Department of Defense projects. As such, it was important that there wouldn’t be a single point of failure – a key element that could be taken out of service to cause the entire network to stop functioning.

Since IP is a logical network, any computer that needs to send out IP packets must do so via the physical network, using the data link layer. Often, this is Ethernet, which uses a 48-bit Ethernet address that is completely unrelated to a 32-bit IP address (or a 128-bit IPv6 address). To send an IP packet out, the system needs to identify the **link layer** destination address (MAC, or Media Access Control address) on the local area network that corresponds to the desired IP destination (it may be the address of a router if the packet is going to a remote network). The **Address Resolution Protocol**, or **ARP**, accomplishes this. It works by broadcasting a request containing an IP address (the message asks, _do you know the corresponding MAC address for this IP address?_) and then waiting for a response from the computer with the corresponding IP address. To avoid doing this for every outgoing packet, ARP maintains a cache of most recently used addresses.

## Transport layer: TCP and UDP

IP is responsible for transporting packets between computers. The transport layer enables applications to communicate with each other by providing logical communication channels so that related messages can be abstracted as a single stream at an application.

There are two transport-layer protocols on top of IP: TCP and UDP.

TCP (**Transmission Control Protocol**) provides **reliable byte stream** (**connection-oriented**) service. This layer of software ensures that packets arrive at the application in order and lost or corrupt packets are retransmitted. The transport layer keeps track of the destination so that the application can have the illusion of a connected data stream.

UDP (**User Datagram Protocol**) provides **datagram** (**connectionless**) service. While UDP drops packets with corrupted data, it does not ensure in-order delivery or reliable delivery.

**Port numbers** in both TCP and UDP are used to allow the operating system to direct the data to the appropriate application (or, more precisely, to the communication endpoint, or **socket**, that is associated with the communication stream).

TCP tries to give a datagram _some_ of the characteristics of a virtual circuit network. The TCP layer will send packet sequence numbers along with the data, buffer received data in memory so they can be presented to the application in order, acknowledge received packets, and request a retransmission of missing or corrupt packets. The software will also keep track of source and destination addresses (this is _state_ that is maintained at the source and destination systems). We now have the _illusion_ of having a network-level virtual circuit with its preset connection and reliable in-order message delivery. What we _do not_ get is constant latency or guaranteed bandwidth. TCP also implements **flow control** to ensure that the sender does not send more data than the receiver can receive. To implement this, the receiver simply sends the amount of free buffer space it has when it sends responses. Finally, TCP tries to be a good network citizen and implements **congestion control**. If the sender gets notification of a certain level of packet loss, it assumes that some router’s queue must be getting congested. It then lowers its transmission rate to relieve the congestion.

The design of the Internet employs the **end-to-end principle**. This is a design philosophy that states that application-specific functions should, whenever possible, reside in the end nodes of a network and not in intermediary nodes, such as routers. Only if the functions cannot be implemented “completely and correctly,” should any logic migrate to the network elements. An example of this philosophy in action is TCP. TCP’s reliable, in-order delivery and flow control are all is implemented via software on the sender and receiver: routers are blissfully unaware of any of this.

A related principle is **fate sharing**, which is also a driving philosophy of Internet design. Fate sharing states that it is acceptable to lose the state information associated with an entity if, at the same time, the entity itself is lost. For example, it is acceptable to lose a TCP connection if a client or server dies. The argument is that the connection has no value in that case and will be reestablished when the computer recovers. However, it is _not_ acceptable to lose a TCP connection if a router in the network dies. As long as alternate paths for packet delivery are available, the connection should remain alive.

## Acknowledgements

To achieve reliable delivery on an unreliable network, we rely on detecting lost or corrupted packets and requesting retransmissions.

The simplest possible mechanism is to send a packet and wait for the receiver to acknowledge it … then send the next one and wait for that to get acknowledged. This, unfortunately, is horribly inefficient since only a single packet is on the network at any time. It is more efficient to use **pipelining** and send multiple packets before receiving any acknowledgements. Acknowledgements can arrive asynchronously and the sender needs to be prepared to retransmit any lost packets.

It would be a waste of network resources for the TCP layer to send back a packet containing nothing an acknowledgement number. While this is inevitable in some cases, if the receiver happens to have data to transmit back to the sender, the acknowledgement number is simply set in the TCP header of the transmitted segment, completely avoiding the need to send a separate acknowledgement. Using an outgoing data segment to transmit an acknowledgement is known as a **piggybacked acknowledgement**.

TCP also uses **cumulative acknowledgements**. Instead of sending an acknowledgement per received message, TCP can acknowledge multiple messages at once.

## Sockets

**Sockets** are a general-purpose interface to the network provided to applications by the operating system. By this, we mean that they were not designed to support one specific network but rather provide a generic mechanism for inter-process communication. They are the only way that an application can interact with the network.

They are created with the **socket** system call and assigned an address and port number with the **bind** system call. For connection-oriented protocols (e.g., TCP), a socket on the server can be set to listen for connections with the **listen** system call. The **accept** call blocks until a connection is received, at which point the server receives a socket dedicated to that connection. A client establishes a connection with the **connect** system call. The “connection” is not a a configuration of routers as with virtual circuits; it is just state that is maintained by the transport layer of the network stack in the operating system at both endpoints. After this, sending and receiving data is compatible with file operations: the same **read/write** system calls can be used. When communication is complete, the socket can be closed with the **shutdown** or **close** system calls.

With sockets that use a connectionless protocol (e.g., UDP), there is no need to establish a connection or to close one. Hence, there is no need for the _connect_, _listen_, or _shutdown_ system calls. The **sendto** and **recvfrom** system calls were created to send and receive datagrams since the _read_ and _write_ system calls do not enable you to specify the remote address. _sendto_ allows you to send a datagram and specify its destination. _recvfrom_ allows you to receive a datagram and identify who sent it.

## Protocol encapsulation

We saw that if we want to send an IP packet out on an Ethernet network (IP is a logical network, so there is no physical IP network), we needed to send out an Ethernet packet. The entire IP packet becomes the payload (data) of an Ethernet packet. Similarly, TCP and UDP have their own headers, distinct from IP headers (they need a port number, for example). A TCP or UDP packet is likewise treated as data by the IP layer. This wrapping process is known as **protocol encapsulation**.
