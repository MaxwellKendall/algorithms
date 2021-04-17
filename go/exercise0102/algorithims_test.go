package exercise0102

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"../testutils"
)

func TestAll(t *testing.T) {
	for _, tt := range []struct {
		name  string
		s1    string
		s2    string
		want  bool
	}{
		{
			name: "re-arranged alphabet returns true",
			s1: "abcdefghijklmnopqrstuvwxyz",
			s2: "zbcdefghijklmnopqrstuvwxya",
			want: true,
		},
		{
			name: "different letters with same number of occurrences returns false",
			s1: "zz",
			s2: "bb",
			want: false,
		},
		{
			name: "unequal length of strings returns false",
			s1: "abcdefghijklmnopqrstuvwxyz",
			s2: "abcdefghijklmnopqrstuvwxyzz",
			want: false,
		},
		{
			name: "empty string returns true",
			s1: "",
			s2: "",
			want: true,
		},
		{
			name: "double a (w/ no t) in alphabet returns false",
			s1: "abcdefghijklmnopqrstuvwxyz",
			s2: "abcdefghijklmnopqrsauvwxyz",
			want: false,
		},
		{
			name: "test for conley",
			s1: "ab",
			s2: "ac",
			want: false,
		},
		{
			name: "string w/ 4 bs and 2 as vs 2bs and 4as",
			s1: "aabbaa",
			s2: "aabbbb",
			want: false,
		},
		{
			name: "empty string test",
			s1: "",
			s2: "",
			want: false,
		},
	} {
		for _, a := range []func(string, string) bool{
			WithMap,
			With2Maps,
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
