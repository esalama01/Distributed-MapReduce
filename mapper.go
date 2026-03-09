package main

import(
	"strings"
)
type KeyValue struct{
	Key string
	Value string
}
func Mapper(input_string string)[]KeyValue{
	var my_list []KeyValue
	words := strings.Fields(input_string)
	for _, word := range words{
		new_object := KeyValue{
			Key : word,
			Value : "1",
		}
		my_list = append(my_list, new_object)
	}
	return my_list
}

