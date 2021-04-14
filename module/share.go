package module

import (
  "fmt"
  "strings"
  "gopkg.in/xmlpath.v2"
  libvirt "libvirt-go"
)

func GetLibvirtConnect() libvirt.Connect {
  conn, err := libvirt.NewConnect("qemu:///system")
  if err != nil {
    fmt.Println("Fail to create libvirt connection.")
  }
  return *conn
}

func ParserXML(xml string, xpath string) []string {
  path := xmlpath.MustCompile(xpath)
  root, err := xmlpath.Parse(strings.NewReader(xml))
  var result []string
  if err == nil {
    values := path.Iter(root)
    for values.Next() {
      node := values.Node()
      result = append(result, node.String())
    }
    /* return first match as string
    if value, ok := path.String(root); ok {
      return value
    }
    */
  }
  return result
}

func GetListAllDomainsFlag(flag string) libvirt.ConnectListAllDomainsFlags {
  var fg libvirt.ConnectListAllDomainsFlags
  switch flag {
  case "active":
    fg = libvirt.CONNECT_LIST_DOMAINS_ACTIVE
  case "inactive":
    fg = libvirt.CONNECT_LIST_DOMAINS_INACTIVE
  case "running":
    fg = libvirt.CONNECT_LIST_DOMAINS_RUNNING
  case "paused":
    fg = libvirt.CONNECT_LIST_DOMAINS_PAUSED
  case "shutoff":
    fg = libvirt.CONNECT_LIST_DOMAINS_SHUTOFF
  default:
    fg = libvirt.CONNECT_LIST_DOMAINS_PERSISTENT
  }
  return fg
}

func GetDomainStateText(state libvirt.DomainState) string {
  var stateText string = "No State"
  switch state {
  case 0:
    stateText = "No State"
  case 1:
    stateText = "Running"
  case 2:
    stateText = "Blocked"
  case 3:
    stateText = "Paused"
  case 4:
    stateText = "Shutdown"
  case 5:
    stateText = "Shutoff"
  case 6:
    stateText = "Crashed"
  case 7:
    stateText = "Pmsuspended"
  default:
    stateText = "No State"
  }
  return stateText
}
