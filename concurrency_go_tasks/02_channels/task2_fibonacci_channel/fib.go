package fibonacci

// Fib возвращает канал, из которого можно читать первые n чисел Фибоначчи.
func Fib(n int) <-chan int {
	ch := make(chan int)

	if n <= 0 {
		close(ch)
		return ch
	}

	go func() {
		a, b := 0, 1

		for range n {
			ch <- a
			a, b = b, a+b
		}
		close(ch)
	}()

	return ch
}
