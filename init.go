// init.go

package networkinterfaceinfo

import (
    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/inputs"
)

func init() {
    inputs.Add("network_interface_info", func() telegraf.Input {
        return &NetworkInterfaceInfo{}
    })
}
