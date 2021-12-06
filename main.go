package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "apa kabar!")
}

func main()  {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		var data = map[string]string{
			"Name": "Abi F",
			"Message": "thank you",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil{
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("starting listening ")
	http.ListenAndServe(":8000", nil)

}
