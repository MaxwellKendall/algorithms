package testutils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func MyFunc() {}

func TestFuncName(t *testing.T) {
	for _, tt := range []struct {
		name string
		f    interface{}
		want string
	}{
		{
			name: "simple func works",
			f:    MyFunc,
			want: "MyFunc",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := FuncName(tt.f)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want/+got): %v\n", diff)
			}

		})
	}
}
