package validUtf8

import "testing"

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Sample 1",
			args: args{
				data: []int{197, 130, 1},
			},
			want: true,
		},
		{
			name: "Sample 2",
			args: args{
				data: []int{235, 140, 4},
			},
			want: false,
		},
		{
			name: "Test 1",
			args: args{
				data: []int{248, 130, 130, 130},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
