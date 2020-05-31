package renderer

import (
	"os"

	"github.com/cbroglie/mustache"
)

func Process(output string, stack *Stack) {

	for {
		stackItem := stack.Pop()

		if stackItem == nil {
			break
		}

		defer render(output, stackItem)

	}
}

func render(output string, stackItem *StackItem) {

	output, err := mustache.RenderFile(stackItem.Template, stackItem.TemplateData)
	check(err)

	f, err := os.Create(stackItem.Output)
	check(err)
	defer f.Close()

	_, err = f.WriteString(output)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
