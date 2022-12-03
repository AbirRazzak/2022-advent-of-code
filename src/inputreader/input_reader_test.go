package inputreader_test

import (
	"github.com/AbirRazzak/2022-advent-of-code/src/inputreader"
	"path/filepath"
	"testing"
)

func TestInputReader_GetFileNameForDay(t *testing.T) {
	type args struct {
		dayNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "single digit day number",
			args: args{
				dayNumber: 1,
			},
			want:    "day01_input.txt",
			wantErr: false,
		},
		{
			name: "double digit day number",
			args: args{
				dayNumber: 10,
			},
			want:    "day10_input.txt",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := inputreader.InputReader{}
			got, err := r.GetFileNameForDay(tt.args.dayNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileNameForDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFileNameForDay() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputReader_GetFilePathForDay(t *testing.T) {
	type args struct {
		dayNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				dayNumber: 1,
			},
			want:    filepath.Join("..", "resources", "day01_input.txt"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := inputreader.InputReader{}
			got, err := r.GetFilePathForDay(tt.args.dayNumber)
			if (err != nil) != tt.wantErr {
				t.Skipf("GetFilePathForDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFilePathForDay() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputReader_GetInputForDay(t *testing.T) {
	type args struct {
		dayNumber int
	}
	tests := []struct {
		name      string
		args      args
		wantInput string
		wantErr   bool
	}{
		{
			name: "happy path",
			args: args{
				dayNumber: 0,
			},
			wantInput: "this is a test file",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := inputreader.InputReader{}
			gotInput, err := r.GetInputForDay(tt.args.dayNumber)
			if (err != nil) != tt.wantErr {
				t.Skipf("GetInputForDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotInput != tt.wantInput {
				t.Errorf("GetInputForDay() gotInput = %v, want %v", gotInput, tt.wantInput)
			}
		})
	}
}
