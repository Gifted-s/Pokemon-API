package helpers

import (
	"strings"
)

func CustomizeQueryParams(queryParams map[string][]string) map[string][]string {
	customParams := map[string][]string{}
	for k, v := range queryParams {
		key := ""
		value := []string{}
		index1 := strings.Index(k, "[")
		if index1 != -1 {
			key = k[0:index1]
			value = []string{
				k[index1+1 : len(k)-1],
				v[0],
			}
		} else if k == "page" {
			key = k
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

	return customParams
}
