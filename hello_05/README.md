自带 测试框架
===========
Golang单元测试对文件名和方法名，参数都有很严格的要求。  
1、文件名必须以xx_test.go命名  
2、方法必须是Test[^a-z]开头  
3、方法参数必须 t *testing.T  

执行
===
cd tools
> $ go test