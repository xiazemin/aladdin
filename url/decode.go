package main
import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"strings"
	"fmt"
)

func main()  {
	encodedUrl:=`www.baidu.com?a=1&b=2&data=%7B%22sessionId%22%3A%220bts0W1DWKm70B4UZq3V1h3r2DpsbhDsc2WD%22%2C%0A%09%22eventId%22%3A%228F2qNf0bts0W1DWKm70B4UZq3V1h3r2Dpsbh%22%2C%0A%09%22androidId%22%3A%22b22f3d41736f748c%22%2C%0A%09%22userAgent%22%3A%22Mozilla%2F5.0+%28Linux%3B+Android+5.1%3B+MI+PAD+2+Build%2FLMY47I%3B+wv%29%0A%09AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0%0A%09Chrome%2F55.0.2883.91+Safari%2F537.36%22%2C%22osVersion%22%3A%225.1.1%22%2C%0A%09%22bundleId%22%3A%22com.brianbaek.popstar%22%2C%22connectionType%22%3A%22wifi%22%2C%0A%09%22deviceMake%22%3A%22Xiaomi%22%2C%22deviceModel%22%3A%22MI+PAD+2%22%2C%0A%09%22language%22%3A%22zh_CN%22%2C%22timeZone%22%3A%22GMT%2B08%3A00%22%2C%22campaignId%22%3A3261%2C%0A%09%22mac%22%3A%2238%3Aa4%3Aed%3Afe%3A99%3Ac8%22%7D`
	l3, err3 := url.Parse(encodedUrl)
	fmt.Println(l3, err3)

}

func export(prefix, key, value string) {
	fmt.Printf("export %s_%s=%s\n", prefix, key, value)
}

//Get is a common http client for httpGet operations and hide UA as GoogleBot.
func Get(urls string) ([]byte, error) {

	req, err := http.NewRequest("GET", urls, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Googlebot/2.1 (+http://www.google.com/bot.html)")

	client := getClient()
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	tempData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err

	}

	return tempData, nil
}

func getClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
		DisableKeepAlives:  true,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}
	return client
}

func RequestWithHeader(
method string,
url1 string,
timeout int,
header map[string]string,
params map[string]string,
) ([]byte, error) {

	client := getClient()

	paramsA := ""
	if method == "POST" {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v)
		}
		paramsA = values.Encode()
	}

	req, _ := http.NewRequest(method, url1, strings.NewReader(paramsA))

	for k, v := range header {
		req.Header.Set(k, v)
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if method == "GET" {
		values := req.URL.Query()
		for k, v := range params {
			values.Add(k, v)
		}
		req.URL.RawQuery = values.Encode()
	}

	getResp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer getResp.Body.Close()
	body, err := ioutil.ReadAll(getResp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
