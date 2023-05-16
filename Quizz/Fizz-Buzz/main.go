package main

func main() {
	for i := 1; i < 101; i++ {
		isFizz := false
		isBuzz := false
		if i%3 == 0 {
			isFizz = true
			print("Fizz")
		}
		if i%5 == 0 {
			isBuzz = true
			print("Buzz")
		}
		if !isFizz && !isBuzz {
			print(i)
		}
		println()
	}
}
