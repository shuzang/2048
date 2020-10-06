package game

import (
	"reflect"
	"testing"
)

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
			b := &board{matrix: tt.fields.matrix}
			b.leftRotate90()
			if !reflect.DeepEqual(b.matrix, tt.want) {
				t.Errorf("b.leftRotate90() = %v, want %v", b.matrix, tt.want)
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
			b := &board{matrix: tt.fields.matrix}
			if b.leftRotate90(); !reflect.DeepEqual(b.matrix, tt.want) {
				t.Errorf("b.leftRotate90() = %v, want %v", b.matrix, tt.want)
			}
		})
	}
}
