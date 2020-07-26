并发获取天气
==========
```go
for _,value := range cmap {
    go getWeather(value)
}
time.Sleep(1 * time.Second)
```
把时间延迟放在go func()后面，是对于函数的延迟