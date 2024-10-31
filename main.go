package main

import (
  "net/http"
  "log"
  "html/template"
  "fmt"
)

func main()  {

  http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles/"))))
  http.HandleFunc("/", home)
  // server starts on port 80
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("ahoy your server is running ")
}

//home function handle requests to home package
func home(w http.ResponseWriter, r *http.Request)  {
  renderTemplate(w, "home.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string)  {
  // parsing the specified template file as input 
  t, err := template.ParseFiles("template/" + tmpl)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  t.Execute(w, nil)
}
