package cpuMetrics

import (
	"strings"
	"strconv"
)

type CpuStruct struct {
	User	int64
	System	int64
}

func Grab(dataSource func(string) []string, containerId string) CpuStruct {
	lines := dataSource(containerId)
	user := int64(-1)
	system := int64(-1)
	for i := 0; i < len(lines); i++ {
		elements := strings.Split(lines[i], " ")
		switch elements[0] {
		case "user" :
			user, _ = strconv.ParseInt(elements[1], 10, 64)
		case "system" :
			system, _ = strconv.ParseInt(elements[1], 10, 64)
		}
	}

	return CpuStruct{User: user, System: system}
}
