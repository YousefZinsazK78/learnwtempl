package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/yousefzinsazk78/templ_test_project/templates"
)

func main() {
	helloComponent := templates.Hello("sina")

	http.Handle("/", templ.Handler(helloComponent))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

}
