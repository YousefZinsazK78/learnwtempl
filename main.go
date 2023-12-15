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
	condComponent := templates.ConditionAttr()
	anchorComponent := templates.AnchorComponent("www.google.com/")
	textbuttonComponent := templates.Button3("text text text")

	http.Handle("/", templ.Handler(helloComponent))
	http.Handle("/btn", templ.Handler(buttonComponent))
	http.Handle("/btnwvalue", templ.Handler(buttonComponent2))
	http.Handle("/attr", templ.Handler(booleanComponent))
	http.Handle("/attrwcond", templ.Handler(condComponent))
	http.Handle("/atag", templ.Handler(anchorComponent))
	http.Handle("/scriptbtn", templ.Handler(textbuttonComponent))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

}
