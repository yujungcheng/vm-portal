package main

import (
	mod "./module"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

func overviewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handler - overview")
	overview := mod.GetOverview()
	tplFiles := []string{
		"template/portal.tpl",
		"template/base.tpl",
		"template/overview.tpl",
	}
	tpl, err := template.ParseFiles(tplFiles...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
	err = tpl.Execute(w, overview)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func domainListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handler - domain list")
	flag := "persistent" // active, inactive, running, paused, shutoff
	domains := mod.GetAllDomain(flag)
	tplFiles := []string{
		"template/portal.tpl",
		"template/base.tpl",
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

func imageListHandler(w http.ResponseWriter, r *http.Request) {
}

func imageInfoHandler(w http.ResponseWriter, r *http.Request) {
}

func volumeListHandler(w http.ResponseWriter, r *http.Request) {
}

func volumeInfoHandler(w http.ResponseWriter, r *http.Request) {
}

func networkListHandler(w http.ResponseWriter, r *http.Request) {
}

func networkInfoHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	mod.StartTime = time.Now()
	mod.ProcessID = os.Getpid()

	mux := http.NewServeMux()
	mux.HandleFunc("/", overviewHandler)
	mux.HandleFunc("/domain", domainListHandler)
	mux.HandleFunc("/domain/info", domainInfoHandler)
	http.ListenAndServe(":3000", mux)
}
