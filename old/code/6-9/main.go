package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumNumbers() (int, error) {
	fmt.Println("Please enter numbers one by one, followed by Enter. Enter 0 to end input:")
	var sum int
	// 標準入力から数値を読み込んでメモリに格納
	reader := bufio.NewReader(os.Stdin)

	for {
		// 入力された数値を文字列として取得
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}
		// 入力された文字列からスペースを削除
		input = strings.TrimSpace(input)

		// 0を入力したらループを抜ける
		if input == "0" {
			break
		}

		// 入力された文字列を数値に変換
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", input)
			continue
		}

		sum += num
	}
	return sum, nil
}

func main() {
	sum, err := sumNumbers()
	if err != nil {
		// 標準エラー出力にエラーメッセージを表示し、プログラムを終了
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sum: %d\n", sum)
}
