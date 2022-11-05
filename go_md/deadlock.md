# deadlock 问题

使用 channel 进行并发子任务执行时，使用不当容易造成 deadlock

### 1.读取空channel产生死锁
```go
func main() {
	ch := make(chan int, 1)
	<-ch
}
```

采用 select 进行阻塞时默认值处理
```go
func main() {
	ch := make(chan int, 1)
	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println("the channel no data")
	}
}
```

### 2.channel阻塞产生死锁
```go
func main() {
	ch := make(chan int)
	// 无缓冲写入数据，没有读数据处理，channel阻塞
	ch <- 1 
	// channel阻塞，无法执行
	fmt.Println(<-ch) 
}
```

1.采用子协程方式，保证channel读写成对数据存取
```go
func main() {
	ch := make(chan string)
	// 开启子goroutine写入数据
	go func() {
		ch <- "hello"
	}()
	// 初始阻塞，goroutine执行时写入数据，则读取成功
	fmt.Println(<-ch)
}
```

2.采用有缓冲channel，在容量范围内不会阻塞
```go
func main() {
	ch := make(chan string, 1)
	ch <- "hello"
	fmt.Println(<-ch)
}
```

### for range channel 产生死锁
```go
func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// for range 一直读取直到channel关闭，channel无close则产生阻塞死锁
	for v := range ch {
		fmt.Println(v)
	}
}
```
显式关闭channel，开启子协程，主协程 sleep 等待时间后退出
```go
func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	time.Sleep(1000)
}
```





