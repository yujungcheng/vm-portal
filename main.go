package main

import (
  "fmt"
  "net/http"
  "html/template"
  mod "./module"
)


func overviewHandler(w http.ResponseWriter, r *http.Request) {
  // overview page

  tplFiles := []string {
    "template/portal.tpl",
    "template/base.tpl",
    "template/header.tpl",
    "template/footer.tpl",
    "template/overview.tpl",
  }
  tpl, err := template.ParseFiles(tplFiles...)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Internal Server Error", 500)
  }
  err = tpl.Execute(w, nil)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Internal Server Error", 500)
  }
}

func domainListHandler(w http.ResponseWriter, r *http.Request) {
  // active, inactive, running, paused, shutoff
  flag := "persistent"
  domains := mod.GetAllDomain(flag)
  tplFiles := []string {
    "template/portal.tpl",
    "template/base.tpl",
    "template/header.tpl",
    "template/footer.tpl",
    "template/domain_list.tpl",
  }
  tpl, err := template.ParseFiles(tplFiles...)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Internal Server Error", 500)
  }
  err = tpl.Execute(w, domains)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Internal Server Error", 500)
  }
}

func domainInfoHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", overviewHandler)
  mux.HandleFunc("/domain", domainListHandler)
  mux.HandleFunc("/domain/info", domainInfoHandler)
  http.ListenAndServe(":3000", mux)
}
