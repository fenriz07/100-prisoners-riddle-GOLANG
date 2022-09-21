package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fenriz07/100-prisoners-riddle-GOLANG/app/usecase"
)

func main() {

	stats := make(map[string]int)
	stats["fail"] = 0
	stats["success"] = 0

	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {
		err := usecase.GameUseCase()

		if err == nil {
			stats["success"] = stats["success"] + 1
		} else {
			stats["fail"] = stats["fail"] + 1
		}
	}

	s := float32(stats["success"]) / 100
	f := float32(stats["fail"]) / 100

	fmt.Println("RESULT")
	fmt.Printf("Success %f \n", s)
	fmt.Printf("Failed %f \n", f)
}
