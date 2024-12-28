package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Генерация случайной сроки для передачи в канал
func RandString(a int) string {
	b := make([]byte, a)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// Конкурентная функция для обработки данных из канала
func worker(id int, ctx context.Context, wg *sync.WaitGroup, chanl <-chan string) {
	wg.Done()

	for {
		select {
		case <-ctx.Done(): // Заершение воркера при получении сигнала
			fmt.Printf("Воркер %d завершает работу \n", id)
			return
		case data, ok := <-chanl: //Получение данных из канала
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт \n", id)
				return
			}
			fmt.Printf("Воркер %d получил данные: %s \n", id, data)
		}
	}
}

func main() {
	//Количество воркеров
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Printf("Errorr %s", err)
		n = 3
	}
	//Канал для передачи данных
	chanl := make(chan string)

	//Контекст для завершения
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	//Запуск воркеров
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, ctx, wg, chanl)
	}

	// Передача в канал случайных данных
	go func() {

		for {
			select {
			case <-ctx.Done():
				close(chanl)
				return
			default:
				chanl <- RandString(rand.Intn(11-2) + 2)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	//Ожидание завершения программы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmt.Println("Завершение программы")
	cancel()  //Завершение контекста
	wg.Wait() //Ожидание завершения всех воркеров
	fmt.Println("Все воркеры завершили работу")

}

/*При нажатии Ctrl+C, программа вызывает cancel() для завершения всех горутин и
канал закрывается чтобы предотвратить ззависание воркеров при чтении. Для
синхронизации завершения воркеров используется WitGroup */
