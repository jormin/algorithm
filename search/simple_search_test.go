package search

import (
	"fmt"
	"testing"
	"time"
)

// 测试简单查找
func TestSimpleSearch(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				list:   []int{1, 4, 7, 9, 10, 15},
				target: 4,
			},
			want: 1,
		},
		{
			name: "02",
			args: args{
				list:   []int{1, 4, 7, 9, 10, 15},
				target: 1,
			},
			want: 0,
		},
		{
			name: "03",
			args: args{
				list:   []int{1, 4, 7, 9, 10, 15},
				target: 15,
			},
			want: 5,
		},
		{
			name: "04",
			args: args{
				list:   []int{1, 4, 7, 9, 10, 15},
				target: 0,
			},
			want: -1,
		},
		{
			name: "05",
			args: args{
				list:   []int{1, 4, 7, 9, 10, 15},
				target: 19,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				startTime := time.Now()
				got := simpleSearch(tt.args.list, tt.args.target)
				fmt.Printf("执行耗时：%v\n", time.Since(startTime))
				if got != tt.want {
					t.Errorf("simpleSearch() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
