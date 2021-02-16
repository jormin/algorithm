package loop

import (
	"fmt"
	"testing"
	"time"
)

// 测试递归版求和
func TestLoopSum(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				list: []int{1, 3, 5, 12, 6, 19},
			},
			want: 46,
		},
		{
			name: "02",
			args: args{
				list: []int{1},
			},
			want: 1,
		},
		{
			name: "03",
			args: args{
				list: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				startTime := time.Now()
				got := loopSum(tt.args.list)
				fmt.Printf("执行耗时：%v\n", time.Since(startTime))
				if got != tt.want {
					t.Errorf("loopSum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
