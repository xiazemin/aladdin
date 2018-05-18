package api
import (
	"net/http"
	"fmt"
	"github.com/xiazemin/aladdin/damon/file"
	"strings"
	"os"
	"io"
	"github.com/xiazemin/aladdin/damon/logFile"
	"text/template"
	"io/ioutil"
	 "github.com/xiazemin/aladdin/damon/netenv"
	"github.com/xiazemin/aladdin/damon/config"
	"time"
	"go/src/encoding/json"
	"strconv"
)

type Data struct {

}
func (this*Data)Handle(uris []string,w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configData string)string  {
	var resp string
	switch uris[3] {
	case "get":
		resp=this.Get(w,r,defaultDir,logDir,viewDir)
	case "update":
		resp=this.Update(w,r,defaultDir,logDir,viewDir)
	case "add":
		resp=this.Add(w,r,defaultDir,logDir,viewDir,configData)
	case "new":
		resp=this.New(w,r,defaultDir,logDir,viewDir,configData)
	default:
		resp=r.RequestURI+fmt.Sprintf("   %d  %+v  %s  %s",len(uris),uris,uris[2],uris[3])

	}
	return  resp
}

func (this *Data)Get(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string)string{
	dir:=file.GetDir(defaultDir+"data/","data/",logDir)
	str:=file.GetPrintDirs(dir,0)
	str=strings.Replace(strings.Replace(str,"\n","<br/>",-1),
		" ","&nbsp;",-1)

	templ, _ := ioutil.ReadFile(viewDir+"data/"+"get.html")
	t := template.New("get log file ")
	t.Parse(string(templ))
	t.Execute(w, str)
	return viewDir+"\n"+str
}

func (this *Data)Update(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string)string{
	return defaultDir
}

func (this *Data)parseDescription(r *http.Request,defaultDir string,logDir string) config.Description{
	model:=string(r.PostFormValue("model"))
	selected,err:=strconv.ParseBool(r.PostFormValue("selected"))
	if err!=nil{
		logFile.LogWarnf(logDir,err)
	}
	user:=string(r.PostFormValue("user"))
	description:=string(r.PostFormValue("description"))
	var des config.Description
	des.Date = time.Now().Format("2006010215")
	des.Description = description
	des.Model = model
	des.Selected = selected
	des.User = user
	return des
}

func (this *Data)updateConfig(r *http.Request,name string,defaultDir string,logDir string,configData string)  {
	if name=="" {
		return
	}
	//更新配置
	configList := config.LoadLogDataDes(defaultDir, configData) //configData里包含"config/"
	des:=this.parseDescription(r,defaultDir,logDir)
	des.LogName = "data/" + name
	configList = append(configList, des)
	str, err := json.Marshal(configList)
	if err == nil {
		file.Write(defaultDir, configData, string(str))
	} else {
		logFile.LogWarnf(logDir, err)
	}
	return
}

func (this *Data)New(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configData string)string{
	name:=string(r.PostFormValue("file_name"))
	content:=string(r.PostFormValue("file_content"))

	logFile.LogNotice(logDir,name)
	logFile.LogNotice(logDir,content)
	//fmt.Fprintln(w,name+content)
	//保存文件
	file.Write(defaultDir+"data/",name,content)

	if  content!="" {
              this.updateConfig(r,name,defaultDir,logDir,configData)
	}else {
		logFile.LogWarnf(logDir," file name or content is empty")
	}
	http.Redirect(w, r, "/file/data/get/", http.StatusFound)
	return "{name:"+name+",content:"+content+"}"
}

func (this *Data)Add(w http.ResponseWriter,r *http.Request,defaultDir string,logDir string,viewDir string,configData string)string{
	//从请求当中判断方法
	if r.Method == "GET" {
		ip:=netenv.GetLocalIp(logDir)
		url:="http://"+ip+":8088/file/data/new/"
		urlCpmmon:="用户名:<input type=\"text\" name=\"user\"/><br/>" +
			"描述:<input type=\"text\" name=\"description\"/><br/>" +
			"模块名:<input type=\"text\" name=\"model\"/><br/>" +
			"默认使用:<input type=\"radio\"  checked=true name=\"selected\"/><br/>"

		html:="<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传日志</label>"+":"+
			"<input type=\"file\" name='file'  /><br/><br/>    "+
			urlCpmmon+
			"<label><input type=\"submit\" value=\"上传日志\"/></label></form>"+
		        //上面是文件上传的form
                          "<hr/><form action='"+url+"' method=\"post\" enctype=\"multipart/form-data\">" +
			"文件名:<input type=\"text\" name=\"file_name\"/><br/>" +
			"日志内容:<textarea name=\"file_content\"></textarea><br/>" +
			urlCpmmon+
			"<input type=\"submit\" value=\"保存日志文件\"/>"+
		         //上么是文本上传的form
			"</body></html>"
		io.WriteString(w,html )
		return  html
	} else {
		//获取文件内容 要这样获取
		file, head, err := r.FormFile("file")
		if err != nil {
			http.Redirect(w, r, "/file/data/get/", http.StatusFound)
			logFile.LogWarnf(logDir,err)
			return  "获取文件类容失败"
		}
		defer file.Close()
		//创建文件
		fW, err := os.Create(defaultDir+"data/" + head.Filename)
		if err != nil {
			http.Redirect(w, r, "/file/data/get/", http.StatusFound)
			logFile.LogWarnf(logDir,err)
			return "文件创建失败"
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			http.Redirect(w, r, "/file/data/get/", http.StatusFound)
			logFile.LogWarnf(logDir,err)
			return "文件保存失败"
		}

		this.updateConfig(r,head.Filename,defaultDir,logDir,configData)
		//io.WriteString(w, head.Filename+" 保存成功")
		http.Redirect(w, r, "/file/data/get/", http.StatusFound)
		//io.WriteString(w, head.Filename)
		return "成功"
	}
}
