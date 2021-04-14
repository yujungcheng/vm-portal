package module

import (
  "fmt"
  libvirt "libvirt-go"
)

type AllDomain struct {
  Cluster string
  Domains []Domain
}

type Domain struct {
  Name string
  ID uint
  UUID string
  State libvirt.DomainState
  StateText string
  MaxMem uint64
  Memory uint64
  Vcpu uint
  CpuTime uint64
  Tags []string  // tags marker
  Disks []string  // disk device name
  Interfaces []string  // interface MAC address & connected Network
  Events []string  // refere to log module, activities history
  Detail DomainDetail
}

type DomainDetail struct {
  Volume []Volume
}

func GetAllDomain(flag string) AllDomain {
  /* flag value:
     active, inactive, running, paused, shutoff
  */
  var conn = GetLibvirtConnect()
  defer conn.Close()

  fg := GetListAllDomainsFlag(flag)
  //domains, err := conn.ListDomains()  // only active domains
  domains, err := conn.ListAllDomains(fg)
  if err != nil {
    fmt.Println("Error: fail to get domains.")
  }
  allDomain := new(AllDomain)
  allDomain.Cluster = "default"  // todo: groupping domans

  for _, domain := range domains {
    uuid, err := domain.GetUUIDString()
    if err != nil {
      fmt.Print("Error: fail to get domain UUID string")
      continue
    }
    name, err := domain.GetName()
    if err != nil {
      fmt.Printf("Error: fail to get domain name")
      continue
    }
    info, err := domain.GetInfo()
    if err != nil {
      fmt.Printf("Error: fail to get domain info")
      continue
    }
    //fmt.Printf("- %s\n", name)
    dom := new(Domain)  // Create domain object
    dom.Name = name
    dom.UUID = uuid
    dom.State = info.State
    dom.MaxMem = info.MaxMem
    dom.Memory = info.Memory
    dom.Vcpu = info.NrVirtCpu
    dom.CpuTime = info.CpuTime
    dom.StateText = GetDomainStateText(dom.State)  // Set readable status
    if (dom.State == 1 || dom.State == 3) {  // Set domain ID (only poweron VM)
      id, err := domain.GetID()
      if err == nil {
        dom.ID = id
      } else {
        fmt.Print("Error: fail to get domain ID")
      }
    }

    xml, err := domain.GetXMLDesc(0)  // dump domain xml

    //devNames = ParserXML(xml, "/domain/devices/disk/target/@dev")
    dom.Disks = ParserXML(xml, "/domain/devices/disk[@device='disk']/target/@dev")
    //for i, diskName := range diskNames {  fmt.Printf("%d, %s\n", i, diskName) }

    //diskPaths =  ParserXML(xml, "/domain/devices/disk/source/@file")
    //diskPaths =  ParserXML(xml, "/domain/devices/disk[@device='disk']/source/@file")
    //for i, diskPath := range diskPaths { fmt.Printf("%d, %s\n", i, diskPath) }

    allDomain.Domains = append(allDomain.Domains, *dom)
    domain.Free()
    fmt.Printf("- %s\n", name)
  }
  return *allDomain
}


func GetAllDomainDetail(flag string) AllDomain {
  allDomain := GetAllDomain(flag)

  var conn = GetLibvirtConnect()
  defer conn.Close()

  for _, domain := range allDomain.Domains {
    //fmt.Printf("%T %T", i, domain)
    dom, err := conn.LookupDomainByUUIDString(domain.UUID)
    if err != nil {
      fmt.Print("Error: fail to LookupDomainByUUIDString")
    }
    //fmt.Printf("%d, %T, %T\n", i, domain.UUID, dom)
    for _, disk := range domain.Disks {
      blockInfo, err := dom.GetBlockInfo(disk, 0)
      if err != nil {
        fmt.Printf("Error: fail to get BlockInfo")
      } else {
        capacity := blockInfo.Capacity
        allocation := blockInfo.Allocation
        physical := blockInfo.Physical
        fmt.Printf("%d | %d | %d\n", capacity, allocation, physical)
      }
    }

  }

  /*
  domain.GetAutostart()
  domain.GetBlkioParameters()
  domain.GetBlockInfo()
  domain.GetBlockIoTunne()
  domain.GetBlockJobInfo()
  domain.GetCPUStats()
  domain.GetControlInfo()
  domain.GetDiskErrors()
  domain.GetEmulatorPinInfo()
  domain.GetFSInfo()
  domain.GetGuestInfo()
  domain.GetGuestVcpus()
  domain.GetHostname()
  domain.GetIOThreadInfo()
  domain.GetInterfaceParameters()
  domain.GetJobInfo()
  domain.GetJobStats()
  domain.GetLaunchSecurityInfo()
  domain.GetMaxMemory()
  domain.GetMaxVcpus()
  domain.GetMemoryParameters()
  domain.GetMetadata()
  domain.GetName()
  domain.GetNumaParameters()
  domain.GetOSType()
  domain.GetPerfEvents()
  domain.GetSchedulerParameters()
  domain.GetSchedulerParametersFlags()
  domain.GetSecurityLabel()
  domain.GetSecurityLabelList()
  domain.GetState()
  domain.GetTime()
  domain.GetUUID()
  domain.GetVcpuPinInfo()
  domain.GetVcpus()
  domain.GetVcpusFlags()
  domain.GetXMLDesc()

  dom.HasCurrentSnapshot()
  dom.HasManagedSaveImage()

  */

  return allDomain
}
