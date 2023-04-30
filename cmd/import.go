package cmd

import (
	"errors"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/gosimple/slug"
	"github.com/kkyr/go-recipe/pkg/recipe"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a recipe from a website",
	Long: `This commmand imports a recipe from a provided website URL.
recipectl import <url>`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("a URL must be provided")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Importing recipe...")
		url := args[0]

		recipe, err := recipe.ScrapeURL(url)
		if err != nil {
			panic(err)
		}

		type Recipe struct {
			Title        string
			Description  string
			Author       string
			CookTime     string
			PrepTime     string
			TotalTime    string
			Yield        string
			Instructions []string
			Ingredients  []string
			ImageURL     string
			Date         string
		}

		title, _ := recipe.Name()
		description, _ := recipe.Description()
		author, _ := recipe.Author()
		cookTime, _ := recipe.CookTime()
		prepTime, _ := recipe.PrepTime()
		totalTime, _ := recipe.TotalTime()
		yield, _ := recipe.Yields()
		instructions, _ := recipe.Instructions()
		ingredients, _ := recipe.Ingredients()
		imageURL, _ := recipe.ImageURL()

		td := Recipe{
			Title:       title,
			Description: description,
			Author:      author,
			// Converting to string here because many recipe sites
			// do not rigidly follow this standard
			CookTime:     cookTime.String(),
			TotalTime:    totalTime.String(),
			PrepTime:     prepTime.String(),
			Yield:        yield,
			Instructions: instructions,
			Ingredients:  ingredients,
			ImageURL:     imageURL,
			Date:         time.Now().Format(time.RFC1123),
		}
		slug := slug.Make(td.Title)
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		tmpl, err := template.ParseFiles(cwd + "/_recipe_template.md.tpl")
		if err != nil {
			panic(err)
		}
		file, err := os.Create(cwd + "/" + slug + ".md")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		err = tmpl.Execute(file, td)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.SetUsageTemplate(`Usage:
recipectl import <url> [flags]
recipectl import https://recipes.com/cheese-tax`)
}
