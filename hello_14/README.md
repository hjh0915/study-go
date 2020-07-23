批量获取天气
==========
利用map
```go
var cmap map[string]string 
cmap = map[string]string {
    "天津": "101030100",
    "南昌": "101240101",
    "北京": "101010100",
    "上海": "101020100",
    "香港": "101320101",
    "澳门": "101330101",
    "杭州": "101210101",
    "苏州": "101190401",
    "南京": "101190101"}
for _,value := range cmap {
    getWeather(value)
}
```
先创建一个map，添加map的key,value，遍历获取value值