package swagger

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SwaggerRoutes() {
	http.Handle(
		"/swagger/",
		httpSwagger.WrapHandler,
	)
}
