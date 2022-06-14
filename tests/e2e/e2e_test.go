package httpexpect

import (
	"log"
	"net/http"
	"pokemon/m/v1/db"
	"pokemon/m/v1/routers"
	"pokemon/m/v1/tests/fixtures"
	"testing"
	"github.com/gavv/httpexpect/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pokemon/m/v1/models"
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
	defer db.DropDB()
	fakeRandomPokemons := fixtures.CreateFakePokemons()
	e := pokmonRouteTester(t)
	db.InsertPokemons(fakeRandomPokemons)
	r := e.GET("/pokemon").
		WithQuery("hp[gte]", "100").WithQuery("defense[gte]", "100").WithQuery("attack[gte]", "100").WithQuery("page", "1").WithQuery("search", "Chaldeans").
		Expect().
		Status(http.StatusOK).JSON().Object()
	r.Value("status").Number().Equal(200)
}

func TestResponseKeys(t *testing.T) {
	defer db.DropDB()
	fakeRandomPokemons := fixtures.CreateFakePokemons()
	e := pokmonRouteTester(t)
	db.InsertPokemons(fakeRandomPokemons)
	r := e.GET("/pokemon").
		WithQuery("hp[gte]", "100").WithQuery("defense[gte]", "100").WithQuery("attack[gte]", "100").WithQuery("page", "1").WithQuery("search", "Chaldeans").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r.Keys().ContainsOnly("status", "pokemons")
}


func TestResponsePageSize(t *testing.T) {
	defer db.DropDB()
	fakeRandomPokemons := fixtures.CreateFakePokemons()
	e := pokmonRouteTester(t)
	db.InsertPokemons(fakeRandomPokemons)
	r := e.GET("/pokemon").
		WithQuery("hp[gte]", "0").WithQuery("defense[gte]", "0").WithQuery("attack[gte]", "0").WithQuery("page", "1").WithQuery("search", "").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r.Value("pokemons").Array().Length().Equal(10)
}

func TestFilterWithHP(t *testing.T) {
	defer db.DropDB()
	e := pokmonRouteTester(t)
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316d")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b3167")
	fakeRandomPokemons := []models.Pokemon{
		{
			ID:           fakeId1,
			Name:         "TestName1",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           50,
			Attack:       400,
			Defense:      200,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId2,
			Name:         "TestName2",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           10,
			Attack:       400,
			Defense:      200,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
	}

	db.InsertPokemons(fakeRandomPokemons)
	// Test HP greater than a particular value
	r := e.GET("/pokemon").
		WithQuery("hp[gt]", "30").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r.Value("pokemons").Array().Length().Equal(1)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("hp", 50)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	// Test HP less than a particular value
	r2 := e.GET("/pokemon").
		WithQuery("hp[lt]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r2.Value("pokemons").Array().Length().Equal(1)
	r2.Value("pokemons").Array().Element(0).Object().ValueEqual("hp", 10)
	r2.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId2)

	// Test HP less than or equal to a particular value
	r3 := e.GET("/pokemon").
		WithQuery("hp[lte]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r3.Value("pokemons").Array().Length().Equal(2)
	r3.Value("pokemons").Array().Element(0).Object().ValueEqual("hp", 50)
	r3.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	r3.Value("pokemons").Array().Element(1).Object().ValueEqual("hp", 10)
	r3.Value("pokemons").Array().Element(1).Object().ValueEqual("_id", fakeId2)

	// Test HP greater than or equal to a particular value
	r4 := e.GET("/pokemon").
		WithQuery("hp[gte]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r4.Value("pokemons").Array().Length().Equal(1)
	r4.Value("pokemons").Array().Element(0).Object().ValueEqual("hp", 50)
	r4.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	// Test HP equal to a particular value
	r5 := e.GET("/pokemon").
		WithQuery("hp[eq]", "10").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r5.Value("pokemons").Array().Length().Equal(1)
	r5.Value("pokemons").Array().Element(0).Object().ValueEqual("hp", 10)
	r5.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId2)

	// Test HP not equal to a particular value
	r6 := e.GET("/pokemon").
		WithQuery("hp[ne]", "10").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r6.Value("pokemons").Array().Length().Equal(1)
	r6.Value("pokemons").Array().Element(0).Object().ValueEqual("hp", 50)
	r6.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

}


func TestFilterWithDefence(t *testing.T) {
	defer db.DropDB()
	e := pokmonRouteTester(t)
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316d")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b3167")
	fakeRandomPokemons := []models.Pokemon{
		{
			ID:           fakeId1,
			Name:         "TestName1",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           60,
			Attack:       400,
			Defense:      50,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId2,
			Name:         "TestName2",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           40,
			Attack:       80,
			Defense:      10,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
	}

	db.InsertPokemons(fakeRandomPokemons)
	// Test Defense greater than a particular value
	r := e.GET("/pokemon").
		WithQuery("defense[gt]", "30").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r.Value("pokemons").Array().Length().Equal(1)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("defense", 50)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	// Test Defense less than a particular value
	r2 := e.GET("/pokemon").
		WithQuery("defense[lt]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r2.Value("pokemons").Array().Length().Equal(1)
	r2.Value("pokemons").Array().Element(0).Object().ValueEqual("defense", 10)
	r2.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId2)

	// Test Defense less than or equal to a particular value
	r3 := e.GET("/pokemon").
		WithQuery("defense[lte]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r3.Value("pokemons").Array().Length().Equal(2)
	r3.Value("pokemons").Array().Element(0).Object().ValueEqual("defense", 50)
	r3.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	r3.Value("pokemons").Array().Element(1).Object().ValueEqual("defense", 10)
	r3.Value("pokemons").Array().Element(1).Object().ValueEqual("_id", fakeId2)

	// Test Defense greater than or equal to a particular value
	r4 := e.GET("/pokemon").
		WithQuery("defense[gte]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r4.Value("pokemons").Array().Length().Equal(1)
	r4.Value("pokemons").Array().Element(0).Object().ValueEqual("defense", 50)
	r4.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	// Test Defense equal to a particular value
	r5 := e.GET("/pokemon").
		WithQuery("defense[eq]", "10").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r5.Value("pokemons").Array().Length().Equal(1)
	r5.Value("pokemons").Array().Element(0).Object().ValueEqual("defense", 10)
	r5.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId2)

	// Test Defense not equal to a particular value
	r6 := e.GET("/pokemon").
		WithQuery("defense[ne]", "10").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r6.Value("pokemons").Array().Length().Equal(1)
	r6.Value("pokemons").Array().Element(0).Object().ValueEqual("defense", 50)
	r6.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

}





func TestFilterWithAttack(t *testing.T) {
	defer db.DropDB()
	e := pokmonRouteTester(t)
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316d")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b3167")
	fakeRandomPokemons := []models.Pokemon{
		{
			ID:           fakeId1,
			Name:         "TestName1",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           60,
			Attack:       50,
			Defense:      50,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId2,
			Name:         "TestName2",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           40,
			Attack:       10,
			Defense:      30,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
	}

	db.InsertPokemons(fakeRandomPokemons)
	// Test Attack greater than a particular value
	r := e.GET("/pokemon").
		WithQuery("attack[gt]", "30").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r.Value("pokemons").Array().Length().Equal(1)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("attack", 50)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	// Test Attack less than a particular value
	r2 := e.GET("/pokemon").
		WithQuery("attack[lt]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r2.Value("pokemons").Array().Length().Equal(1)
	r2.Value("pokemons").Array().Element(0).Object().ValueEqual("attack", 10)
	r2.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId2)

	// Test Attack less than or equal to a particular value
	r3 := e.GET("/pokemon").
		WithQuery("attack[lte]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r3.Value("pokemons").Array().Length().Equal(2)
	r3.Value("pokemons").Array().Element(0).Object().ValueEqual("attack", 50)
	r3.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	r3.Value("pokemons").Array().Element(1).Object().ValueEqual("attack", 10)
	r3.Value("pokemons").Array().Element(1).Object().ValueEqual("_id", fakeId2)

	// Test Attack greater than or equal to a particular value
	r4 := e.GET("/pokemon").
		WithQuery("attack[gte]", "50").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r4.Value("pokemons").Array().Length().Equal(1)
	r4.Value("pokemons").Array().Element(0).Object().ValueEqual("attack", 50)
	r4.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	// Test Attack equal to a particular value
	r5 := e.GET("/pokemon").
		WithQuery("attack[eq]", "10").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r5.Value("pokemons").Array().Length().Equal(1)
	r5.Value("pokemons").Array().Element(0).Object().ValueEqual("attack", 10)
	r5.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId2)

	// Test Attack not equal to a particular value
	r6 := e.GET("/pokemon").
		WithQuery("attack[ne]", "10").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r6.Value("pokemons").Array().Length().Equal(1)
	r6.Value("pokemons").Array().Element(0).Object().ValueEqual("attack", 50)
	r6.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

}




func TestSearchFilter(t *testing.T) {
	defer db.DropDB()
	e := pokmonRouteTester(t)
	searchText:= "Wartortle"
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316d")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b3167")
	fakeRandomPokemons := []models.Pokemon{
		{
			ID:           fakeId1,
			Name:         "Wartortle",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           60,
			Attack:       50,
			Defense:      50,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId2,
			Name:         "Wartoralla",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           40,
			Attack:       10,
			Defense:      30,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
	}

	db.InsertPokemons(fakeRandomPokemons)
	// Test search for a name
	r := e.GET("/pokemon").
	WithQuery("search", searchText).WithQuery("attack[gt]", "0").
		Expect().
		Status(http.StatusOK).JSON().Object()

	r.Value("pokemons").Array().Length().Equal(2)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("name", "Wartortle")
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	r.Value("pokemons").Array().Element(1).Object().ValueEqual("name", "Wartoralla")
	r.Value("pokemons").Array().Element(1).Object().ValueEqual("_id", fakeId2)

}


func TestResponseSortBasedOnEditDistance(t *testing.T) {
	defer db.DropDB()
	e := pokmonRouteTester(t)
	searchText:= "Wartortle"
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316d")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b3167")
	fakeId3, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316c")
	fakeRandomPokemons := []models.Pokemon{
		{
			ID:           fakeId1,
			Name:         "Wartortle",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           60,
			Attack:       50,
			Defense:      50,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId2,
			Name:         "Wartoralla",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           40,
			Attack:       10,
			Defense:      30,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId3,
			Name:         "Wartorallaaa",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           40,
			Attack:       10,
			Defense:      30,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
	}

	db.InsertPokemons(fakeRandomPokemons)
	r := e.GET("/pokemon").
	WithQuery("search", searchText).WithQuery("attack[gt]", "0").
		Expect().
		Status(http.StatusOK).JSON().Object()
 
	// Edit Distance is 0 since searchText is thesame as name
	r.Value("pokemons").Array().Length().Equal(3)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("name", "Wartortle")
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)

	// Edit Distance is 3 since it takes three operations to convert Wartortle to Wartoralla
	r.Value("pokemons").Array().Element(1).Object().ValueEqual("name", "Wartoralla")
	r.Value("pokemons").Array().Element(1).Object().ValueEqual("_id", fakeId2)

    // Edit Distance is 5 since it takes five operations to convert Wartortle to Wartorallaaa
	r.Value("pokemons").Array().Element(2).Object().ValueEqual("name", "Wartorallaaa")
	r.Value("pokemons").Array().Element(2).Object().ValueEqual("_id", fakeId3)
}


func TestAllFilters(t *testing.T) {
	defer db.DropDB()
	e := pokmonRouteTester(t)
	searchText:= "Wartortle"
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316d")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b3167")
	fakeId3, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316c")
	fakeRandomPokemons := []models.Pokemon{
		{
			ID:           fakeId1,
			Name:         "Wartortle",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           25,
			Attack:       50,
			Defense:      40,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId2,
			Name:         "Wartoralla",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           40,
			Attack:       60,
			Defense:      30,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId3,
			Name:         "Wartorallaaa",
			Type1:        "Fire",
			Type2:        "Grass",
			Total:        100,
			HP:           60,
			Attack:       40,
			Defense:      70,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
	}

	db.InsertPokemons(fakeRandomPokemons)
	r := e.GET("/pokemon").
	WithQuery("search", searchText).WithQuery("attack[gte]", "40").WithQuery("defense[lte]", "70").WithQuery("hp[gt]", "20").
		Expect().
		Status(http.StatusOK).JSON().Object()
 
	
	r.Value("pokemons").Array().Length().Equal(3)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("_id", fakeId1)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("name", "Wartortle")
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("hp", 25)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("attack", 50)
	r.Value("pokemons").Array().Element(0).Object().ValueEqual("defense", 40)


	r.Value("pokemons").Array().Element(1).Object().ValueEqual("_id", fakeId2)
	r.Value("pokemons").Array().Element(1).Object().ValueEqual("name", "Wartoralla")
	r.Value("pokemons").Array().Element(1).Object().ValueEqual("hp", 40)
	r.Value("pokemons").Array().Element(1).Object().ValueEqual("attack", 60)
	r.Value("pokemons").Array().Element(1).Object().ValueEqual("defense", 30)
   
	r.Value("pokemons").Array().Element(2).Object().ValueEqual("_id", fakeId3)
	r.Value("pokemons").Array().Element(2).Object().ValueEqual("name", "Wartorallaaa")
	r.Value("pokemons").Array().Element(2).Object().ValueEqual("hp", 60)
	r.Value("pokemons").Array().Element(2).Object().ValueEqual("attack", 40)
	r.Value("pokemons").Array().Element(2).Object().ValueEqual("defense", 70)
}
