解析json文件
===========
利用结构体struct解析
------------------
```go
type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}
```
Name 是 json文件中的 name 映射而来

需要引用的包
----------
```go
import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)
```