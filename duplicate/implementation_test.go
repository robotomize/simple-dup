package duplicate

import (
	"reflect"
	"testing"
)

func asChan(data [][]string) chan []string {
	out := make(chan []string, len(data))
	defer close(out)

	for _, v := range data {
		out <- v
	}

	return out
}

func Test_dup_Find(t *testing.T) {
	tests := []struct {
		name string
		args [][]string
		want []Item
	}{
		{
			name: "test vector 0",
			args: [][]string{
				{"1", "1", "0000", "1"},
				{"2", "1", "4708", "2"},
				{"3", "1", "4708", "3"},
			},
			want: []Item{
				{1, 0, "1"},
				{1, 1, "1"},
				{1, 2, "1"},

				{2, 1, "4708"},
				{2, 2, "4708"},
			},
		},

		{
			name: "test vector 1",
			args: [][]string{
				{"a", "b", "c"},
				{"a", "b", "c"},
			},
			want: []Item{
				{0, 0, "a"},
				{0, 1, "a"},

				{1, 0, "b"},
				{1, 1, "b"},

				{2, 0, "c"},
				{2, 1, "c"},
			},
		},

		{
			name: "test vector 2",
			args: [][]string{
				{"a", "b", "c"},
				{"d", "d", "d"},
				{"a", "b", "c"},
			},
			want: []Item{
				{0, 0, "a"},
				{0, 2, "a"},

				{1, 0, "b"},
				{1, 2, "b"},

				{2, 0, "c"},
				{2, 2, "c"},
			},
		},
	}

	finder := NewFinder()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := make(map[Item]struct{})
			for _, v := range tt.want {
				want[v] = struct{}{}
			}

			got := make(map[Item]struct{})
			for _, v := range finder.Find(asChan(tt.args)) {
				got[v] = struct{}{}
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
