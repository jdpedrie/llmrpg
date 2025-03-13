package game

import "github.com/jdpedrie/llmrpg/models"

type Action struct {
	SystemPrompt string `json:"system_prompt"`
	Choice       string `json:"choice"`
	Success      bool   `json:"success"`

	Characters []models.Character `json:"characters"`
	History    []string           `json:"history"`
	Context    []string           `json:"context"`
}
