package hStrings

import (
	"testing"
)

func TestConcatenate(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{str: []string{"hello", "world", "123"}},
			want: "helloworld123",
		},
		{
			name: "No String",
			args: args{str: []string{}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concatenate(tt.args.str...); got != tt.want {
				t.Errorf("Concatenate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name:       "OK",
			args:       args{str: "hello world"},
			wantResult: "dlrow olleh",
		},
		{
			name:       "No String",
			args:       args{str: ""},
			wantResult: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Reverse(tt.args.str); gotResult != tt.wantResult {
				t.Errorf("Reverse() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSlug(t *testing.T) {
	type args struct {
		str       string
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{str: "hello world", separator: "_"},
			want: "hello_world",
		},
		{
			name: "Default Separator",
			args: args{str: "hello world", separator: ""},
			want: "hello-world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slug(tt.args.str, tt.args.separator); got != tt.want {
				t.Errorf("Slug() = %v, want %v", got, tt.want)
			}
		})
	}
}
