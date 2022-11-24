package fizzbuzz

func calculateFizzbuzz(num int) string {
	if num % 5 == 0 {
		return  "buzz"
	}
	return "fizz"
}
