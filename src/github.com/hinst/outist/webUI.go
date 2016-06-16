package outist

import "net/http"

const FileDirectory = "github.com/hinst/outist_page"

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
	http.HandleFunc(url,
		func(response http.ResponseWriter, request *http.Request) {
			http.ServeFile(response, request, FileDirectory+"/"+file)
		})
}
