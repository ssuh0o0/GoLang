//ch29/ex29.7/ex29.7.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World") //  웹 핸들러 등록
	})

	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil) // 호출인수로 인증서와 비밀키 파일을 넣어준다.
	if err != nil {
		log.Fatal(err)
	}
}
