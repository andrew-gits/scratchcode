package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(os.Getenv("PLANETSCALE_DB"))
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
