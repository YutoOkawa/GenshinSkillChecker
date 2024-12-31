package main

import (
	"context"
	"fmt"

	"github.com/YutoOkawa/GenshinSkillChecker/pkg/skillchecker"
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

	for _, character := range characters {
		fmt.Print(character)
	}
}
