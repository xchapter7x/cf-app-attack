package vegetaclihelper

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/xchapter7x/lo"
)

//ReplaceAppHost -- a function to replace a host in a given template string
func ReplaceAppHost(templateReader io.Reader, apphost string) (appHostTemplate io.Reader) {
	templateBytes, _ := ioutil.ReadAll(templateReader)
	templateString := string(templateBytes[:])
	vegeta := vegetaRequest{AppHost: apphost}

	if tmpl, err := template.New("vegetaRequest").Parse(templateString); err != nil {
		lo.G.Error("could not generate template for given request: ", err.Error())
		panic(err)

	} else {
		var buf bytes.Buffer

		if err = tmpl.Execute(&buf, vegeta); err != nil {
			lo.G.Error("could not execute template: ", err.Error())
			panic(err)

		} else {
			appHostTemplateString := buf.String()
			appHostTemplate = strings.NewReader(appHostTemplateString)
		}
	}
	return
}

type vegetaRequest struct {
	AppHost string
}
