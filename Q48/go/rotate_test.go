package rotate

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
		},
		{
			name: "Example 2",
			args: args{
				matrix: [][]int{
					{5, 1, 9, 11},
					{2, 4, 8, 10},
					{13, 3, 6, 7},
					{15, 14, 12, 16},
				},
			},
		},
		{
			name: "Example 3",
			args: args{
				matrix: [][]int{
					{5, 1, 9, 11, 25, 66, 17},
					{2, 4, 8, 10, 54, 22, 97},
					{13, 3, 6, 7, 19, 52, 46},
					{15, 14, 12, 16, 25, 61, 23},
					{28, 42, 43, 44, 45, 46, 51},
					{53, 55, 57, 67, 68, 69, 70},
					{81, 72, 74, 76, 82, 85, 88},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
		})
	}
}
