# Go 实用密码学
[阅读原文](https://leanpub.com/gocrypto/read#leanpub-auto-rsa)
## 第一章：简介

这是一本有关密码学的书：如何安全地进行通信。密码学旨在解决以下几个目标：**机密性**， **完整性**和**真实性**。它还有助于解决安全通信中出现的其他一些问题，但是请记住，它不是安全问题的完整解决方案，这一点很重要。在本书中，我们将研究如何构建安全的系统。还将指出密码学无法解决的一些问题。本书将引导您尝试了解如何使用加密技术保护服务安全，并使用 Go 编程语言进行说明。

如上所述，密码安全性的基础是机密性，完整性和真实性这三个目标。机密性是只有预定的一方才能阅读给定消息的要求；完整性是不能篡改消息内容的要求；真实性是必须信任消息的**来源**（或来源）的要求。信任将在我们的安全系统中扮演重要角色，但是没有唯一的解决方案。在构建安全系统方面将面临许多挑战。密码算法对数据应用了一些转换，以实现这些目标，并且应用了各种算法来实现不同的目标。

为了讨论密码学，需要基本词汇表。以下术语具有特定含义：

- **plaintext**是原始信息。
- **ciphertext**传统上是原始信息转变为提供机密性的密文信息。
- **cipher**是用于加密或解密的消息的密码变换。
- **message authentication code**（或**MAC**）是一个提供真实性和完整性的数据。MAC 算法用于生成和验证此代码。
- **encrypt**是应用了保密转换的消息，但经常被用来描述满足所有密码学的三个目标的一个转换。
- **decrypt**逆转了保密转换，并经常表示其他两个属性已经核实。
- **hash**或**hash algorithm**将一些任意大小的输入转化为固定大小的输出，也被称为`消化`或`散列`。密码哈希是满足某些特定安全目标的算法。
- **peer**或**party**描述参与通信过程的实体。可能是一个人或另一台机器。

一个安全的通信系统将防御**被动**和 **主动**攻击。被动攻击是指无意向其发送消息的一方正在监听通信。主动攻击是指一些对抗方篡改消息并可以注入，更改或回应消息的攻击。

加密技术应用于解决特定问题，例如

- 窃听：与面对面的对话一样，攻击者可以侦听进出的流量，有可能窃取来回传递的秘密。机密性的安全目标将在一定程度上减轻这种攻击；尽管加密技术会掩盖消息的内容，但它本身并不能掩盖两方正在通信的事实。攻击者也许还可以根据消息的大小来确定信息。
- 篡改：可以在途中修改进出应用程序的流量；系统需要确保接收到的消息未被篡改。完整性目标用于确保消息未被篡改。
- 欺骗：攻击者可以通过伪造消息的某些详细信息来伪装成合法用户。攻击者可以使用欺骗手段窃取敏感信息，伪造用户请求或接管合法会话。身份验证通过验证用户和消息的身份来帮助防御这种攻击。

在本书中，我们将研究加密技术的上下文以及构建安全系统，对称和非对称（或公共密钥）加密技术，如何交换密钥，存储机密，信任和常见用例时涉及的一些工程问题。

本书在https://github.com/kisom/gocrypto/上的Github上具有相关的示例代码存储库。该代码包含本书中的代码以及将要提到的一些补充材料。通常，本书中的代码没有注释。该代码将在文本中进行解释。但是，示例代码已注释。

## 第 2 章：工程问题和平台安全性

如果密码学是众所周知的保险库门，那么在决定保险库门的细节之前，先评估建筑物的其余部分和其基础是有意义的。经常重复说，安全系统仅作为其最弱的组成部分才是安全的，这意味着在使用密码学之前，系统的其他部分必须达到标准。在本章中，我们将研究有关安全系统的工程问题，特别是对于基于 Unix 的系统。

### 基本安全

安全应提供**身份验证**，**授权**和 **审核**。认证是指系统验证与系统交互的各方的身份；授权验证是否应允许他们进行此交互；审核会创建安全事件日志，可以对其进行验证和检查以确保系统提供安全性。最终目标是**确保**系统安全。

#### 验证

身份验证问一个问题，“我在跟谁说话？”；它尝试验证某方的**身份**。密码是身份验证的一种方法。它们不是强大的身份验证机制，因为知道密码的任何人（无论是因为他们选择了密码，被给予密码还是被猜中了）都将通过身份验证。多因素身份验证试图基于以下三个因素为身份保证提供更坚实的基础：

1. 您知道的一些信息（例如密码）
2. 您拥有的东西（例如某种身份验证令牌）
3. 您是某物（例如生物识别）

普遍使用的最常见的多因素身份验证配置是采用前两个因素的两个因素。可能要求用户从手机上的应用程序输入密码和基于时间的一次性密码（例如 TOTP）。这里的假设是，用于生成 TOTP 的密钥仅出现在验证者（即邮件提供者）和用户的电话上，并且用户的电话并未受到损害或从中窃取了密钥。双方共享此一次性密码（OTP）的密钥，并且在设置密码后无需交流任何其他信息。

这与两步验证相反。一个示例是发送到手机的 SMS 代码。用户和服务器必须通过与浏览器不同的渠道进行通信，以共享此代码。这仍然提供了拦截代码的渠道。

#### 授权

授权提出问题，“您应该这样做吗？” 授权依赖于某种访问控制机制。这可能像访问控制列表一样简单，在访问控制列表中，系统具有应有权访问资源或应被允许执行某些操作的各方的列表。Unix 安全模型使用一组访问控制列表，以供所有者，资源所属的组以及环境进行读取，写入和执行。它采用“自由访问控制”：用户可以显式更改那些访问控制列表的值，并酌情授予其他用户和组权限。强制访问控制模型（例如 SELinux 或 AppArmor 提供）在安全级别或标签上运行；每个标签都具有一组功能。为用户或流程赋予标签或安全级别，

例如，用户可以创建一个文本文件，并选择使其在 DAC 模型中世界可读：现在，任何进程或用户都可以访问它。在 MAC 模型中，对该文件的访问将受到标签的限制。如果某个进程或用户没有基于其标签的权限，则他们将无法访问它，并且原始用户根本无法以这种方式共享文本文件。标签是由管理员或安全员分配的，用户无法更改。访问控制不再由用户决定，而是强制性的。

访问控制列表的替代方法是基于角色的访问控制安全模型，它是基于角色的安全性。在 Unix 系统上，root 具有对系统的完全控制权。在基于角色的系统中，此控件分为多个角色，每个角色都具有执行该角色的最小权限集。与访问控制列表相比，此模型的粒度也更细，因为它可以为特定操作指定授予或权限。

#### 稽核

如果没有人正在审核系统以确保其正常运行，那么安全工作就白费了。审核日志应该可用，并且访问权限仅限于记录安全事件的审核员。记录的事件和出现的详细信息将根据系统要求而有所不同。审核员还应确信审核日志不会被篡改。

成功进行身份验证的攻击者可能不会留下任何表明系统已受到威胁的迹象。找出折衷方案的唯一方法是通过积极的审核：即审核成功事件的记录。需要考虑这种妥协的风险是否超过维护使用隐私的需要。

#### 政策

应该有一组策略来明确指定系统中的身份验证，授权和审核。对于大型组织而言，这可能相当复杂。对于提供给用户的应用程序，它可能很简单，就像基于密码的应用程序一样，并在系统日志中显示安全故障。本质上，这不是工程问题，但实际上必须将其纳入安全应用程序中。

### 规格

规范是构建安全系统以了解其行为的关键部分。从安全的角度来看，了解其行为一定不是必须的也很重要。安全模型是系统规范的一部分，应注意正确构建它。

测试也是规范的重要组成部分。这是系统按照规范运行的保证。单元测试验证系统每个单元内的代码路径，功能测试验证组件和系统的行为，并进行回归测试，以确保不会重新引入错误。集成测试也可能对验证兼容性很有用。

构建安全系统取决于编写正确的代码。出厂的不正确系统带有安全漏洞，这些漏洞可能会破坏内置的任何加密安全性。系统中存在的代码越多，攻击面就越大：只有实现符合规格的系统的最少代码才应使用。这包括系统使用的所有库：如果可能，请删除所有未使用的功能。这也降低了系统成本：必须编写更少的测试代码，从而减少了时间和财务成本。

#### 安全模型

设计安全系统的关键步骤之一是建立**安全模型**。安全模型描述了有关系统的假设，可以提供安全性的条件，并确定了对系统及其功能的威胁。除非了解其特性和要解决的问题，否则无法构建安全的系统。安全模型应包括对系统攻击面（可能受到攻击的组件），现实威胁向量（攻击来自何处），可以攻击系统的各方及其功能以及需要提供哪些对策的分析。安全。安全模型不应仅涵盖加密组件：还必须考虑系统将在其中运行的环境和平台。考虑哪些问题本质上也很重要，以及哪些问题本质上是社会问题。信任也是此模型中的关键考虑因素。也就是说，了解合法方的角色，预期功能和交互以及受损害的受信方的影响。

根据经验，在设计或构建完系统后，要将安全性附加到系统上是极其困难的。出于同样的原因，在系统设计的初始阶段开始讨论安全性要求很重要，考虑其他技术要求也很重要。例如，如果不考虑系统的负载水平，可能会导致做出糟糕的体系结构决策，从而增加大量技术负担，从而阻碍了系统的稳定，可靠。同样，如果不考虑安全要求，可能会导致类似的架构决策。一个安全的系统必须是可靠的并且必须是正确的；大多数安全漏洞是由利用系统中行为不正确的部分引起的。适当的工程设计是确保系统安全的关键；明确的规范以及正面和负面的测试（测试系统是否正常运行以及是否正常运行）将极大地提高系统实现其安全目标的能力。将安全性视为性能指标很有用。安全系统的性能与其在不受到损害的情况下运行的能力以及从损害中恢复过来的能力有关。还必须考虑安全系统的非安全性能：如果系统太慢或太难使用，将无法使用。未使用的安全系统是不安全的系统，因为它无法提供消息安全性。将安全性视为性能指标很有用。安全系统的性能与其在不受到损害的情况下运行的能力以及从损害中恢复过来的能力有关。还必须考虑安全系统的非安全性能：如果系统太慢或太难使用，将无法使用。未使用的安全系统是不安全的系统，因为它无法提供消息安全性。将安全性视为性能指标很有用。安全系统的性能与其在不受到损害的情况下运行的能力以及从损害中恢复过来的能力有关。还必须考虑安全系统的非安全性能：如果系统太慢或太难使用，将无法使用。未使用的安全系统是不安全的系统，因为它无法提供消息安全性。

安全组件必须服务于系统的其他目标；他们必须在系统规范内做一些有用的事情。

作为安全系统规范的一部分，还必须选择密码算法。在本书中，我们将首选 NaCl 算法用于未开发的领域的设计：它们是由受人尊敬的密码学家设计的，他以编写精心设计的代码而闻名，它们具有易于使用的简单界面，并且并非由 NIST。它们提供高性能和强大的安全性。在其他情况下，需要与现有系统或标准（例如 FIPS）兼容；在这种情况下，应选择兼容的算法。

### 关于错误

攻击者掌握的有关加密操作失败原因的信息越多，他们破解系统的机会就越大。例如，有些攻击通过区分解密失败和填充失败来进行操作。在本书中，我们要么使用布尔值发出错误信号，要么使用通用错误值（例如“加密失败”）发出信号。

我们还将尽早检查假设，并在发现问题后尽快保释。

### 输入卫生

安全系统还必须仔细检查其输入和输出，以确保它们不会降低安全性或为攻击者提供立足之地。在软件工程中众所周知，必须对来自外部世界的输入进行过滤。应该对数据进行完整性检查，并且系统应拒绝处理无效数据。

有两种方法可以执行此操作：将黑名单（默认为允许）和白名单（默认为拒绝）。黑名单是一种被动措施，涉及对已知的不良输入做出响应。黑名单系统将始终对它检测到的新的不良输入（可能是通过攻击）做出反应。白名单决定了一组正确的输入，并且仅允许这些输入。确定正确的输入看起来需要更多的工作，但是它可以为系统提供更高的保证。在测试有关输入的假设时，它也可能很有用：如果有效输入经常被白名单击中，则有关传入数据的假设可能看起来是错误的。

### 记忆

从某种意义上说，考虑到当前的技术，必须将敏感数据载入内存。Go 是一种托管内存语言，这意味着用户几乎无法控制内存，这给确保系统安全性带来了其他挑战。最近的漏洞（如 Heartbleed）表明，内存中的任何内容都可以通过访问该内存的方式泄露给攻击者。以 Heartbleed 为例，它是攻击者，可以通过网络访问内存中包含机密的进程。**流程隔离**这是一种对策：阻止攻击者访问进程的内存空间将有助于减轻对系统的成功攻击。但是，可以通过物理控制台或远程 SSH 会话访问计算机的攻击者现在可以访问该计算机上运行的任何进程的内存空间。这是其他安全机制对安全系统至关重要的地方：它们阻止攻击者到达该内存空间。

但是，易受攻击的不仅仅是在计算机上运行的任何进程的内存空间。现在可以通过文件系统访问交换到磁盘的所有内存。交换到磁盘的密钥现在在两个位置存在。如果该进程在处于睡眠状态的便携式计算机上运行，则该内存通常会写入磁盘。如果外围设备具有直接内存访问（DMA），并且许多设备都具有直接内存访问功能，则该外围设备可以访问计算机中的所有内存，包括每个进程的内存空间。如果程序崩溃并转储核心，则该内存通常会写入核心文件。CPU 缓存也可以存储机密信息，这可能是额外的攻击面，尤其是在共享环境（例如 VPS）上。

有几种方法可以减轻这种情况：使用堆栈来防止秘密进入堆，并在不再需要内存中的敏感数据时将其归零（尽管这并不总是有效的，例如[5]）。在本书中，我们将在有意义的地方进行此操作，但应注意这一点。

也不能保证可以完全安全地擦除磁盘上存储的机密信息（除非使用健康剂量的铝热剂）。如果磁盘上的某个扇区发生故障，则磁盘控制器可能将该块标记为坏块，并尝试将数据复制到另一个扇区，从而使该数据仍保留在硬件上。磁盘控制器可能被颠覆，因为磁盘驱动器包含的驱动器控制器的审核固件质量较差（如果有的话）。

简而言之，考虑到我们当前的技术，内存是难以保护的攻击面。对每个机密询问以下以下问题会有所帮助：

1. 它是否驻留在磁盘上以进行长期存储？如果是这样，谁有权使用它？哪些授权机制可确保只有经过身份验证的方可以访问？
2. 当它加载到内存中时，谁拥有它？它在内存中保留了多长时间？不再使用该怎么办？
3. 如果这些秘密存在于虚拟机上，那么可以访问主机的各方可以信任多少呢？其他租户（即其他虚拟机的用户）可以找到一种访问机密的方法吗？停用机器时会发生什么？

### 随机性

密码系统依赖于足够随机数据的来源。我们希望这些来源的数据与理想的随机数据（在可能值范围内的均匀分布）没有区别。历史上，Unix 平台上可用的选项之间存在很多混淆，但是正确的答案（例如[6]）是使用`/dev/urandom`。幸运的是，`crypto/rand.Reader`在 Go 标准库中，Unix 系统上使用了它。

确保平台具有足够的随机性是另一个问题，主要归结为确保内核的 PRNG 在用于加密目的之前已正确植入种子。对于虚拟机而言，这尤其是一个问题，虚拟机可能会在其他位置复制或从已知或通用种子开始。在这种情况下，在内核的 PRNG 中包括其他熵源可能是有用的，例如，写入内核 PRNG 的硬件 RNG。主机也可以通过磁盘或内存访问 PRNG，以允许主机对其进行观察，这也必须加以考虑。

### 时间

一些协议依赖于在对等体之间同步的时钟。从历史上看，这一直是一个具有挑战性的问题。例如，审核日志通常依靠时钟来确定事件发生的时间。密码系统的主要挑战之一是检查密钥是否已过期。如果时间到了，系统可能会错误地拒绝使用尚未过期的密钥或使用已经过期的密钥。有时，时钟用于唯一值，不应依赖该值。另一个用例是单调递增计数器。时钟回归（例如通过 NTP）使其不那么单调。依赖于基于时间的一次性密码的身份验证也需要准确的时钟。

拥有实时时钟很有用，但并非每个系统都有一个。实时时钟也会根据硬件的物理属性而漂移。网络时间同步在大多数时间都有效，但是它们会受到网络故障的影响。虚拟机可能会受到主机上时钟的影响。

将时钟本身用作单调计数器也会导致问题。可以将已向前漂移的时钟设置回正确的时间（即通过 NTP），这会导致计数器向后退。自启动以来，有一个包含滴答声的 CPU 时钟；可能是用当前时间戳引导的，并且最新的计数器值被永久存储（如果最新的计数器值被更早的值替换或删除，会发生什么情况？）。

它有助于怀疑时钟值。我们会尽力使用计数器而不是合理的时钟。

### 侧通道

旁道是完全基于物理实现的加密系统上的攻击面；虽然算法可能是正确正确的，但实现过程可能会由于物理现象而泄漏信息。攻击者可以观察两次操作之间的时间间隔或用电量的差异，以推断出有关私钥或原始消息的信息。

这些类型的旁通道包括：

- 定时：对系统某些部分执行操作所花费的时间的观察。攻击者使用它甚至可以成功地通过网络攻击系统（[2]）。
- 功耗：通常用于智能卡；攻击者观察到各种操作的用电量如何变化。
- 电源故障：系统电源故障，或接近 CPU 的关闭值。有时，这会导致系统以无法预料的方式发生故障，从而泄露有关密钥或消息的信息。
- EM 泄漏：某些电路会泄漏电磁辐射（例如 RF 波），可以观察到。

这些攻击可能出奇地有效且具有破坏性。设计加密实现时必须考虑这些通道（例如使用中的恒定时间函数`crypto/subtle`）；安全模型应考虑这些攻击的可能性和可能的对策。

### 隐私和匿名

在设计系统时，应确定应提供何种程度的隐私和匿名性。在需要匿名的系统中，也许审核日志不应记录事件成功的时间，而仅记录事件失败的时间。甚至这些失败也可能泄漏信息：如果用户输入错误的密码，会损害他们的身份吗？如果记录了诸如 IP 地址之类的信息，则当它们与其他数据（例如来自用户计算机的活动日志）结合使用时，可用于取消匿名用户。仔细考虑系统在此方面的行为。

### 可信计算

基础平台的一个问题是确保它未被颠覆。恶意软件和 Rootkit 可能会使其他安全措施失效。最好能保证所涉及的各方在具有安全配置的平台上运行。诸如可信计算小组的可信计算倡议之类的努力旨在证明某种程度的平台完整性和系统参与者的真实性，但是解决方案很复杂并且充满了警告。

### 虚拟环境

如今，云计算风靡一时，这有充分的理由：它提供了一种经济高效的方式来部署和管理服务器。但是，计算机安全方面有一句古老的格言：对计算机进行物理访问的攻击者可以破坏计算机上的任何安全性，而云计算则使对其中某些计算机的“物理”访问变得更加容易。硬件是在软件中模拟的，因此获得了对主机访问权限（甚至通过远程 SSH 会话或类似权限）的攻击者也具有同等访问权限。鉴于当前的技术，这使得在云中保护敏感数据（如加密密钥）的任务成为可疑的前景。如果不信任主机，如何信任虚拟机？这不仅意味着对拥有或运营主机的各方的信任：他们的管理软件安全吗？攻击者访问主机有多困难？另一个租户（或主机上其他虚拟机上的用户）是否可以访问他们不应该访问的主机或其他虚拟机？如果取消了虚拟机的使用，是否已充分擦除驱动器，以使其永远不会落入其他租户手中？在虚拟环境中部署的系统的安全模型除了要开发的系统外，还需要考虑主机提供商和基础架构的安全性，包括正在运行的映像的完整性。驱动器是否已充分擦拭，以至于永远不会被其他租户所用？在虚拟环境中部署的系统的安全模型除了要开发的系统外，还需要考虑主机提供商和基础架构的安全性，包括正在运行的映像的完整性。驱动器是否已充分擦拭，以至于永远不会被其他租户所用？在虚拟环境中部署的系统的安全模型除了要开发的系统外，还需要考虑主机提供商和基础架构的安全性，包括正在运行的映像的完整性。

### 公钥基础设施

在部署使用公共密钥密码术的系统时，确定如何信任和分发公共密钥成为一项挑战，这将给系统增加额外的工程复杂性和成本。公钥本身不包含任何信息。需要指定包含其他任何必需的身份和元数据信息的某种格式。为此，有一些标准，例如可怕的 X.509 证书格式（将公钥与有关私钥持有者和证明该公钥的公钥持有者的信息配对）。应该考虑确定要包括哪些身份信息以及如何对其进行验证，以及在需要时还应考虑密钥的生存期以及如何强制密钥过期。需要考虑行政和政策因素；

密钥轮换是 PKI 的挑战之一。它要求确定密钥的 **加密**期限（**有效期**为多长时间）；一个给定的密钥通常只能加密或签名那么多数据，然后才能将其替换（例如，这样就不会重复消息）。在使用 TLS 的情况下，许多组织都在使用寿命较短的证书。这意味着，如果密钥被盗用并且撤销无效，那么损害将是有限的。关键的轮换问题也可能会成为 DoS 攻击：如果轮换被破坏，则在修复之前，它可能使系统无法使用。

密钥撤销是密钥轮换问题的一部分：如何将密钥标记为已泄露或丢失？事实证明，以这种方式标记密钥是一件容易的事，但让其他人知道却不是。TLS 中有几种解决方法：证书吊销列表，其中包含吊销密钥的列表；OCSP（在线证书状态协议），它提供了一种查询权威人士有关密钥是否有效的方法；TACK 和证书透明性尚未得到大规模采用。CRL 和 OCSP 都存在问题：如果针对 CRL 或 OCSP 服务器的密钥折衷与 DDoS 相结合，该怎么办？用户可能看不到密钥已被撤销。如果无法访问 OCSP 服务器，有些人选择拒绝接受证书。在正常的网络中断情况下会发生什么？CRL 通常按设定的时间表发布，用户必须经常请求 CRL 进行更新。他们应该多久检查一次？即使他们每小时检查一次，也会留下一个小时的窗口，在这个窗口中，仍然可以信任泄露的密钥。

由于这些问题以及提供有用的公用密钥基础结构的困难，PKI 在安全和加密社区中往往是个肮脏的词。

### 密码不提供什么

考虑到我们将讨论的加密方法提供了强有力的安全保证，没有一种提供任何消息长度模糊性；取决于该系统，即使在强有力的安全保证的情况下，这也可以使纯文本可预测。如果攻击者正在监视通信通道，则在发送消息时，加密中也不会隐藏任何内容。许多密码系统也不会隐藏谁在通信。很多时候，仅通过观察通信通道（例如跟踪 IP 地址）就可以证明这一点。就其本身而言，加密不会提供强大的匿名性，但它可能会成为此类系统中构建模块的一部分。这种通信通道监视被称为流量分析，要克服该挑战具有挑战性。

此外，尽管我们将提供不可伪造的保证，但加密不会采取任何措施来防止重放攻击。重播攻击类似于欺骗攻击，其中攻击者捕获先前发送的消息并重播它们。一个示例是记录金融交易，然后重播此交易以窃取金钱。消息号是我们如何解决此问题的方法；系统切勿重复发送消息，并且系统应丢弃重复发送的消息。那是系统需要处理的事情，并且不能通过密码学解决。

### 数据寿命

在本书中，当我们发送加密的消息时，我们更喜欢使用短暂的消息密钥来完成通信，这些消息密钥会在通信完成后删除。这意味着消息以后无法使用与之加密的相同密钥解密。虽然这对安全性有好处，但是这意味着弄清楚如何存储消息（如果需要）的负担落在了系统上。某些系统（例如 Pond（https://pond.imperialviolet.org/））会强制执行一周的邮件寿命。这种强制删除消息被认为是社会规范。这些因素将影响到如何存储解密的消息以及存储多长时间的决策。

数据存储和消息流量之间也有过渡：消息流量使用永不存储的临时密钥进行加密，而存储的数据则需要使用长期密钥进行加密。系统体系结构应考虑这些不同类型的加密，并确保适当保护存储的数据。

### 选项，旋钮和转盘

系统选择使用其使用的密码的选项越多，犯密码错误的机会就越大。因此，我们将在此处避免这样做。通常，我们更喜欢 NaCl 加密，它具有一个简单的界面，没有任何选项，并且高效且安全。在设计系统时，它有助于对所使用的密码学做出明智的选择。**加密敏捷性**的属性或能够切换出加密技术的属性对于从可疑故障中恢复可能会很有用。但是，最好还是退后一步，考虑一下为什么会发生故障，并将其合并到将来的修订版中。

### 兼容性

加密实现的质量范围很广。在 Go 中可以很好地实现某些东西这一事实并不意味着将在用于构建系统其他部分的语言中存在良好的实现。这必须纳入系统的设计中：与其他组件集成或用另一种语言构建客户端库有多容易？幸运的是，氯化钠被广泛使用。这是我们更喜欢它的另一个原因。

### 结论

密码学通常被视为构建安全系统的有趣部分。但是，在密码学进入图像之前，还需要完成许多其他工作。这不是万能的安全解决方案；它不是万能的。我们不能只是在不安全的系统上撒些魔术的加密粉尘，然后使其突然变得安全。我们还必须确保了解我们的问题，并确保我们尝试解决的问题实际上是使用密码术解决的正确问题。在虚拟环境中，构建安全系统的挑战更加艰巨。安全模型必须成为规范的一部分，并且要遵循正确的软件工程技术以确保系统的正确性，这一点至关重要。记住，一个缺陷会导致整个系统崩溃。

### 进一步阅读

1. [Anders08] R.安德森。_安全工程：构建可靠的分布式系统指南，第二版_。威利（Wiley），2008 年 4 月。
2. [Brum03] D. Brumley，D。Boneh。“远程定时攻击是可行的。” 在 2003 年第 12 届 USENIX 安全研讨会的论文集中。
3. [Ferg10] N. Ferguson，B。Schneier，T。Kohno。_密码工程_。威利（Wiley），2010 年 3 月。
4. [Graff03] MG Graff，KR van Wyk。_安全编码：原则和实践_。OReily Media，2003 年 6 月。
5. [Perc14] C. Percival。“如何将缓冲区归零。” http://www.daemonology.net/blog/2014-09-04-how-to-zero-a-buffer.html 2014-09-04。
6. [Ptac14] T. Ptacek。“如何安全地生成一个随机数。” http://sockpuppet.org/blog/2014/02/25/safely-generate-random-numbers/ 2014-02-25。
7. [Viega03] J. Viega 和 M. Messier。_C 和 C ++安全编程手册_。OReilly Media，2003 年 7 月。

## 第三章：对称安全

对称密码学是最简单的密码学形式：各方共享相同的密钥。它也往往是最快的加密类型。从根本上讲，对称算法使用的密钥是字节序列，用作操作位的转换算法的输入。与对称密码学相比，与非对称密码学相比，密钥分发更加困难，因为敏感密钥材料的传输需要安全通道。在下一章中，我们将介绍交换密钥的方法。

对称加密有两个组件：提供机密性的算法（是块或流密码），以及提供完整性和真实性的组件（MAC 算法）。大多数密码不会在同一算法中同时提供这两种密码，但确实有两种密码被称为身份验证加密（AE）或带有附加数据的身份验证加密（AEAD）密码。在本章中，我们将考虑四个**密码套件**：NaCl，AES-GCM，带有 HMAC 的 AES-CTR 和带有 HMAC 的 AES-CBC。密码套件是我们将用来提供安全性的一系列算法。本章中的代码可以`chapter3`在示例代码包中找到。

### 不可区分性

安全加密系统的特性之一是它提供了**不可区分性**。这里有两种特殊的不可区分性：IND-CPA，即在选择的明文攻击下的不可区分性，以及 IND-CCA，即在可选项的密文攻击下的不可区分性。

在 IND-CPA 中，攻击者将一对长度相同的消息发送到要加密的服务器。服务器选择消息之一，对其进行加密，然后发回密文。攻击者应该无法确定哪个消息已加密。此财产保持机密。考虑要求消息长度相同是很有用的：在大多数密码中，加密消息的长度与原始消息的长度有关。也就是说，加密消息不会隐藏其长度。

在 IND-CCA 中，攻击者提交自己选择的密文，然后由服务器解密。经过一些观察，攻击者提交了两个质询密文，服务器随机选择了一个以进行解密并发送回攻击者。攻击者应该不能区分该明文对应于哪个密文。在对称安全性中，这种攻击通常被视为填充 Oracle 攻击，其中所使用的加密方案不包含消息身份验证代码（例如不具有 HMAC 的 AES-CBC），并且可以使攻击者恢复用于加密的密钥。IND-CCA 有两种变体。第一个（IND-CCA1）表示攻击者在发送挑战后无法提交新的密文。第二种（IND-CCA2 或自适应 CCA）允许攻击者在挑战后继续提交密文。

安全性既需要机密性组件（例如 AES-CBC）又需要完整性和真实性组件（例如 HMAC-SHA-256）。

另一个不可区分的要求是，我们的密钥材料必须与随机（尤其是均匀分布）不可区分。

### 真实性和完整性

为什么需要完整性和真实性组件？人们尝试构建安全系统时出现的一件事是，仅使用机密性方面：程序员将仅使用具有某些未经身份验证的模式（例如 CBC，OFB 或 CTR 模式）的 AES。AES 密文具有延展性：可以对其进行修改，而接收者通常都不是更明智的选择。在加密的即时消息，电子邮件或网页的上下文中，这种修改似乎很明显。但是，攻击者可以利用无效密文（解密失败时）和无效消息（明文错误）之间的不同响应来提取密钥。也许最著名的此类攻击是 padding oracle 攻击。在其他情况下，无效的纯文本可能被用来利用消息处理程序中的错误。

附加 HMAC 或使用身份验证模式（例如 GCM）要求攻击者证明他们具有用于身份验证消息的密钥。拒绝失败的 MAC 消息可减少出现无效消息的可能性。这也意味着，刚刚发送无效数据的攻击者甚至在浪费处理器时间进行解密之前，都会丢弃其消息。

为了有效地通过 HMAC 进行身份验证，HMAC 密钥应该是与 AES 密钥不同的密钥。在本书中，我们将 HMAC 与 AES 结合使用，并将 HMAC 密钥附加到 AES 密钥上以获得完整的加密或解密密钥。对于具有 HMAC-SHA-256 的 CBC 或 CTR 模式下的 AES-256，这意味着 32 字节的 AES 密钥和 32 字节的 HMAC 密钥，总密钥大小为 64 字节；HMAC-SHA-256 的选择将在后面的部分中阐明。

关于如何应用 MAC，有几种选择。正确的答案是先加密然后再加密 MAC：

1. 加密和 MAC：在这种情况下，我们将对明文应用 MAC，然后发送加密的明文和 MAC。为了验证 MAC，接收者必须解密消息；这仍然允许攻击者提交修改过的密文，并且存在与前面所述相同的问题。这为 IND-CCA 攻击提供了一个表面。
2. MAC-then-encrypt：将 MAC 应用并附加到明文中，并且两者都被加密。请注意，接收方仍然必须解密消息，并且可以通过修改结果密文来修改 MAC，这也是 IND-CCA 攻击的表面。
3. Encrypt-then-MAC：加密消息并附加密文的 MAC。接收方将验证 MAC，如果 MAC 无效，则不会继续解密。这将去除 IND-CCA 表面。

Moxie Marlinspike 的“加密厄运原理” [Moxie11]对此进行了更详尽的讨论。

**始终**使用身份验证模式（如 GCM）或“加密-然后-MAC”。

### 氯化钠

[NaCl](https://nacl.cr.yp.to/)是网络和密码学库，具有对称库（secretbox）和非对称库（box），由 Daniel J. Bernstein 设计。附加的 Go 加密程序包包含 NaCl 的实现。它使用 32 字节密钥和 24 字节随机数。随机数是一次使用的数字：随机数永远不能在使用相同密钥加密的一组消息中重用，否则可能会造成妥协。在某些情况下，随机生成的随机数是合适的。在其他情况下，它将成为有状态系统的一部分；可能是消息计数器或序列号。

密码箱系统使用称为“ XSalsa20”的流密码来提供机密性，并使用称为“ Poly1305”的 MAC。程序包将数据类型`*[32]byte`用于密钥和`*[24]byte`随机数。使用这些数据类型可能有点陌生。以下代码演示了如何生成随机密钥和随机随机数，以及如何与期望使用的函数进行互操作`[]byte`。

```
 1 const (
 2         KeySize   = 32
 3         NonceSize = 24
 4 )
 5
 6 // GenerateKey creates a new random secret key.
 7 func GenerateKey() (*[KeySize]byte, error) {
 8 	    key := new([KeySize]byte)
 9 	    _, err := io.ReadFull(rand.Reader, key[:])
10 	    if err != nil {
11 	    	return nil, err
12 	    }
13
14 	    return key, nil
15 }
16
17 // GenerateNonce creates a new random nonce.
18 func GenerateNonce() (*[NonceSize]byte, error) {
19 	    nonce := new([NonceSize]byte)
20 	    _, err := io.ReadFull(rand.Reader, nonce[:])
21 	    if err != nil {
22 	    	return nil, err
23 	    }
24
25 	    return nonce, nil
26 }
```

NaCl 使用“**盖章**”一词来表示保护消息（以使它现在是机密的，并且可以验证其完整性和真实性），而**打开则**表示恢复消息（验证消息的完整性和真实性并解密消息）。

在此代码中，将使用随机生成的随机数。在密钥交换一章中，将阐明此选择。值得注意的是，所选择的密钥交换方法将允许随机选择的随机数作为获得随机数的安全手段。在其他用例中，情况可能并非如此！收件人将需要某种方式来恢复随机数，因此它将被添加到邮件中。如果使用另一种获取随机数的方法，则可能会有不同的方法来确保收件人具有与用于密封邮件相同的随机数。

```
 1 var (
 2         ErrEncrypt = errors.New("secret: encryption failed")
 3         ErrDecrypt = errors.New("secret: decryption failed")
 4 )
 5
 6 // Encrypt generates a random nonce and encrypts the input using
 7 // NaCl's secretbox package. The nonce is prepended to the ciphertext.
 8 // A sealed message will the same size as the original message plus
 9 // secretbox.Overhead bytes long.
10 func Encrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
11 	    nonce, err := GenerateNonce()
12 	    if err != nil {
13 	    	return nil, ErrEncrypt
14 	    }
15
16 	    out := make([]byte, len(nonce))
17 	    copy(out, nonce[:])
18 	    out = secretbox.Seal(out, message, nonce, key)
19 	    return out, nil
20 }
```

解密期望消息包含一个前置的随机数，我们通过检查消息的长度来验证这一假设。太短而无法成为有效加密消息的消息将立即被丢弃。

```
 1 // Decrypt extracts the nonce from the ciphertext, and attempts to
 2 // decrypt with NaCl's secretbox.
 3 func Decrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
 4 	    if len(message) < (NonceSize + secretbox.Overhead) {
 5 	    	return nil, ErrDecrypt
 6 	    }
 7
 8 	    var nonce [NonceSize]byte
 9 	    copy(nonce[:], message[:NonceSize])
10 	    out, ok := secretbox.Open(nil, message[NonceSize:], &nonce, key)
11 	    if !ok {
12 	    	return nil, ErrDecrypt
13 	    }
14
15 	    return out, nil
16 }
```

请记住，随机随机数并不总是正确的选择。我们将在有关密钥交换的一章中进一步讨论这一点，在该章中，我们将讨论如何实际获取和共享正在使用的密钥。

在后台，NaCl 将加密一条消息，然后对该密文应用 MAC 算法以获得最终消息。此“先加密然后再 MAC”过程是如何正确组合加密密码和 MAC。

### AES-GCM

如果需要或选择了 AES，则 AES-GCM 通常是最佳选择。它将 AES 分组密码与 GCM 分组密码模式配对。它是 AEAD 密码：带有其他数据的经过身份验证的加密。它会加密一些数据，这些数据将与一些未加密的可选附加数据一起进行身份验证。对于 AES-128，密钥长度是 16 个字节；对于 AES-192，密钥长度是 24 个字节；对于 AES-256，密钥长度是 32 个字节。它还将一个随机数作为输入，并且相同的警告适用于此处的随机数选择。另一个警告是，GCM 难以正确实施，因此，审查使用 AES-GCM 的系统中可能使用的软件包的质量非常重要。

您应该选择哪种密钥大小？这取决于应用程序。通常，如果有规范，请使用指定的密钥大小。密码学工程学（[Ferg10]）建议使用 256 位密钥。这就是我们在这里使用的。同样，系统的安全模型应规定这些参数。在本章的 AES 示例中，将密钥大小更改为 16 就足以切换到 AES-128（对于 AES-192，则为 24）。随机数大小在三个版本中均不变。

与大多数分组密码模式不同，GCM 提供身份验证。它还允许对其他一些未加密的数据以及密文进行身份验证。鉴于它是 AEAD 模式（提供完整性和真实性），因此无需为该模式附加 HMAC。

包装中的`AEAD`类型`crypto/cipher`使用与 NaCl 相同的“打开”和“密封”术语。上面的 NaCl 加密的 AES-GCM 类似物为：

```
 1 // Encrypt secures a message using AES-GCM.
 2 func Encrypt(key, message []byte) ([]byte, error) {
 3         c, err := aes.NewCipher(key)
 4         if err != nil {
 5                 return nil, ErrEncrypt
 6         }
 7
 8         gcm, err := cipher.NewGCM(c)
 9         if err != nil {
10                 return nil, ErrEncrypt
11         }
12
13         nonce, err := GenerateNonce()
14         if err != nil {
15                 return nil, ErrEncrypt
16         }
17
18         // Seal will append the output to the first argument; the usage
19         // here appends the ciphertext to the nonce. The final parameter
20         // is any additional data to be authenticated.
21         out := gcm.Seal(nonce, nonce, message, nil)
22         return out, nil
23 }
```

此版本不以密文形式提供任何其他（未加密但经过身份验证的）数据。

也许有一个系统，其中消息以 32 位发送者 ID 开头，这使接收者可以选择适当的解密密钥。以下示例将对此发件人 ID 进行身份验证：

```
 1 // EncryptWithID secures a message and prepends a 4-byte sender ID
 2 // to the message.
 3 func EncryptWithID(key, message []byte, sender uint32) ([]byte, error) {
 4         buf := make([]byte, 4)
 5         binary.BigEndian.PutUint32(buf, sender)
 6
 7         c, err := aes.NewCipher(key)
 8         if err != nil {
 9                 return nil, ErrEncrypt
10         }
11
12         gcm, err := cipher.NewGCM(c)
13         if err != nil {
14                 return nil, ErrEncrypt
15         }
16
17         nonce, err := GenerateNonce()
18         if err != nil {
19                 return nil, ErrEncrypt
20         }
21
22         buf = append(buf, nonce)
23         buf := gcm.Seal(buf, nonce, message, message[:4])
24         return buf, nil
25 }
```

为了解密消息，接收者也将需要提供适当的发送者 ID。和以前一样，我们首先检查关于消息长度的一些基本假设，并考虑到前置的消息 ID 和随机数。

```
 1 func DecryptWithID(message []byte) ([]byte, error) {
 2         if len(message) <= NonceSize+4 {
 3                 return nil, ErrDecrypt
 4         }
 5
 6     // SelectKeyForID is a mock call to a database or key cache.
 7         id := binary.BigEndian.Uint32(message[:4])
 8         key, ok := SelectKeyForID(id)
 9         if !ok {
10                 return nil, ErrDecrypt
11         }
12
13         c, err := aes.NewCipher(key)
14         if err != nil {
15                 return nil, ErrDecrypt
16         }
17
18         gcm, err := cipher.NewGCM(c)
19         if err != nil {
20                 return nil, ErrDecrypt
21         }
22
23         nonce := make([]byte, NonceSize)
24         copy(nonce, message[4:])
25
26         // Decrypt the message, using the sender ID as the additional
27         // data requiring authentication.
28         out, err := gcm.Open(nil, nonce, message[4+NonceSize:], message[:4])
29         if err != nil {
30                 return nil, ErrDecrypt
31         }
32         return out, nil
33 }
```

如果完全更改了邮件头，即使新的发件人 ID 返回相同的密钥，邮件也将无法解密：对其他数据的任何更改都将导致解密失败。

### 带有 HMAC 的 AES-CTR

如果可以选择的话，您应该考虑的最后一个选项是 AES-CTR 和带有 HMAC 的 AES-CBC。在这些密码套件中，首先以适当的模式用 AES 加密数据，然后附加 HMAC。在本书中，我们假定仅在需要作为规范的一部分或出于兼容性时才使用这些密码套件。

点击率也使用随机数；同样，nonce 必须仅使用相同的密钥使用一次。重用随机数可能会造成灾难性的后果，并且会泄漏有关该消息的信息。该系统现在将无法通过不可区分性要求，因此将变得不安全。如果对随机数是否唯一存在任何疑问，应生成随机随机数。如果将其用于与现有系统兼容，则需要考虑该系统如何处理随机数。

如果使用的是 AES-CTR，则可能会遵循某种规范，该规范应指定要使用的 HMAC 构造。FIPS 准则中的一般经验法则是：对于 AES-128，使用 HMAC-SHA-256；对于 AES-256，使用 HMAC-SHA-384。密码学工程学（[Ferg10]）和[Perc09]推荐使用 HMAC-SHA-256。我们将使用 HMAC-SHA-256 和 AES-256。

在这里，我们将通过选择随机随机数，加密数据并计算密文的 MAC 来进行加密。随机数将被添加到消息之前，并添加 MAC。该消息将被就地加密。该密钥应该是附加在 AES 密钥上的 HMAC 密钥。

```
 1 const (
 2         NonceSize = aes.BlockSize
 3         MACSize = 32 // Output size of HMAC-SHA-256
 4         CKeySize = 32 // Cipher key size - AES-256
 5         MKeySize = 32 // HMAC key size - HMAC-SHA-256
 6 )
 7
 8 var KeySize = CKeySize + MKeySize
 9
10 func Encrypt(key, message []byte) ([]byte, error) {
11         if len(key) != KeySize {
12                 return nil, ErrEncrypt
13         }
14
15         nonce, err := util.RandBytes(NonceSize)
16         if err != nil {
17                 return nil, ErrEncrypt
18         }
19
20         ct := make([]byte, len(message))
21
22         // NewCipher only returns an error with an invalid key size,
23         // but the key size was checked at the beginning of the function.
24         c, _ := aes.NewCipher(key[:CKeySize])
25         ctr := cipher.NewCTR(c, nonce)
26         ctr.XORKeyStream(ct, message)
27
28         h := hmac.New(sha256.New, key[CKeySize:])
29         ct = append(nonce, ct...)
30         h.Write(ct)
31         ct = h.Sum(ct)
32         return ct, nil
33 }
```

为了解密，检查消息长度以确保其具有随机数，MAC 和非零的消息大小。然后，检查 MAC。如果有效，则对消息进行解密。

```
 1 func Decrypt(key, message []byte) ([]byte, error) {
 2         if len(key) != KeySize {
 3                 return nil, ErrDecrypt
 4         }
 5
 6         if len(message) <= (NonceSize + MACSize) {
 7                 return nil, ErrDecrypt
 8         }
 9
10         macStart := len(message) - MACSize
11         tag := message[macStart:]
12         out := make([]byte, macStart-NonceSize)
13         message = message[:macStart]
14
15         h := hmac.New(sha256.New, key[CKeySize:])
16         h.Write(message)
17         mac := h.Sum(nil)
18         if !hmac.Equal(mac, tag) {
19                 return nil, ErrDecrypt
20         }
21
22         c, _ := aes.NewCipher(key[:CKeySize])
23         ctr := cipher.NewCTR(c, message[:NonceSize])
24         ctr.XORKeyStream(out, message[NonceSize:])
25         return out, nil
26 }
```

### AES-CBC

先前的模式掩盖了分组密码的基本性质：AES 对数据块进行操作，并且需要一个完整的块来进行加密或解密。先前的模式充当流密码，其中消息长度不必是块大小的倍数。但是，CBC 不能以这种方式起作用，并且需要将消息填充到适当的长度。CBC 也不以相同的方式使用随机数。

在 CBC 模式下，每个密文块都与前一个块进行异或。这就引出了第一个块与 XOR 运算的问题。在 CBC 中，我们使用一种称为初始化向量的虚拟块。它可能是随机生成的，这通常是正确的选择。我们还注意到，在其他加密方案中，可以将消息或序列号用作 IV：此类数字不应直接与 CBC 一起使用。应该使用**单独的**IV 加密密钥对它们进行加密（使用 AES-ECB）。IV 绝不能与相同的消息和密钥一起重复使用。

使用的标准填充方案是 PKCS＃7 填充方案。我们用包含填充字节数的字节填充其余字节：如果必须添加三个填充字节，则将附加`0x03 0x03 0x03` 到纯文本的末尾。

```
1 func pad(in []byte) []byte {
2         padding := 16 - (len(in) % 16)
3         for i := 0; i < padding; i++ {
4                 in = append(in, byte(padding))
5         }
6         return in
7 }
```

解压后，我们将获取最后一个字节，检查它是否有意义（它是否表示填充比消息长？填充是否大于块长度？），然后确保所有填充字节都为展示。接受消息之前，请务必检查您对消息的假设。完成后，我们将删除填充字符并返回纯文本。

```
 1 func unpad(in []byte) []byte {
 2         if len(in) == 0 {
 3                 return nil
 4         }
 5
 6         padding := in[len(in)-1]
 7         if int(padding) > len(in) || padding > aes.BlockSize {
 8                 return nil
 9         } else if padding == 0 {
10                 return nil
11         }
12
13         for i := len(in) - 1; i > len(in)-int(padding)-1; i-- {
14                 if in[i] != padding {
15                         return nil
16                 }
17         }
18         return in[:len(in)-int(padding)]
19 }
```

填充发生在加密之外：我们在加密数据之前填充，在解密之后取消填充。

加密是通过填充消息并生成随机 IV 来完成的。

```
 1 func Encrypt(key, message []byte) ([]byte, error) {
 2         if len(key) != KeySize {
 3                 return nil, ErrEncrypt
 4         }
 5
 6         iv, err := util.RandBytes(NonceSize)
 7         if err != nil {
 8                 return nil, ErrEncrypt
 9         }
10
11         pmessage := pad(message)
12         ct := make([]byte, len(pmessage))
13
14         // NewCipher only returns an error with an invalid key size,
15         // but the key size was checked at the beginning of the function.
16         c, _ := aes.NewCipher(key[:CKeySize])
17         ctr := cipher.NewCBCEncrypter(c, iv)
18         ctr.CryptBlocks(ct, pmessage)
19
20         h := hmac.New(sha256.New, key[CKeySize:])
21         ct = append(iv, ct...)
22         h.Write(ct)
23         ct = h.Sum(ct)
24         return ct, nil
25 }
```

加密与 CTR 模式一样进行，并增加了邮件填充。

在解密中，我们验证了两个假设：

1. 消息长度应为 AES 块大小的倍数（为 16）。HMAC-SHA-256 产生一个 32 字节的 MAC，它也是块大小的倍数。我们可以检查整个邮件的长度，而不是仅检查密文。加密之前未填充不是块大小倍数的消息，因此是无效消息。
2. 该消息必须至少有四个块长：IV 的一个块，消息的一个块和 HMAC 的两个块。如果使用具有较大输出大小的 HMAC 函数，则需要重新考虑该假设。

解密还会在实际解密消息之前检查 HMAC，并验证是否正确填充了明文。

```
 1 func Decrypt(key, message []byte) ([]byte, error) {
 2         if len(key) != KeySize {
 3                 return nil, ErrEncrypt
 4         }
 5
 6         // HMAC-SHA-256 returns a MAC that is also a multiple of the
 7         // block size.
 8         if (len(message) % aes.BlockSize) != 0 {
 9                 return nil, ErrDecrypt
10         }
11
12         // A message must have at least an IV block, a message block,
13         // and two blocks of HMAC.
14         if len(message) < (4 * aes.BlockSize) {
15                 return nil, ErrDecrypt
16         }
17
18         macStart := len(message) - MACSize
19         tag := message[macStart:]
20         out := make([]byte, macStart-NonceSize)
21         message = message[:macStart]
22
23         h := hmac.New(sha256.New, key[CKeySize:])
24         h.Write(message)
25         mac := h.Sum(nil)
26         if !hmac.Equal(mac, tag) {
27                 return nil, ErrDecrypt
28         }
29
30         // NewCipher only returns an error with an invalid key size,
31         // but the key size was checked at the beginning of the function.
32         c, _ := aes.NewCipher(key[:CKeySize])
33         ctr := cipher.NewCBCDecrypter(c, message[:NonceSize])
34         ctr.CryptBlocks(out, message[NonceSize:])
35
36         pt := unpad(out)
37         if pt == nil {
38                 return nil, ErrDecrypt
39         }
40
41         return pt, nil
42 }
```

### 消息与流

在本书中，我们处理**消息**：离散大小的数据块。由于真实性要求，处理数据流更加困难。您如何提供身份验证信息？让我们考虑对流进行加密，以尝试提供与本章相同的安全性属性。

我们不能先加密然后再加密 MAC：从本质上讲，我们通常不知道流的大小。流完成后，我们将无法发送 MAC，这通常由流关闭表示。我们不能动态解密流，因为我们必须查看整个密文才能检查 MAC。尝试保护流会增加问题的复杂性，而没有好的答案。解决方案是将流分成离散的块，并将其视为消息。不幸的是，这意味着我们不能轻松地对`io.Reader`s 和`io.Writer`s 进行加密或解密，并且必须对`[]byte`消息进行操作。丢弃 MAC 根本不是一种选择。

### 结论

在本章中，我们省略了有关如何实际获取密钥的讨论（通常，生成随机密钥没有用）。这是一个足够大的主题，值得在其单独的章节中进行讨论。

一些关键点：

1. 尽可能使用 NaCl。如果需要 AES，并且可以选择，请使用 AES-GCM。使用 AES-CBC 和 AES-CTR 进行兼容性。
2. 总是先加密然后再加密 MAC。永远不要只是加密。
3. 解密消息之前，请始终检查有关消息的假设，包括消息的真实性。
4. 考虑一下您如何获得随机数和静脉输液，以及它是否是合适的方法。

### 进一步阅读

1. [Ferg10] N. Ferguson，B。Schneier，T。Kohno。_密码工程_。威利（Wiley），2010 年 3 月。
2. [Moxie11] Moxie Marlinspike。“密码厄运原理”。http://www.thoughtcrime.org/blog/the-cryptographic-doom-principle/
3. [Perc09] C. Percival。“密码正确答案”。http://www.daemonology.net/blog/2009-06-11-cryptographic-right-answers.html，2009-06-11。
4. [Rizzo10] J. Rizzo，T。Duong。“实用的 oracle 填充攻击。” 在第四届 USENIX 进攻技术会议论文集（WOOT'10）中。USENIX 协会，美国加利福尼亚州伯克利，1-8。2010 年。
5. [Vaud02] S. Vaudenay。“ CBC 填充引起的安全缺陷-SSL，IPSEC，WTLS 的应用……” 在国际密码技术理论与应用会议论文集：密码学的进展（EUROCRYPT '02）中，拉尔斯
6. 努森（编辑）。Springer-Verlag，英国伦敦，英国，534-546。2002 年。

## 第 4 章：安全通道和密钥交换

在上一章中，我们介绍了对称安全性。对称加密虽然速度很快，但要求双方共享同一密钥。但是，假定如果需要加密，则通信介质是不安全的。在本章中，我们将介绍一些在不安全通道上交换密钥的机制。

### 安全通道

本章的目的是在两个对等点之间建立**安全通道**：攻击者无法在其中窃听或伪造消息。通常，我们将使用对称安全性来提供安全通道，并使用一些密钥交换机制来设置初始密钥。该通道将在两个方向上建立，这意味着应该有单独的一组密钥（组合的加密和身份验证密钥）来在两个方向上进行通信。

安全通道是什么样的？出于上一章中所述的相同原因，它将是面向消息的。它将建立在无法信任的不安全通道（如 Internet）的顶部。例如，攻击者可能能够破坏或操纵这种不安全的渠道；我们也无法隐藏大型消息或正在讲话的人。理想情况下，消息将不可重播或更改其顺序（并且在出现任何一种时都将是显而易见的）。

防止重放的最简单方法是跟踪消息号。此号码可能会序列化为消息的一部分。例如，一条消息可能被认为是消息编号和某些消息内容的配对。我们将跟踪已发送和已接收消息的编号。

```
1 type Message struct {
2         Number   uint32
3         Contents []byte
4 }
```

序列化消息会将内容附加到 4 字节消息号上。该`out`变量仅用四个字节初始化，但具有可容纳消息内容的容量。

```
1 func MarshalMessage(m Message) []byte {
2         out := make([]byte, 4, len(m.Contents) + 4)
3         binary.BigEndian.PutUint32(out[:4], m.Number)
4         return append(out, m.Contents...)
5 }
```

首先，对消息进行解组检查将检查消息是否包含序列号和至少一个内容字节的假设。然后，提取消息编号和内容。

```
 1 func UnmarshalMessage(in []byte) (Message, bool) {
 2         m := Message{}
 3         if len(in) <= 4 {
 4                 return m, false
 5         }
 6
 7         m.Number = binary.BigEndian.Uint32(in[:4])
 8         m.Contents = in[4:]
 9         return m, true
10 }
```

仅在检查消息时包括消息号才有用。我们将跟踪给定会话的消息号：如本章所述，我们为每个会话交换唯一的密钥，因此我们仅跟踪会话范围内的消息号。在会话之外重播的消息将不可解密，因此我们不必担心。我们还希望跟踪分别发送的消息和已接收的消息的消息数。我们还为每个方向保留单独的键。

```
 1 type Channel io.ReadWriter
 2
 3 type Session struct {
 4         lastSent uint32
 5         sendKey *[32]byte
 6
 7         lastRecv uint32
 8         recvKey *[32]byte
 9
10         Channel Channel
11 }
12
13 func (s *Session) LastSent() uint32 {
14         return s.lastSent
15 }
16
17 func (s *Session) LastRecv() uint32 {
18         return s.lastRecv
19 }
```

当我们将消息加密为会话的一部分时，我们会将消息号设置为最后发送的消息号加 1。序列化的消息号和内容然后被加密，并传递给发送。

```
1 func (s *Session) Encrypt(message []byte) ([]byte, error) {
2         if len(message) == 0 {
3                 return nil, secret.ErrEncrypt
4         }
5
6         s.lastSent++
7         m := MarshalMessage(Message{s.lastSent, message})
8         return secret.Encrypt(s.sendKey, m)
9 }
```

现在，确保我们具有递增的消息号是解密的要求。如果消息号没有增加，我们假定它是重播的消息，并认为它是解密失败。

```
 1 func (s *Session) Decrypt(message []byte) ([]byte, error) {
 2         out, err := secret.Decrypt(s.recvKey, message)
 3         if err != nil {
 4                 return nil, err
 5         }
 6
 7         m, ok := UnmarshalMessage(out)
 8         if !ok {
 9                 return nil, secret.ErrDecrypt
10         }
11
12         if m.Number <= s.lastRecv {
13                 return nil, secret.ErrDecrypt
14         }
15
16         s.lastRecv = m.Number
17
18         return m.Contents, nil
19 }
```

这个例子可以简单地扩展为包括其他消息元数据。使用像 GCM 这样的 AEAD 意味着不需要加密此元数据。但是，我们宁愿尽可能多地加密，以将有关消息的信息量限制到窃听者。

可以在`chapter4/session/`包中的示例源代码中找到会话的更完整示例。

### 基于密码的密钥派生

交换密钥的最简单机制是使用密码派生密钥。这样做有几种选择，但是我们更喜欢的是 scrypt。它在`golang.org/x/crypto/scrypt`包装中提供。在内部，scrypt 使用另一个称为 PBKDF2 的密钥派生函数（KDF），但它增加了资源需求。这要求实现在使用大量内存或大量 CPU 之间进行选择，以降低硬件实现的有效性，从而使它们的生产成本更高。

scrypt 的资源需求由其参数控制：

- `N`：迭代次数。在该算法的原始演示文稿[Perc09b]中，对于交互式登录，建议使用 16384（2 14），对于文件加密，建议使用 1048576（2 20）。在本书中，我们使用 2 20 作为默认值来保护密码秘密。对于我们将使用它的应用程序，一次性获得密钥的初始成本是可以接受的。
- `r`：相对内存成本参数（控制基础哈希中的块大小）。原始演示文稿建议值为 8。
- `p`：相对 CPU 成本参数。原始演示文稿建议值为 1。

在我写这本书的机器上，一台具有 2.67 GHz 四核 Intel Core i5 M560 和 6G 内存的 2010 ThinkPad，scrypt 平均需要 6.3s 来获得密钥。由于密钥材料的获取方式，生成 32 字节的 NaCl 或 AES-GCM 密钥与生成 80 字节的 AES-CBC 或 AES-CTR 密钥之间没有明显的时序差异。

为了生成密钥，scrypt 需要输入密码和密码。盐不是秘密，我们将其添加到使用 scrypt 派生的密钥加密的数据之前。盐传递给 PBKDF2，用于防止攻击者仅存储`(password,key)`对，因为不同的盐会产生与 scrypt 不同的输出。每个密码都必须与用于导出密钥的盐一起检查。我们将使用[Ferg10]中推荐的随机生成的 256 位盐。

在上一章中，提到了我们的密钥交换方法允许我们使用随机随机数。当我们使用 scrypt 进行密钥交换时，每种加密都使用新的盐。这实际上意味着即使我们使用相同的密码短语，每个加密也将使用不同的密钥。这在天文学上意味着我们不太可能重复使用具有相同密钥的随机数。

在 Go 中使用 scrypt 是一种方法（加上错误检查），因此让我们尝试使用通行密码用 NaCl 加密某些内容。第一步是编写密钥派生函数，该函数主要将返回的字节片转换为包`scrypt.Key`所`*[32]byte`使用 的字节片`secretbox`。

```
 1 // deriveKey generates a new NaCl key from a passphrase and salt.
 2 func deriveKey(pass, salt []byte) (*[secret.KeySize]byte, error) {
 3         var naclKey = new([secret.KeySize]byte)
 4         key, err := scrypt.Key(pass, salt, 1048576, 8, 1, secret.KeySize)
 5         if err != nil {
 6                 return nil, err
 7         }
 8
 9         copy(naclKey[:], key)
10         util.Zero(key)
11         return naclKey, nil
12 }
```

加密功能将使用密码短语和一条消息，生成一个随机盐，导出密钥，对该消息进行加密，然后将盐添加到最终的密文中。

```
 1 func Encrypt(pass, message []byte) ([]byte, error) {
 2         salt, err := util.RandBytes(SaltSize)
 3         if err != nil {
 4                 return nil, ErrEncrypt
 5         }
 6
 7         key, err := deriveKey(pass, salt)
 8         if err != nil {
 9                 return nil, ErrEncrypt
10         }
11
12         out, err := secret.Encrypt(key, message)
13         util.Zero(key[:]) // Zero key immediately after
14         if err != nil {
15                 return nil, ErrEncrypt
16         }
17
18         out = append(salt, out...)
19         return out, nil
20 }
```

仅在调用时才需要派生密钥`secret.Encrypt`：它在调用之前立即派生，并在调用之后立即归零。这是一种仅在需要时才将秘密保留在内存中的尝试。同样，调用此函数后，调用者应立即将密码短语清零。错误处理可以等到密钥归零后再进行，这有助于防止密钥材料从此功能中泄漏出来的可能性。

为了解密，我们检查关于消息长度的假设（也就是说，正确加密的密码短语保护的消息将加上盐，再加上 NaCl 随机数和开销）。然后，我们导出密钥，解密并返回纯文本。

```
 1 const Overhead = SaltSize + secretbox.Overhead + secret.NonceSize
 2
 3 func Decrypt(pass, message []byte) ([]byte, error) {
 4         if len(message) < Overhead {
 5                 return nil, ErrDecrypt
 6         }
 7
 8         key, err := deriveKey(pass, message[:SaltSize])
 9         if err != nil {
10                 return nil, ErrDecrypt
11         }
12
13         out, err := secret.Decrypt(key, message[SaltSize:])
14         util.Zero(key[:]) // Zero key immediately after
15         if err != nil {
16                 return nil, ErrDecrypt
17         }
18
19         return out, nil
20 }
```

再次尝试限制内存中密钥材料的范围。这不能保证，但这是我们能做到的最好的。

使用密码也称为“预共享机密”；预共享意味着必须通过其他一些安全通道来交换密码。例如，您的无线网络可能使用带有预共享机密（您的无线密码）的加密。该密码用于保护网络通信的安全，并且您可以告诉您的朋友，该密码是相当安全的，前提是假设以攻击者为目标的网络攻击者不会物理存在（除非他们在您的计算机上安装了远程访问工具并且在麦克风上收听）。通常，可以窃听此语言密钥交换的攻击者并不是您担心要窃听网络流量的攻击者。

的`chapter4/passcrypt`示例代码包中包含的基于密码的加密的例子。

### 非对称密钥交换：ECDH

另一种选择是通过使用非对称密钥执行密钥交换来就对称密钥达成一致。特别是，我们将使用一种基于椭圆曲线（[Sull13]）的非对称算法，称为椭圆曲线 Diffie-Hellman 密钥协商协议，它是 Diffie-Hellman 密钥交换中椭圆曲线的一种变体。

非对称加密算法或公共密钥加密算法是一种用于加密的密钥与用于解密的密钥不同的算法。用于加密的密钥是公开的，拥有此公共密钥的任何人都可以对发送给它的消息进行加密。解密密钥由持有者保密。在对称安全性中，所有通信方都必须共享密钥，这意味着`(N * (N - 1)) / 2`密钥。使用非对称密钥，仅需要交换 N 个公钥。

在非对称密钥交换中，双方都有一个由私有部分和公共部分组成的密钥对。它们甚至在不安全的通道上也互相发送其公共部分，并且 ECDH 意味着组合您的私钥和其公钥将产生与组合其私钥和您的公钥相同的对称密钥。但是，只能在使用相同曲线的关键点之间执行 ECDH。

生成椭圆曲线密钥对通常是一种快速的操作，在本书中，我们将**临时**密钥对用于会话：在每个会话开始时生成的密钥对，如果会话结束则至少丢弃一次（如果不是）较早。使用临时密钥限制了密钥泄露的危害：设法恢复临时密钥的攻击者只能解密该会话的消息。

如工程关注点一章所述，将非对称密码学引入系统会带来很多其他问题。归结为信任，在某个时候，必须有信任根，特定密钥交换中的所有参与者都同意信任。这不是一个小问题，应该仔细考虑。尽管它确实提供了在不预先共享秘密的情况下通过不安全的渠道执行关键协议的能力，但需要将其好处与其他因素进行权衡。例如，我们将在下一章中介绍数字签名，但现在足以注意到通常使用长期身份密钥来对会话密钥进行签名以证明其所有权。现在双方都必须确定如何获取对方的身份密钥，如何知道自己可以信任它，以及他们如何将公钥与正在与之对话的对等项进行匹配。必须确定一些分配和信任公共密钥的方法。

#### NaCl：Curve25519

在 NaCl 中，生成新的密钥对和执行密钥交换非常简单，并且由中的函数提供`golang.org/x/crypto/nacl/box`。

```
1 pub, priv, err := box.GenerateKey(rand.Reader)
```

我们可以使用`Session`前面定义的设置来执行密钥交换。首先，在上创建一个新会话，`Channel`并分配密钥。

```
1 func NewSession(ch Channel) *Session {
2         return &Session{
3                 recvKey: new([32]byte),
4                 sendKey: new([32]byte),
5 	    Channe: ch,
6         }
7 }
```

该`keyExchange`函数获取字节片并填充适当的密钥材料。在不再需要密钥材料时，它还会尝试将其归零。

```
 1 func keyExchange(shared *[32]byte, priv, pub []byte) {
 2         var kexPriv [32]byte
 3         copy(kexPriv[:], priv)
 4         util.Zero(priv)
 5
 6         var kexPub [32]byte
 7         copy(kexPub[:], pub)
 8
 9         box.Precompute(shared, &kexPub, &kexPriv)
10         util.Zero(kexPriv[:])
11 }
```

最后，密钥交换将计算发送和接收密钥。该`dialer` 选项应`true`适用于发起对话的一方（通过拨打另一方）。

```
 1 func (s *Session) KeyExchange(priv, peer *[64]byte, dialer bool) {
 2         if dialer {
 3                 keyExchange(s.sendKey, priv[:32], peer[:32])
 4                 keyExchange(s.recvKey, priv[32:], peer[32:])
 5         } else {
 6                 keyExchange(s.recvKey, priv[:32], peer[:32])
 7                 keyExchange(s.sendKey, priv[32:], peer[32:])
 8         }
 9         s.lastSent = 0
10         s.lastRecv = 0
11 }
```

该`session`示例包含`Dial`和`Listen`函数，该函数在通道上设置密钥交换。

该`Precompute`功能实际执行密钥交换；但是，NaCl 包也可以为每个消息进行密钥交换。取决于应用程序，这可能效率较低。

```
1 out := box.Seal(nil, message, nonce, peerPublic, priv)
```

这对于一次性消息可能很有用。在这种情况下，通常为每个用于密钥交换的消息生成一个新的临时密钥对，并在开始时打包临时公共密钥，以便接收者可以解密。该`chapter4/naclbox`软件包包含一个以这种方式保护消息的示例。

在 NaCl 中，所有键对都使用相同的椭圆曲线 Curve25519（[DJB05]）。这使接口变得简单，并简化了公共密钥的交换：公共密钥始终为 32 个字节。只要使用的是 Curve25519，我们就不必担心发件人可能正在使用哪条曲线。

### NIST 曲线

Go 标准库支持常用的 NIST 曲线，其中有几种。为了使用这些 NIST 曲线进行通信，双方必须就使用哪条曲线达成共识。这给前面提到的注意事项增加了额外的开销：现在，系统设计人员不仅必须考虑密钥来自何处以及如何被信任，还必须考虑使用哪种类型的曲线。有[严重关切](https://www.schneier.com/blog/archives/2013/09/the_nsa_is_brea.html#c1675929) 在这些曲线的来源，以及他们是否已经设计了一个 NSA 后门; 仅在出于兼容性考虑或作为规范的一部分时才应使用它们。它们可以使用生成

```
1 priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
```

感兴趣的三个曲线为 P256，P384 和 P521（或 secp256r1 / prime256v1，secp384r1 和 secp521r1）。P521 被认为具有与 AES-256 等效的安全级别（尽管对曲线的完整性存在疑问），因此它适合与 AES-256 和 HMAC 密码套件一起使用。

使用 Go`crypto/ecdsa`软件包执行此操作更加复杂。我们首先必须验证曲线是否匹配，然后执行标量乘法（或点乘法）。返回的数字是共享密钥，但不会统一随机。我们确实希望共享密钥具有这种不可区分的属性，因此我们将计算该数字的 SHA-512 摘要，该摘要会产生一个 64 字节的值：此大小足以与具有 HMAC-SHA-的 AES-256-CBC 一起使用 256 和带有 HMAC-SHA-256 的 AES-256-CTR。

```
 1 var ErrKeyExchange = errors.New("key exchange failed")
 2
 3 func ECDH(priv *ecdsa.PrivateKey, pub *ecdsa.PublicKey) ([]byte, error) {
 4         if prv.PublicKey.Curve != pub.Curve {
 5                 return nil, ErrKeyExchange
 6         } else if !prv.PublicKey.Curve.IsOnCurve(pub.X, pub.Y) {
 7                 return nil, ErrKeyExchange
 8         }
 9
10         x, _ := pub.Curve.ScalarMult(pub.X, pub.Y, prv.D.Bytes())
11         if x == nil {
12                 return nil, ErrKeyExchange
13         }
14
15         shared := sha512.Sum512(x.Bytes())
16         return shared[:secret.KeySize], nil
17 }
```

请注意检查以确保公钥有效：不检查此项可能导致公开私钥（[Antip03]）。我们不信任解析器为我们执行此验证，因此我们将在此处进行检查。

由于支持对软件包中的 ECDSA 密钥进行序列化（即使用和 功能）`crypto/ecdsa`，因此此处使用的密钥代替了`crypto/elliptic`软件包中定义的 密钥。` crypto/x509``MarshalPKIXPublicKey``ParsePKIXPublicKey `

```
 1 func ParseECPublicKey(in []byte) (*ecdsa.PublicKey, error) {
 2         // UnmarshalPKIXPublicKey returns an interface{}.
 3         pub, err := x509.ParsePKIXPublicKey(in)
 4         if err != nil {
 5                 return nil, err
 6         }
 7
 8         ecpub, ok := pub.(*ecdsa.PublicKey)
 9         if !ok {
10                 return nil, errors.New("invalid EC public key")
11         }
12
13         return ecpub, nil
14 }
```

由于对 NIST 曲线的担心，我们在本书中通常会避免使用它们（尽管它们是为将要查看的其他系统定义的），并且主要会坚持使用 Curve25519。

该`chapter4/nistecdh`示例代码包中包含了相关的例子。

### 其他密钥交换方法

密码验证的密钥交换是用于协商也使用公共密钥的对称密钥的另一种机制。我没有在 Go 中看到任何实现，因此在本书中我们不会使用它们。

另一种机制是原始的 Diffie-Hellman 算法。一种实现是在“ github.com/dchest/dhgroup14”。但是，在本书中，除了 ECDH 之外，我们不会使用 DH。

还使用了更复杂的机制，例如 Kerberos 中发现的机制。

稍后，我们将使用称为 RSA 的非对称算法来交换密钥，尽管这不是我们首选的机制。椭圆曲线键可以足够快地生成，以使临时键实用。RSA 密钥生成速度要慢得多，并且 RSA 实现很难正确实现。另一种方法是使用 RSA 对临时密钥进行签名。这通常与 Diffie-Hellman 一起使用。

### 实用：文件加密器

为了付诸实践，请尝试构建一个使用密码生成密钥的文件加密程序。首先，考虑程序应如何工作：应加密多个文件还是单个文件？安全模型是什么样的？它需要在哪种平台上运行？有兼容性要求吗？

我对此的解决方案是在 [filecrypt](https://github.com/kisom/filecrypt/)存储库中。您的要求和规范（以及安全模型）可能会有所不同，因此最终可能会得到不同的解决方案。

### 进一步阅读

1. [Antip03] A. Antipa，DRL 布朗，A。Menezes，R。Struik，
2. 1. 凡斯通。椭圆曲线公钥的验证，《计算机科学讲座》，第 2567 卷，第 211-223 页，Springer。2003。
3. [DJB05] DJ 伯恩斯坦。“ Curve25519：新的 Diffie-Hellman 速度记录。” PKC 2006 会议录出现。2005 年 11 月 15 日。
4. [Ferg10] N. Ferguson，B。Schneier，T。Kohno。_密码工程_。威利（Wiley），2010 年 3 月。
5. [Hao2008] F. Hao，P。Ryan。“通过杂耍通过密码验证的密钥交换。” 第 16 届国际安全协议国际研讨会论文集，2008 年。
6. [Perc09a] C. Percival。“通过顺序存储-硬函数导出更强的密钥。” BSDCan'09，2009 年 5 月。
7. [Perc09b] C. Percival。“ scrypt：一种新的密钥派生功能。” BSDCan'09，2009 年 5 月。
8. [Sull13] N. Sullivan。“椭圆曲线密码学上的（相对容易理解的）入门书。” https://blog.cloudflare.com/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/ 2013 年 10 月 24 日。
9. [Wu2000] T. Wu。“ SRP 身份验证和密钥交换系统。” RFC 2945，IETF 和 Intenret Society，2000 年 9 月。

## 第 5 章：数字签名

非对称算法的另一个用途是对消息进行数字签名。它基于这样的假设：只有持有者才能为给定消息生成签名（即签名不可伪造），并且签名仅对单个消息有效。在本章中，我们将讨论使用 Ed25519 进行数字签名，并将其与 Curve25519 加密密钥一起使用。ECDSA，将与 NIST 曲线一起使用；和 RSA。在我们将在此处使用的算法中，消息不是直接签名的。而是对消息的哈希进行签名。有时，我们必须自己提供消息的哈希值，而其他时候，这将成为我们的计算机。

### 加密哈希算法

我们已经看到了哈希函数：在对称安全性章节中，SHA-256 被用作 HMAC 的一部分，而 SHA-512 被用作 NIST ECDH 的一部分。我们尚未讨论它们的作用或何时使用它们。

哈希函数将一些任意输入映射到固定大小的输出。例如，不管是什么`m`，`SHA-256(m)`总是产生 32 字节的输出。在非加密用途中，有时称为校验和。在密码学中，对哈希函数有更严格的要求，我们仅使用满足这些要求的密码哈希函数。在这本书中，术语“哈希函数”指的是密码哈希函数。

最基本的属性是**抗冲突性**：对于 m 1 和 m 2 中的任何两个消息，其中 m 1 不等于 m 2，m 1 的哈希值不应该等于 m 2。如果两个输入不相同，我们不希望两个输入发生冲突并产生相同的哈希值。

我们还希望我们的哈希函数具有原**像抵抗性**。这意味着如果我们有一些哈希函数的输出`x`，我们就不能找到一些`m`散列到的哈希值`x`。哈希函数应该是单向的：易于从消息中计算哈希，但是从哈希中计算消息不可行。

另一个特性是，更改`m`应导致）的更改 `h(m`。`h(m)`应该与随机是没有区别的，这意味着攻击者只要看到`h(m)`和`h(m')`（在哪里`m'`进行了任何修改，甚至是单个位`m`），都将无法确定它们是否来自相关消息。

考虑到散列函数输出固定大小的输出并接受任意输入，将发生一些冲突。我们接受计算上不可行的要求。必须及时发生成功的攻击才能利用该攻击。另一方面，散列函数必须在计算上有效，并且必须从其输入快速生成散列。

我们在本书中使用的哈希函数属于 SHA-2 哈希函数系列，它们根据生成的输出大小进行命名：SHA-256 生成 32 字节（256 位）的哈希值，SHA-384 生成 48 字节（384 位）的哈希，依此类推。有时会指定 SHA-1 散列，但是对其进行了攻击，使其不适合用于安全目的。请注意，对 SHA-1 的攻击不会扩展到 HMAC-SHA-1。这种结构仍然[被认为是安全的](https://www.schneier.com/blog/archives/2005/02/sha1_broken.html)。出于安全考虑，我们将使用 SHA-2 系列来保持一致性。

哈希通常被滥用；它们已被用于尝试保护密码或进行身份验证。我们将使用它们的目的是从某些数据中删除结构或从任意输入中产生一个 ID。在对称安全性一章中，我们使用具有 HMAC 结构的 SHA-256 来生成经过身份验证的消息 ID。在密钥交换一章中，我们使用 SHA-512 从产生的数字中删除了一些结构，以使其更适合用作密钥。

使用散列来保护密码的问题又回到了散列必须高效的要求。具有一组哈希密码的攻击者可以快速发起旨在恢复这些密码的攻击。在身份验证中，SHA-2 散列的构造方式使它们容易受到长度扩展攻击（例如，请参见[Duong09]）。

散列的另一个滥用是用于完整性检查，例如下载页面上文件的 SHA-256 摘要。这不提供任何真实性，可以替换下载文件的攻击者很可能能够替换摘要：无法保证摘要是由合法方提供的。HMAC 也不是正确的选择，因为任何拥有 HMAC 密钥来验证它的人都可以产生自己的密钥。只要公钥是众所周知的，在这种情况下数字签名就非常有用。如果仅在下载页面上提供了公共密钥，并且攻击者可以交换页面上的文件，则他们可以交换出公共密钥。

### 前向保密

ECDH 密钥交换给我们带来了一个问题：我们如何信任我们收到的公钥？假设我们选择了非对称路由并具有密钥分发和信任机制，那么我们有一套用于建立这种信任的工具。我们的信任机制不会用于会话密钥；这些仅持续一个会话的生命周期。取而代之的是，我们使用长期的身份密钥，然后将其用于对会话密钥进行签名。在会话开始时，将对会话密钥进行签名，而另一端将验证签名。

这种安排的关键特征是会话密钥和身份密钥是不相关的，在我们构建的系统中将始终是单独的密钥。在 NaCl 的情况下，它们甚至会变成不同的曲线。这遵循使用单独的密钥进行身份验证和保密的安全性原则。如果攻击者能够破解身份密钥，则他们将无法使用该密钥解密以前的任何会话。但是，他们可以使用身份密钥来签名任何将来的会话。

身份密钥本身就是数字。它们不携带有关承载的先天信息。使用身份密钥依赖于定义良好的密钥分配计划，该计划将公共密钥映射到身份，并为参与者提供一种机制，以获取适当的公共密钥（通常是任何关联的身份元数据）。他们还需要某种方法来确定密钥已被信任的程度。我们将在以后的章节中进一步讨论该问题，并且将看到一些现实世界的系统如何解决该问题。这绝对不是一个性感的问题，但对于成功的安全系统而言至关重要。

让我们考虑会话加密密钥被泄露的情况以及身份密钥被泄露的情况。在第一种情况下，密钥仅在单个会话的范围内。此处的妥协破坏了该会话的安全性，并保证检查失败以确定对策和适当的措施。但是，关键损害仅限于受影响的会话。如果身份密钥被泄露，则必须同样注意。但是，通常在与其他对等方进行通信时，会产生额外的开销，使他们将受到破坏的密钥标记为不可信。取决于密钥分发机制，更换该密钥可能是困难的和/或昂贵的。

### 编号 25519

Adam Langley[在 Github 上](https://github.com/agl/ed25519)具有[Ed25519](https://github.com/agl/ed25519)的实现，这是我们将在本书中使用的实现。Ed25519 私钥和签名的长度为 64 字节，而公钥的长度为 32 字节。此实现在内部使用 SHA-512 对消息进行哈希处理，因此我们将消息直接传递给它。

生成密钥的方法与 curve25519 密钥相同：

```
1 pub, priv, err := ed25519.GenerateKey(rand.Reader)n
```

签名是使用`Sign`函数进行的，并通过函数进行了验证 `Verify`：

```
1 sig, err := ed25519.Sign(priv, message)
2
3 if !ed25519.Verify(pub, message) {
4         // Perform signature verification failure handling.
5 }
```

在本书中，我们将首选 Ed25519 签名：接口简单，密钥小并且算法高效。签名也是 **确定性的，**并且不依赖于随机性。这意味着在 PRNG 失败的情况下，它们不会损害安全性，我们将在 ECDSA 的部分中详细讨论。此属性是开发此算法的主要动机之一。

### ECDSA

椭圆曲线数字签名算法或 ECDSA 是一种签名算法，将与前面提到的 NIST 曲线一起使用。正如 ECDH 是 DH 的椭圆曲线变体一样，ECDSA 是原始 DSA 的椭圆曲线变体。

DSA 中存在一个严重的缺陷（扩展到 ECDSA），已被多个实际系统（包括 Android 比特币钱包和 PS3）利用。签名算法依赖于质量随机性（与随机性没有区别的比特）；PRNG 进入可预测状态后，签名可能会泄漏私钥。使用 ECDSA 的系统必须意识到这一问题，并特别注意其 PRNG。

ECDSA 签名通常在签名中提供一对数字，称为`r`和`s`。序列化这两个数字的最常见方法是使用[SEC1]中定义的 ASN.1，在使用 ECDSA 签名时将使用此序列号。

```
 1 type ECDSASignature struct {
 2         R, S *big.Int
 3 }
 4
 5 func SignMessage(priv *ecdsa.PrivateKey, message []byte) {
 6         hashed := sha256.Sum256(message)
 7         r, s, err := ecdsa.Sign(rand.Reader, priv, hashed[:])
 8         if err != nil {
 9                 return nil, err
10         }
11
12         return asn1.Marshal(ECDSASignature{r, s})
13 }
14
15 func VerifyMessage(pub *ecdsa.PublicKey, message []byte, signature []byte) bool {
16         var rs ECDSASignature
17
18         if _, err := asn.Unmarshal(signature, &rs); err != nil {
19                 return false
20         }
21
22         hashed := sha256.Sum256(message)
23         return ecdsa.Verify(pub, hashed[:], rs.R, rs.S)
24 }
```

### RSA

RSA 是与迄今为止所见算法不同类型的非对称算法。它不是基于椭圆曲线的，并且在历史上比椭圆曲线加密更广泛地被使用。值得注意的是，尽管椭圆曲线密码术越来越多地使用，但 TLS 和 PGP 都大量使用了 RSA。

在 PKCS＃1 标准（[PKCS1]）中为 RSA 定义了两种签名方案：PKCS＃1 v1.5 和概率签名方案。这些都是带有附录的签名方案，签名和消息是分开的。通常，签名会附加到邮件中。相反，有一些具有消息恢复功能的签名方案（尽管未为 RSA 定义），其中签名验证步骤会生成原始消息，作为验证的副产品。

在 RSA 签名的各种实现中发现了许多漏洞。最近的是 [BERserk](http://www.intelsecurity.com/resources/wp-berserk-analysis-part-1.pdf)。即使 Go 实施编写得不错，其他组件中使用的库也可能不正确。Go 还可以做出其他一些适当的选择，例如选择可接受的公共指数。RSA 加密情况甚至更糟。我们将尽可能避免使用 RSA。有很多地方可能会出错，而且很难做到正确。当我们需要兼容性（例如与 TLS）时，将使用它。在这种情况下，我们将遵循规范。我们还将在完全用 Go 编写的应用程序中使用 PSS，因为它具有更强的安全性证明（[Lind06]，[Jons01]）。

使用此`SignPKCS1v15` 功能可以完成 PKCS＃1 v1.5 的 RSA 签名。同样，使用来执行 PSS 签名`SignPSS`。在这两种情况下，都必须先对消息进行哈希处理，然后再将其发送给函数。该`SignPSS`函数还带有一个`rsa/PSSOptions`值；这应该在您的程序中定义为

```
1 var opts = &rsa.PSSOptions{}
```

`opts`除非规范明确要求，否则您不应更改任何值。

以下是使用 SHA-256 作为哈希函数进行签名的示例函数。

```
1 func SignPKCS1v15(priv *rsa.PrivateKey, message []byte) ([]byte, error) {
2         hashed := sha256.Sum256(message)
3         return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
4 }
5
6 func SignPSS(priv *rsa.PrivateKey, message []byte) ([]byte, error) {
7         hash := sha256.Sum256(message)
8         return rsa.SignPSS(rand.Reader, priv, crypto.SHA256, hashed[:], opts)
9 }
```

### 结论

在某些情况下，数字签名提供了一种有用的机制来提供真实性和完整性，但应将其使用与对公钥基础结构的最终要求进行权衡。我们首选 Ed25519 作为 NaCl 套件的一部分，因为它对 PRNG 失败具有鲁棒性，简单性，安全性和效率。还应检查用于构建系统组件的其他语言使用的任何库的质量，因为它们很难正确使用。

### 实用：具有身份的会议

从上一个示例扩展会话示例以对会话密钥进行签名。请记住，需要考虑用于分发和验证签名密钥的机制。在编写代码之前，您应该编写描述签名密钥分发和验证的安全模型。该`chapter5/sessions`子包中包含了一个解决这个问题。

在 Github 上的[go-schannel](https://github.com/kisom/go-schannel)包中也可以找到真实的解决方案 。

### 进一步阅读

1. [Duong09] T. Duong，J。Rizzo。“ Flickr 的 API 签名伪造漏洞。” http://netifera.com/research/flickr_api_signature_forgery.pdf，2009年9月28日。
2. [Ferg10] N. Ferguson，B。Schneier，T。Kohno。_密码工程_。威利（Wiley），2010 年 3 月。
3. [Jons01] J. Jonsson，“ RSA-PSS 签名的安全证明”。欧洲 RSA 实验室，瑞典斯德哥尔摩，2001 年 7 月。
4. [Lind06] C. Lindenberg，K。Wirt，J。Buchmann。“ RSA-PSS 正确性的形式证明。”，达姆施塔特工业大学计算机科学系，德国达姆施塔特，2006 年。
5. [PKCS1] J. Jonsson，B。Kaliski。“公钥密码标准（PKCS）＃1：RSA 密码规范 2.1 版。” RFC 3447，IETF 和互联网协会，2003 年 2 月。
6. [SEC1] DRL 布朗。高效密码技术的标准 1：椭圆曲线密码技术。Certicom Corp，2009 年 5 月。

## 附录：各章的加密审查

本附录列出了其他人进行过密码学审查的章节。我希望获得对此的反馈，并且愿意接受任何评论。

- 第 1 章：**未审查**
- 第 2 章：**未审核**
- 第三章：**未审查**
- 第 4 章：**未审查**
- 第 5 章：**未审查**

阅读本书时请牢记这一点。
