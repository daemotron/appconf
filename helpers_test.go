package appconf

import (
	"math"
	"reflect"
	"testing"
)

func TestAppConf_almostEqual(t *testing.T) {
	type args struct {
		f1 float64
		f2 float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "zero equal sign", args: args{0.0, 0.0}, want: true},
		{name: "zero different sign", args: args{0.0, -0.0}, want: true},
		{name: "one nan value", args: args{0.0, math.NaN()}, want: false},
		{name: "two nan values", args: args{math.NaN(), math.NaN()}, want: false},
		{name: "one inf value", args: args{0.0, math.Inf(0)}, want: false},
		{name: "two inf values", args: args{math.Inf(0), math.Inf(0)}, want: false},
		{name: "one nan value", args: args{0.0, math.NaN()}, want: false},
		{name: "two nan values", args: args{math.NaN(), math.NaN()}, want: false},
		{name: "two similar values", args: args{1.000000000001, 1.000000000001}, want: true},
		{name: "two non-similar values", args: args{1.000000000001, 1.000000000002}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := almostEqual(tt.args.f1, tt.args.f2); got != tt.want {
				t.Errorf("almostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppConf_mergeMaps(t *testing.T) {
	type args[K comparable, V any] struct {
		maps []map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K]V
	}
	tests := []testCase[string, string]{
		{
			name: "map with different keys",
			args: args[string, string]{
				maps: []map[string]string{
					{"foo": "bar"},
					{"boo": "far"},
				},
			},
			want: map[string]string{"foo": "bar", "boo": "far"},
		},
		{
			name: "map with same keys",
			args: args[string, string]{
				maps: []map[string]string{
					{"foo": "bar"},
					{"foo": "far"},
				},
			},
			want: map[string]string{"foo": "far"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeMaps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppConf_getEpsilon(t *testing.T) {
	epsilon := getEpsilon()
	if !((1.0+epsilon) != 1.0 && (1.0+epsilon/2) == 1.0) {
		t.Errorf("Epsilon doesn't fulfil requirement 1.0+e != 1.0 && 1.0+e/2 == 1.0 (computed: %v)", epsilon)
	}
}
