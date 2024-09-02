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

type DataRouter struct {
	Name           string
	ModName        string
	ControllerName string
}

func CreateRouter(name string) {
	var console = helper.Console{}

	path := helper.GetWD()
	modName := helper.GetModuleName(filepath.Join(path, `go.mod`))
	data := DataRouter{
		ModName: modName,
		Name:    name,
	}

	fileExist := helper.FileExist(filepath.Join(path, `routers`, name+`.go`))
	if fileExist {
		console.Log("Error", name+" Already Exist!")
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join(path, `leet-gin`, `templates`, `RouterTemplate.tmpl`))
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

	file, _ := os.Create(filepath.Join(path, `routers`, name+`.go`))
	file.Write(formatted)

	console.Log("Success", name+" Created Successfully")

}

func CreateRouterWithController(name string, controllerName string) {
	var console = helper.Console{}

	path := helper.GetWD()
	modName := helper.GetModuleName(filepath.Join(path, `go.mod`))
	data := DataRouter{
		ModName:        modName,
		Name:           name,
		ControllerName: controllerName,
	}

	fileExist := helper.FileExist(filepath.Join(path, `routers`, name+`.go`))
	if fileExist {
		console.Log("Error", name+" Already Exist!")
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join(path, `leet-gin`, `templates`, `RouterTemplateWithController.tmpl`))
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

	file, _ := os.Create(filepath.Join(path, `routers`, name+`.go`))
	file.Write(formatted)

	console.Log("Success", name+" Created Successfully")

}
