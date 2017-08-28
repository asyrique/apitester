package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func catchall(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", r)
}
func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/*"), catchall)
	mux.HandleFunc(pat.Post("/*"), catchall)
	http.ListenAndServe("localhost:8080", mux)
}
