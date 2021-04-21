package bench

import (
	"fmt"
	"testing"
)

type Data struct {
	Name string
}

type DataMap map[string]*Data

type Flavor struct {
	Flavors FlavorMap
	Entry   *Data
}

type FlavorMap map[string]*Flavor

func gen(depths []int) *Flavor {
	depth := len(depths)

	if depth == 0 {
		return &Flavor{
			Entry: &Data{Name: "somedata"},
		}
	}

	flavor := &Flavor{
		Flavors: make(FlavorMap, depths[0]),
	}
	for d := 0; d < depths[0]; d++ {
		name := fmt.Sprintf("%df%d", depth, d+1)
		flavor.Flavors[name] = gen(depths[1:])
	}

	return flavor
}

func GetDataMap(f *Flavor) DataMap {
	fm := make(DataMap)
	if f == nil {
		return fm
	}

	return buildFlavor("", f)
}

func buildFlavor(s string, f *Flavor) DataMap {
	dm := make(DataMap)
	for name, flavor := range f.Flavors {
		if flavor.Entry != nil {
			dm[s+"/"+name] = flavor.Entry
			continue
		}
		for hName, hData := range buildFlavor(name, flavor) {
			// Depth-0
			if s == "" {
				dm[hName] = hData
				continue
			}
			dm[s+"/"+hName] = hData
		}
	}
	return dm
}

func BenchmarkGetDataMap(b *testing.B) {
	testCases := [][]int{
		{2, 10, 10, 10, 10},
		{10, 2, 10, 10, 10},
		{10, 10, 2, 10, 10},
		{10, 10, 10, 2, 10},
		{10, 10, 10, 10, 2},
		{2, 20, 20, 20, 20},
		{20, 2, 20, 20, 20},
		{20, 20, 2, 20, 20},
		{20, 20, 20, 2, 20},
		{20, 20, 20, 20, 2},
	}
	for _, tc := range testCases {
		name := ""
		for i, d := range tc {
			if i == 0 {
				name = fmt.Sprintf("%02d", d)
				continue
			}
			name = fmt.Sprintf("%s_%02d", name, d)
		}

		b.Run(fmt.Sprintf("%s", name), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				f := gen(tc)
				b.StartTimer()

				GetDataMap(f)
			}
		})
	}
}
