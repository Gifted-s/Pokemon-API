package customize_query_params_test

import (
	"pokemon/m/v1/helpers"
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)



func Test_Customze_Query_Params(t *testing.T) {
	assert := assert.New(t)
	queryParams := map[string][]string{}
	queryParams["hp[gte]"]= []string{"100"}
	queryParams["defense[lte]"]=  []string{"200"}
	queryParams["attack[eq]"]=  []string{"300"}
	queryParams["page"]=[]string{"1"}
	queryParams["search"]=[]string{"Meta"}
	customParams, err := helpers.CustomizeQueryParams(queryParams)
    expectedCustomparams := map[string][]string{
		"hp": {"$gte","100"},
		"defense": {"$lte","200"},
		"attack": {"$eq","300"},
		"page": {"1"},
		"search":{"Meta"},
	}
	areMapsEqual := reflect.DeepEqual(customParams, expectedCustomparams)
	assert.Equal(true, areMapsEqual, "The two maps must be equal")
	assert.Equal(nil, err, "No error should occur since params input is valid")

}

func Test_Customze_Query_Params_When_Search_Param_Is_Ommited(t *testing.T) {
	assert := assert.New(t)
	queryParams := map[string][]string{}
	queryParams["hp[gte]"]= []string{"100"}
	queryParams["defense[lte]"]=  []string{"200"}
	queryParams["attack[eq]"]=  []string{"300"}
	queryParams["page"]=[]string{"1"}
	// queryParams["search"]=[]string{"Meta"}
	customParams, err := helpers.CustomizeQueryParams(queryParams)
    expectedCustomparams := map[string][]string{
		"hp": {"$gte","100"},
		"defense": {"$lte","200"},
		"attack": {"$eq","300"},
		"page": {"1"},
		// "search":{"Meta"},
	}
	areMapsEqual := reflect.DeepEqual(customParams, expectedCustomparams)
	assert.Equal(true, areMapsEqual, "expectedCustomparams must not contain search param")
	assert.Equal(nil, err, "No error should occur since params input is valid")

}



func Test_Customze_Query_Params_When_Page_Is_Ommited(t *testing.T) {
	assert := assert.New(t)
	queryParams := map[string][]string{}
	queryParams["hp[gte]"]= []string{"100"}
	queryParams["defense[lte]"]=  []string{"200"}
	queryParams["attack[eq]"]=  []string{"300"}
	//queryParams["page"]=[]string{"1"}
    queryParams["search"]=[]string{"Meta"}
	customParams, err := helpers.CustomizeQueryParams(queryParams)
    expectedCustomparams := map[string][]string{
		"hp": {"$gte","100"},
		"defense": {"$lte","200"},
		"attack": {"$eq","300"},
		//"page": {"1"},
	    "search":{"Meta"},
	}
	areMapsEqual := reflect.DeepEqual(customParams, expectedCustomparams)
	assert.Equal(true, areMapsEqual, "expectedCustomparams must not contain page param")
	assert.Equal(nil, err, "No error should occur since params input is valid")

}


func Test_Customze_Query_Params_When_Operator_Is_Invalid(t *testing.T) {
	assert := assert.New(t)
	queryParams := map[string][]string{}
	queryParams["defense[wrong_operator]"]=  []string{"200"}
	queryParams["hp[gte]"]= []string{"100"}
	queryParams["attack[eq]"]=  []string{"300"}
	queryParams["page"]=[]string{"1"}
    queryParams["search"]=[]string{"Meta"}
	customParams, err := helpers.CustomizeQueryParams(queryParams)
	assert.Equal(err.Error(), "wrong_operator" + ": is not a valid operator please select operator from these [gte | lte | gt | lt | eq | ne ]", "expects an error message")
	assert.Equal(map[string][]string(map[string][]string(nil)), customParams, "customParams should be nil since input was invalid")

}


func Test_Customze_Query_Params_When_Page_Type_Is_Invalid(t *testing.T) {
	assert := assert.New(t)
	queryParams := map[string][]string{}
	queryParams["page"]=[]string{"wrong_pagination"}
	queryParams["defense[lte]"]=  []string{"200"}
	queryParams["hp[gte]"]= []string{"100"}
	queryParams["attack[eq]"]=  []string{"300"}
    queryParams["search"]=[]string{"Meta"}
	customParams, err := helpers.CustomizeQueryParams(queryParams)
	assert.Equal(err.Error(), "page value is invalid. page type should be of type int e.g 1")
	assert.Equal(map[string][]string(map[string][]string(nil)), customParams, "customParams should be nil since input was invalid")
}
