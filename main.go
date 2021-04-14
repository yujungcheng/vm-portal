package main

import (
  "net/http"
  "html/template"
  mod "./module"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  // overview page

}

func domainListHandler(w http.ResponseWriter, r *http.Request) {
  // active, inactive, running, paused, shutoff
  flag := "persistent"
  //domains := mod.GetAllDomain(flag)
  domains := mod.GetAllDomainDetail(flag)  
  tpl := template.Must(template.ParseFiles("template/domain_list.html"))
  tpl.Execute(w, domains)
}

func domainListDetailHandler(w http.ResponseWriter, r *http.Request) {

}



func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", indexHandler)
  mux.HandleFunc("/domain", domainListHandler)
  mux.HandleFunc("/domain/detail", domainListDetailHandler)
  http.ListenAndServe(":3000", mux)
}
