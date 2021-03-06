package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)
// 使用html/template中的函数对客户端输入的信息进行转义过滤，使得javaScript等脚本注入执行
// func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
// func HTMLEscapeString(s string) string //转义s之后返回结果字符串
// func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		if len(r.Form["username"][0]) == 0 {
			t, _ := template.ParseFiles("login.gtpl")
			log.Println(t.Execute(w, "用户名为空"))
		}
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println(r.Form["fruit"])
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		u := strings.Join(r.Form["username"], " ")
		template.HTMLEscape(w, []byte(u))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
