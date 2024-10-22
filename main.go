package main

import (
  "net/http"
  "log"
  "html/template"
  "fmt"
)

func main()  {
  http.HandleFunc("/", home)
}

//home function handle requests to home package
func home(w http.ResponseWritter, r *http.Request)  {
  renderTemplate(w, "home.html")
}

func renderTemplate(w http.ResponseWritter, tmpl string)  {
  // parsing the specified template file as input 
  t,err := template.ParseFiles("template/" + tmpl)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  t.Execute(w, nil)
}
