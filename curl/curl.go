package main
import (
	"fmt"
	"github.com/mikemintang/go-curl"
)
func main() {
	url := "www.baidu.com"
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
	// 链式操作
	req := curl.NewRequest()
	resp, err := req.
	SetUrl(url).
		SetHeaders(headers).
		SetCookies(cookies).
		SetQueries(queries).
		SetPostData(postData).
		Post()
	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
		}
	}
}