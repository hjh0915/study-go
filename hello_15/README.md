并发获取天气
==========
```go
var wg sync.WaitGroup
for i := 0; i < len(cmap); i++ {
    wg.Add(1)
    
    go func()
    wg.Done()
}
wg.Wait()
```