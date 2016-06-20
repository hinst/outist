package main

import (
	"net/http"

	"github.com/hinst/outist"
)

func main() {
	outist.StartGlobalLog(outist.AppDirectory + "/userLog") // optional
	outist.CreateWebUI().Start()
	http.ListenAndServe(":9000", nil)
}
