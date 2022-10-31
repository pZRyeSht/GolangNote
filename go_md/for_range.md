#for range 问题

**遍历取不到所有元素指针**
```go
func main() {
	arr := []int{1, 2, 3}
	var res []*int
	for _, v := range arr {
		res = append(res, &v)
	}
    // expect: 1 2 3 
	fmt.Println(*res[0], *res[1], *res[2]) 
	// but output: 3 3 3
}
```
以上遍历对于 array 或者 map 得到的结果同样时达不到预期的遍历。
go for-range是语法糖，其内部的调用还是for循环，初始化会拷贝带遍历的列表（array、slice、map）然后每次遍历的 v 都是对同一个元素的遍历赋值。 
也就是说如果直接对 v 取地址，最终只会拿到一个地址，而对应的值就是最后遍历的那个元素所附给v的值。 go for-range是值拷贝，且只会声明初始化一次.
* 使用局部变量拷贝
```go
for _, v := range arr {
// 局部变量v替换了v
v := v
res = append(res, &v)
}
```
* 直接索引获取原来的元素
```go
for k := range arr {
    res = append(res, &arr[k])
}
```

**map 遍历中的 goroutine**
```go
var m = []int{1, 2, 3}
	for i := range m {
		go func() {
			fmt.Print(i)
		}()
	}
	// block main 1ms to wait goroutine finished
	time.Sleep(time.Millisecond)
```
同 slice 遍历一致，都是拷贝问题
* 以参数方式传入
```go
for i := range m {
    go func(i int) {
        fmt.Print(i)
    }(i)
}
```
* 使用局部变量拷贝
```go
for i := range m {
    i := i
    go func() {
        fmt.Print(i)
    }()
}
```

**在循环体内修改遍历的对象**
```go
func main() {
    arr := []int{1,2,3}
    for _, v := range arr {
        arr = append(arr, v)
    }
    fmt.Println(arr)
}
// 输出 [1 2 3 1 2 3]
```
之所以只输出两遍1 2 3而不是一直循环下去，是因为for range在编译期间，就会把arr赋值给一个新的变量，对原来arr的修改不会反映到遍历中。


