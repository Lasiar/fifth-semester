package main

import (
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		slc        []int
		searchElem int
		offset     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[] find 0",
			args: struct {
				slc        []int
				searchElem int
				offset     int
			}{
				slc:        []int{},
				searchElem: 0,
			},
			want: -1,
		},
		{
			name: "[0] find 0",
			args: struct {
				slc        []int
				searchElem int
				offset     int
			}{
				slc:        []int{0},
				searchElem: 0,
				offset:     0,
			},
			want: 0,
		},
		{
			name: "[0,0,1,1,2,2] find 0",
			args: struct {
				slc        []int
				searchElem int
				offset     int
			}{
				slc:        []int{0},
				searchElem: 0,
				offset:     0,
			},
			want: 0,
		},
		{
			name: "[1,2,3] find 0",
			args: struct {
				slc        []int
				searchElem int
				offset     int
			}{
				slc:        []int{1, 2, 3},
				searchElem: 0,
				offset:     0,
			},
			want: -1,
		},
		{
			name: "[1,2,3] find 3",
			args: struct {
				slc        []int
				searchElem int
				offset     int
			}{
				slc:        []int{1, 2, 3},
				searchElem: 3,
				offset:     0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.slc, tt.args.searchElem, tt.args.offset); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchFloat(t *testing.T) {
	type args struct {
		slc        []float64
		searchElem float64
		offset     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[0.02,0.2,0.223] find 0.02",
			args: struct {
				slc        []float64
				searchElem float64
				offset     int
			}{
				slc:        []float64{0.02, 0.2, 0.223},
				searchElem: 0.02,
				offset:     0,
			},
			want: 0,
		},
		{
			name: "[0.02,0.2,0.223] find 0.2",
			args: struct {
				slc        []float64
				searchElem float64
				offset     int
			}{
				slc:        []float64{0.02, 0.2, 0.233},
				searchElem: 0.2,
				offset:     0,
			},
			want: 1,
		},
		{
				name: "[0.02,0.2,50] find 50",
			args: struct {
				slc        []float64
				searchElem float64
				offset     int
			}{
				slc:        []float64{0.02, 0.2, 50},
				searchElem: 50,
				offset:     0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchFloat(tt.args.slc, tt.args.searchElem, tt.args.offset); got != tt.want {
				t.Errorf("binarySearchFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
