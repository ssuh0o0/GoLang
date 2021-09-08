//ch29/ex29.2/ex29.2.go
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    // 쿼리 인수 가져오기 //map[string]string 타입의 별칭
	name := values.Get("name") // name 키값 가져오기 //기본적으로 반환형이 string인듯
	if name == "" {
		name = "world"
	}
	id, _ := strconv.Atoi(values.Get("id")) //아이디 키값을 가져와서 char -> int 형으로 바꾸기
	fmt.Fprintf(w, "hello %s! id: %d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}
