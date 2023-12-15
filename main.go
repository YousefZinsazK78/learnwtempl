package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/yousefzinsazk78/templ_test_project/templates"
)

func main() {

	var items = []templates.Item{
		templates.Item{1, "yousef"},
		templates.Item{2, "sina"},
	}

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
	forloopComponent := templates.ForLoopDisplay(items)
	compositeComponent := templates.ShowAll()
	compositeComponent2 := templates.ShowAllChildren()
	compositeParamsComponent := templates.Layout(templates.Paraghraph("string content from main.go"))
	compositeParamsComponent2 := templates.Root()
	cssStyleComponent := templates.StyleButton("click me!", []string{"borderColor", "bgblue"})
	cssStyleButtonComponent := templates.StyleButton2("click me!", false)
	cssStyleButtonComponent3 := templates.StyleButton33("click meeee!", true)

	var c1 = templates.ClassName()
	var c2 = templates.PrimaryClassName()

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
	http.Handle("/for", templ.Handler(forloopComponent))
	http.Handle("/compo", templ.Handler(compositeComponent))
	http.Handle("/compo2", templ.Handler(compositeComponent2))
	http.Handle("/compo3", templ.Handler(compositeParamsComponent))
	http.Handle("/compo4", templ.Handler(compositeParamsComponent2))
	http.Handle("/css1", templ.Handler(cssStyleComponent))
	http.Handle("/css2", templ.Handler(cssStyleButtonComponent))
	http.Handle("/css3", templ.Handler(cssStyleButtonComponent3))

	handlerCss := templ.NewCSSMiddleware(templ.Handler(cssStyleButtonComponent3), c1, c2)

	if err := http.ListenAndServe(":8000", handlerCss); err != nil {
		panic(err)
	}

}
