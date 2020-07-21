package model

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Testing dictionary initialization function",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
			if len(dictionary) < 7 {
				t.Errorf("CheckInit() dictionary length = %v, want at least %v", len(dictionary), 7)
			}
		})
	}
}

func TestCheckNumber(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "In boundaries number check",
			args: args{
				i: 19,
			},
			want: true,
		},
		{
			name: "Zero number check",
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Out of boundaries number check",
			args: args{
				i: NumberConversionLimit + 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
			if got := CheckNumber(tt.args.i); got != tt.want {
				t.Errorf("CheckNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderedKeys(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "Test ordering dictionary keys",
			want: []int{1000, 500, 100, 50, 10, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
			if got := orderKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test number 101 corresponding to IC",
			args: args{
				n: 101,
			},
			want: "IC",
		},
		{
			name: "Test out of boundaries number",
			args: args{
				n: NumberConversionLimit + 1,
			},
			want: "",
		},
		{
			name: "Test negative number",
			args: args{
				n: -10,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.n); got != tt.want {
				fmt.Println(got)
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
