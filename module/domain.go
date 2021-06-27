package module

import (
  "fmt"
  "time"
  libvirt "libvirt-go"
)

type AllDomain struct {
  CheckedAt time.Time  // check time
  Domains []Domain
}

type Domain struct {
  Name string
  ID uint
  UUID string
  State libvirt.DomainState
  StateStr string
  MaxMem uint64
  MaxMemStr string
  Memory uint64
  MemoryStr string
  Vcpu uint
  CpuTime uint64
  Tags []string  // tags marker
  Disks []string  // disk device name
  Interfaces []string  // interface MAC address & connected Network
  Events []string  // refere to log module, activities history
}


/* ------------------------------------------------------------------------ */

func GetAllDomain(flag string) AllDomain {
  /* flag value:
     active, inactive, running, paused, shutoff
  */
  var conn = GetLibvirtConnect()
  defer conn.Close()

  /* Get All domain object
  ---------------------------------------- */
  fg := GetListAllDomainsFlag(flag)
  //domains, err := conn.ListDomains()  // only active domains
  domains, err := conn.ListAllDomains(fg)
  if err != nil {
    fmt.Println("Error: fail to get domains")
  }
  allDomain := new(AllDomain)
  allDomain.CheckedAt = time.Now()

  /* Gat all domain data
  ---------------------------------------- */
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
    fmt.Printf("- %s\n", name)

    dom := new(Domain)  // Create domain object
    dom.Name = name
    dom.UUID = uuid
    dom.State = info.State
    dom.MaxMem = info.MaxMem
    dom.Memory = info.Memory
    dom.MemoryStr = ConvertSizeToString(info.Memory*1024, "MB")
    dom.Vcpu = info.NrVirtCpu
    dom.CpuTime = info.CpuTime
    dom.StateStr = GetDomainStateStr(dom.State)  // Set readable status
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
    //diskPaths :=  ParserXML(xml, "/domain/devices/disk/source/@file")
    diskPaths :=  ParserXML(xml, "/domain/devices/disk[@device='disk']/source/@file")
    //for i, diskPath := range diskPaths { fmt.Printf("%d, %s\n", i, diskPath) }

    var disksInfo []string
    var domDiskPaths map[string]string
    domDiskPaths = make(map[string]string)
    for i, disk := range dom.Disks {
      // Set disk path
      path := diskPaths[i]
      domDiskPaths[disk] = path
      //fmt.Printf("%s %s\n", disk, path)

      blockInfo, err := domain.GetBlockInfo(disk, 0)
      if err != nil {
        fmt.Printf("Error: fail to get BlockInfo")
      } else {
        capacityStr := ConvertSizeToString(blockInfo.Capacity, "GB")
        /* Unused
        allocationStr := ConvertSizeToString(blockInfo.Allocation, "GB")
        physicalStr := ConvertSizeToString(blockInfo.Physical, "GB")
        */
        disksInfo = append(disksInfo, disk+"("+capacityStr+")")
        //fmt.Printf("%s\n", disksInfo)
      }
    }
    dom.Disks = disksInfo // update disks data

    var interfaces []string
    macs := ParserXML(xml, "/domain/devices/interface[@type='network']/mac/@address")
    nets := ParserXML(xml, "/domain/devices/interface[@type='network']/source/@network")
    models := ParserXML(xml, "/domain/devices/interface[@type='network']/model/@type")
    for i, mac := range macs {
      interfaces = append(interfaces, mac+" "+models[i]+" "+nets[i])
    }
    dom.Interfaces = interfaces

    allDomain.Domains = append(allDomain.Domains, *dom)
    domain.Free()
    //fmt.Printf("\n")
  }
  return *allDomain
}


func GetDomainDetail(uuid string) Domain {
  domain := new(Domain)

  var conn = GetLibvirtConnect()
  defer conn.Close()

  dom, err := conn.LookupDomainByUUIDString(uuid)
  if err != nil {
    fmt.Print("Error: fail to LookupDomainByUUIDString")
  }
  fmt.Printf("%s, %T\n", uuid, dom)

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
  return *domain
}
