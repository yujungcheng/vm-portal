package module

import (
  libvirt "libvirt-go"
)

type Volume struct {
  Device string
  Path string
  Type string
  Capacity uint64
  Allocation uint64
  Physical uint64
}

func GetDiskInfo(domain libvirt.Domain, disk string) {
//  blockInfo, err := domain.GetBlockInfo(disk, 0)

}
