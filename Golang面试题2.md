# Golang面试题2

遇到过的Golang题目。（选择题）

### 1、关于channel的特性，下面说法正确的是（）[不定项]

```
A.给一个 nil channel 发送数据，造成永远阻塞
B.从一个 nil channel 接收数据，造成永远阻塞
C.给一个已经关闭的 channel 发送数据，引起 panic
D.从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
```

答案：ABCD

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