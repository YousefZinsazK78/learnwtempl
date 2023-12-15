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
	expressionComponent := templates.ExpressionComp()
	greetComponent := templates.Greet("yousef and sina")
	statementComponent := templates.ShowIf(false)
	statementComponent2 := templates.Display(2003.2, 23)
	statementComponent3 := templates.Login(true)
	usertypeComponent := templates.UserTypeDisplay("Test")

	http.Handle("/", templ.Handler(helloComponent))
	http.Handle("/btn", templ.Handler(buttonComponent))
	http.Handle("/btnwvalue", templ.Handler(buttonComponent2))
	http.Handle("/attr", templ.Handler(booleanComponent))
	http.Handle("/attrwcond", templ.Handler(condComponent))
	http.Handle("/atag", templ.Handler(anchorComponent))
	http.Handle("/scriptbtn", templ.Handler(textbuttonComponent))
	http.Handle("/expr", templ.Handler(expressionComponent))
	http.Handle("/greet", templ.Handler(greetComponent))
	http.Handle("/state", templ.Handler(statementComponent))
	http.Handle("/display", templ.Handler(statementComponent2))
	http.Handle("/login", templ.Handler(statementComponent3))
	http.Handle("/type", templ.Handler(usertypeComponent))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

}
