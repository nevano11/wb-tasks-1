package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.

func workerWork(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d: job=%d\n", id, job)
	}
}

func main1() {
	// Число воркеров
	n := 0
	fmt.Scan(&n)
	// Задачи для воркеров
	jobs := make(chan int, n)

	// Ловим ctrl+C
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	// Запускаем работяг
	for i := 0; i < n; i++ {
		go workerWork(i+1, jobs)
	}

	for {
		// Если не Ctri+C, то пишем в канал рандомное число
		if len(shutdown) == 0 {
			jobs <- rand.Intn(9999)
		} else {
			// Ctrl+C = закрываем канал, тормозим программу
			fmt.Printf("Shutdown signal\n")
			close(jobs)
			break
		}
	}
}
