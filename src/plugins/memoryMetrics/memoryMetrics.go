package memoryMetrics

import (
	"strings"
	"strconv"
)

type MemoryStruct struct {
	Total	int64
	Rss		int64
	Cache	int64
	Swap	int64
}

func Grab(dataSource func(string) []string, containerId string) MemoryStruct {
	lines := dataSource(containerId)
	total := int64(-1)
	rss := int64(-1)
	cache := int64(-1)
	swap := int64(-1)
	for i := 0; i < len(lines); i++ {
		elements := strings.Split(lines[i], " ")
		switch elements[0] {
		case "cache" :
			cache, _ = strconv.ParseInt(elements[1], 10, 64)
		case "rss" :
			rss, _ = strconv.ParseInt(elements[1], 10, 64)
		case "swap" :
			swap, _ = strconv.ParseInt(elements[1], 10, 64)
		case "hierarchical_memory_limit" :
			total, _ = strconv.ParseInt(elements[1], 10, 64)
		}
	}
	return MemoryStruct{total, rss, cache, swap}
}
