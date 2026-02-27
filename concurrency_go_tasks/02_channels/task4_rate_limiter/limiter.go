package limiter

import (
	"sync"
	"time"
)

// Limiter ограничивает количество событий до 5 в секунду.
type Limiter struct {
	tokens chan struct{}
	ticker *time.Ticker
	stop   chan struct{}
	once   sync.Once
}

// NewLimiter создаёт новый лимитер с ёмкостью 5 токенов.
func NewLimiter() *Limiter {
	ch := make(chan struct{}, 5)

	for i := 0; i < 5; i++ {
		ch <- struct{}{}
	}

	ticker := time.NewTicker(time.Second / 5)
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				select {
				case ch <- struct{}{}:
				default:
				}
			case <-stop:
				return
			}
		}
	}()

	return &Limiter{
		tokens: ch,
		ticker: ticker,
		stop:   stop,
	}
}

// Allow возвращает true, если событие разрешено в текущий момент.
func (l *Limiter) Allow() bool {
	select {
	case <-l.tokens:
		return true
	default:
		return false
	}
}

// Stop останавливает лимитер.
func (l *Limiter) Stop() {
	l.once.Do(func() {
		l.ticker.Stop()
		close(l.stop)
	})
}
