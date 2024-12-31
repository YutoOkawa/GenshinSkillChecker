package skillchecker

import (
	"context"
	"fmt"

	"github.com/YutoOkawa/EnkaNetworkGoClient/pkg/client"
	"github.com/YutoOkawa/EnkaNetworkGoClient/pkg/model"
)

type SkillChecker struct {
	Client           *client.Client
	characterData    map[string]model.CharacterData
	localizationData map[string]map[string]string
}

func NewSkillChecker() *SkillChecker {
	return &SkillChecker{
		Client: client.NewClient(),
	}
}

func (s *SkillChecker) InitializeData(ctx context.Context) error {
	characterData, err := s.Client.GetCharacterData(ctx)
	if err != nil {
		return err
	}
	s.characterData = characterData

	localizationData, err := s.Client.GetLocalizationData(ctx)
	if err != nil {
		return err
	}
	s.localizationData = localizationData

	return nil
}

func (s *SkillChecker) GetCharacters(ctx context.Context, uid string) ([]Character, error) {
	allData, err := s.Client.GetAllData(ctx, uid)
	if err != nil {
		return nil, err
	}

	characters := make([]Character, len(allData.AvatarInfoList))
	language := "ja"

	for avatarNum, avatarInfo := range allData.AvatarInfoList {
		avatarId := fmt.Sprintf("%d", avatarInfo.AvatarID)
		characterNameHash := fmt.Sprintf("%d", s.characterData[avatarId].NameTextMapHash)
		characters[avatarNum].CharacterName = s.localizationData[language][characterNameHash]
		for skillNum, skillOrder := range s.characterData[avatarId].SkillOrder {
			skillHash := fmt.Sprintf("%d", skillOrder)
			skillLevel := avatarInfo.SkillLevelMap[skillHash]
			if skillNum == 0 {
				characters[avatarNum].NormalAttackLevel = skillLevel
			}
			if skillNum == 1 {
				characters[avatarNum].SkillLevel = skillLevel
			}
			if skillNum == 2 {
				characters[avatarNum].UltimateSkillLevel = skillLevel
			}
		}
	}

	return characters, nil
}
