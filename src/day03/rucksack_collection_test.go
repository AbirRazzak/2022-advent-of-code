package day03

import "testing"

func TestRuckSackCollection_CalculateItemPriority(t *testing.T) {
	type fields struct {
		Sacks []*RuckSack
	}
	type args struct {
		item string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "a",
			args: args{item: "a"},
			want: 1,
		},
		{
			name: "c",
			args: args{item: "c"},
			want: 3,
		},
		{
			name: "z",
			args: args{item: "z"},
			want: 26,
		},
		{
			name: "A",
			args: args{item: "A"},
			want: 27,
		},
		{
			name: "F",
			args: args{item: "F"},
			want: 32,
		},
		{
			name: "Z",
			args: args{item: "Z"},
			want: 52,
		},
		{
			name:    "error returned when given not a letter",
			args:    args{item: "2"},
			wantErr: true,
		},
		{
			name:    "error returned when given long string",
			args:    args{item: "qwerty"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := RuckSackCollection{
				Sacks: tt.fields.Sacks,
			}
			got, err := c.CalculateItemPriority(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateItemPriority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateItemPriority() got = %v, want %v", got, tt.want)
			}
		})
	}
}
