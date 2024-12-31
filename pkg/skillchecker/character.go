package skillchecker

import "fmt"

type Character struct {
	CharacterName      string
	NormalAttackLevel  int
	SkillLevel         int
	UltimateSkillLevel int
}

func (c Character) String() string {
	return fmt.Sprintf("%s\n* 通常攻撃:   %d\n* 元素スキル: %d\n* 元素爆発:   %d\n", c.CharacterName, c.NormalAttackLevel, c.SkillLevel, c.UltimateSkillLevel)
}
