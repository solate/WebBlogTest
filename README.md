
# WebBlogTest
这是一个博客系统的测试程序，主要是为了学习

使用beego 来进行开发和学习， orm 用 gorm , 数据库为 mysql , 前端 bootstrap 


## 一些笔记

一般我们写http 使用的是 http.ListenAndServe(":8080", nil), handler 这里写的是nil

更底层， handler 使用的是 mux := http.NewServeMux(), mux.Handle() 这个需要实现一个 ServeHTTP()






## 原生http


### http.ListenAndServe

handler 参数(w http.ResponseWriter, r *http.Request)

go参数传递为值传递，request长用来获取参数等，所以直接传递指针比较好，而 ResponseWriter 是个接口，只要实现接口就行 无所谓传不传指针的问题。


```
package main

import (
	"net/http"
	"log"
	"io"
)

func main() {

	http.Handle("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello world, version 1")
}



//访问 http://localhost:8080/

```


### http.ServeMux 更底层路由设置

[http.ServeMux解析](http://studygolang.com/articles/4890)

ServeMux 其实是路由表，主要使用map结构，其实例必须实现 **ServeHTTP()** 方法

`mux.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern} `

```

package main

import (
	"net/http"
	"io"
	"log"
)

func main() {

	mux := http.NewServeMux() //路由表结构
	mux.Handle("/", &MyHandler{}) //这里注册的是处理的指针,默认根路径"/"
	mux.HandleFunc("/hello", sayHello) //注册/hello 

	err := http.ListenAndServe(":8080", mux) // 将mux 放入
	if err != nil {
		log.Fatal(err)
	}
}


type MyHandler struct {
	
}

func (_ * MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "URL: " + r.URL.String() )
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello world, version 2")
}


//访问 http://localhost:8080/
//访问 http://localhost:8080/hello


```



### 更进一步实现http.Server

自己实现Server 最重要的是需要自己在 ServeHTTP() 中实现路由转发

```

package main

import (
	"net/http"
	"io"
	"log"
	"time"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {

	server := http.Server{
		Addr : ":8080", //设置地址
		Handler : &MyHandler{}, //设置处理handler
		ReadTimeout : 5 * time.Second, //设置超时时间 5S
	}

	//因为没有提供方法，所以需要自己实现路由，然后在ServeHTTP中进行路由转发
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye

	err := server.ListenAndServe() // 使用自己实例化的server
	if err != nil {
		log.Fatal(err)
	}
}

type MyHandler struct {

}

func (_ *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//因为这里没有方法，所以需要进行路由转发
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r) //如果存在就转发
		return
	}
	io.WriteString(w, "URL: " + r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world, version 3")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bye bye, version 3")
}


//访问 http://localhost:8080/
//访问 http://localhost:8080/hello
//访问 http://localhost:8080/bye

```



### 静态文件

静态文件就需要使用到http.FileServer

[使用Golang 搭建http web服务器](http://www.cnblogs.com/yjf512/archive/2012/09/03/2668384.html)


```

package main

import (
	"net/http"
	"io"
	"log"
	"os"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", &MyHandler{})
	mux.HandleFunc("/hello", sayHello)

	wd, err := os.Getwd()// 获取当前路径
	if err != nil {
		log.Fatal(err)
	}

	//http.Dir(wd) //获取相对路径
	//http.FileServer 静态处理
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd)))) //设置静态文件路径

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}


type MyHandler struct {

}

func (_ * MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "URL: " + r.URL.String() )
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello world, version 4")
}



// http://localhost:8080/static/

```
