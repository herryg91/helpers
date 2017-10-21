package hArrays

import "testing"

func TestInArray(t *testing.T) {
	type args struct {
		array interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "String Value Found",
			args: args{array: []string{"test1", "test2", "test3"}, value: "test1"},
			want: true,
		},
		{
			name: "String Value Not Found",
			args: args{array: []string{"test1", "test2", "test3"}, value: "test4"},
			want: false,
		},
		{
			name: "Int Value Found",
			args: args{array: []int{16, 32, 64}, value: 16},
			want: true,
		},
		{
			name: "Int Value Not Found",
			args: args{array: []int{16, 32, 64}, value: 128},
			want: false,
		},
		{
			name: "Not Array",
			args: args{array: 16, value: 128},
			want: false,
		},
		{
			name: "Value different type",
			args: args{array: []int{16, 32, 64}, value: "128"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.array, tt.args.value); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
