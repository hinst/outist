package outist

import (
	"bytes"
	"html/template"
	"net/http"
)

const FileDirectory = "src/github.com/hinst/outist_page"

type TWebUI struct {
	URL string
}

func CreateWebUI() *TWebUI {
	var result = &TWebUI{}
	return result
}

func (this *TWebUI) Prepare() {
	this.registerFile("concise.css")
}

func (this *TWebUI) registerFile(file string) {
	var url = this.URL + "/" + file
	var filePath = AppDirectory + "/" + FileDirectory + "/" + file
	GlobalLog.Write(url + " -> " + filePath)
	http.HandleFunc(url,
		func(response http.ResponseWriter, request *http.Request) {
			http.ServeFile(response, request, filePath)
		})
}

func (this *TWebUI) loadPage(data TPageData) string {
	var result = ""
	data.Body = this.substituteBody(data)
	var lTemplate, parseResult = template.ParseFiles(AppDirectory + "/" + FileDirectory + "/template.html")
	if parseResult == nil {
		var output bytes.Buffer
		var executeResult = lTemplate.Execute(&output, data)
		if executeResult == nil {
			result = output.String()
		} else {
			GlobalLog.Write("Error: could not execute template: " + executeResult.Error())
		}
	} else {
		GlobalLog.Write("Error: could not parse template: " + parseResult.Error())
	}
	return result
}

func (this *TWebUI) substituteBody(data TPageData) string {
	var result = data.Body
	var lTemplate, templateParseResult = template.New("").Parse(data.Body)
	if templateParseResult == nil {
		var output bytes.Buffer
		var executeResult = lTemplate.Execute(&output, data)
		if executeResult == nil {
			result = output.String()
		} else {
			GlobalLog.Write("Error: could not execute template: " + executeResult.Error())
		}
	} else {
		GlobalLog.Write("Error: could not parse template: " + templateParseResult.Error())
	}
	return result
}
