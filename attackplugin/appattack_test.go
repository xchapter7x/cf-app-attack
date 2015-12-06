package attackplugin_test

import (
	"errors"
	"fmt"

	"github.com/cloudfoundry/cli/plugin/fakes"
	"github.com/cloudfoundry/cli/plugin/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/xchapter7x/cf-app-attack/attackplugin"
)

var _ = Describe("AppAttack", func() {
	Describe("given GetMetadata method", func() {
		Context("when called", func() {
			It("then it should return valid metadata", func() {
				appAttack := &AppAttack{Version: "v1.1.1"}
				meta := appAttack.GetMetadata()
				Ω(meta.Name).Should(Equal(PluginName))
			})
		})
	})
	Describe("given Run method", func() {
		var (
			appAttack = &AppAttack{Version: "v1.1.1"}
			argSpy    []string
			hostSpy   string
		)
		VegetaRunner = func(a []string, h string) {
			hostSpy = h
			argSpy = a
		}

		BeforeEach(func() {
			argSpy = nil
			hostSpy = ""
		})
		Context("when called for a invalid app", func() {

			var fakeCli *fakes.FakeCliConnection

			BeforeEach(func() {
				fakeCli = new(fakes.FakeCliConnection)
				fakeCli.GetAppReturns(plugin_models.GetAppModel{
					Routes: []plugin_models.GetApp_RouteSummary{
						plugin_models.GetApp_RouteSummary{
							Host: "fakehost",
							Domain: plugin_models.GetApp_DomainFields{
								Name: "fakedomain.com",
							},
						},
					},
				}, errors.New("fake invalid app error"))
			})
			It("then it should panic", func() {
				Ω(func() {
					appAttack.Run(fakeCli, []string{CmdBench, "badapp", "blah", "blah"})
				}).Should(Panic())
			})
		})

		Context("when called for a valid app", func() {
			var (
				fakeCli       *fakes.FakeCliConnection
				controlHost   = "fakehost"
				controlDomain = "fakedomain.com"
			)

			BeforeEach(func() {
				fakeCli = new(fakes.FakeCliConnection)
				fakeCli.GetAppReturns(plugin_models.GetAppModel{
					Routes: []plugin_models.GetApp_RouteSummary{
						plugin_models.GetApp_RouteSummary{
							Host: controlHost,
							Domain: plugin_models.GetApp_DomainFields{
								Name: controlDomain,
							},
						},
					},
				}, nil)
			})

			It("then it should use the proper host for the given app", func() {
				expectedHost := fmt.Sprintf("%s.%s", controlHost, controlDomain)
				appAttack.Run(fakeCli, []string{CmdBench, "goodapp", "blah", "blah"})
				Ω(hostSpy).Should(Equal(expectedHost))
			})
		})
		Context("when called with a invalid command", func() {
			It("then it should not call any vegeta actions", func() {
				appAttack.Run(new(fakes.FakeCliConnection), []string{"bad-command", "blah", "blah"})
				Ω(argSpy).Should(BeNil())
			})
		})
		Context("when called with 'bench' command", func() {
			var fakeCli *fakes.FakeCliConnection

			BeforeEach(func() {
				fakeCli = new(fakes.FakeCliConnection)
				fakeCli.GetAppReturns(plugin_models.GetAppModel{
					Routes: []plugin_models.GetApp_RouteSummary{
						plugin_models.GetApp_RouteSummary{
							Host: "fakehost",
							Domain: plugin_models.GetApp_DomainFields{
								Name: "fakedomain.com",
							},
						},
					},
				}, nil)
			})
			It("then it should create a properly structured attack vegeta call", func() {
				controlVegetaArgs := []string{"blah", "blah", "hithere"}
				appAttack.Run(fakeCli, append([]string{CmdBench, "myappname"}, controlVegetaArgs...))
				Ω(argSpy).Should(Equal(controlVegetaArgs))
			})
		})
	})
})
