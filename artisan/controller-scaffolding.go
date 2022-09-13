package artisan

import (
	"bytes"
	"go/format"
	"html/template"
	"log"
	"os"

	"github.com/usama-tariq1/leet-astro/helper"
)

type DataController struct {
	Name    string
	ModName string
}

func CreateController(name string) {
	var console = helper.Console{}

	path := helper.GetWD()
	modName := helper.GetModuleName(path + `\go.mod`)
	data := DataController{
		Name:    name,
		ModName: modName + `/models`,
	}

	fileExist := helper.FileExist(path + `\controllers\` + name + `.go`)
	if fileExist {
		console.Log("Error", name+" Already Exist!")
		return
	}

	tmpl, err := template.ParseFiles(path + `\leet-gin\templates\ControllerTemplate.tmpl`)
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

	file, _ := os.Create(path + `\controllers\` + name + `.go`)
	file.Write(formatted)

	console.Log("Success", name+" Created Successfully")

}
