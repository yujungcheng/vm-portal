package module

import (
  "fmt"
  "time"
  "strings"
  "strconv"
  "gopkg.in/xmlpath.v2"
  libvirt "libvirt-go"
)

var StartTime time.Time
var ProcessID int

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

func Uptime(startTime time.Time) time.Duration {
  return time.Since(startTime)
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

func GetDomainStateStr(state libvirt.DomainState) string {
  var stateStr string = "No State"
  switch state {
  case 0:
    stateStr = "No State"
  case 1:
    stateStr = "Running"
  case 2:
    stateStr = "Blocked"
  case 3:
    stateStr = "Paused"
  case 4:
    stateStr = "Shutdown"
  case 5:
    stateStr = "Shutoff"
  case 6:
    stateStr = "Crashed"
  case 7:
    stateStr = "Pmsuspended"
  default:
    stateStr = "No State"
  }
  return stateStr
}

func GetStoragePoolStateStr(state libvirt.StoragePoolState) string {
  var stateStr string = "Unknown"
  switch state {
  case 0:
    stateStr = "Inactive"
  case 1:
    stateStr = "Building"
  case 2:
    stateStr = "Running"
  case 3:
    stateStr = "Degraded"
  case 4:
    stateStr = "Inaccessible"
  }
  return stateStr
}

func GetStorageVolTypeStr(volType libvirt.StorageVolType) string {
  var typeStr string = "Unknown"
  switch volType {
  case 0:
    typeStr = "File"
  case 1:
    typeStr = "Block"
  case 2:
    typeStr = "Dir"
  case 3:
    typeStr = "Network"
  case 4:
    typeStr = "Network Dir"
  case 5:
    typeStr = "Ploop"
  case 6:
    typeStr = "Last"  // unclear what this type is
  }
  return typeStr
}

func ConvertSizeToString(size uint64, unit string) string {
  var newSize uint64
  switch unit {
  case "KB":
    newSize = size/1024
  case "MB":
    newSize = size/1048576
  case "GB":
    newSize = size/1073741824
  default:
    newSize = size/1024
  }
  return strconv.FormatUint(newSize, 10)+unit
}
