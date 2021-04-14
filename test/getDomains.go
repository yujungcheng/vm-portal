package main

import (
  //"os"
  "fmt"
  //"log"
  libvirt "libvirt-go"
  "html/template"
  "net/http"
)

type Domain struct {
  Name string
  State libvirt.DomainState
  MaxMem uint64
  Memory uint64
  Vcpu uint
  CpuTime uint64
}

type AllDomain struct {
  Group string
  Domains []Domain
}

func getActiveDomain() AllDomain {
  conn, err := libvirt.NewConnect("qemu:///system")
  if err != nil {
    fmt.Println("Fail to create libvirt connection.")
  }
  defer conn.Close()

  domains, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
  if err != nil {
    fmt.Println("Fail to get active domains.")
  }
  allDomain := new(AllDomain)
  allDomain.Group = "default"

  for _, domain := range domains {
    name, err := domain.GetName()
    if err != nil {
      fmt.Printf("Error to get domain name")
      continue
    }
    info, err := domain.GetInfo()
    if err != nil {
      fmt.Printf("Error to get domain info")
      continue
    }

    dom := new(Domain)
    dom.Name = name
    dom.State = info.State
    dom.MaxMem = info.MaxMem
    dom.Memory = info.Memory
    dom.Vcpu = info.NrVirtCpu
    dom.CpuTime = info.CpuTime

    allDomain.Domains = append(allDomain.Domains, *dom)
    domain.Free()
    fmt.Printf("%s\n", name)
  }
  return *allDomain
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  domains := getActiveDomain()
  tpl := template.Must(template.ParseFiles("../template/domain_list.html"))
  tpl.Execute(w, domains)
}


func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", indexHandler)
  http.ListenAndServe(":3000", mux)
}
