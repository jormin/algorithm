package graph

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 测试狄克斯特拉算法
func TestDijkstra(t *testing.T) {
	type args struct {
		graph   map[string]map[string]int
		costs   map[string]int
		parents map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "01",
			args: args{
				graph: map[string]map[string]int{
					"start": {
						"a": 6,
						"b": 2,
					},
					"a": {
						"final": 1,
					},
					"b": {
						"a":     3,
						"final": 5,
					},
					"final": {},
				},
				costs: map[string]int{
					"a":     6,
					"b":     2,
					"final": -1,
				},
				parents: map[string]string{
					"a":     "start",
					"b":     "start",
					"final": "",
				},
			},
			want: []string{"start", "b", "a", "final"},
		},
		{
			name: "02",
			args: args{
				graph: map[string]map[string]int{
					"start": {
						"a": 5,
						"b": 2,
					},
					"a": {
						"b": 1,
						"c": 4,
					},
					"b": {
						"c": 3,
					},
					"c": {
						"d": 1,
						"e": 2,
					},
					"d": {
						"final": 3,
					},
					"e": {
						"final": 1,
					},
					"final": {},
				},
				costs: map[string]int{
					"a":     5,
					"b":     2,
					"c":     -1,
					"d":     -1,
					"e":     -1,
					"final": -1,
				},
				parents: map[string]string{
					"a":     "start",
					"b":     "start",
					"c":     "",
					"d":     "c",
					"e":     "c",
					"final": "",
				},
			},
			want: []string{"start", "b", "c", "e", "final"},
		},
		{
			name: "02",
			args: args{
				graph: map[string]map[string]int{
					"start": {
						"a": 5,
						"b": 2,
						"c": 3,
						"d": 2,
					},
					"a": {
						"e": 2,
					},
					"b": {
						"a": 1,
						"f": 2,
					},
					"c": {
						"e": 1,
					},
					"d": {
						"c": 1,
						"f": 5,
					},
					"e": {
						"h": 3,
					},
					"f": {
						"g": 4,
					},
					"g": {
						"final": 1,
					},
					"h": {
						"final": 3,
					},
					"final": {},
				},
				costs: map[string]int{
					"a":     5,
					"b":     2,
					"c":     3,
					"d":     2,
					"e":     -1,
					"f":     -1,
					"g":     -1,
					"h":     -1,
					"final": -1,
				},
				parents: map[string]string{
					"a":     "start",
					"b":     "start",
					"c":     "start",
					"d":     "start",
					"e":     "",
					"f":     "",
					"g":     "f",
					"h":     "e",
					"final": "",
				},
			},
			want: []string{"start", "b", "f", "g", "final"},
		},
		{
			name: "03",
			args: args{
				graph: map[string]map[string]int{
					"start": {
						"a": 5,
						"b": 2,
					},
					"a": {
						"c": 4,
						"d": 2,
					},
					"b": {
						"a": 8,
						"d": 7,
					},
					"c": {
						"d":     6,
						"final": 3,
					},
					"d": {
						"final": 2,
					},
					"final": {},
				},
				costs: map[string]int{
					"a":     5,
					"b":     2,
					"c":     -1,
					"d":     -1,
					"final": -1,
				},
				parents: map[string]string{
					"a":     "start",
					"b":     "start",
					"c":     "a",
					"d":     "",
					"final": "",
				},
			},
			want: []string{"start", "a", "d", "final"},
		},
		{
			name: "04",
			args: args{
				graph: map[string]map[string]int{
					"start": {
						"a": 10,
					},
					"a": {
						"c": 20,
					},
					"b": {
						"a": 1,
					},
					"c": {
						"b":     1,
						"final": 30,
					},
					"final": {},
				},
				costs: map[string]int{
					"a":     10,
					"b":     -1,
					"c":     -1,
					"final": -1,
				},
				parents: map[string]string{
					"a":     "start",
					"b":     "c",
					"c":     "a",
					"final": "c",
				},
			},
			want: []string{"start", "a", "c", "final"},
		},
		{
			name: "05",
			args: args{
				graph: map[string]map[string]int{
					"start": {
						"a": 2,
						"b": 2,
					},
					"a": {
						"b": 2,
					},
					"b": {
						"c":     2,
						"final": 2,
					},
					"c": {
						"a":     -1,
						"final": 2,
					},
					"final": {},
				},
				costs: map[string]int{
					"a":     2,
					"b":     2,
					"c":     -1,
					"final": -1,
				},
				parents: map[string]string{
					"a":     "start",
					"b":     "start",
					"c":     "",
					"final": "",
				},
			},
			want: []string{"start", "b", "final"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				startTime := time.Now()
				got := dijkstra(tt.args.graph, tt.args.costs, tt.args.parents)
				fmt.Printf("执行耗时：%v\n", time.Since(startTime))
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("dijkstra() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
