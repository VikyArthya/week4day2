package main

import (
	"fmt"
	"time"
)

//buatkan 3 function dimana function ini ada parameter chanel dan msg, kemudian dijalankan di goroutine berbeda
// buatkan juga masing" punya chanel sendiri, tampilkan pesan dari 3 func ini yg paling cepat responnya

func msg1(message1 string, ch1 chan string) {
	time.Sleep(3 * time.Millisecond)
	ch1 <- message1
}

func msg2(message2 string, ch2 chan string) {
	time.Sleep(7 * time.Millisecond)
	ch2 <- message2
}

func msg3(message3 string, ch3 chan string) {
	time.Sleep(5 * time.Millisecond)
	ch3 <- message3
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go msg1("chanel 1", ch1)
	go msg2("chanel 2", ch2)
	go msg3("chanel 3", ch3)

	select {
	case data := <-ch1:
		fmt.Println("data diterima dari ", data)
	case data := <-ch2:
		fmt.Println("data diterima dari ", data)
	case data := <-ch3:
		fmt.Println("data diterima dari ", data)

	}
}
