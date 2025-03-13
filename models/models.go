package models

import (
	"time"

	"github.com/geldata/gel-go/geltypes"
)

type Character struct {
	ID              geltypes.UUID `gel:"id"`
	Name            string        `gel:"id"`
	Description     string        `gel:"description"`
	Context         []string      `gel:"context"`
	Active          bool          `gel:"active"`
	MainCharacter   bool          `gel:"main_character"`
	Stats           []Stat        `gel:"stats"`
	Characteristics []Stat        `gel:"characteristice"`
	Relationship    []Stat        `gel:"relationship"`
}

type Stat struct {
	Name  string `gel:"name"`
	Value int    `gel:"value"`
}

type Game struct {
	ID                   geltypes.UUID `gel:"id"`
	Name                 string        `gel:"name"`
	Description          string        `gel:"description"`
	Characters           []Character   `gel:"characters"`
	StartingMessage      string        `gel:"starting_message"`
	Scenario             string        `gel:"scenario"`
	Objectives           string        `gel:"objectives"`
	IsTemplate           bool          `gel:"is_template"`
	IsRunning            bool          `gel:"is_running"`
	PlaythroughStartTime *time.Time    `gel:"playthrough_start_time"`
	PlaythroughEndTime   *time.Time    `gel:"playthrough_end_time"`
	LastActivityTime     *time.Time    `gel:"last_activity_time"`
	CreateTime           *time.Time    `gel:"create_time"`
}
