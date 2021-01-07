package exercise0101

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"../testutils"
)

func TestAll(t *testing.T) {
	for _, tt := range []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string returns true",
			s:    "",
			want: true,
		},
		{
			name: "one character string returns true",
			s:    "",
			want: true,
		},
		{
			name: "alphabet returns true",
			s:    "abcdefghijklmnopqrstuvwxyz",
			want: true,
		},
		{
			name: "double a returns false",
			s:    "aabcdefghijklmnopqrstuvwxyz",
			want: false,
		},
		{
			name: "double z returns false",
			s:    "abcdefghijklmnopqrstuvwxyzz",
			want: false,
		},
		{
			name: "double p returns false",
			s:    "abcdefghijklmnoppqrstuvwxyz",
			want: false,
		},
	} {
		for _, a := range []func(string) bool{
			WithMap,
			NestedLoop,
			SortFirst,
		} {
			t.Run(fmt.Sprintf("%v %v", testutils.FuncName(a), tt.name), func(t *testing.T) {
				got := a(tt.s)
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("(-want/+got): %v\n", diff)
				}

			})
		}
	}
}
