package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func helloHandler(rw http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(rw,
		`<h1>Hello Codefresh!</h1>

		<p>Greetings from %s with a %s CPU`,
			runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":3000", nil)
}
