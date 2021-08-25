//ch17/ex17.2/ex17.2.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Money      = 1000
	plusMoney  = 500
	minusMoney = 100
	min        = 0
	max        = 5000
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	money := Money
	for {

		fmt.Printf("1-5 사이의 숫자를 입력하세요 : ")
		n, err := InputIntValue()

		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(5) + 1

		if n < 1 || n > 5 || err != nil {
			fmt.Println(" 숫자를 다시 입력하세요 ")
		} else {
			if n == r {
				money += plusMoney
				fmt.Println(" 축하합니다~~ 잔액은 : %d", money)
			} else {
				money -= minusMoney
				fmt.Println(" 아쉬워요.. 잔액은 : %d", money)
			}
			if 5000 < money {
				fmt.Println(" 게임을 승리하셨습니다. ")
				break
			} else if 0 > money {
				fmt.Println(" 돈이 0원이 되어 게임을 종료합니다. ")
				break
			}
		}
	}

}
