package curl

import (
         goCurl "github.com/mikemintang/go-curl"
	"github.com/xiazemin/aladdin/damon/log"
	"github.com/xiazemin/aladdin/damon/logFile"
)

func QueryJson(dir string,reqParam log.Request) string{
	url := reqParam.Url
	headers := map[string]string{
		"User-Agent":    "Sublime",
		"Authorization": "Bearer access_token",
		"Content-Type":  "application/json",
	}
	cookies := map[string]string{
		"userId":    "12",
		"loginTime": "15045682199",
	}
	queries := map[string]string{
		"page": "2",
		"act":  "update",
	}
	postData := map[string]interface{}{
		"name":      "mike",
		"age":       24,
		"interests": []string{"basketball", "reading", "coding"},
		"isAdmin":   true,
	}

	for k,v:=range (reqParam.Arguments){
		postData[k]=v
	}
	// 链式操作
	req := goCurl.NewRequest()
	var(
		resp *goCurl.Response
		err error
	)
	if reqParam.Method=="POST" {
		resp, err= req.
		SetUrl(url).
			SetHeaders(headers).
			SetCookies(cookies).
			SetQueries(queries).
			SetPostData(postData).
			Post()
	}else if reqParam.Method=="GET" {
		resp, err = req.
		SetUrl(url).
			SetHeaders(headers).
			SetCookies(cookies).
			SetPostData(postData).
			Get()
	}else{
		logFile.LogWarnf(dir,reqParam)
	}

	logFile.LogDebug(dir,"\n-------------------")
	logFile.LogNotice(dir,req)
	if err != nil {
		logFile.LogWarnf(dir,err)
		return ""
	} else {
		if resp.IsOk() {
			logFile.LogNotice(dir,resp.Body)
			return resp.Body
		} else {
			logFile.LogWarnf(dir,resp.Raw)
			return ""
		}
	}
	return ""
}


func QueryForm(dir string,reqParam log.Request) string{
	url := reqParam.Url
	headers := map[string]string{
		"User-Agent":    "Sublime",
		"Authorization": "Bearer access_token",
		"Content-Type":  "Content-Type: application/x-www-form-urlencoded",
	}
	cookies := map[string]string{
		"userId":    "12",
		"loginTime": "15045682199",
	}
	queries := map[string]string{
		"page": "2",
		"act":  "update",
	}

	for k,v:=range (reqParam.Arguments){
		if val,ok:=v.(string) ;ok{
			queries[k]=val
		}
	}

	postData := map[string]interface{}{
		"name":      "mike",
		"age":       24,
		"interests": []string{"basketball", "reading", "coding"},
		"isAdmin":   true,
	}

	for k,v:=range (reqParam.Arguments){
		postData[k]=v
	}

	// 链式操作
	req := goCurl.NewRequest()
	var(
		resp *goCurl.Response
		err error
	)
	if reqParam.Method=="POST" {
		resp, err= req.
		SetUrl(url).
			SetHeaders(headers).
			SetCookies(cookies).
			SetQueries(queries).
			SetPostData(postData).
			Post()
	}else if reqParam.Method=="GET" {
		resp, err = req.
		SetUrl(url).
			SetHeaders(headers).
			SetCookies(cookies).
			SetQueries(queries).
			Get()
	}else{
		logFile.LogWarnf(dir,reqParam)
	}

	logFile.LogDebug(dir,"\n-------------------")
	logFile.LogNotice(dir,req)
	if err != nil {
		logFile.LogWarnf(dir,err)
		return ""
	} else {
		if resp.IsOk() {
			logFile.LogNotice(dir,resp.Body)
			return resp.Body
		} else {
			logFile.LogWarnf(dir,resp.Raw)
			return ""
		}
	}
	return ""
}