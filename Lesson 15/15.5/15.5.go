// Необходимо реализовать интерфейс
// type Meteo interface {
// ReadTemp() string
// ChangeTemp(v string)
// }
// Речь про температуру окружающей среды. ReadTemp читает
// температуру, ChangeTemp изменяет температуру.
//Код должен быть потокобезопасным,
//т.е. при работе с температурой нескольких параллельных горутин не должно возникать состояние гонки.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type currencyTemp struct {
	mu    sync.Mutex
	value int
}

func (ce *currencyTemp) read() int {
	ce.mu.Lock()
	defer ce.mu.Unlock()
	return ce.value
}
func (ce *currencyTemp) add() {
	ce.mu.Lock()
	defer ce.mu.Unlock()
	ce.value++
}
func main() {
	ce := currencyTemp{}
	for i := 0; i < 1000; i++ {
		go func() {
			ce.add()
		}()
		go func() {
			ce.read()
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(ce.read())
}
