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

	characterIndex, characterName, err := characterPrompt.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(characterName + "を選択しました。")

	skillPronmpt := promptui.Select{
		Label: "育成したいスキルを選択してください",
		Items: []string{"通常攻撃", "元素スキル", "元素爆発"},
	}

	skillIndex, skillName, err := skillPronmpt.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(skillName + "を選択しました。")

	currentLevel := 0
	if skillIndex == 0 {
		currentLevel = characters[characterIndex].NormalAttackLevel
	} else if skillIndex == 1 {
		currentLevel = characters[characterIndex].SkillLevel
	} else {
		currentLevel = characters[characterIndex].UltimateSkillLevel
	}

	fmt.Println("現在のレベル: ", currentLevel)

	// 育成したいレベルを入力してください
	levelPrompt := promptui.Prompt{
		Label:   "育成したいレベルを入力してください",
		Default: fmt.Sprintf("%d", currentLevel),
	}

	levelInput, err := levelPrompt.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("育成したいレベル: ", levelInput)

	// 各レベルに必要な素材の計算

	// 必要な素材を提示
}
