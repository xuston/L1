package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var flag bool
var mu sync.Mutex

func Hellocont(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутины завершают работу (контекст)")
			return
		default:
			fmt.Println("Hello")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func Hellochan(stopxhan chan struct{}) {
	for {
		select {
		case <-stopxhan:
			fmt.Println("Горутины завершают работу (канал)")
			return
		default:
			fmt.Println("Hello")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func Hellloflag() {
	for {
		mu.Lock()
		if flag {
			mu.Unlock()
			fmt.Println("Горутины завершают работу (флаг)")
			return
		}
		mu.Unlock()
		fmt.Println("Hello")
		time.Sleep(250 * time.Millisecond)
	}
}

func Hellotimeafter(schan <-chan time.Time) {
	for {
		select {
		case <-schan:
			fmt.Println("Горутины завершают работу (таймер)")
			return
		default:
			fmt.Println("Hello")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func Hellorecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Горутины завершают работу (рековер)")
		}
	}()
	for {
		fmt.Println("Hello")
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	//с помощью контекста
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go Hellocont(ctx)
	time.Sleep(3 * time.Second)

	//мспользование канала для сигналов завершения
	stopChan := make(chan struct{})
	go Hellochan(stopChan)
	time.Sleep(2 * time.Second)
	close(stopChan)
	time.Sleep(1 * time.Second)

	//Использование глобального флага
	go Hellloflag()
	time.Sleep(2 * time.Second)
	mu.Lock()
	flag = true
	mu.Unlock()
	time.Sleep(1 * time.Second)

	//Использование таймера
	schan := time.After(2 * time.Second)
	go Hellotimeafter(schan)
	time.Sleep(3 * time.Second)

	//Прерывание через recover
	go func() {
		time.Sleep(2 * time.Second)
		panic("Остановка горутин")
	}()
	go Hellorecover()
	time.Sleep(3 * time.Second)

	fmt.Println("Програма завершена")
}
