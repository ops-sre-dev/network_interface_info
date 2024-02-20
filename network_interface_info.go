// network_interface_info.go

package networkinterfaceinfo

import (
    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/inputs"
)

type NetworkInterfaceInfo struct {
    Interfaces []string `toml:"interfaces"`
    IncludeIPv6 bool `toml:"include_ipv6"`
    ExcludeInterfaces []string `toml:"exclude_interfaces"`
}

func (_ *NetworkInterfaceInfo) Description() string {
    // 实现描述方法
}

func (_ *NetworkInterfaceInfo) SampleConfig() string {
    // 实现示例配置方法
}

// 其他方法实现...

