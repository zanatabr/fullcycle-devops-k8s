package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", fakeWebServer)
	http.ListenAndServe(":8000", nil)
}

func fakeWebServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func greeting(msg string) string {
	return ("<b>" + msg + "</b>")
}
