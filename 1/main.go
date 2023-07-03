package main

import "fmt"

// 課題.1 「構造体のスライスを扱ってみよう」
// 変数`games`の最後に適当なデータを入力して出力させよう
// 標準入力してください fmt.Scan(&変数名)

type game struct {
	GameID   uint8
	GameName string
}

func main() {
	games := []game{
		{
			GameID:   0,
			GameName: "mario",
		},
		{
			GameID:   1,
			GameName: "pokemon",
		},
	}

	insertGame()

	fmt.Println(games)
	//fmt.Println()
}

func insertGame() {

}
