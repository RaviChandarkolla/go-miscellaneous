package ultimate_go_programming_course

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	println()
	tmpl2, err := template.New("test").Parse("{{23 -}} < {{- 45}}")
	if err != nil {
		panic(err)
	}
	err = tmpl2.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
