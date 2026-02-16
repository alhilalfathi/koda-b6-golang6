package main

import (
	"fmt"
	"koda-b6-golang6/model"
	"time"
)

func main() {
	data := []model.Data{
		{
			Name: "Adit",
			Wait: 2,
		},
		{
			Name: "Anton",
			Wait: 3,
		},
		{
			Name: "Reza",
			Wait: 2,
		},
		{
			Name: "Arif",
			Wait: 4,
		},
		{
			Name: "Aska",
			Wait: 2,
		},
	}

	var ch = make(chan string)

	go func() {
		fmt.Println("Menunggu Antrian...")
		for _, d := range data {
			time.Sleep(time.Second * time.Duration(d.Wait))
			ch <- d.Name
		}
		close(ch)
	}()

	for name := range ch {
		fmt.Printf("Pesanan %s selesai\n", name)
	}
	fmt.Println("\nAntrian selesai")
}
