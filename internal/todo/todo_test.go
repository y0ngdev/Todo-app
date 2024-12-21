package todo

import (
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {

	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields []string
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name:   "Given a todo of shop and a search of sh, i should get shop back",
			fields: []string{"shop"},
			args:   args{"sh"},
			want:   []string{"shop"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService()

			for _, add := range tt.fields {
				err := svc.Add(add)
				if err != nil {
					t.Error(err)
				}
			}
			if got := svc.Search(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
