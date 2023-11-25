package main

import (
	"fmt"
	"reflect"
)

//Struct tags add metadata to fields within a struct

type Animal struct {
	Name string `required max_len:10`
	Age  int
}

func main() {
	fmt.Println("Struct Tags")

	a := Animal{"Cat", 12}
	fmt.Printf("Type = %T, Value = %+v \n", a, a)

	t := reflect.TypeOf(a)
	fmt.Printf("Type = %T, Value = %+v \n", t, t)

	field, _ := t.FieldByName("Name")
	fmt.Printf("Type = %T, Value = %+v \n", field, field)
	fmt.Printf("Type = %T, Value = %+v \n", field.Tag, field.Tag)

	//===========================================================

	encodingJson()
}

/*
====================================================================
OUTPUT
====================================================================
Struct Tags
Type = main.Animal, Value = {Name:Cat Age:12}
Type = *reflect.rtype, Value = main.Animal
Type = reflect.StructField, Value = {Name:Name PkgPath: Type:string Tag:required max_len:10 Offset:0 Index:[0] Anonymous:false}
Type = reflect.StructTag, Value = required max_len:10

from encodingJson function
Struc Tags : encoding json
{"filed1":"v1","filed2":"","filed3":"v3","filed4":"v4"}

====================================================================
*/
