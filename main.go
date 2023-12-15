package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/yousefzinsazk78/templ_test_project/templates"
)

func main() {
	helloComponent := templates.Hello("sina")
	buttonComponent := templates.Button("Click Me!")
	buttonComponent2 := templates.ButtonWithAttr("some value", "Click Me!")
	booleanComponent := templates.BooleanAttr()

	http.Handle("/", templ.Handler(helloComponent))
	http.Handle("/btn", templ.Handler(buttonComponent))
	http.Handle("/btnwvalue", templ.Handler(buttonComponent2))
	http.Handle("/attr", templ.Handler(booleanComponent))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

}
