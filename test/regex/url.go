package main

import "regexp"
import "fmt"
func main() {
	//r, _ := regexp.Compile("^/(\\w+(-\\w+)*)(\\.(\\w+(-\\w+)*))*(\\?\\s*)?$")
	//r, _ := regexp.Compile("^/[A-Za-z0-9]+\\.[A-Za-z0-9]+[=\\?%\\-&_~`@\\[\\]\\':+!]*([^<>\"])*$")
	r, _ := regexp.Compile("^(/[A-Za-z0-9]+)+(\\?[A-Za-z0-9_\\-]+=[A-Za-z0-9_%\\-.]+)?(&[A-Za-z0-9_\\-]+=[A-Za-z0-9_%\\-.]+)*$")
	fmt.Println(r.MatchString("/home/test/test"))
	fmt.Println(r.MatchString("/home/test?"))
	fmt.Println(r.MatchString("/home/test?aa12"))
	fmt.Println(r.MatchString("/home/test?aa12=134&f_12=1.2300"))
	fmt.Println(r.MatchString("/home/test?abc=14242&def=12324"))
	if !r.MatchString("/home/test?abc=14242&def=12324"){
		fmt.Print(r)
	}
	fmt.Println(r.MatchString("/s?ie=utf-8&f=8&rsv_bp=1&rsv_idx=1&tn=baidu&wd=golang%20%E6%AD%A3%E5%88%99%E8%A1%A8%E8%BE%BE%E5%BC%8F%20%E6%8B%AC%E5%8F%B7&oq=golang%2520%25E6%25AD%25A3%25E5%2588%2599%25E8%25A1%25A8%25E8%25BE%25BE%25E5%25BC%258F%2520url&rsv_pq=c7d31268000303cf&rsv_t=ba45nw77u7VXRpmUFyOdT3FfqgI9CMwVnxJbHcxZ8SKnn414ILG0rNWFlbs&rqlang=cn&rsv_enter=0&rsv_sug3=238&rsv_sug1=159&rsv_sug7=000&rsv_sug2=0&inputT=1806&rsv_sug4=4382&rsv_sug=1"))
}
