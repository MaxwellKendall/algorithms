package exercise0109

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"../testutils"
)

func TestAll(t *testing.T) {
	for _, tt := range []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "mismatched lengths returns false",
			s1:   "t",
			s2:   "erbottlewat",
			want: false,
		},
		{
			name: "mispelling returns false",
			s1:   "waterbottle",
			s2:   "ewrbottleat",
			want: false,
		},
		{
			name: "example rotation returns true",
			s1:   "waterbottle",
			s2:   "erbottlewat",
			want: true,
		},
	} {
		for _, a := range []func(string, string) bool{
			Double,
			Loop,
		} {
			t.Run(fmt.Sprintf("%v %v", testutils.FuncName(a), tt.name), func(t *testing.T) {
				got := a(tt.s1, tt.s2)
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("(-want/+got): %v\n", diff)
				}

			})
		}
	}
}
