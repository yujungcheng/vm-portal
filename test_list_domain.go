package main

import (
  "fmt"
  libvirt "libvirt-go"
)

func main() {

  conn, err := libvirt.NewConnect("qemu:///system")
  if err != nil {
    fmt.Println("Fail to create libvirt connection.")
  }
  defer conn.Close()

  doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
  if err != nil {
    fmt.Println("Fail to get active domains.")
  }
  fmt.Printf("%d running domains:\n", len(doms))
  for _, dom := range doms {
    name, err := dom.GetName()
    if err == nil {
      fmt.Printf("  %s\n", name)
    }
    dom.Free()
  }

  doms, err = conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
  if err != nil {
    fmt.Println("Fail to get inactive domains.")
  }
  fmt.Printf("%d stop domains:\n", len(doms))
  for _, dom := range doms {
    name, err := dom.GetName()
    if err == nil {
      fmt.Printf("  %s\n", name)
    }
    dom.Free()
  }

}
