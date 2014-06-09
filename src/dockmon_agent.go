package main

import (
	"plugins/containers"
	"plugins/memoryMetrics"
	"encoding/json"
	"fmt"
	"plugins/cpuMetrics"
	"plugins/networkMetrics"
	"flag"
	"io/ioutil"
	"services"
	"strings"
)


type dataStruct struct {
	Memory		memoryMetrics.MemoryStruct
	Cpu			cpuMetrics.CpuStruct
	Info		containers.ContainerStruct
	Network		networkMetrics.NetworkStruct
}

type agentConf struct {
	DockerUrl			string
	CgroupDockerPrefix	string
}

func readConf() agentConf {
	var confFile string
	flag.StringVar(&confFile, "c", "/etc/dockmon_agent.conf", "Configuration file location (absolute path)")
	flag.Parse()
	var conf agentConf
	confJson, _ := ioutil.ReadFile(confFile)
	json.Unmarshal(confJson, &conf)
	return conf
}

func getJsonData() map[string] *dataStruct {
	conf := readConf()
	dockerService := &services.DockerService{Url: conf.DockerUrl}
	cgroupService := &services.CgroupService{DockerPrefix: conf.CgroupDockerPrefix}
	containersMap := containers.List(dockerService.List)
	name2Id := make(map[string]string, 0)
	data := make(map[string]*dataStruct)
	for id, names := range containersMap {
		data[id] = &dataStruct{}
		data[id].Info = containers.Grab(dockerService.Inspect, id)
		name2Id[data[id].Info.Name] = id
		for i := 0; i < len(names); i++ {
			name := names[i]
			if strings.LastIndex(name, "/") > 0 {
				parent := strings.Split(name, "/")
				data[id].Info.Parents = append(data[id].Info.Parents, parent[1])
			}
		}
		if data[id].Info.Running {
			data[id].Memory = memoryMetrics.Grab(cgroupService.GetMemLines, id)
			data[id].Cpu = cpuMetrics.Grab(cgroupService.GetCpuLines, id)
			data[id].Network = networkMetrics.Grab(cgroupService.GetNetworkLines, id)
		}
	}
	for _, value := range data {
		if value.Info.Parents != nil {
			for i := 0; i < len(value.Info.Parents); i++ {
				parentId := name2Id[value.Info.Parents[i]]
				data[parentId].Info.Children = append(data[parentId].Info.Children, value.Info.Name)
			}
		}
	}
	return data
}


func main() {
	jsonData, _ := json.Marshal(getJsonData())
	fmt.Println(string(jsonData))
}
