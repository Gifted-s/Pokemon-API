package operator_matcher_test

import (
	"pokemon/m/v1/helpers"
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)



func Test_Operator_Matcher(t *testing.T) {
	assert := assert.New(t)
	
	customParams := helpers.URLOperatorToMongoDBOperatorMatcher()
    expectedCustomparams := map[string]string{
	    "gte": "$gte",
		"lte": "$lte",
		"gt": "$gt",
		"lt": "$lt",
		"eq": "$eq",
		"ne": "$ne",
	}
	areMapsEqual := reflect.DeepEqual(customParams, expectedCustomparams)
	assert.Equal(true, areMapsEqual, "The two maps must be equal")
}
