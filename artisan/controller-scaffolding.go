package artisan

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/usama-tariq1/leet-astro/helper"
)

type DataController struct {
	Name    string
	ModName string
}

func CreateController(name string) {
	var console = helper.Console{}

	fileSrc := strings.Split(name, "/")
	fileSrcPath := strings.Join(fileSrc, `\`)

	path := helper.GetWD()
	modName := helper.GetModuleName(path + `\go.mod`)
	data := DataController{
		Name:    fileSrc[len(fileSrc)-1],
		ModName: modName,
	}

	fmt.Println(fileSrc[len(fileSrc)-1])

	fileExist := helper.FileExist(path + `\controllers\` + fileSrcPath + `.go`)
	if fileExist {
		console.Log("Error", fileSrcPath+" Already Exist!")
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
	file, _ := os.Create(path + `\controllers\` + fileSrcPath + `.go`)
	file.Write(formatted)

	console.Log("Success", fileSrcPath+" Created Successfully")

}
