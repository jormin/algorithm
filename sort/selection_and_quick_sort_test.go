package sort

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 测试选择排序和快速排序
func TestSelectionAndQuickSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "01",
			args: args{
				list: []int{99, 10, 27, 38, 11, 9, 12, 29, 56, 64, 37, 81, 19},
			},
			want: []int{9, 10, 11, 12, 19, 27, 29, 37, 38, 56, 64, 81, 99},
		},
		{
			name: "02",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "03",
			args: args{
				list: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "04",
			args: args{
				list: readListFromFile("./random_list.txt"),
			},
			want: readListFromFile("./sort_list.txt"),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				fmt.Printf("数组长度：%d\n", len(tt.args.list))
				// 由于切片本质是底层数组的顶层映射，所以需要拷贝为新的切片
				sl := make([]int, len(tt.args.list))
				copy(sl, tt.args.list)
				startTime := time.Now()
				got := selectionSort(sl)
				fmt.Printf("选择排序执行耗时：%v\n", time.Since(startTime))
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("selectionSort() = %v, want %v", got, tt.want)
				}
				// 由于切片本质是底层数组的顶层映射，所以需要拷贝为新的切片
				ql := make([]int, len(tt.args.list))
				copy(ql, tt.args.list)
				startTime = time.Now()
				got = quickSort(ql)
				fmt.Printf("快速排序执行耗时：%v\n", time.Since(startTime))
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("quickSort() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
