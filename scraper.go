package main

import (
	"html/template"
	"log"
	"os"

	"github.com/kkyr/go-recipe/pkg/recipe"
)

func main() {
	url := "https://damndelicious.net/2022/10/14/perfect-pot-roast/"

	recipe, err := recipe.ScrapeURL(url)
	if err != nil {
		panic(err)
	}

	ingredients, _ := recipe.Ingredients()
	instructions, _ := recipe.Instructions()

	log.Print(recipe.Description())

	log.Print(ingredients, instructions)

	type Todo struct {
		Name        string
		Description string
	}

	td := Todo{"Test templates", "Let's get down to business."}
	t, err := template.New("todos").Parse(`# {{.Name}}
{{.Description}}`)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(("test.md"))
	t.Execute(file, td)
	file.Close()
	if err != nil {
		panic(err)
	}
}
