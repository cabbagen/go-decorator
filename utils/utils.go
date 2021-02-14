package utils

import (
	"fmt"
	"crypto/md5"
)

func StringSliceToInterfaces(list []string) []interface{} {
	result := make([]interface{}, len(list))

	for index, value := range list {
		result[index] = value
	}
	return result
}

func SliceFind(list []interface{}, callback func (item interface{}, index int) bool) interface{} {
	for index, item := range list {
		if isExist := callback(item, index); isExist {
			return item
		}
	}
	return nil
}

func Md5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}
