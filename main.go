package main

import (
	"first-api/library"
	"net/http"
)

func main()  {
	fmt.Println("Name", library.Student.Name)
	fmt.Println("Grade", library.Student.Grade)
}
