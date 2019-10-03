package password

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		possibleIndexLetters map[int][]rune
		possibleWords        []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Returns an empty slice",
			args{
				map[int][]rune{},
				[]string{},
			},
			[]string{},
		},
		{
			"Returns the word map when not filtered",
			args{
				map[int][]rune{},
				[]string{"a", "b"},
			},
			[]string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.possibleIndexLetters, tt.args.possibleWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
