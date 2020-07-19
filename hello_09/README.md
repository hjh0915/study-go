信道 channel
===========
```go
dch := make(chan int)
go digits(number, dch)
```
使用make 声明

dchnl <- digit
信道传值

自动阻塞，不需要自定义设置延迟时间