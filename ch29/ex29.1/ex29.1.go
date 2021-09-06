//ch29/ex29.1/ex29.1.go
package main

import (
	"fmt"
	"net/http"
)

<<<<<<< HEAD
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
	http.ListenAndServe(":3000", nil)
}
=======
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "hello world")
	})

	http.ListenAndServe(":3000", nil)
}
>>>>>>> 485bae0bf0040cbb67299f19d4f0b72510de3bbb
