/* Alta3 Research | RZFeeser
   text/template - rendering a simple text string   */

package main

import (
	"os"
	"text/template"
)

// create a struct that matches the "blanks" in our template string
type Todo struct {
	Name        string
	Description string
}

func main() {

	td := Todo{"Test templates", "Let's test a template to see the magic."}

	// Create a template
	// Reading the template from an external file is probably more realistic
	t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"\n")

	// error checking
	if err != nil {
		panic(err)
	}

	// this is where we take the template (t) and merge it with our fill-in-the blank data (td)
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}

	// it is possible to reuse the same template, without
	// needing to create or parse it again by providing the struct you want
	// to use to the "Execute" function again
	rzfNew := Todo{"Make breakfast", "Grind and brew coffee, and find a croissant"}

	err = t.Execute(os.Stdout, rzfNew) // reuse our same template
	// error checking
	if err != nil {
		panic(err)
	}
}
