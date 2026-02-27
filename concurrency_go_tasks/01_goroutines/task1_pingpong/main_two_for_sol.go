package main

/* import (
	"io"
	"os"
	"sync"
) */

// PingPong должен запускать две горутины "ping" и "pong",
// которые поочередно выводят строки пять раз каждая.
// Реализуйте синхронизацию через каналы и ожидание завершения.
/* func PingPong(w io.Writer) {
	// TODO: реализовать обмен сообщениями между горутинами

	var wg sync.WaitGroup

	toPingCh := make(chan struct{})
	toPongCh := make(chan struct{})

	// Ping
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-toPingCh
			w.Write([]byte("ping\n"))

			toPongCh <- struct{}{}
		}
	}()

	// Pong
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-toPongCh
			w.Write([]byte("pong\n"))

			if i == 4 {
				break
			}

			toPingCh <- struct{}{}
		}
	}()

	toPingCh <- struct{}{}

	wg.Wait()
}

func main() {
	PingPong(os.Stdout)
} */
