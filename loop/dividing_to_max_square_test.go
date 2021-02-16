package loop

import (
	"fmt"
	"testing"
	"time"
)

// 测试长方形分割成最大正方形
func TestDividing_to_max_square(t *testing.T) {
	type args struct {
		length int
		width  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				length: 1680,
				width:  640,
			},
			want: 80,
		},
		{
			name: "02",
			args: args{
				length: 2190,
				width:  310,
			},
			want: 10,
		},
		{
			name: "03",
			args: args{
				length: 7688,
				width:  431,
			},
			want: 1,
		},
		{
			name: "04",
			args: args{
				length: 640,
				width:  640,
			},
			want: 640,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				startTime := time.Now()
				got := dividing_to_max_square(tt.args.length, tt.args.width)
				fmt.Printf("执行耗时：%v\n", time.Since(startTime))
				if got != tt.want {
					t.Errorf("dividing_to_max_square() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
