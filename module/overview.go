package module

import (
  "fmt"
  "time"
  "strconv"
  "strings"
  libvirt "libvirt-go"
)

type Overview struct {
  PortalStatus string
  PortalVersion string
  PortalUptime time.Duration
  PortalPID int

  CpuModel string
  CpuMHz uint
  CpuCoreNum uint
  CpuCorePerSocket uint32
  CpuThreadPerCore uint32
  CpuSocketsNum uint32
  MemorySize uint64
  NumaCellNum uint32

  DomainNames []string
  StoragePools []string
  Networks []string

  TemplateNames []string
  ClusterNames []string
  BackupNames []string

  Hostname string
  KernelVersion string
  Distribution string
  LibvirtVersion string
}


func GetOverview() Overview {
  var conn = GetLibvirtConnect()
  defer conn.Close()

  overview := new(Overview)

  overview.PortalUptime = Uptime(StartTime)
  overview.PortalPID = ProcessID

  hostname, err := conn.GetHostname()
  if err != nil {
    fmt.Printf("Err: fail to get hostname")
  } else {
    overview.Hostname = hostname
  }

  libvirtVersion, err := conn.GetLibVersion()
  if err != nil {
    fmt.Printf("Err: fail to get libvirt version")
  } else {
    major := libvirtVersion/1000000
    minor := (libvirtVersion-(major*1000000))/1000
    release := (libvirtVersion-(major*1000000)-(minor*1009))
    overview.LibvirtVersion = fmt.Sprintf("%d.%d.%d", major, minor, release)
    //overview.LibvirtVersion = strconv.FormatUint(uint64(libvirtVersion), 10)
  }

  /* -------- node info -------- */
  nodeInfo, err := conn.GetNodeInfo()
  if err != nil {
    fmt.Printf("Err: fail to get node info")
  } else {
    overview.CpuModel = nodeInfo.Model
    overview.CpuMHz = nodeInfo.MHz
    overview.CpuCoreNum = nodeInfo.Cpus
    overview.CpuCorePerSocket = nodeInfo.Cores
    overview.CpuThreadPerCore = nodeInfo.Threads
    overview.CpuSocketsNum = nodeInfo.Sockets
    overview.MemorySize = nodeInfo.Memory
    overview.NumaCellNum = nodeInfo.Nodes
  }

  /* -------- domain name list -------- */
  domains, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
  if err != nil {
    fmt.Println("Error: fail to get all domains")
  } else {
    for _, domain := range domains {
      name, err := domain.GetName()
      if err != nil {
        fmt.Printf("Error: fail to get domain name")
        continue
      }
      state, reason, err := domain.GetState()
      if err != nil {
        fmt.Printf("Error: fail to get domain state")
        continue
      }
      _ = reason
      stateStr := GetDomainStateStr(state)
      fmt.Println("- " + name + " " + stateStr)
      overview.DomainNames = append(overview.DomainNames, name+" ("+stateStr+")")
    }
  }

  /* -------- storage pool info -------- */
  storagePools, err := conn.ListAllStoragePools(0)
  if err != nil {
    fmt.Printf("Error: fail to get all storage pool")
  } else {
    for _, storagePool := range storagePools {
      storagePoolName, err := storagePool.GetName()
      if err != nil {
        fmt.Printf("Error: fail to get storage pool name")
      }
      storagePoolInfo, err := storagePool.GetInfo()
      if err != nil {
        fmt.Printf("Error: fail to get storage pool info")
      }
      storagePoolIsActive, err := storagePool.IsActive()
      if err != nil {
        fmt.Printf("Error: fail to get storage pool active state")
        storagePoolIsActive = false
      }
      var storageNumOfVolumes int
      if storagePoolIsActive == true {
        storageNumOfVolumes, err = storagePool.NumOfStorageVolumes()
        if err != nil {
          fmt.Printf("Error: fail to get storage pool volume count")
        }
      }

      storagePoolState := GetStoragePoolStateStr(storagePoolInfo.State)
      storagePoolCapacity := ConvertSizeToString(storagePoolInfo.Capacity, "GB")
      storagePoolAllocation := ConvertSizeToString(storagePoolInfo.Allocation, "GB")
      storagePoolAvailable := ConvertSizeToString(storagePoolInfo.Available, "GB")

      var storagePoolPath string
      storagePoolXMLDesc, err := storagePool.GetXMLDesc(0)
      if err != nil {
        fmt.Printf("Error: fail to get storage pool XML Description")
      } else {
        path := ParserXML(storagePoolXMLDesc, "/pool/target/path")
        storagePoolPath = path[0]
      }

      fmt.Printf("%s - %s - %s - %s - %s - %d volumes - %s\n",
        storagePoolName,
        storagePoolState,
        storagePoolCapacity,
        storagePoolAllocation,
        storagePoolAvailable,
        storageNumOfVolumes,
        storagePoolPath)
/*
      pool_info := storagePoolName+" ("+
        storagePoolAllocation+"/"+
        storagePoolCapacity+") "+
        storagePoolPath+" Active="+
        strconv.FormatBool(storagePoolIsActive)+" #"+
        strconv.Itoa(storageNumOfVolumes)
*/
      pool_info_slice := []string{
        storagePoolName,
        "("+storagePoolAllocation+"/"+storagePoolCapacity+")",
        storagePoolPath,
        "Active="+strconv.FormatBool(storagePoolIsActive),
        "#"+strconv.Itoa(storageNumOfVolumes)}
      pool_info := strings.Join(pool_info_slice, " ")
      overview.StoragePools = append(overview.StoragePools, pool_info)

      /*
      storageVolumes, err := storagePool.ListAllStorageVolumes(0)
      if err != nil {
        fmt.Printf("Error: fail to get storage pool volumes")
      } else {
        for _, storageVolume := range storageVolumes {
          storageVolumeName, err := storageVolume.GetName()
          if err != nil {
            fmt.Printf("Error: fail to get storage pool volume name")
          }
          storageVolumeInfo, err := storageVolume.GetInfo()
          if err != nil {
            fmt.Printf("Error: fail to get storage pool volume info")
          }
          storageVolumeType := GetStorageVolTypeStr(storageVolumeInfo.Type)
          storageVolumeCapacity := ConvertSizeToString(storageVolumeInfo.Capacity, "GB")
          storageVolumeAllocation := ConvertSizeToString(storageVolumeInfo.Allocation, "GB")
          storageVolumePath, err := storageVolume.GetPath()
          if err != nil {
            fmt.Printf("Error: fail to get storage pool volume path")
          }
          fmt.Printf("%s - %s - %s - %s - %s\n", storageVolumeName,
                                                 storageVolumeType,
                                                 storageVolumeCapacity,
                                                 storageVolumeAllocation,
                                                 storageVolumePath)
        }
      }
      */
    }
  }

  /* -------- network info -------- */
  networks, err := conn.ListAllNetworks(0)
  if err != nil {
    fmt.Printf("Error: fail to get all networks")
  } else {
    //fmt.Println(networks)
    for _, network := range networks {
      networkName, err := network.GetName()
      if err != nil {
        fmt.Printf("Error: fail to get network name")
      }
      networkUUID, err := network.GetUUIDString()
      if err != nil {
        fmt.Printf("Error: fail to get network UUID")
      }
      bridgeName, err := network.GetBridgeName()
      if err != nil {
        fmt.Printf("Error: fail to get network bridge name")
      }
      var networkMAC []string
      networkXMLDesc, err := network.GetXMLDesc(0)
      if err != nil {
        fmt.Printf("Error: fail to get network XML Description")
      } else {
        networkMAC = ParserXML(networkXMLDesc, "/network/mac/@address")
        networkIP := ParserXML(networkXMLDesc, "/network/ip/@address")
        networkIPNetMask := ParserXML(networkXMLDesc, "/network/ip/@netmask")
        networkDHCPStart := ParserXML(networkXMLDesc, "/network/ip/dhcp/range/@start")
        networkDHCPEnd := ParserXML(networkXMLDesc, "/network/ip/dhcp/range/@end")
        fmt.Printf("- %s - %s - %s - %s - %s\n", networkMAC, networkIP,
          networkIPNetMask, networkDHCPStart, networkDHCPEnd)
      }
      /*
      networkDHCPLeases, err := network.GetDHCPLeases()
      if err != nil {
        fmt.Printf("Error: fail to get network DHCP leases")
      }
      fmt.Printf("%s\n", networkDHCPLeases)
      */

      /* Function 'virNetworkListAllPorts' not available in the libvirt library
       used during Go build'))
      networkAllPorts, err := network.ListAllPorts(0)
      if err != nil {
        fmt.Printf("Error: fail to get network ports", err)
      }*/

      fmt.Printf("%s - %s - %s\n", networkName, networkUUID, bridgeName)
      network_info := networkName+" ("+bridgeName+" ... "+networkMAC[0]+")"
      overview.Networks = append(overview.Networks, network_info)
    }
  }

  return *overview
}
