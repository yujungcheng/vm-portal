package module

import (
  "fmt"
  "time"
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

  StoragePools []string
  Networks []string
  DomainNames []string

  ImageNames []string
  TemplateNames []string

  CpuCoreUsed uint
  MemoryUsed uint
  StorageUsed []string
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
  //fmt.Printf("StartTime: %s", StartTime)
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
    libvirtVersion = 4001020
    major := libvirtVersion/1000000
    minor := (libvirtVersion-(major*1000000))/1000
    release := (libvirtVersion-(major*1000000)-(minor*1009))
    //fmt.Printf("\n %d %d %d\n", major, minor, release)
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
      fmt.Println("- " + name)
      overview.DomainNames = append(overview.DomainNames, name)
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
      storagePoolState := GetStoragePoolStateStr(storagePoolInfo.State)
      storagePoolCapacity := ConvertSizeToString(storagePoolInfo.Capacity, "GB")
      storagePoolAllocation := ConvertSizeToString(storagePoolInfo.Allocation, "GB")
      storagePoolAvailable := ConvertSizeToString(storagePoolInfo.Available, "GB")
      fmt.Printf("%s - %s - %s - %s - %s\n", storagePoolName,
                                             storagePoolState,
                                             storagePoolCapacity,
                                             storagePoolAllocation,
                                             storagePoolAvailable)

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

      storageNumOfVolumes, err := storagePool.NumOfStorageVolumes()
      fmt.Printf("number of volumes: %s\n", storageNumOfVolumes)
      pool_info := storagePoolName+"("+storagePoolAllocation+"/"+storagePoolCapacity+")"
      overview.StoragePools = append(overview.StoragePools, pool_info)
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
      bridgeName, err := network.GetBridgeName()
      if err != nil {
        fmt.Printf("Error: fail to get bridge name")
      }
      fmt.Printf("%s - %s\n", networkName, bridgeName)

      network_info := networkName+"("+bridgeName+")"
      overview.Networks = append(overview.Networks, network_info)
    }
  }

  return *overview
}
