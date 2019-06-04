# Golang面试题2

遇到过的Golang题目。（选择题）

### 1、关于channel的特性，下面说法正确的是（）[不定项]

```
A.给一个 nil channel 发送数据，造成永远阻塞
B.从一个 nil channel 接收数据，造成永远阻塞
C.给一个已经关闭的 channel 发送数据，引起 panic
D.从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
```

答案：A B C D

解答：nil channel代表channel未初始化，向未初始化的channel读写数据会造成永久阻塞。Ps.关闭(close)未初始化的channel会引起panic。

通道为空时，会报deadlock错误。fatal error: all goroutines are asleep - deadlock!

### 2、可以给任意类型添加相应的方法。这一说法是否正确。

```
A.true
B.false
```

答案：B

解答：必须是自定义的类型。正确说法：任意自定义类型(包括内置类型，但不包括指针类型)添加相应的方法。

### 3、关于类型转化，下面语法正确的是（）[单项]

![img](https://uploadfiles.nowcoder.com/images/20171203/3367369_1512277518304_28E6FBEBA19DFA1C862F3891BF21A5AB)

答案：C

解答：Go语言类型转换语法：Type(expression)  D选项是类型断言，类型断言语法为：expression.(Type) 对于类型断言，首先 expression 必须是接口类型，但D选项中 i 是 int 类型，无法进行类型断言；其次 i 是 int 类型，无法通过类型断言转换成 MyInt，只有类型相符时，类型断言才会成功。

### 4、关于select机制，下面说法正确的是（）[不定项]

```
A.select机制用来处理异步IO问题

B.select机制最大的一条限制就是每个case语句里必须是一个IO操作

C.golang在语言级别支持select关键字

D.select关键字的用法与switch语句非常类似，后面要带判断条件
```

答案：A B C

解答：golang 的 select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。  

在执行select语句的时候，运行时系统会随机地判断每个case中的发送或接收操作是否可以被立即执行(立即执行：意思是当前Goroutine不会因此操作而被阻塞) 

select的用法与switch非常类似，由select开始一个新的选择块，每个选择条件由case语句来描述。与switch语句可以选择任何可使用相等比较的条件相比，select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作。

### 5、关于slice或map操作，下面正确的是（）[不定项]

![img](https://uploadfiles.nowcoder.com/images/20171203/3367369_1512279828612_E135777046538A1F75605878FC012839)

答案：A C D

解答：Make只用来创建slice,map,channel。 其中map使用前必须初始化。 append可直接动态扩容slice，而map不行。

### 6、Golang可以复用C/C++的模块，这个功能叫Cgo,这一说法是否正确。

```
A.true

B.false
```

答案：B

解答： CGO是C语言和Go语言之间的桥梁，原则上无法直接支持C++的类。CGO不支持C++语法的根本原因是C++至今为止还没有一个二进制接口规范(ABI)。CGO只支持C语言中值类型的数据类型，所以我们是无法直接使用C++的引用参数等特性的。

###   7、关于内存泄露，下面说法正确的是（）[不定项]

```
A.golang有自动垃圾回收，不存在内存泄露
```

```
B.golang中检测内存泄露主要依靠的是pprof包
```

```
C.内存泄露可以在编译阶段发现
```

```
D.应定期使用浏览器来查看系统的实时内存信息，及时发现内存泄露问题
```

答案：B D

解答：内存泄漏（Memory Leak）是指程序中己动态分配的堆内存由于某种原因程序未释放或无法释放，造成系统内存的浪费，导致程序运行速度减慢甚至系统崩溃等严重后果。 

在影响软件系统稳定性的因素里，我们最担心的一个问题是内存泄漏，随着系统的运行，系统消耗的内存越来越多，直到最后整个操作系统越来越慢，甚至还会导致系统崩溃。在Go语言里，我们检测内存泄漏主要依靠的是go里面的[pprof包](http://golang.org/pkg/runtime/pprof/)，除此之外，我们还可以使用浏览器来查看系统的实时内存信息(包括CPU、goroutine等的信息)。

### 8、关于布尔变量b的赋值，下面错误的用法是（）[不定项]

```
A.b = true
```

```
B.b = 1
```

```
C.b = bool(1)
```

```
D.b = (1 == 2)
```

答案：B C

解答：bool类型和int类型没法强制转换，go语言的bool类型值只有false和true。

### 9、下面的程序的运行结果是:

```go
package main
 func main() {
     if true {
        defer fmt.Printf("1")
     } else {
         defer fmt.Printf("2")
     }
     fmt.Printf("3")
}
```

```
A. 321

B. 32

C. 31

D. 13
```

答案：C

解答：Go语言的defer（延迟执行）语句，是在函数结束前执行，而如果在函数中有多个defer语句时，会像一个堆栈一样，先进后出，后进先出 。

### 10、关于switch语句，下面说法正确的有:

```
A . 条件表达式必须为常量或者整数.

B.  单个case中，可以出现多个结果选项.

C.  需要用break来明确退出一个case

D.  只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case
```

答案：B D

解答：Go 语言改进了 switch 的语法设计，避免人为造成失误。Go 语言的 switch 中的每一个 case 与 case 间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行。 在 Go 语言中 case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着下一个 case 执行。但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能:

```go
var s = "hello"
switch {
case s == "hello":
    fmt.Println("hello")
    fallthrough
case s != "world":
    fmt.Println("world")
}
```

结果：

```
hello
world
```

11、[多选题]golang中的引用类型包括（）

```
A.数组切片
B.map
C.channel
D.interface
```

答案：A B C D
解答：golang中分为值类型和引用类型

值类型分别有：int系列、float系列、bool、string、数组和结构体

引用类型有：指针、slice切片、管道channel、接口interface、map、函数等

值类型的特点是：变量直接存储值，内存通常在栈中分配

引用类型的特点是：变量存储的是一个地址，这个地址对应的空间里才是真正存储的值，内存通常在堆中分配。

12、关于select机制，下面说法正确的是（）

```
A.select机制用来处理异步IO问题

B.select机制最大的一条限制就是每个case语句里必须是一个IO操作

C.golang在语言级别支持select关键字

D.select关键字的用法与switch语句非常类似，后面要带判断条件
```

答案：A B C

解答：golang 的 select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。  在执行select语句的时候，运行时系统会自上而下地判断每个case中的发送或接收操作是否可以被立即执行(立即执行：意思是当前Goroutine不会因此操作而被阻塞) 。  select的用法与switch非常类似，由select开始一个新的选择块，每个选择条件由case语句来描述。与switch语句可以选择任何可使用相等比较的条件相比，select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作。 

## 持续更新...