package go_vkapi

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
	"testing"
)

func TestSliceToString(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Int slice",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: "1,2,3,4,5",
		},
		{
			name: "BaseUserGroupFields slice",
			args: args{[]objects.BaseUserGroupFields{"1", "2", "3", "4", "5", "6"}},
			want: "1,2,3,4,5,6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToString(tt.args.slice); got != tt.want {
				t.Errorf("SliceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
