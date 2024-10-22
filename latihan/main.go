package main

import (
	"fmt"
)

// Fungsi countdown
func countdown(seconds int, done chan bool) {
	for i := seconds; i >= 0; i-- {
		fmt.Printf("Countdown: %d\n", i)
	}
	done <- true // Mengirim sinyal bahwa countdown selesai
}

func main() {
	// Channel untuk menandakan bahwa countdown selesai
	done := make(chan bool)

	// Menjalankan countdown di dalam goroutine
	go countdown(10, done) // Misalnya countdown 10 detik

	// Menunggu sinyal dari goroutine
	<-done

	// Countdown selesai
	fmt.Println("Countdown selesai!")
}
