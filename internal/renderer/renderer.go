package renderer

import (
	"os"
	"path"

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

	renderedContent, err := mustache.RenderFile(stackItem.Template, stackItem.TemplateData)
	check(err)

	filename := path.Join(output, stackItem.Output)
	dir := path.Dir(filename)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	_, err = f.WriteString(renderedContent)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
