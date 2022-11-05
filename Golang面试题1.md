# Golang面试题1

网上遇到的面试题。（作答题）

------

### 1、写出下面代码的输出。

```go
 
package main
 
import (
    "fmt"
)
 
func main() {
    deferCall()
}
 
func deferCall() {
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印后") }()
    panic("触发异常")
}
```

考点：defer的执行顺序，panic的执行方式。

解答：defer 是后进先出。
panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。

```
打印后
打印中
打印前
panic: 触发异常
```

### 2、以下代码有什么问题，说明原因。

```go
type student struct {
    Name string
    Age  int
}
 
func main() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    for _, stu := range stus {
        m[stu.Name] = &stu
    }
 
}
```

考点：foreach

解答： 与Java的foreach一样，都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝。 就像想修改切片元素的属性：

```
for _, stu := range stus {
    stu.Age = stu.Age+10
}
```

也是不可行的。 大家可以试试打印出来：

```go
func main() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    // 错误写法
    for _, stu := range stus {
        m[stu.Name] = &stu
    }
 
    for k,v:=range m{
        println(k,"=>",v.Name)
    }
 
    // 正确
    for i:=0;i<len(stus);i++  {
        m[stus[i].Name] = &stus[i]
    }
    for k,v:=range m{
        println(k,"=>",v.Name)
    }
}
```

### 3、下面的代码会输出什么，并说明原因

```go
func main() {
    runtime.GOMAXPROCS(1)
    wg := sync.WaitGroup{}
    wg.Add(20)
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println("A: ", i)
            wg.Done()
        }()
    }
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println("B: ", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

考点：go执行的随机性和闭包

解答：WaitGroup的用途：它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
WaitGroup总共有三个方法：Add(delta int),Done(),Wait()。简单的说一下这三个方法的作用。
Add:添加或者减少等待goroutine的数量
Done:相当于Add(-1)
Wait:执行阻塞，直到所有的WaitGroup数量变成0

谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。 但是`A:`均为输出10，`B:`从0~9输出(顺序不定)。 第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。 故go func执行时，i的值始终是10。

第二个go func中i是函数参数，与外部for中的i完全是两个变量。 尾部(i)将发生值拷贝，go func内部指向值拷贝地址。

### 4. 下面代码会输出什么？

```go
type People struct{}
 
func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p *People) ShowB() {
    fmt.Println("showB")
}
 
type Teacher struct {
    People
}
 
func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}
 
func main() {
    t := Teacher{}
    t.ShowA()
}
```

考点：go的组合继承

解答：这是Golang的组合模式，可以实现OOP的继承。 被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），但它们的方法(ShowA())调用时接受者并没有发生变化。 此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。

```go
showA
showB
```

引申一点：

首先明确一点 go中没有继承关系。也不应该提及“继承”这个词，其中Trecher并没有继承Propler，而是嵌套People,
而t.ShowA()是一个语法糖，其实t.ShowA() = t.people.ShowA(),也就是说在嵌套结构中，go会优先调用本身方法，
如果本身没有此方法，就回去调用其所包含结构的方法。

本题中，showA()是Teacher不具有的，但是它所嵌套的People具有，因此回调用People.showA(),People.showA()
中调用了*People 的showB()当然会展示“shwoB”，而不是“teacher showB”

```
如果嵌套有两个结构，并且两个结构具有相同的方法，如何执行的？
```

```go
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Human struct{}

func (h *Human) ShowA() {
    fmt.Println("Human showA")
}

type Teacher struct {
	Human
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
```

答案是 编译报错，不支持这种情况的。

### 5. 下面代码会触发异常吗？请详细说明

```go
func main() {
    runtime.GOMAXPROCS(1)
    int_chan := make(chan int, 1)
    string_chan := make(chan string, 1)
    int_chan <- 1
    string_chan <- "hello"
    select {
    case value := <-int_chan:
        fmt.Println(value)
    case value := <-string_chan:
        panic(value)
    }
}
```

考点：select随机性

解答：select会随机选择一个可用通用做收发操作。 所以代码是有肯触发异常，也有可能不会。 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：

- select 中只要有一个case能return，则立刻执行。
- 当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
- 如果没有一个case能return则可以执行”default”块。

### 6. 下面代码输出什么？

```go
func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)
    return ret
}
 
func main() {
    a := 1
    b := 2
    defer calc("1", a, calc("10", a, b))
    a = 0
    defer calc("2", a, calc("20", a, b))
    b = 1
}
```

考点：defer执行顺序

解答：这道题类似第1题 需要注意到defer执行顺序和值传递 index:1肯定是最后执行的，但是index:1的第三个参数是一个函数，所以最先被调用calc("10",1,2)==>10,1,2,3 执行index:2时,与之前一样，需要先调用calc("20",0,2)==>20,0,2,2 执行到b=1时候开始调用，index:2==>calc("2",0,2)==>2,0,2,2 最后执行index:1==>calc("1",1,3)==>1,1,3,4

```go
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
```

### 7. 请写出以下输入内容

```go
func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}
```

考点：make默认值和append

解答：make初始化是由默认值的哦，此处默认值为0

```go
[0 0 0 0 0 1 2 3]
```

大家试试改为:

```
s := make([]int, 0)
s = append(s, 1, 2, 3)
fmt.Println(s)//[1 2 3]
```

### 8. 下面的代码有什么问题?

```go
type UserAges struct {
	ages map[string]int
	sync.Mutex
}
 
func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
 
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
```

考点：map线程安全

解答：可能会出现`fatal error: concurrent map read and map write`. 修改一下看看效果

```go
func (ua *UserAges) Get(name string) int {
    ua.Lock()
    defer ua.Unlock()
    if age, ok := ua.ages[name]; ok {
        return age
    }
    return -1
}
```

### 9. 下面的迭代会有什么问题？

```go
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
 
		for elem := range set.s {
			ch <- elem
		}
 
		close(ch)
		set.RUnlock()
 
	}()
	return ch
}
```

考点：chan缓存池

解答：看到这道题，我也在猜想出题者的意图在哪里。 chan?sync.RWMutex?go?chan缓存池?迭代? 所以只能再读一次题目，就从迭代入手看看。 既然是迭代就会要求set.s全部可以遍历一次。但是chan是为缓存的，那就代表这写入一次就会阻塞。 我们把代码恢复为可以运行的方式，看看效果

```go
package main
 
import (
    "sync"
    "fmt"
)
 
//下面的迭代会有什么问题？
 
type threadSafeSet struct {
    sync.RWMutex
    s []interface{}
}
 
func (set *threadSafeSet) Iter() <-chan interface{} {
    // ch := make(chan interface{}) // 解除注释看看！
    ch := make(chan interface{},len(set.s))
    go func() {
        set.RLock()
 
        for elem,value := range set.s {
            ch <- elem
            println("Iter:",elem,value)
        }
 
        close(ch)
        set.RUnlock()
 
    }()
    return ch
}
 
func main()  {
 
    th:=threadSafeSet{
        s:[]interface{}{"1","2"},
    }
    v:=<-th.Iter()
    fmt.Sprintf("%s%v","ch",v)
}
```

### 10. 以下代码能编译过去吗？为什么？

```go
package main
 
import (
	"fmt"
)
 
type People interface {
	Speak(string) string
}
 
type Stduent struct{}
 
func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
 
func main() {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
```

考点：golang的方法集

解答：不通过， 一句话：golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关。Student并没有继承People，people中有Speak(string)string方法，而Student类型中没有Speak()方法，代码中的Student方法是*Student类型的，所有  var peo People = Student{}是不符合规范的。

### 11. 以下代码打印出来什么内容，说出为什么。

```go
package main
 
import (
	"fmt"
)
 
type People interface {
	Show()
}
 
type Student struct{}
 
func (stu *Student) Show() {
 
}
 
func live() People {
	var stu *Student
	return stu
}
 
func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
```

考点：interface内部结构

解答：很经典的题！ 这个考点是很多人忽略的interface内部结构。 go中的接口分为两种一种是空的接口类似这样：

```
var in interface{}
```

另一种如题目：

```
type People interface {
    Show()
}
```

他们的底层结构如下：

```go
type eface struct {      //空接口
    _type *_type         //类型信息
    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type iface struct {      //带有方法的接口
    tab  *itab           //存储type信息还有结构实现方法的集合
    data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type _type struct {
    size       uintptr  //类型大小
    ptrdata    uintptr  //前缀持有所有指针的内存大小
    hash       uint32   //数据hash值
    tflag      tflag
    align      uint8    //对齐
    fieldalign uint8    //嵌入结构体时的对齐
    kind       uint8    //kind 有些枚举值kind等于0是无效的
    alg        *typeAlg //函数指针数组，类型实现的所有方法
    gcdata    *byte
    str       nameOff
    ptrToThis typeOff
}
type itab struct {
    inter  *interfacetype  //接口类型
    _type  *_type          //结构类型
    link   *itab
    bad    int32
    inhash int32
    fun    [1]uintptr      //可变大小 方法集合
}
```

可以看出iface比eface 中间多了一层itab结构。 itab 存储_type信息和[]fun方法集，从上面的结构我们就可得出，因为data指向了nil 并不代表interface 是nil， 所以返回值并不为空，这里的fun(方法集)定义了接口的接收规则，在编译的过程中需要验证是否实现接口 结果：

```go
BBBBBBB
```

### 12.是否可以编译通过？如果通过，输出什么？

```go
func main() {
	i := GetValue()
 
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
 
}
 
func GetValue() int {
	return 1
}
```

考点：type

解答：编译失败，因为type只能使用在interface

### 13.下面函数有什么问题？

```go
func funcMui(x,y int)(sum int,error){
    return x+y,nil
}
```

考点：函数返回值命名

解答：在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。 如果返回值有有多个返回值必须加上括号； 如果只有一个返回值并且有命名也需要加上括号； 此处函数第一个返回值有sum名称，第二个未命名，所以错误。

### 14.是否可以编译通过？如果通过，输出什么？

```go
package main
 
func main() {
 
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}
 
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}
 
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
 
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
```

考点:defer和函数返回值

解答：需要明确一点是defer需要在函数结束前执行。 函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数 DeferFunc1有函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4; DeferFunc2函数中t的作用域为函数，返回1; DeferFunc3返回3

### 15.是否可以编译通过？如果通过，输出什么？

```go
func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
```

考点：new

解答：list:=make([]int,0)

### 16.是否可以编译通过？如果通过，输出什么？

```go
package main
 
import "fmt"
 
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
```

考点：append

解答：当append切片的时候别漏了'...'，append(s1, s2...)

### 17.是否可以编译通过？如果通过，输出什么？

```go
func main() {
 
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
 
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
 
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
 
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
```

考点:结构体比较

解答：进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。

```
sn3:= struct {
    name string
    age  int
}{age:11,name:"qq"}
```

sn3与sn1就不是相同的结构体了，不能比较。 还有一点需要注意的是结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。

可以使用reflect.DeepEqual进行比较

```go
if reflect.DeepEqual(sn1, sm) {
    fmt.Println("sn1 ==sm")
}else {
    fmt.Println("sn1 !=sm")
}
```

所以编译不通过： invalid operation: sm1 == sm2

### 18.是否可以编译通过？如果通过，输出什么？

```go
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	Foo(x)
}
```

考点：interface内部结构

```
non-empty interface
```

### 19.是否可以编译通过？如果通过，输出什么？

```go
func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return nil, false
}
func main()  {
	intmap:=map[int]string{
		1:"a",
		2:"bb",
		3:"ccc",
	}
 
	v,err:=GetValue(intmap,3)
	fmt.Println(v,err)
}
```

考点：函数返回值类型

解答：nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。报:`cannot use nil as type string in return argument`.

### 20.是否可以编译通过？如果通过，输出什么？

```go
const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)
 
func main()  {
	fmt.Println(x,y,z,k,p)
}
```

考点：iota

结果:

```
0 1 zz zz 4
```

### 21.编译执行下面代码会出现什么?

```go
package main
var(
    size :=1024
    max_size = size*2
)
func main()  {
    println(size,max_size)
}
```

考点:变量简短模式

解答：变量简短模式限制：

- 定义变量同时显式初始化
- 不能提供数据类型
- 只能在函数内部使用

结果：

```
syntax error: unexpected :=
```

### 22.下面函数有什么问题？

```go
package main
const cl  = 100
 
var bl    = 123
 
func main()  {
    println(&bl,bl)
    println(&cl,cl)
}
```

考点:常量

解答：常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，

```
cannot take the address of cl
```

### 23.编译执行下面代码会出现什么?

```go
package main
 
func main()  {
 
    for i:=0;i<10 ;i++  {
    loop:
        println(i)
    }
    goto loop
}
```

考点：goto

解答：goto不能跳转到其他函数或者内层代码

```
goto loop jumps into block starting at
```

### 24.编译执行下面代码会出现什么?

```go
package main
import "fmt"
 
func main()  {
    type MyInt1 int
    type MyInt2 = int
    var i int =9
    var i1 MyInt1 = i
    var i2 MyInt2 = i
    fmt.Println(i1,i2)
}
```

考点：**Go 1.9 新特性 Type Alias **

解答：基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias。 MyInt1为称之为defintion，虽然底层类型为int类型，但是不能直接赋值，需要强转； MyInt2称之为alias，可以直接赋值。

结果:

```
cannot use i (type int) as type MyInt1 in assignment
```

### 25.编译执行下面代码会出现什么?

```go
package main
import "fmt"
 
type User struct {
}
type MyUser1 User
type MyUser2 = User
func (i MyUser1) m1(){
    fmt.Println("MyUser1.m1")
}
func (i User) m2(){
    fmt.Println("User.m2")
}
 
func main() {
    var i1 MyUser1
    var i2 MyUser2
    i1.m1()
    i2.m2()
}
```

考点：**Go 1.9 新特性 Type Alias **

解答：因为MyUser2完全等价于User，所以具有其所有的方法，并且其中一个新增了方法，另外一个也会有。 但是

```
i1.m2()
```

是不能执行的，因为MyUser1没有定义该方法。 结果:

```
MyUser1.m1
User.m2
```

### 26.编译执行下面代码会出现什么?

```go
package main
 
import "fmt"
 
type T1 struct {
}
func (t T1) m1(){
    fmt.Println("T1.m1")
}
type T2 = T1
type MyStruct struct {
    T1
    T2
}
func main() {
    my:=MyStruct{}
    my.m1()
}
```

考点：**Go 1.9 新特性 Type Alias **

解答：是不能正常编译的,异常：

```
ambiguous selector my.m1
```

结果不限于方法，字段也也一样；也不限于type alias，type defintion也是一样的，只要有重复的方法、字段，就会有这种提示，因为不知道该选择哪个。 改为:

```
my.T1.m1()
my.T2.m1()
```

type alias的定义，本质上是一样的类型，只是起了一个别名，源类型怎么用，别名类型也怎么用，保留源类型的所有方法、字段等。

### 27.编译执行下面代码会出现什么?

```go
package main
 
import (
    "errors"
    "fmt"
)
 
var ErrDidNotWork = errors.New("did not work")
 
func DoTheThing(reallyDoIt bool) (err error) {
    if reallyDoIt {
        result, err := tryTheThing()
        if err != nil || result != "it worked" {
            err = ErrDidNotWork
        }
    }
    return err
}
 
func tryTheThing() (string,error)  {
    return "",ErrDidNotWork
}
 
func main() {
    fmt.Println(DoTheThing(true))
    fmt.Println(DoTheThing(false))
}
```

考点：变量作用域

解答：因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量，结果：

```
<nil>
<nil>
```

改为：

```go
func DoTheThing(reallyDoIt bool) (err error) {
    var result string
    if reallyDoIt {
        result, err = tryTheThing()
        if err != nil || result != "it worked" {
            err = ErrDidNotWork
        }
    }
    return err
}
```

### 28.编译执行下面代码会出现什么?

```go
package main
 
func test() []func()  {
    var funs []func()
    for i:=0;i<2 ;i++  {
        funs = append(funs, func() {
            println(&i,i)
        })
    }
    return funs
}
 
func main(){
    funs:=test()
    for _,f:=range funs{
        f()
    }
}
```

考点：闭包延迟求值

解答：for循环复用局部变量i，每一次放入匿名函数的应用都是想一个变量。 结果：

```
0xc042046000 2
0xc042046000 2
```

如果想不一样可以改为：

```go
func test() []func()  {
    var funs []func()
    for i:=0;i<2 ;i++  {
        x:=i
        funs = append(funs, func() {
            println(&x,x)
        })
    }
    return funs
}
```

### 29.编译执行下面代码会出现什么?

```go
package main
 
func test(x int) (func(),func())  {
    return func() {
        println(x)
        x+=10
    }, func() {
        println(x)
    }
}
 
func main()  {
    a,b:=test(100)
    a()
    b()
}
```

考点：闭包引用相同变量*

解答：结果：

```
100
110
```

### 30.编译执行下面代码会出现什么?

```go
package main
 
import (
    "fmt"
    "reflect"
)
 
func main1()  {
    defer func() {
       if err:=recover();err!=nil{
           fmt.Println(err)
       }else {
           fmt.Println("fatal")
       }
    }()
 
    defer func() {
        panic("defer panic")
    }()
    panic("panic")
}
 
func main()  {
    defer func() {
        if err:=recover();err!=nil{
            fmt.Println("++++")
            f:=err.(func()string)
            fmt.Println(err,f(),reflect.TypeOf(err).Kind().String())
        }else {
            fmt.Println("fatal")
        }
    }()
 
    defer func() {
        panic(func() string {
            return  "defer panic"
        })
    }()
    panic("panic")
}
```

考点：panic仅有最后一个可以被revover捕获

解答：触发`panic("panic")`后顺序执行defer，但是defer中还有一个panic，所以覆盖了之前的`panic("panic")`

```
defer panic
```



## 持续更新...

部分引用来源：<https://blog.csdn.net/weiyuefei/article/details/77963810>
