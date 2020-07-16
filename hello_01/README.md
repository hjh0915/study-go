下载配置go环境
============
https://golang.google.cn/doc/install  
go1.14.5.linux-amd64.tar.gz

在下载文件夹中
> $ tar -C /usr/local -xzf go1.14.5.linux-amd64.tar.gz

在.bashrc增加配置

```sh
export PATH=$PATH:/usr/local/go/bin
```

执行
===
> $ go build hello.go

> $ ./hello

fmt
===
fmt 包使用函数实现 I/O 格式化（类似于 C 的 printf 和 scanf 的函数）, 格式化参数源自C
fmt.Errorf(format, args)
fmt.Fprint(writer, args)
fmt.Fprintf(writer, format, args)
fmt.Println(args)
fmt.Sprintln(args)

电子书
=====
https://studygolang.com/subject/2  