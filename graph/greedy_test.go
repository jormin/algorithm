package graph

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set"
)

// 测试贪婪算法
func TestGreedy(t *testing.T) {
	type args struct {
		states   mapset.Set
		stations map[string]mapset.Set
	}

	tests := []struct {
		name string
		args args
		want mapset.Set
	}{
		{
			name: "01",
			args: args{
				states: mapset.NewSetFromSlice([]interface{}{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}),
				stations: map[string]mapset.Set{
					"kone":   mapset.NewSetFromSlice([]interface{}{"id", "nv", "ut"}),
					"ktwo":   mapset.NewSetFromSlice([]interface{}{"wa", "id", "mt"}),
					"kthree": mapset.NewSetFromSlice([]interface{}{"or", "nv", "ca"}),
					"kfour":  mapset.NewSetFromSlice([]interface{}{"nv", "ut"}),
					"kfive":  mapset.NewSetFromSlice([]interface{}{"ca", "az"}),
				},
			},
			want: mapset.NewSetFromSlice([]interface{}{"kone", "ktwo", "kthree", "kfive"}),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				startTime := time.Now()
				res := greedy(tt.args.states, tt.args.stations)
				fmt.Printf("执行耗时：%v\n", time.Since(startTime))
				if !reflect.DeepEqual(res, tt.want) {
					t.Errorf("greedy() = %v, want %v", res, tt.want)
				}
			},
		)
	}
}
