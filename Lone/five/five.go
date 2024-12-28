package main

import (
	"context"
	"fmt"
	"math/rand"
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

// Запись данных в канал
func Write(ctx context.Context, channel chan string) {
	defer close(channel)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Запись в канал остановлена")
			return
		case channel <- RandString(rand.Intn(11-2) + 2):
			fmt.Println("Информация отправлена")
		}
	}
}

// Чтение данных из канала
func Read(ctx context.Context, channel chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Чтение канала остановлено")
			return
		case data, ok := <-channel:
			if !ok {
				fmt.Println("Канал закрыт")
			}
			fmt.Printf("Получено: %s\n", data)
		}
	}
}

func main() {

	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Printf("Error %e", err)
		n = 10
	}

	channel := make(chan string)

	// Задание контекста после завершения программы
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	//Зпуск горутин для чтения и записи
	go Write(ctx, channel)
	go Read(ctx, channel)

	// Ожидание завершения работы
	<-ctx.Done()
	fmt.Println("Программа завершена")
}
