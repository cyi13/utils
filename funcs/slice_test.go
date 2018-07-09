package funcs

import (
	"reflect"
	"testing"
)

func TestDeleteInts(t *testing.T) {
	type args struct {
		s   []int
		key int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "TestDeleteInts_1",
			args: args{
				s:   []int{1, 2, 3, 4},
				key: 0,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "TestDeleteInts_2",
			args: args{
				s:   []int{1, 2, 3, 4},
				key: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "TestDeleteInts_3",
			args: args{
				s:   []int{1, 2, 3, 4},
				key: 1,
			},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteInts(tt.args.s, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSlice(t *testing.T) {
	type args struct {
		s   interface{}
		key int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "TestDeleteSlice_1",
			args: args{
				s:   []int{1, 2, 3, 4},
				key: 0,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "TestDeleteSlice_2",
			args: args{
				s:   []int{1, 2, 3, 4},
				key: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "TestDeleteSlice_3",
			args: args{
				s:   []int{1, 2, 3, 4},
				key: 1,
			},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSlice(tt.args.s, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDeleteSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
		DeleteSlice(s, 6)
	}
}

func BenchmarkDeleteInts(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
		DeleteInts(s, 6)
	}
}

func BenchmarkDeleteStrings(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "10"}
		DeleteStrings(s, 6)
	}
}
