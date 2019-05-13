package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"

	helper "github.com/shouva/dailyhelper"
)

// Route :
type Route struct {
	Name string
	URL  string
}

func createRoutes(models []Route) string {
	strtemplate := `
	package main
	func loadrouter(rgin *gin.Engine) {
		{{range $index, $route := .}}initRouters{{$route.Name}}(rgin, "{{$route.URL}}")
		{{end}}}
	`
	tmpl := template.New("create api template")
	tmpl, err := tmpl.Parse(strtemplate)
	if err != nil {
		log.Fatal("Parse: ", err)
		return ""
	}

	// openfile
	filename := helper.GetCurrentPath(false) + "/out/routes.go"
	f, err := os.Create(filename)
	if err != nil {
		log.Println("create file: ", err)
		return ""
	}

	// var strout string
	err = tmpl.Execute(f, models)

	if err != nil {
		log.Fatal("Execute: ", err)
		return ""
	}
	f.Close()
	out, err := exec.Command("goimports", filename).Output()
	// out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Create(filename)
	f.Write(out)
	f.Close()
	return ""
}