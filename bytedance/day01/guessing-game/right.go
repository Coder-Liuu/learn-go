package main

import (
	"fmt"
	"math/rand"
	"time"

	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 使用时间戳作为随机数种子
	rand.Seed(time.Now().UnixNano())

	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("secretNumber is", secretNumber)

	fmt.Println("Please input your guess")

	for {
		reader := bufio.NewReader(os.Stdin)
		// 读取一行，这里必须是单引号
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		// 去掉换行符
		input = strings.TrimSuffix(input, "\r\n")
		// 转换为数字
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an interage value")
			continue
		}

		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number, Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number, Please try again")
		} else {
			fmt.Println("Correct, You Legend")
			break
		}
	}
}
