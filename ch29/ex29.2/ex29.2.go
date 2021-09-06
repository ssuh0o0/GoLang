//ch29/ex29.2/ex29.2.go
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	values := r.URL.Query()    // 쿼리 인수 가져오기 //map[string]string 타입의 별칭
	name := values.Get("name") // name 키값 가져오기 //기본적으로 반환형이 string인듯
	if name == "" {
		name = "world"
	}
	id, _ := strconv.Atoi(values.Get("id")) //아이디 키값을 가져와서 char -> int 형으로 바꾸기
	fmt.Fprintf(w, "hello %s! id: %d", name, id)
=======
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "world"
	}
	id, _ = strconv.Atoi(values.Get("id"))
	fmt.Fprint(w, "Hello %s ! id: %d", name, id)
>>>>>>> 485bae0bf0040cbb67299f19d4f0b72510de3bbb
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}
