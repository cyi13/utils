package funcs

import (
	"testing"
)

func TestIp2long(t *testing.T) {
	type args struct {
		ipAddr string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "TestIp2long_1",
			args: args{
				ipAddr: "192.168.1.1",
			},
			want: 3232235777,
		},
		{
			name: "TestIp2long_2",
			args: args{
				ipAddr: "192.168.1.",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ip2long(tt.args.ipAddr); got != tt.want {
				t.Errorf("Ip2long() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLong2ip(t *testing.T) {
	type args struct {
		ipLong uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "TestLong2ip_1",
			args: args{
				ipLong: 3232235777,
			},
			want: "192.168.1.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Long2ip(tt.args.ipLong); got != tt.want {
				t.Errorf("Long2ip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIp2long(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ip2long("192.168.1.1")
	}
}

func BenchmarkLong2ip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Long2ip(3232235777)
	}
}
