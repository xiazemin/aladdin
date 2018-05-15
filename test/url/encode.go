package main
/*
import "fmt"
import "net/http"
import "net/url"
import "io/ioutil"

func main() {
	AA := "hello word"
	req := " https://www.baidu.com?query=" + AA
	u, _ := url.Parse(req)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	println("u:%+v",u)
	println(q.Encode())
	println(u.RawQuery)
	res, _ := http.Get(u.String())

	result, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", result)
}
*/
import (
	"crypto/md5"
	"flag"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func main() {

	u := url.Values{}
	u.Set("a", "1")
	u.Set("b", "2")
	u.Set("data", `{"sessionId":"0bts0W1DWKm70B4UZq3V1h3r2DpsbhDsc2WD",
	"eventId":"8F2qNf0bts0W1DWKm70B4UZq3V1h3r2Dpsbh",
	"androidId":"b22f3d41736f748c",
	"userAgent":"Mozilla/5.0 (Linux; Android 5.1; MI PAD 2 Build/LMY47I; wv)
	AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0
	Chrome/55.0.2883.91 Safari/537.36","osVersion":"5.1.1",
	"bundleId":"com.brianbaek.popstar","connectionType":"wifi",
	"deviceMake":"Xiaomi","deviceModel":"MI PAD 2",
	"language":"zh_CN","timeZone":"GMT+08:00","campaignId":3261,
	"mac":"38:a4:ed:fe:99:c8"}`)
	fmt.Println(u.Encode())

	fmt.Println("./timetoken -t 3600 -key key  -url url")

	var t int64
	var key string
	var resUrl string
	flag.Int64Var(&t, "t", 0, "expire timestamp")
	flag.StringVar(&key, "key", "", "encrypt key")
	flag.StringVar(&resUrl, "url", "", "resource url")
	flag.Parse()

	if t == 0 || key == "" || resUrl == "" {
		return
	}

	expireTime := fmt.Sprintf("%x", time.Now().Unix()+t)

	resUri, pErr := url.Parse(resUrl)
	if pErr != nil {
		return
	}
	fmt.Println(resUri)
	path := resUri.EscapedPath()
	fmt.Println(path)
	rawStr := fmt.Sprintf("%s%s%s", key, path, expireTime)
	fmt.Println(rawStr)
	md5H := md5.New()
	md5H.Write([]byte(rawStr))

	sign := fmt.Sprintf("%x", md5H.Sum(nil))
	//sign := hex.EncodeToString(md5H.Sum(nil))

	var newUrl string
	if strings.Contains(resUrl, "?") {
		newUrl = fmt.Sprintf("%s&sign=%s&t=%s", resUrl, sign, expireTime)
	} else {
		newUrl = fmt.Sprintf("%s?sign=%s&t=%s", resUrl, sign, expireTime)
	}

	fmt.Println(newUrl)



}