并发获取天气(done)
================
```go
done := make(chan bool, 9)
for _, value := range cmap {
    go getWeather(value, done)
}

for v := range done {
    fmt.Println(v)
}
```
尚存问题