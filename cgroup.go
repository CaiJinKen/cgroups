package cgroups

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type Cgroup struct {
	Name        string
	Path        string
	CPUSet      string
	CPURate     string
	Memory      string
	TcpMemory   string
	DevicesDeny string
	DeviceAllow string
	BlkReadBps  string
	BlkWriteBps string
}

var cgroupPath = "/sys/fs/cgroup"

func (c *Cgroup) SetCPUNum() error {
	return c.basePath("cpuset", "cpuset.cpus", []byte(c.CPUSet))
}

func (c *Cgroup) SetCPURate() error {
	return c.basePath("cpu", "cpu.cfs_quota_us", []byte(c.CPURate))
}

func (c *Cgroup) SetMemory() error {
	return c.basePath("memory", "memory.limit_in_bytes", []byte(c.Memory))
}

func (c *Cgroup) SetTcpMemory() error {
	return c.basePath("memory", "memory.kmem.tcp.limit_in_bytes", []byte(c.Memory))
}

func (c *Cgroup) SetDeviceDeny() error {
	return c.basePath("devices", "devices.deny", []byte(c.DevicesDeny))
}

func (c *Cgroup) SetDeviceAllow() error {
	return c.basePath("devices", "devices.allow", []byte(c.DeviceAllow))
}

func (c *Cgroup) SetBlkReadBps() error {
	return c.basePath("blkio", "blkio.throttle.read_bps_device", []byte(c.BlkReadBps))
}

func (c *Cgroup) SetBlkWriteBps() error {
	return c.basePath("blkio", "blkio.throttle.write_bps_device", []byte(c.BlkWriteBps))
}

func (c *Cgroup) basePath(subp, f string, data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}
	var path string
	if c.Path == "" {
		path = cgroupPath
	}
	cp := filepath.Join(path, subp, c.Name)
	if err := checkPath(cp); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(cp, f), data, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(cp, "tasks"), []byte(strconv.Itoa(os.Getpid())), 0600); err != nil {
		return err
	}
	return nil
}
