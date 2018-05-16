package main
import (
"fmt"
"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "hello world")
}

func main() {
http.HandleFunc("/", IndexHandler)
err:=http.ListenAndServe("127.0.0.1:8000", nil)
	fmt.Print(err)
}