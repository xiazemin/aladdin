package main
import (
	"os"
	"text/template"
)
type x struct {
	A姓名, B级别, C性别 string
}
const M = `{{range .}}{{if .B级别}}姓名：{{.A姓名}}  性别：{{.C性别}}{{end}}
{{end}}`
func main() {
	var di = []x{{"曦晨", "1", "男"}, {"晨曦", "2", "女"}, {"曦love晨", "", "男love女"}}
	t := template.New("")
	template.Must(t.Parse(M))
	t.Execute(os.Stdout, di)
}