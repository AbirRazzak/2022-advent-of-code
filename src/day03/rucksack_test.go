package day03

import "testing"

func TestRuckSack_GetCompartments(t *testing.T) {
	type fields struct {
		Contents string
	}
	tests := []struct {
		name    string
		fields  fields
		want1   string
		want2   string
		wantErr bool
	}{
		{
			name:   "even sized string",
			fields: fields{Contents: "vJrwpWtwJgWrhcsFMMfFFhFp"},
			want1:  "vJrwpWtwJgWr",
			want2:  "hcsFMMfFFhFp",
		},
		{
			name:    "returns error when contents is unset",
			fields:  fields{Contents: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RuckSack{
				Contents: tt.fields.Contents,
			}
			got, got1, err := r.GetCompartments()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompartments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want1 {
				t.Errorf("GetCompartments() got = %v, want %v", got, tt.want1)
			}
			if got1 != tt.want2 {
				t.Errorf("GetCompartments() got1 = %v, want %v", got1, tt.want2)
			}
		})
	}
}

func TestRuckSack_FindDuplicateItem(t *testing.T) {
	type fields struct {
		Contents     string
		Compartment1 string
		Compartment2 string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:   "duplicate lowercase item",
			fields: fields{Compartment1: "vJrwpWtwJgWr", Compartment2: "hcsFMMfFFhFp"},
			want:   "p",
		},
		{
			name:   "matching upper and lowercase item does not match",
			fields: fields{Compartment1: "vJrwpWtwJgWr", Compartment2: "VcsFMMfFFhFp"}, // comp1 has v, comp2 has V
			want:   "p",
		},
		{
			name:    "error returned when compartments unset",
			fields:  fields{},
			wantErr: true,
		},
		{
			name:    "error returned when no items match",
			fields:  fields{Compartment1: "asdfgh", Compartment2: "qwerty"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RuckSack{
				Contents:     tt.fields.Contents,
				Compartment1: tt.fields.Compartment1,
				Compartment2: tt.fields.Compartment2,
			}
			got, err := r.FindDuplicateItem()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindDuplicateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindDuplicateItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}
