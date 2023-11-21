package main

/*
Реализовать собственную функцию sleep.
*/
import (
	"fmt"
	"time"
)

// Функция mySleep задерживает выполнение программы на указанное количество времени d.
func mySleep(d time.Duration) {
	// С помощью select и time.After дожидаемся истечения времени d.
	select {
	case <-time.After(d):
		// После истечения времени, продолжаем выполнение программы.
	}
}

func main() {
	start := time.Now() // Запоминаем текущее время.
	// time.Sleep(time.Second * 3)
	// Объявляем и задаем продолжительность времени (3 секунды).
	mySleep(time.Duration(time.Second * 3))

	elapsed := time.Since(start)        // Вычисляем прошедшее время с момента start.
	fmt.Printf("page took %s", elapsed) // Выводим результат, сколько времени прошло.
}
