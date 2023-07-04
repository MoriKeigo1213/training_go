package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 課題2 構造体のスライスを扱ってみよう2
// 入力したゲームIDを担当した開発者を出力してください。
// 出力例:[{0 Ichiro [0]} {1 Jiro [2]} {3 Shiro [0 2]} {5 Rokuro [2 3]}]
// ↑ 改行されていても、配列出力でなくてもOK ただし重複出力はNG
// memo: ループ文のネストは浅く（2階層まで）

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
	fmt.Println("ゲームIDを入力してください")
	var input string
	fmt.Scan(&input)

	inputs := strings.Split(input, " ")
	gameIDs := make([]uint8, len(inputs))
	for i, input := range inputs {
		id, _ := strconv.Atoi(input)
		gameIDs[i] = uint8(id)
	}

	games := getGamesByGameID(gameIDs)
	developers := getDevelopers()

	developersWithGames := findDevelopersByGames(games, developers)

	fmt.Println(developersWithGames)
}

func findDevelopersByGames(games []Game, developers []Developer) []Developer {
	var result []Developer
	for _, game := range games {
		for _, developer := range developers {
			for _, gameID := range developer.GameIDs {
				if game.GameID == gameID && !contains(result, developer) {
					result = append(result, developer)
					break
				}
			}
		}
	}

	return result
}

func contains(developers []Developer, developer Developer) bool {
	for _, d := range developers {
		if d.DeveloperID == developer.DeveloperID {
			return true
		}
	}

	return false
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
