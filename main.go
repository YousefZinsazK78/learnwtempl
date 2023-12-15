package main

import (
	"context"
	"fmt"
	"os"

	"github.com/yousefzinsazk78/templ_test_project/templates"
)

func main() {
	fmt.Println("this is test from yousef...")
	templates.Hello("yousef").Render(context.TODO(), os.Stdout)
}
