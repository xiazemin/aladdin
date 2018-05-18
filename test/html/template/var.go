package main

import (
	"os"
	"text/template"
)

type x struct {
	A姓名, B级别, C性别 string
}
const M = `{{range $k,$v := .}}信息：{{$v.A姓名}}
{{end}}`
func main() {
	var di = []x{{"曦晨", "1", "男"}, {"晨曦", "2", "女"}}
	t := template.New("")
	t.Parse(M)
	t.Execute(os.Stdout, di)
}