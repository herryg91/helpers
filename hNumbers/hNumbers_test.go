package hNumbers

import (
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		number  float64
		decimal int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive Number To Int Round Down",
			args: args{number: 102.4, decimal: 0},
			want: 102,
		},
		{
			name: "Positive Number To Int Round Up",
			args: args{number: 102.5, decimal: 0},
			want: 103,
		},
		{
			name: "Positive Number One Dec Round Down",
			args: args{number: 102.44, decimal: 1},
			want: 102.4,
		},
		{
			name: "Positive Number Two Dec Round Up",
			args: args{number: 102.588, decimal: 2},
			want: 102.59,
		},
		{
			name: "Negative Number To Int Round Down",
			args: args{number: -102.4, decimal: 0},
			want: -102,
		},
		{
			name: "Negative Number To Int Round Up",
			args: args{number: -102.5, decimal: 0},
			want: -103,
		},
		{
			name: "Negative Number One Dec Round Down",
			args: args{number: -102.44, decimal: 1},
			want: -102.4,
		},
		{
			name: "Negative Number Two Dec Round Up",
			args: args{number: -102.588, decimal: 2},
			want: -102.59,
		},
		{
			name: "Negative decimal",
			args: args{number: 102.4, decimal: -1},
			want: 102,
		},
		{
			name: "Overflow decimal",
			args: args{number: 102.4, decimal: 8},
			want: 102.4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.number, tt.args.decimal); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberFormat(t *testing.T) {
	type args struct {
		number            float64
		decimals          int
		decimalPoint      string
		thousandSeparator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "OK",
			args: args{123456.789, 2, ",", "."},
			want: "123.456,79",
		},
		{
			name: "Different Separator",
			args: args{123456.789, 2, ".", ","},
			want: "123,456.79",
		},
		{
			name: "No Decimal",
			args: args{123456, 0, ",", "."},
			want: "123.456",
		},
		{
			name: "One Head",
			args: args{1234567.9646, 3, ",", "."},
			want: "1.234.567,965",
		},
		{
			name: "Two Head",
			args: args{12345678.9646, 3, ",", "."},
			want: "12.345.678,965",
		},
		{
			name: "Small Number",
			args: args{123.44, 1, ",", "."},
			want: "123,4",
		},
		{
			name: "Small Number No Decimal",
			args: args{123.5, 0, ",", "."},
			want: "124",
		},
		{
			name: "Minus",
			args: args{-123456.789, 2, ",", "."},
			want: "-123.456,79",
		},
		{
			name: "Default Separator",
			args: args{-123456.789, 2, "", ""},
			want: "-123.456,79",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberFormat(tt.args.number, tt.args.decimals, tt.args.decimalPoint, tt.args.thousandSeparator); got != tt.want {
				t.Errorf("NumberFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMinus(t *testing.T) {
	type args struct {
		number interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name:       "Int Positive",
			args:       args{number: int(10)},
			wantResult: false,
		},
		{
			name:       "Int8 Negative",
			args:       args{number: int8(-10)},
			wantResult: true,
		},
		{
			name:       "Int16 Positive",
			args:       args{number: int16(10)},
			wantResult: false,
		},
		{
			name:       "Int32 Negative",
			args:       args{number: int32(-10)},
			wantResult: true,
		},
		{
			name:       "Int64 Zero",
			args:       args{number: int64(0)},
			wantResult: false,
		},
		{
			name:       "Float32 positive",
			args:       args{number: float32(10)},
			wantResult: false,
		},
		{
			name:       "Float64 negative",
			args:       args{number: float64(-10)},
			wantResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := IsMinus(tt.args.number); gotResult != tt.wantResult {
				t.Errorf("IsMinus() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
