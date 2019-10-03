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
		{
			"Filters on the first letter",
			args{
				map[int][]rune{
					0: {'a', 'c'},
				},
				[]string{"a", "b"},
			},
			[]string{"a"},
		},
		{
			"Filters on the second letter",
			args{
				map[int][]rune{
					1: {'a', 'c'},
				},
				[]string{"xc", "xb"},
			},
			[]string{"xc"},
		},
		{
			"Filters on all letters",
			args{
				map[int][]rune{
					0: {'a', 'd'},
					1: {'b', 'd'},
					2: {'c', 'd'},
				},
				[]string{"aaa", "bbb", "ccc", "ddd"},
			},
			[]string{"ddd"},
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
