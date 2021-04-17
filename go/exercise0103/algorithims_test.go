package exercise0103

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"../testutils"
)

func TestAll(t *testing.T) {
	for _, tt := range []struct {
		name  string
		str    string
		want  string
	}{
		{
			name: "Prepended Space",
			str: " t",
			want: "%20t",
		},
		{
			name: "Appended Space",
			str: "t ",
			want: "t%20",
		},
		{
			name: "Prepended and Appended Space",
			str: " t ",
			want: "%20t%20",
		},
		{
			name: "Empty String",
			str: "",
			want: "",
		},
		{
			name: "Only a space",
			str: " ",
			want: "%20",
		},
	} {
		for _, a := range []func(string) string {
			ReplaceSpace,
		} {
			t.Run(fmt.Sprintf("%v %v", testutils.FuncName(a), tt.name), func(t *testing.T) {
				got := a(tt.str)
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("(-want/+got): %v\n", diff)
				}
			})
		}
	}
}
