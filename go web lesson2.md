模板语法

[李文周Gin-framework](https://www.liwenzhou.com/posts/Go/Gin_framework/)

要使用模板得做三件事：

传统方式使用模板：

1. 定义模板
2. 解析模板
3. 渲染模板

```go
func handle (w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	//渲染模板
	//template.ParseFiles默认查询路径是项目路径因此template_lecture不可缺少
	t,err := template.ParseFiles("template_lecture/hello.html")
	if err != nil{
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	steve :=map[string]interface{}{
		"name":"steve",
		"age":24,
		"gender":"男",
	}

	jack := map[string]interface{}{
		"name":"jack",
		"age":21,
		"gender":"男",
	}
	data := map[string]interface{}{
		"steve":steve,
		"jack":jack,
	}
  //传递复杂数据
	t.Execute(w,map[string]interface{}{
		"people":data,
	})

}

func main(){
	http.HandleFunc("/",handle)
	err := http.ListenAndServe(":9002",nil)
	if err != nil{
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
}
```

HTML的写法：

```
 <ul>{{.people.steve.name}}
   <li>年龄: {{.people.steve.age}}</li>
   <li>性别: {{people.steve.gender}}</li>
 </ul>
```

模板中可以写判断语句

```
<body>
    <ul>{{.people.steve.name}}
        <li>年龄: {{.people.steve.age}}</li>
        <li>性别: {{.people.steve.gender}}</li>
    </ul>
    <hr>
    {{$jack := .people.jack}}
    {{if $jack}}
    <p>你好杰克</p>
    {{else}}
    <p>杰克不在</p>
    {{/*下面的end是必须的*/}}
    {{end}}
</body>
```

### 在模板引擎中自定义函数





















