# Golang基础

**Go语言是一种并发的、带垃圾回收的、快速编译的新语言**。它具有以下特点：
● 可以在一台计算机上仅用几秒钟的时间编译一个大型的Go语言程序。

● Go 语言为软件构造提供了一种模型，它使依赖分析更加容易，且避免了大部分C语言风格include文件与库的开头。

● Go 语言是静态类型的语言，它的类型系统没有层级。因此，用户不需要在定义类型之间的关系上花费时间，看似比典型的面向对象语言更轻量级。

● Go语言完全是垃圾回收型的语言，而且为并发执行与通信提供了基本的支持。

● Go 语言是一种云计算时代的语言，它能够充分利用计算机的多核，通过轻量级别的goroutine就可以实现多并发。

Go 语言是一种编译型语言，它结合了解释型语言的游刃有余，动态类型语言的开发效率，以及静态类型的安全性。

------

## 1. 关键字

Golang的25个关键字或保留字：

|  **break**   |   **default**   |  **func**  | **interface** | **select** |
| :----------: | :-------------: | :--------: | :-----------: | :--------: |
|   **case**   |    **defer**    |   **go**   |    **map**    | **struct** |
|   **chan**   |    **else**     |  **goto**  |  **package**  | **switch** |
|  **const**   | **fallthrough** |   **if**   |   **range**   |  **type**  |
| **continue** |     **for**     | **import** |  **return**   |  **var**   |

Golangde 36个预定义标识符：

| **append** |  **imag**   |  **int**  |   **new**   |  **uint**   |  **complex**   |
| :--------: | :---------: | :-------: | :---------: | :---------: | :------------: |
|  **bool**  |  **false**  | **int8**  |   **nil**   |  **uint8**  | **complex64**  |
|  **byte**  | **float32** | **int16** |  **panic**  | **uint16**  | **complex128** |
| **close**  | **float64** | **int32** |  **print**  | **uint32**  |  **recover**   |
|  **copy**  |   **len**   | **int64** | **println** | **uint64**  |   **string**   |
| **false**  |  **make**   | **iota**  |  **real**   | **uintptr** |    **true**    |

------

## 2. channel

### 1.什么是channel

channel在golang语言中主要应用于routine之间的通信机制，我们学习计算机知道进程间（线程间）通信的机制里面也有类似的方法：管道，消息，socket（套接字），共享内存等。channel就类似管道机制。

### 2.channel有什么特性

channel分为有缓冲和无缓冲的划分，也有方向的划分。channel是可以关闭的（close方法），关闭的channle就不能继续使用了。

从无缓存的 channel 中读取消息会阻塞，直到有 goroutine 向该 channel 中发送消息；同理，向无缓存的 channel 中发送消息也会阻塞，直到有 goroutine 从 channel 中读取消息。

> 通过无缓存的 channel 进行通信时，接收者收到数据 happens before 发送者 goroutine 唤醒

有缓存的 channel 的声明方式为指定 make 函数的第二个参数，该参数为 channel 缓存的容量

```
ch := make(chan int, 10)
```

有缓存的 channel 类似一个阻塞队列(采用环形数组实现)。当缓存未满时，向 channel 中发送消息时不会阻塞，当缓存满时，发送操作将被阻塞，直到有其他 goroutine 从中读取消息；相应的，当 channel 中消息不为空时，读取消息不会出现阻塞，当 channel 为空时，读取操作会造成阻塞，直到有 goroutine 向 channel 中写入消息。

通过 len 函数可以获得 chan 中的元素个数，通过 cap 函数可以得到 channel 的缓存长度。

### 3.channel的应用场景

channel一般是结合goroutine进行使用，在并发控制的时候进行消息机制，可以有以下的编程模式以及业务实现：

- 异步调用的通知机制
- goroutine的结合select可以做到多路复用
- 单向的消息传递
- 生产者和消费者的实现
- 数据缓存
- goroutine的负载均衡
- 转线程（协程）等

### 4.channel使用注意问题

因为可以实现的业务的设计模式比较多，不同的业务设计模式需要结合不同的做法，channel本身实现的管道功能比较简单。
所以业务代码需要有很清晰的实现思路，那么就需要业务上对以下方面特别注意：
有缓冲的channel（就是定义的时候给了channel大小），对于生成者和消费者的协同，要考虑消费者的效率要高于生产者，否则会出现写入的等待，如果是做网络数据接收，这个时候会影响数据接收，引发连锁反应。
在多routine间的channel数据传递，需要业务层明确定义每个channel的状态和业务，确保状态机是明确的有向状态机。莫要出现复杂的状态转换