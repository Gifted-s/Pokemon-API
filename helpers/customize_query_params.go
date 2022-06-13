package helpers

import (
	"errors"
	"pokemon/m/v1/utils"
	"strings"
	"strconv"
)

func CustomizeQueryParams(queryParams map[string][]string) (map[string][]string, error) {
	customParams := map[string][]string{}
	for k, v := range queryParams {
		key := ""
		value := []string{}
		index1 := strings.Index(k, "[")
		if index1 != -1 {
			key = k[0:index1]
			operator := k[index1+1 : len(k)-1]
			if customizedKey, ok := utils.URLOperatorToMongoDBOperatorMatcher()[operator]; ok {
				value = []string{
					customizedKey,
					v[0],
				}
			} else {
				return nil, errors.New(operator + ": is not a valid operator please select operator from these [gte | lte | gt | lt | eq | ne ]")
			}

		} else if k == "page" {
			key = k
			_, err := strconv.Atoi(v[0])
			if(err != nil){
				return nil, errors.New("page value is invalid. page type should be of type int e.g 1")
			}
			value = []string{
				v[0],
			}
		} else if k == "search" {
			key = k
			value = []string{
				v[0],
			}
		}
		customParams[key] = value
	}

	return customParams, nil
}
