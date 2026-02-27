package producerconsumer

import (
	"io"
	"strconv"
	"sync"
)

// Run запускает продюсера, который отправляет числа от 1 до 10, и консюмера,
// который выводит их в writer. Используйте небуферизованный канал и ожидание
// завершения горутин.
func Run(w io.Writer) {
	// TODO: реализовать продюсер и консюмер

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	// producer
	go producer(&wg, ch)
	// consumer
	go consumer(&wg, w, ch)

	wg.Wait()
}

func producer(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for i := range 10 {
		ch <- i + 1
	}
	close(ch)
}

func consumer(wg *sync.WaitGroup, w io.Writer, ch <-chan int) {
	defer wg.Done()

	for v := range ch {
		w.Write(
			[]byte(strconv.Itoa(v) + "\n"),
		)
	}
}
