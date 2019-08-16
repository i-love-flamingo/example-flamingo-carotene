package interfaces

import (
	"context"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	HelloController struct {
		responder *web.Responder
	}

	helloViewData struct {
		Name     string
		Nickname string
	}
)

func (controller *HelloController) Inject(responder *web.Responder) {
	controller.responder = responder
}

func (controller *HelloController) Get(ctx context.Context, r *web.Request) web.Result {
	// Calling the Render method from the response helper and render the template "hello"
	return controller.responder.Render("hello/hello", helloViewData{
		Name: "World",
	})
}

func (controller *HelloController) GreetMe(ctx context.Context, r *web.Request) web.Result {
	name, err := r.Query1("name")
	if err != nil {
		name = "World (default)"
	}

	nick, _ := r.Params["nickname"]

	return controller.responder.Render("hello/hello", helloViewData{
		Name:     name,
		Nickname: nick,
	})
}

// CurrentTime is a DataAction that handles data calls from templates
func (controller *HelloController) CurrentTime(ctx context.Context, r *web.Request, params web.RequestParams) interface{} {
	return time.Now().Format(time.RFC822)
}
