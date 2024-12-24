package main

import (
	"fmt"
	"sync"
)

// Вычисление квадрата числа
func sqare(num int, wg *sync.WaitGroup, chanel chan int) {
	defer wg.Done() //уменьшает счетчик ожидания
	chanel <- num * num
}

func main() {
	array := []int{2, 4, 6, 8, 10}

	//канал для передачи результата
	chanel := make(chan int, len(array))

	var wg sync.WaitGroup

	//запуск горутины
	for _, num := range array {
		wg.Add(1)
		go sqare(num, &wg, chanel)
	}
	//ожидание завершения горутин
	wg.Wait()

	close(chanel)
	sum := 0
	//вывод результата
	for chanel := range chanel {
		sum += chanel
		fmt.Printf("%d,  ", sum)
	}
}
