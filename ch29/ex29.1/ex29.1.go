//ch29/ex29.1/ex29.1.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
	http.ListenAndServe(":3000", nil)
}
