package main

import (
	"fmt"

	"html/template"
	"log"
	"net/http"
)

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
		t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		t.ExecuteTemplate(w, "T", template.HTML(r.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
