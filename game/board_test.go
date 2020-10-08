package game

import (
	"reflect"
	"testing"
)

func TestMoveRow(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "one",
			input: []int{2, 2, 0, 0},
			want:  []int{4, 0, 0, 0},
		},
		{
			name:  "two",
			input: []int{2, 2, 4, 8},
			want:  []int{4, 4, 8, 0},
		},
		{
			name:  "three",
			input: []int{2, 4, 4, 8},
			want:  []int{2, 8, 8, 0},
		},
		{
			name:  "four",
			input: []int{2, 4, 8, 8},
			want:  []int{2, 4, 16, 0},
		},
		{
			name:  "five",
			input: []int{2, 2, 2, 2},
			want:  []int{4, 4, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveRow(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftRotate90(t *testing.T) {
	type fields struct {
		matrix [][]int
	}

	tests := []struct {
		name   string
		fields fields
		want   [][]int
	}{
		{
			name: "one",
			fields: fields{
				matrix: [][]int{
					{1, 2, 3, 9},
					{4, 5, 6, 10},
					{6, 7, 8, 11},
					{16, 17, 18, 111},
				},
			},
			want: [][]int{
				{9, 10, 11, 111},
				{3, 6, 8, 18},
				{2, 5, 7, 17},
				{1, 4, 6, 16},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{board: tt.fields.matrix}
			b.leftRotate90()
			if !reflect.DeepEqual(b.board, tt.want) {
				t.Errorf("b.leftRotate90() = %v, want %v", b.board, tt.want)
			}
		})
	}
}

func TestRightRotate90(t *testing.T) {
	type fields struct {
		matrix [][]int
	}

	tests := []struct {
		name   string
		fields fields
		want   [][]int
	}{
		{
			name: "one",
			fields: fields{
				matrix: [][]int{
					{1, 2, 3, 9},
					{4, 5, 6, 10},
					{6, 7, 8, 11},
					{16, 17, 18, 111},
				},
			},
			want: [][]int{
				{16, 6, 4, 1},
				{17, 7, 5, 2},
				{18, 8, 6, 3},
				{111, 11, 10, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{board: tt.fields.matrix}
			if b.rightRotate90(); !reflect.DeepEqual(b.board, tt.want) {
				t.Errorf("b.rightRotate90() = %v, want %v", b.board, tt.want)
			}
		})
	}
}
