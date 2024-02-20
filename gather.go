// gather.go

package networkinterfaceinfo

// 导入必要的包

func (ni *NetworkInterfaceInfo) Gather(acc telegraf.Accumulator) error {
    // 实现数据收集逻辑
}

func (ni *NetworkInterfaceInfo) shouldGather(interfaceName string) bool {
    // 实现判断接口是否被采集的逻辑
}
