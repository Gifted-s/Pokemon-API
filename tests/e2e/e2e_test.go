package httpexpect

import (
	"log"
	"net/http"
	"pokemon/m/v1/routers"
	"testing"
	"github.com/gavv/httpexpect/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func pokmonRouteTester(t *testing.T) *httpexpect.Expect {
	err := godotenv.Load("../../.test.env")

	if err != nil {
	log.Panic("Error loading .env file", err)
	}
	r := mux.NewRouter()

	pokemonRouteHandler := routers.PokemonRouter(r)

	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(pokemonRouteHandler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func TestRouteParams(t *testing.T) {
	e := pokmonRouteTester(t)

	
	r := e.GET("/pokemon").
		WithQuery("hp[gte]", "100").WithQuery("defense[gte]", "200").
		Expect().
		Status(http.StatusOK).JSON().Object()
		r.Keys().ContainsOnly("status", "pokemons")
}
