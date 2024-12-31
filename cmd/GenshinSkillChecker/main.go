package main

import (
	"context"
	"fmt"

	"github.com/YutoOkawa/GenshinSkillChecker/pkg/skillchecker"
	"github.com/manifoldco/promptui"
)

func main() {
	ctx := context.Background()
	skillchecker := skillchecker.NewSkillChecker()
	if err := skillchecker.InitializeData(ctx); err != nil {
		panic(err)
	}

	characters, err := skillchecker.GetCharacters(ctx, "862615394")
	if err != nil {
		panic(err)
	}

	characterPrompt := promptui.Select{
		Label: "キャラクターを選択してください",
		Items: func() []string {
			var items []string
			for _, character := range characters {
				items = append(items, character.CharacterName)
			}
			return items
		}(),
	}

	_, characterName, err := characterPrompt.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(characterName)
}
