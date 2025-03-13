package game

import (
	"encoding/json"
	"strings"

	"github.com/jdpedrie/llmrpg/models"
)

type Gamemaster struct {
	Choices          []Choice `json:"choice"`
	AllowAnyResponse bool     `json:"allow_any_response"`
	GameOver         bool     `json:"game_over"`

	ModifiedCharacters []models.Character `json:"modified_characters"`
	AddedHistory       []string           `json:"added_history"`
	AddedContext       []string           `json:"added_context"`
}

type Choice struct {
	Message       string        `json:"message"`
	Hint          string        `json:"hint"`
	SuccessChance SuccessChance `json:"success_chance"`
}

type SuccessChance int8

const (
	SuccessChanceNone SuccessChance = iota
	SuccessChangeVeryLow
	SuccessChanceLow
	SuccessChanceMedium
	SuccessChanceHigh
	SuccessChanceVeryHigh
	SuccessChanceGuaranteed
)

var (
	successChanceToString = map[SuccessChance]string{
		SuccessChanceNone:       "none",
		SuccessChangeVeryLow:    "very_low",
		SuccessChanceLow:        "low",
		SuccessChanceMedium:     "medium",
		SuccessChanceHigh:       "high",
		SuccessChanceVeryHigh:   "very_high",
		SuccessChanceGuaranteed: "guaranteed",
	}
	stringToSuccessChance = map[string]SuccessChance{}
)

func init() {
	for k, v := range successChanceToString {
		stringToSuccessChance[v] = k
	}
}

func (s *SuccessChance) UnmarshalJSON(in []byte) error {
	str := strings.Trim(string(in), `"`)
	*s = stringToSuccessChance[str]
	return nil
}

func (s SuccessChance) MarshalJSON() ([]byte, error) {
	return json.Marshal(successChanceToString[s])
}
