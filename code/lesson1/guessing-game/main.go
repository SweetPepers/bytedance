package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("the secret number is ", secretNumber)
	fmt.Println("the secret number have been generated ")
	for {
		fmt.Println("please input your guess")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("try again, err : ", err)
			continue
		}

		input = strings.TrimSuffix(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("try again, err : ", err)
			continue
		}

		fmt.Println("your guess is ", guess)

		if guess > secretNumber {
			fmt.Println("your guess is greater than the secret num. Please try again ")
		} else if guess < secretNumber {
			fmt.Println("your guess is less than the secret num. Please try again ")
		} else {
			fmt.Println("success!!")
			break
		}
	}

}
