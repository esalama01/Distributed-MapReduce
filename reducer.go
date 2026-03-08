package main

import(
	"strconv"
)

func Reducer(key string, values []string)string{
	n := len(values)
	str := strconv.Itoa(n)
	return str
}