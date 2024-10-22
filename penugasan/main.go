package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numSensors   = 3
	sensorDelay  = 1 * time.Second // Delay pengiriman data sensor
	readInterval = 2 * time.Second // Interval untuk membaca data
	timeout      = 5 * time.Second // Timeout untuk sensor
)

// Sensor struct untuk menyimpan informasi sensor
type Sensor struct {
	name string
	data chan float64
}

// Fungsi untuk mensimulasikan pengiriman data sensor
func (s *Sensor) generateData(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Simulasi pengiriman data suhu, kelembaban, atau tekanan
		s.data <- rand.Float64() * 100 // Mengirim data random antara 0 - 100
		time.Sleep(sensorDelay)
	}
}

// Fungsi untuk mengambil data dari sensor
func readSensorData(sensor *Sensor) {
	ticker := time.NewTicker(readInterval)
	defer ticker.Stop()

	for {
		select {
		case value := <-sensor.data:
			fmt.Printf("Data dari %s: %.2f\n", sensor.name, value)
		case <-time.After(timeout):
			fmt.Printf("Sensor %s timeout\n", sensor.name)
			return
		case <-ticker.C:
			// Jika kita tidak mendapatkan data dalam waktu timeout, kita akan mengeluarkan pesan
			fmt.Printf("Tidak ada pembaruan dari %s dalam interval 2 detik.\n", sensor.name)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Inisialisasi seed untuk random
	var wg sync.WaitGroup

	sensors := []Sensor{
		{name: "Sensor Suhu", data: make(chan float64, 1)},
		{name: "Sensor Kelembaban", data: make(chan float64, 1)},
		{name: "Sensor Tekanan", data: make(chan float64, 1)},
	}

	// Mulai goroutine untuk setiap sensor
	for i := range sensors {
		wg.Add(1)
		go sensors[i].generateData(&wg)
		go readSensorData(&sensors[i])
	}

	// Tunggu goroutine selesai
	go func() {
		wg.Wait()
		for _, sensor := range sensors {
			close(sensor.data) // Menutup channel data untuk setiap sensor
		}
	}()

	// Tunggu 30 detik untuk melihat hasil (akan menghentikan program setelah itu)
	time.Sleep(30 * time.Second)
	fmt.Println("Program selesai.")
}
