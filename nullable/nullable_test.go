package nullable

import (
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	var b1 = true
	b1 = false // always waring 피하기 위해서
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		{
			name: "전달된 bool의 포인터를 반환한다",
			args: args{b: b1},
			want: &b1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	var f float32 = 1.0
	type args struct {
		num float32
	}
	tests := []struct {
		name string
		args args
		want *float32
	}{
		{
			name: "전달된 float32의 포인터를 반환한다",
			args: args{num: f},
			want: &f,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	var f float64 = 2.0
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		{
			name: "전달된 float64의 포인터를 반환한다",
			args: args{num: f},
			want: &f,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	var i int = 3
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "전달된 int의 포인터를 반환한다",
			args: args{num: i},
			want: &i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	var s string = "test"
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "전달된 string의 포인터를 반환한다",
			args: args{str: s},
			want: &s,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
