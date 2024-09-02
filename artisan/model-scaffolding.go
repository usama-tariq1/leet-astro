package artisan

import (
	"bytes"
	"go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/usama-tariq1/leet-astro/helper"
)

type DataModel struct {
	Name    string
	ModName string
}

func CreateModel(name string, shouldCreateController bool, shouldCreateRouter bool) {
	var console = helper.Console{}

	path := helper.GetWD()
	modName := helper.GetModuleName(filepath.Join(path, `go.mod`))

	data := DataModel{
		Name:    name,
		ModName: modName,
	}

	fileExist := helper.FileExist(filepath.Join(path, `models`, name+`.go`))
	if fileExist {
		console.Log("Error", name+" Already Exist!")
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join(path, `leet-gin`, `templates`, `ModelTemplate.tmpl`))
	if err != nil {
		log.Print(err)
		return
	}

	var processed bytes.Buffer
	err = tmpl.Execute(&processed, data)
	if err != nil {
		log.Fatalf("unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}

	file, _ := os.Create(filepath.Join(path, `models`, name+`.go`))
	file.Write(formatted)

	console.Log("Success", name+" Created Successfully")

	if shouldCreateController {
		CreateResourceController(name+`Controller`, name)
	}

	if shouldCreateRouter {
		CreateRouterWithController(name+`Router`, name+`Controller`)
	}

}
