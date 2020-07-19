协程并发初步
==========
Go 编程语言原生支持并发。Go 使用 Go 协程（Goroutine） 和信道（Channel）来处理并发。

我们调用了 time 包里的函数 Sleep，该函数会休眠执行它的 Go 协程。在这里，我们使 Go 主协程休眠了 3 秒

```go
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
```

现在，这两个协程并发地运行。numbers 协程首先休眠 250 微秒，接着打印 1，然后再次休眠，打印 2，
依此类推，一直到打印 5 结束。alphabete 协程同样打印从 a 到 e 的字母，并且每次有 400 微秒的休眠时间。 
Go 主协程启动了 numbers 和 alphabete 两个 Go 协程，休眠了 3000 微秒后终止程序。


执行
===
> $ go build hello.go

> $ ./hello