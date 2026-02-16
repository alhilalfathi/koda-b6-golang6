package main

import (
	"fmt"
	"time"
)

type Data struct {
	name string
	wait int
}

func main() {
	data := []Data{
		{
			name: "Adit",
			wait: 5,
		},
		{
			name: "Anton",
			wait: 3,
		},
		{
			name: "Reza",
			wait: 2,
		},
		{
			name: "Arif",
			wait: 4,
		},
		{
			name: "Aska",
			wait: 2,
		},
	}

	var ch = make(chan string)

	go func() {
		for _, d := range data {
			time.Sleep(time.Second * time.Duration(d.wait))
			ch <- d.name
		}

	}()

	for name := range ch {
		fmt.Printf("Pesanan %s selesai\n", name)
	}
}
