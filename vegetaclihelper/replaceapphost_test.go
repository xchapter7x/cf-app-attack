package vegetaclihelper_test

import (
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/xchapter7x/civet-coffee/vegetaclihelper"
)

var _ = Describe("ReplaceAppHost", func() {
	Describe("given a template and a host string", func() {
		controlHost := "127.0.0.1"
		Context("when called w/ a template containing '{{.AppHost}}' ", func() {
			controlFormatString := "%s items are made of %s"
			controlTemplateString := fmt.Sprintf(controlFormatString, "{{.AppHost}}", "{{.AppHost}}")

			It("then it should return a string injects the host value in the proper place", func() {
				expectedResponse := fmt.Sprintf(controlFormatString, controlHost, controlHost)
				actualResponse := ReplaceAppHost(strings.NewReader(controlTemplateString), controlHost)
				actualResponseBytes, _ := ioutil.ReadAll(actualResponse)
				actualResponseString := string(actualResponseBytes[:])
				Ω(actualResponseString).Should(Equal(expectedResponse))
			})
		})

		Context("when called w/ a invalid template string ", func() {
			controlTemplateString := "{{."

			It("then it should error and panic", func() {
				Ω(func() {
					ReplaceAppHost(strings.NewReader(controlTemplateString), controlHost)
				}).Should(Panic())
			})
		})

		Context("when called with unknown template identifiers ", func() {
			controlTemplateString := "here and {{.NotReal}} something"

			It("then it should error and panic", func() {
				Ω(func() {
					ReplaceAppHost(strings.NewReader(controlTemplateString), controlHost)
				}).Should(Panic())
			})
		})
	})
})
