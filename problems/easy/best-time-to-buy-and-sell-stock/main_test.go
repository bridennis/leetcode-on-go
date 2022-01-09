package easy_best_time_to_buy_and_sell_stock

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Max profit #1",
			args: struct{ prices []int }{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "Max profit #2",
			args: struct{ prices []int }{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "Max profit #3",
			args: struct{ prices []int }{prices: []int{2, 1}},
			want: 0,
		},
		{
			name: "Max profit #4",
			args: struct{ prices []int }{prices: []int{3, 2, 6, 5, 0, 3}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
