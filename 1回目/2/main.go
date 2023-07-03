package main

import "fmt"

// 課題2 構造体のスライスを扱ってみよう2
// 入力したゲームIDを担当した開発者を出力してください。
// 出力例:[{0 Ichiro [0]} {1 Jiro [2]} {3 Shiro [0 2]} {5 Rokuro [2 3]}]
// ↑ 改行されていても、配列出力でなくてもOK ただし重複出力はNG

type Game struct {
	GameID   uint8
	GameName string
}

type Developer struct {
	DeveloperID   uint8
	DeveloperName string
	GameIDs       []uint8
}

func main() {
	games := getGamesByGameID([]uint8{0, 2})
	developers := getDevelopers()

}

// この関数の処理はイジらないでください。
func getGamesByGameID(gameIDs []uint8) []Game {
	games := []Game{
		{
			GameID:   0,
			GameName: "mario",
		}, {
			GameID:   1,
			GameName: "pokemon",
		},
		{
			GameID:   2,
			GameName: "pikmin",
		},
	}

	selectedGames := make([]Game, 0, len(gameIDs))

	for _, gameID := range gameIDs {
		for _, game := range games {
			if game.GameID == gameID {
				selectedGames = append(selectedGames, Game{
					GameID:   game.GameID,
					GameName: game.GameName,
				})
			}
		}
	}

	return selectedGames
}

// この関数の処理はイジらないでください。
func getDevelopers() []Developer {
	developers := []Developer{
		{
			DeveloperID:   0,
			DeveloperName: "Ichiro",
			GameIDs:       []uint8{0},
		},
		{
			DeveloperID:   1,
			DeveloperName: "Jiro",
			GameIDs:       []uint8{2},
		},
		{
			DeveloperID:   2,
			DeveloperName: "Saburo",
			GameIDs:       []uint8{1},
		},
		{
			DeveloperID:   3,
			DeveloperName: "Shiro",
			GameIDs:       []uint8{0, 2},
		},
		{
			DeveloperID:   4,
			DeveloperName: "Goro",
			GameIDs:       []uint8{3, 4},
		},
		{
			DeveloperID:   5,
			DeveloperName: "Rokuro",
			GameIDs:       []uint8{2, 3},
		},
	}

	return developers
}
