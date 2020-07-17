模块机制及包机制分离代码
====================
采用不同包的形式

Go Module

> $ go mod init hello_04

生成 go.mod  

    module hello_04  

    go 1.14

mkdir tools

├── hello_04  
│   ├── go.mod  
│   ├── hello.go  
│   └── tools  
│       └── utils.go  

utils.go  
因为属于不同的包，函数名称必须大写开头，才是可见的，改成 Sum 才能访问到  
同一个包内，小写开头的函数，内部访问。