第三方包集成
==========
go第三方包下载代理

https://goproxy.io/zh/

Go 版本是 1.13 及以上

```sh
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go build
```

go.mod 文件
===========

    module hello_05

    go 1.14

    require github.com/stretchr/testify v1.5.1

*新增了go.sum 文件*

执行
===
cd tools
> $ go test