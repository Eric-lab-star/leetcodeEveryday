package removeduplicatefromsortedarray;

import "testing";

func TestCode(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0,0,1,1,1,2,2,3,3,4}, 5},
	};

	for _, tt := range tests {
		got := removeDuplicates(tt.nums)
		if got != tt.want {
			t.Errorf("wrong match got %v want %v", got, tt.want)
		}
	}
}
