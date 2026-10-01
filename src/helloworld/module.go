package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/example-hello-flamingo-carotene/src/helloworld/interfaces"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	// Module is our helloWorld Module
	Module struct {
	}
)

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

// routes struct - our internal struct that gets the interface methods for router.Module
type routes struct {
	// helloController - we will defined routes that are handled by our HelloController - so we need this as a dependency
	helloController *interfaces.HelloController
}

// Inject method - this is called by Dingo and gets a initializes instance of the HelloController passed automatically
func (r *routes) Inject(controller *interfaces.HelloController) {
	r.helloController = controller
}

// Routes method which defines all routes handlers in module
func (r *routes) Routes(registry *web.RouterRegistry) {
	// Bind the controller.Action to the handle "hello":
	registry.HandleGet("helloWorld.hello", r.helloController.Get)
	// Bind the path /hello to a handle with the name "hello"
	registry.Route("/hello", "helloWorld.hello")

	registry.HandleGet("helloWorld.greetme", r.helloController.GreetMe)
	registry.Route("/greetme", "helloWorld.greetme")
	registry.Route("/greetme/:nickname", "helloWorld.greetme")
	registry.Route("/greetflamingo", `helloWorld.greetme(nickname="Flamingo")`)

	registry.HandleData("currenttime", r.helloController.CurrentTime)

}
