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