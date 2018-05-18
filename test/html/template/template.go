package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"time"
	"path/filepath"
	"fmt"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Print(dir)
	fmt.Println("\n------------")
	fmt.Print(err)
	templDir:="/Users/didi/goLang/src/github.com/xiazemin/aladdin/test/html/template/"
	t := template.New("第一个模板").Delims("[[", "]]") //创建一个模板,设置模板边界
	t, _ = t.Parse("hello,[[.UserName]]\n")       //解析模板文件
	data := map[string]interface{}{"UserName": template.HTML("<script>alert('you have been pwned')</script>")}
	t.Execute(os.Stdout, data) //执行模板的merger操作，并输出到控制台
fmt.Println("\nt1")

	t2 := template.New("新的模板")                         //创建模板
	t2.Funcs(map[string]interface{}{"tihuan": tihuan}) //向模板中注入函数
	bytes, _ := ioutil.ReadFile(templDir+"test2.html")          //读文件
	fmt.Print(string(bytes))
	template.Must(t2.Parse(string(bytes)))             //将字符串读作模板
	t2.Execute(os.Stdout, map[string]interface{}{"UserName": "你好世界"})
	fmt.Println("\n", t2.Name(), "\n")
	fmt.Println("\nt2")

	t3, _ := template.ParseFiles(templDir+"test1.html") //将一个文件读作模板
	t3.Execute(os.Stdout, data)
	fmt.Println(t3.Name(), "\n") //模板名称
	fmt.Println("\nt3")

	t4, _ := template.ParseGlob(templDir+"test1.html") //将一个文件读作模板
	t4.Execute(os.Stdout, data)
	fmt.Println(t4.Name())
	fmt.Println("\nt4")
}

//注入模板的函数
func tihuan(str string) string {
	return str + "-------" + time.Now().Format("2006-01-02")
}