package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10000)
	for {
		if num > 999 && num < 10000 {
			break
		}
	}
	*p = num
}

func GetNum(num int, s []int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func OnGame(randSlice []int) {
	var keyNum int
	keySlice := make([]int, 4)
	for {
		fmt.Printf("请输入一个4位数：")
		fmt.Scan(&keyNum)

		GetNum(keyNum, keySlice)

		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了一些！\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了一些！\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了！\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Printf("恭喜你猜对了！\n")
			break
		}
	}
}

func main() {
	var randNum int

	CreateNum(&randNum)

	randSlice := make([]int, 4)
	GetNum(randNum, randSlice)

	OnGame(randSlice)
}
