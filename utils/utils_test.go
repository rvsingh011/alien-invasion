package utils

import (
	"math/rand"
	"testing"
)

func TestGetOppositeDirection(t *testing.T) {
	type args struct {
		direction string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test direction north",
			args: args{"north"},
			want: "south",
		},
		{
			name: "Test direction east",
			args: args{"east"},
			want: "west",
		},
		{
			name: "Test direction south",
			args: args{"south"},
			want: "north",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOppositeDirection(tt.args.direction); got != tt.want {
				t.Errorf("GetOppositeDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomNumber(t *testing.T) {
	type args struct {
		min int
		max int
		r   *rand.Rand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test random with seed 3",
			args: args{0, 3, rand.New(rand.NewSource(3))},
			want: 0,
		},
		{
			name: "Test random with seed 3",
			args: args{0, 4, rand.New(rand.NewSource(3))},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomNumber(tt.args.min, tt.args.max, tt.args.r); got != tt.want {
				t.Errorf("GetRandomNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
