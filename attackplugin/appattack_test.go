package attackplugin_test

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/xchapter7x/cf-app-attack/attackplugin"
)

var _ = Describe("AppAttack", func() {
	Describe("given GetMetadata method", func() {
		Context("when called", func() {
			It("then it should return valid metadata", func() {
				appAttack := new(AppAttack)
				meta := appAttack.GetMetadata()
				Ω(meta.Name).Should(Equal(PluginName))
			})
		})
	})
	Describe("given Run method", func() {
		var (
			appAttack = new(AppAttack)
			argSpy    []string
		)
		VegetaRunner = func(a []string) {
			argSpy = a
		}
		Context("when called with a invalid command", func() {
			It("then it should not call any vegeta actions", func() {
				appAttack.Run(new(fakes.FakeCliConnection), []string{"bad-command", "blah", "blah"})
				Ω(argSpy).Should(BeNil())
			})
		})
		Context("when called with 'bench' command", func() {
			It("then it should create a properly structured attack vegeta call", func() {
				controlVegetaArgs := []string{"blah", "blah", "hithere"}
				appAttack.Run(new(fakes.FakeCliConnection), append([]string{CmdBench, "myappname"}, controlVegetaArgs...))
				Ω(argSpy).Should(Equal(controlVegetaArgs))
			})
		})
	})
})
