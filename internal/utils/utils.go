package utils

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"
)

func Contains(list interface{}, elem interface{}) bool {
	defer func() {
		recover()
	}()

	listV := reflect.ValueOf(list)

	if listV.Kind() == reflect.Slice {
		for i := 0; i < listV.Len(); i++ {
			item := listV.Index(i).Interface()

			target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()

			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}
	return false
}

func ExpandIntList(list []interface{}) []int {
	vi := make([]int, 0, len(list))
	for _, v := range list {
		if val, ok := v.(int); ok {
			vi = append(vi, val)
		}
	}
	return vi
}

func FlattenIntList(list []int) []interface{} {
	vi := make([]interface{}, 0, len(list))
	for _, v := range list {
		vi = append(vi, v)
	}
	return vi
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func StringToSlice(str string) []string {
	r := regexp.MustCompile(`[\[\]"]`)
	replacedStr := r.ReplaceAllString(str, "")

	if replacedStr == "" {
		return []string{}
	} else {
		return strings.Split(replacedStr, ", ")
	}
}
