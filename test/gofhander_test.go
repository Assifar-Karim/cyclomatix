package fs_explorer

import (
	"reflect"
	"testing"

	fsinfo "github.com/Assifar-Karim/cyclomatix/internal/fctinfo"
	"github.com/Assifar-Karim/cyclomatix/internal/fsexplorer"
)

type TestInput struct {
	path     string
	fctTable *[]fsinfo.FctInfo
}

type Test struct {
	name  string
	input TestInput
	want  []map[int32][]int32
}

func TestHandleFile(t *testing.T) {
	tests := []Test{
		{
			"Basic Functions",
			TestInput{
				path:     "../examples/basic.go",
				fctTable: &[]fsinfo.FctInfo{},
			},
			[]map[int32][]int32{
				{0: []int32{1}, 1: []int32{2}, 2: []int32{}},
				{0: []int32{1}, 1: []int32{2}, 2: []int32{}},
			},
		},
		{
			"Conditional Functions 1",
			TestInput{
				path:     "../examples/if.go",
				fctTable: &[]fsinfo.FctInfo{},
			},
			[]map[int32][]int32{
				{0: []int32{1}, 1: []int32{2, 4}, 2: []int32{3}, 3: []int32{4},
					4: []int32{5}, 5: []int32{}},
				{0: []int32{1}, 1: []int32{2, 4}, 2: []int32{3}, 3: []int32{6},
					4: []int32{5}, 5: []int32{6}, 6: []int32{7}, 7: []int32{}},
				{0: []int32{1}, 1: []int32{2, 4}, 2: []int32{3}, 3: []int32{11},
					4: []int32{5, 8}, 5: []int32{6}, 6: []int32{7}, 7: []int32{10},
					8: []int32{9}, 9: []int32{10}, 10: []int32{11}, 11: []int32{12},
					12: []int32{13}, 13: []int32{}},
				{0: []int32{1}, 1: []int32{2, 3}, 2: []int32{5}, 3: []int32{4},
					4: []int32{5}, 5: []int32{}},
				{0: []int32{1}, 1: []int32{2}, 2: []int32{3}, 3: []int32{4, 5},
					4: []int32{5}, 5: []int32{6}, 6: []int32{}},
			},
		},
		{
			"Conditional Functions 2",
			TestInput{
				path:     "../examples/switch.go",
				fctTable: &[]fsinfo.FctInfo{},
			},
			[]map[int32][]int32{
				{0: []int32{1}, 1: []int32{2, 4, 6, 8, 10, 12, 14, 16},
					2: []int32{3}, 3: []int32{19}, 4: []int32{5},
					5: []int32{19}, 6: []int32{7}, 7: []int32{19},
					8: []int32{9}, 9: []int32{19}, 10: []int32{11},
					11: []int32{19}, 12: []int32{13}, 13: []int32{19},
					14: []int32{15}, 15: []int32{19}, 16: []int32{17},
					17: []int32{19}, 18: []int32{19}, 19: []int32{}},
				{0: []int32{1}, 1: []int32{2}, 2: []int32{3}, 3: []int32{4, 6, 8, 10},
					4: []int32{5}, 5: []int32{12}, 6: []int32{7}, 7: []int32{12},
					8: []int32{9}, 9: []int32{12}, 10: []int32{11},
					11: []int32{12}, 12: []int32{13}, 13: []int32{}},
			},
		},
		{
			"Iterative Functions",
			TestInput{
				path:     "../examples/for.go",
				fctTable: &[]fsinfo.FctInfo{},
			},
			[]map[int32][]int32{
				{0: []int32{1}, 1: []int32{2, 3}, 2: []int32{3},
					3: []int32{1, 4}, 4: []int32{}},
				{0: []int32{1}, 1: []int32{2}, 2: []int32{3, 5},
					3: []int32{4}, 4: []int32{2}, 5: []int32{6}, 6: []int32{}},
				{0: []int32{1}, 1: []int32{2, 6}, 2: []int32{3, 4},
					3: []int32{6}, 4: []int32{5}, 5: []int32{6},
					6: []int32{1, 7}, 7: []int32{}},
				{0: []int32{1}, 1: []int32{2, 6}, 2: []int32{3, 4},
					3: []int32{1}, 4: []int32{5}, 5: []int32{6},
					6: []int32{1, 7}, 7: []int32{}},
				{0: []int32{1}, 1: []int32{2, 6}, 2: []int32{3, 4},
					3: []int32{8}, 4: []int32{5}, 5: []int32{6},
					6: []int32{1, 7}, 7: []int32{8}, 8: []int32{}},
				{0: []int32{1}, 1: []int32{2}, 2: []int32{3, 4},
					3: []int32{4}, 4: []int32{2, 5}, 5: []int32{}},
			},
		},
	}
	g := fsexplorer.NewGoFileHandler(4)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g.HandleFile(tt.input.path, tt.input.fctTable)
		})
		ft := *tt.input.fctTable
		for i, fs := range ft {
			ans := fs.GetCfg().AdjList
			if !reflect.DeepEqual(ans, tt.want[i]) {
				t.Errorf("got %v, want %v", ans, tt.want[i])
			}
		}
	}
}
