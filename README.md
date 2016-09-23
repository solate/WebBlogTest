
# WebBlogTest
这是一个博客系统的测试程序，主要是为了学习

使用beego 来进行开发和学习， orm 用 gorm , 数据库为 mysql , 前端 bootstrap 


## 一些笔记

一般我们写http 使用的是 http.ListenAndServe(":8080", nil), handler 这里写的是nil

更底层， handler 使用的是 mux := http.NewServeMux(), mux.Handle() 这个需要实现一个 ServeHTTP()






## beego 使用


### bee 工具

一般我们运行go程序使用的 `go run main.go` 这样形式的

但是使用 bee 工具的时候需要使用 `bee run appname`, appname 是应用名称，在conf 文件中配置

我在刚开始使用的时候输入 `bee run main.go` 是错误的



1. bee new WebBlogTest
2. bee run WebBlogTest  



## 模板

普通使用和go模板类似


### with

在输出结构的时候可以使用with 来进行分组，将前面相同的部分进行省略。


### beego内置模板函数

[模板函数](http://beego.me/docs/mvc/view/template.md)























































