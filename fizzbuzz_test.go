package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculateFizzbuzz(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "Number divisible by 3 gives fizz",
			num:  3,
			want: "fizz",
		},
		{
			name: "Number divisible by 5 gives buzz",
			num:  5,
			want: "buzz",
		},
		{
			name: "Number divisible by 3 AND 5 gives fizzbuzz",
			num:  15,
			want: "fizzbuzz",
		},
		{
			name: "Numbers not divisible by 3 OR 5 gives empty string",
			num:  2,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calculateFizzbuzz(tt.num), "calculateFizzbuzz(%v)", tt.num)
		})
	}
}
