package httpexpect

import (
	"log"
	"net/http"
	"pokemon/m/v1/routers"
	"testing"
	"github.com/gavv/httpexpect/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	// "github.com/stretchr/testify/assert"
)


func pokmonRouteTester(t *testing.T) *httpexpect.Expect {
	err := godotenv.Load("../../test.env")

	if err != nil {
	log.Panic("Error loading test.env file", err)
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

func TestResponseStatus(t *testing.T) {
	e := pokmonRouteTester(t)
	r := e.GET("/pokemon").
		WithQuery("hp[gte]", "100").WithQuery("defense[gte]", "100").WithQuery("attack[gte]", "100").WithQuery("page", "1").WithQuery("search", "Squirtle").
		Expect().
		Status(http.StatusOK).JSON().Object()
		r.Value("status").Number().Equal(200)

}

func TestResponseKeys(t *testing.T) {
	e := pokmonRouteTester(t)
	r := e.GET("/pokemon").
		WithQuery("hp[gte]", "100").WithQuery("defense[gte]", "100").WithQuery("attack[gte]", "100").WithQuery("page", "1").WithQuery("search", "Squirtle").
		Expect().
		Status(http.StatusOK).JSON().Object()
		r.Keys().ContainsOnly("status", "pokemons")
}

