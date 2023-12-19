package jokesapi

import (
	"reflect"
	"testing"
)

func TestNewJokesResp(t *testing.T) {
	tests := []struct {
		name string
		want *JokesResp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJokesResp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJokesResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
