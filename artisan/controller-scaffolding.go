package artisan

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/usama-tariq1/leet-astro/helper"
)

type DataController struct {
	Name      string
	ModName   string
	ModelName string
}

func CreateController(name string) {
	var console = helper.Console{}

	fileSrc := strings.Split(name, "/")
	fileSrcPath := strings.Join(fileSrc, `\`)

	path := helper.GetWD()
	modName := helper.GetModuleName(filepath.Join(path, `go.mod`))
	data := DataController{
		Name:    fileSrc[len(fileSrc)-1],
		ModName: modName,
	}

	fmt.Println(fileSrc[len(fileSrc)-1])

	fileExist := helper.FileExist(filepath.Join(path, `controllers`, filepath.Clean(fileSrcPath)+`.go`))
	if fileExist {
		console.Log("Error", fileSrcPath+" Already Exist!")
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join(path, `leet-gin`, `templates`, `ControllerTemplate.tmpl`))
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
	file, _ := os.Create(filepath.Join(path, `controllers`, fileSrcPath+".go"))
	file.Write(formatted)

	console.Log("Success", fileSrcPath+" Created Successfully")

}

func CreateResourceController(ControllerName string, ModelName string) {
	var console = helper.Console{}

	fileSrc := strings.Split(ControllerName, "/")
	fileSrcPath := strings.Join(fileSrc, `\`)

	path := helper.GetWD()
	modName := helper.GetModuleName(filepath.Join(path, `go.mod`))
	data := DataController{
		Name:      fileSrc[len(fileSrc)-1],
		ModName:   modName,
		ModelName: ModelName,
	}

	fmt.Println(fileSrc[len(fileSrc)-1])

	fileExist := helper.FileExist(filepath.Join(path, `controllers`, filepath.Clean(fileSrcPath)+`.go`))
	if fileExist {
		console.Log("Error", fileSrcPath+" Already Exist!")
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join(path, `leet-gin`, `templates`, `ResourceControllerTemplate.tmpl`))
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
	file, _ := os.Create(filepath.Join(path, `controllers`, fileSrcPath+".go"))
	file.Write(formatted)

	console.Log("Success", fileSrcPath+" Created Successfully")

}
