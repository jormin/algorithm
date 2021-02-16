package algorithm

import (
	"fmt"
	"testing"
	"time"
)

// 测试简单查找和二分查找
func TestSimpleAndBinarySearch(t *testing.T) {
	shortList := []int{
		1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57,
		59, 61, 63, 65, 67, 69, 71, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99,
	}
	var midList []int
	for i := 1; i <= 9999999; i += 2 {
		midList = append(midList, i)
	}
	var longList []int
	for i := 1; i <= 999999999; i += 2 {
		longList = append(longList, i)
	}
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
			name: "short",
			args: args{
				list:   shortList,
				target: 99,
			},
			want: 49,
		},
		{
			name: "mid",
			args: args{
				list:   midList,
				target: 9999999,
			},
			want: 4999999,
		},
		{
			name: "long",
			args: args{
				list:   longList,
				target: 999999999,
			},
			want: 499999999,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				fmt.Printf("数组长度：%d\n", len(tt.args.list))
				startTime := time.Now()
				got := binarySearch(tt.args.list, tt.args.target)
				fmt.Printf("二分搜索执行耗时：%v\n", time.Since(startTime))
				if got != tt.want {
					t.Errorf("binarySearch() = %v, want %v", got, tt.want)
				}
				startTime = time.Now()
				got = simpleSearch(tt.args.list, tt.args.target)
				fmt.Printf("简单搜索执行耗时：%v\n", time.Since(startTime))
				if got != tt.want {
					t.Errorf("simpleSearch() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
