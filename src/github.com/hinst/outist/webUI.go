package outist

import (
	"bytes"
	"html/template"
	"net/http"
)

const PageDirectory = "src/github.com/hinst/outist_page"

type TWebUI struct {
	URL     string
	Pages   []string
	UserMan *TUserMan
}

func CreateWebUI() *TWebUI {
	var result = &TWebUI{}
	result.Pages = []string{"login"}
	return result
}

func (this *TWebUI) Start() {
	if this.URL == "" {
		this.URL = "/users"
	}
	if nil == this.UserMan {
		this.prepareNewUserMan()
	}
	this.registerFile("concise.css")
	http.HandleFunc(this.URL+"/page", this.ProcessPageRequest)
}

func (this *TWebUI) prepareNewUserMan() {
	this.UserMan = CreateUserMan()
	this.UserMan.Directory = AppDirectory + "/data/users"
	this.UserMan.Start()
}

func (this *TWebUI) registerFile(file string) {
	var url = this.URL + "/" + file
	var filePath = AppDirectory + "/" + PageDirectory + "/" + file
	GlobalLog.Write(url + " -> " + filePath)
	http.HandleFunc(url,
		func(response http.ResponseWriter, request *http.Request) {
			http.ServeFile(response, request, filePath)
		})
}

func (this *TWebUI) loadPage(data TPageData) string {
	var result = ""
	data.Body = this.substituteBody(data)
	var lTemplate, parseResult = template.ParseFiles(AppDirectory + "/" + PageDirectory + "/template.html")
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

func (this *TWebUI) CreatePageData() TPageData {
	var result TPageData
	result.AppURL = this.URL
	return result
}

func (this *TWebUI) ProcessPageRequest(response http.ResponseWriter, request *http.Request) {
	var pageName = request.URL.Query().Get("page")
	if CheckStringArrayContains(this.Pages, pageName) {
		var page = this.CreatePageData()
		page.Body = ReadStringFromFile(AppDirectory + "/" + PageDirectory + "/" + pageName + ".html")
		var content = this.loadPage(page)
		response.Write([]byte(content))
	}
}
