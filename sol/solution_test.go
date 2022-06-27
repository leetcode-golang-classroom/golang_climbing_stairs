package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	n := 10
	for idx := 0; idx < b.N; idx++ {
		climbStairs(n)
	}
}
func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "n=3",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "n=4",
			args: args{n: 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
