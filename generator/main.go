package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/chanced/caps"
	flag "github.com/spf13/pflag"
)

const dirPerms = 0o750

func main() {
	var err error

	nameOpt := flag.String(
		"name",
		"",
		"Name of the data source to generate stubs for. MUST be in snake_case, beginning with corefunc_.",
	)

	flag.Parse()

	varMap := createVarMap(nameOpt)

	preview := newTemplate("preview.gotmpl")

	err = preview.Execute(os.Stdout, varMap)
	if err != nil {
		panic(err)
	}

	fmt.Println("")

	// ../bats
	writeFileFromTemplate(
		varMap,
		getAbs("bats/test.tmpl"),
		getAbs("../bats/"+varMap["SnakeStrip"]+".bats.sh"),
	)

	// ../testfixtures
	writeFileFromTemplate(
		varMap,
		getAbs("testfixtures.gotmpl"),
		getAbs("../testfixtures/"+varMap["SnakeStrip"]+".go"),
	)

	// ../corefuncprovider
	cfpPath := "../corefuncprovider/" + varMap["SnakeStrip"]
	writeFileFromTemplate(
		varMap,
		getAbs("./corefuncprovider/data_source.gotmpl"),
		getAbs(cfpPath+"_data_source.go"),
	)
	writeFileFromTemplate(
		varMap,
		getAbs("./corefuncprovider/data_source_test.gotmpl"),
		getAbs(cfpPath+"_data_source_test.go"),
	)
	writeFileFromTemplate(
		varMap,
		getAbs("./corefuncprovider/data_source_fixture.tftpl"),
		getAbs(cfpPath+"_data_source_fixture.tftpl"),
	)

	// ../templates
	writeFileFromTemplate(
		varMap,
		getAbs("./templates/data-source.md.gotmpl"),
		getAbs("../templates/data-sources/"+varMap["SnakeStrip"]+".md.tmpl"),
	)

	// ../examples
	edsPath := "../examples/data-sources/" + varMap["Snake"]

	err = os.MkdirAll(getAbs(edsPath), dirPerms)
	if err != nil {
		panic(err)
	}

	writeFileFromTemplate(
		varMap,
		getAbs("./examples/data-source.tftpl"),
		getAbs(edsPath+"/data-source.tf"),
	)
	writeFileFromTemplate(
		varMap,
		getAbs("./examples/terraform-plan.example.gotmpl"),
		getAbs(edsPath+"/terraform-plan.example"),
	)
	writeFileFromTemplate(
		varMap,
		getAbs("./examples/versions.tftpl"),
		getAbs(edsPath+"/versions.tftpl"),
	)

	fmt.Println("")
	fmt.Println("Add " + varMap["PascalStrip"] + "DataSource" + " to " + getAbs("../corefuncprovider/provider.go"))
}

func getAbs(path string) string {
	s, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return s
}

func writeFileFromTemplate(varMap map[string]string, templatePath, writePath string) {
	fmt.Println("Generating " + writePath)

	tmpl := newTemplate(templatePath)

	f, err := os.Create(writePath) // lint:allow_possible_insecure
	if err != nil {
		panic(err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = tmpl.Execute(f, varMap)
	if err != nil {
		panic(err)
	}
}

func newTemplate(filename string) *template.Template {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func createVarMap(nameOpt *string) map[string]string {
	original := *nameOpt
	snake := caps.ToSnake(original)
	snakeStrip := strings.Replace(snake, "corefunc_", "", 1)
	camel := caps.ToLowerCamel(original)
	camelStrip := strings.Replace(camel, "corefunc", "", 1)
	camelStrip = caps.ToLowerCamel(camelStrip)
	pascal := caps.ToCamel(original)
	pascalStrip := strings.Replace(pascal, "Corefunc", "", 1)

	return map[string]string{
		"Original":    original,
		"Snake":       snake,
		"SnakeStrip":  snakeStrip,
		"Camel":       camel,
		"CamelStrip":  camelStrip,
		"Pascal":      pascal,
		"PascalStrip": pascalStrip,
	}
}
