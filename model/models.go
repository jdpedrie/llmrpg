package model

import (
	"time"

	"github.com/google/uuid"
	v1 "github.com/jdpedrie/llmrpg/genproto/jdpedrie/llmrpg/v1"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Character represents a character in the game
type Character struct {
	ID              uuid.UUID
	Name            string
	Description     string
	Context         []string
	Active          bool
	MainCharacter   bool
	Skills          []CharacterAttribute
	Characteristics []CharacterAttribute
	Relationship    []CharacterAttribute
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// FromDBCharacter converts a database character to a model character
func FromDBCharacter(c postgres.Character, skills, characteristics, relationships []postgres.CharacterAttribute) Character {
	char := Character{
		ID:              c.ID,
		Name:            c.Name,
		Description:     postgres.StringFromText(c.Description),
		Context:         c.Context,
		Active:          c.Active,
		MainCharacter:   c.MainCharacter,
		CreatedAt:       c.CreatedAt.Time,
		UpdatedAt:       c.UpdatedAt.Time,
		Skills:          make([]CharacterAttribute, 0, len(skills)),
		Characteristics: make([]CharacterAttribute, 0, len(characteristics)),
		Relationship:    make([]CharacterAttribute, 0, len(relationships)),
	}

	for _, s := range skills {
		char.Skills = append(char.Skills, FromDBCharacterAttribute(s))
	}

	for _, c := range characteristics {
		char.Characteristics = append(char.Characteristics, FromDBCharacterAttribute(c))
	}

	for _, r := range relationships {
		char.Relationship = append(char.Relationship, FromDBCharacterAttribute(r))
	}

	return char
}

// ToProto converts a Character to its Proto equivalent
func (c *Character) ToProto() *v1.Character {
	return &v1.Character{
		Id:              c.ID.String(),
		Name:            c.Name,
		Description:     c.Description,
		Context:         c.Context,
		Active:          c.Active,
		MainCharacter:   c.MainCharacter,
		Skills:          CharacterAttributesToProto(c.Skills),
		Characteristics: CharacterAttributesToProto(c.Characteristics),
		Relationship:    CharacterAttributesToProto(c.Relationship),
	}
}

// FromProto populates a Character from its Proto equivalent
func (c *Character) FromProto(in *v1.Character) error {
	var err error
	if in.Id != "" {
		c.ID, err = uuid.Parse(in.Id)
		if err != nil {
			return err
		}
	}

	c.Name = in.Name
	c.Description = in.Description
	c.Context = in.Context
	c.Active = in.Active
	c.MainCharacter = in.MainCharacter

	// Populate skills
	c.Skills = make([]CharacterAttribute, 0, len(in.Skills))
	for _, s := range in.Skills {
		var attr CharacterAttribute
		if err := attr.FromProto(s); err != nil {
			return err
		}
		c.Skills = append(c.Skills, attr)
	}

	// Populate characteristics
	c.Characteristics = make([]CharacterAttribute, 0, len(in.Characteristics))
	for _, ch := range in.Characteristics {
		var attr CharacterAttribute
		if err := attr.FromProto(ch); err != nil {
			return err
		}
		c.Characteristics = append(c.Characteristics, attr)
	}

	// Populate relationships
	c.Relationship = make([]CharacterAttribute, 0, len(in.Relationship))
	for _, r := range in.Relationship {
		var attr CharacterAttribute
		if err := attr.FromProto(r); err != nil {
			return err
		}
		c.Relationship = append(c.Relationship, attr)
	}

	return nil
}

// CharacterAttribute represents an attribute of a character
type CharacterAttribute struct {
	ID    uuid.UUID
	Name  string
	Value int16
	Type  string
}

// FromDBCharacterAttribute converts a database character attribute to a model character attribute
func FromDBCharacterAttribute(ca postgres.CharacterAttribute) CharacterAttribute {
	return CharacterAttribute{
		ID:    ca.ID,
		Name:  ca.Name,
		Value: ca.Value,
		Type:  ca.AttributeType,
	}
}

// ToProto converts a CharacterAttribute to its Proto equivalent
func (a *CharacterAttribute) ToProto() *v1.CharacterAttribute {
	return &v1.CharacterAttribute{
		Id:    a.ID.String(),
		Name:  a.Name,
		Value: int32(a.Value),
	}
}

// FromProto populates a CharacterAttribute from its Proto equivalent
func (a *CharacterAttribute) FromProto(in *v1.CharacterAttribute) error {
	var err error
	if in.Id != "" {
		a.ID, err = uuid.Parse(in.Id)
		if err != nil {
			return err
		}
	}

	a.Name = in.Name
	a.Value = int16(in.Value)
	return nil
}

// CharacterAttributesToProto converts a slice of CharacterAttribute to a slice of proto CharacterAttribute
func CharacterAttributesToProto(in []CharacterAttribute) []*v1.CharacterAttribute {
	out := make([]*v1.CharacterAttribute, 0, len(in))
	for _, attr := range in {
		out = append(out, attr.ToProto())
	}
	return out
}

// Game represents a game in the system
type Game struct {
	ID                   uuid.UUID
	Name                 string
	Description          string
	StartingMessage      string
	Scenario             string
	Objectives           string
	Characters           []Character
	Inventory            []InventoryItem
	Skills               []string
	Characteristics      []string
	Relationship         []string
	IsTemplate           bool
	IsRunning            bool
	PlaythroughStartTime *time.Time
	PlaythroughEndTime   *time.Time
	LastActivityTime     *time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

// FromDBGame converts a database game to a model game
func FromDBGame(g postgres.Game, characters []Character, inventory []InventoryItem) Game {
	return Game{
		ID:                   g.ID,
		Name:                 g.Name,
		Description:          postgres.StringFromText(g.Description),
		StartingMessage:      postgres.StringFromText(g.StartingMessage),
		Scenario:             postgres.StringFromText(g.Scenario),
		Objectives:           postgres.StringFromText(g.Objectives),
		Characters:           characters,
		Inventory:            inventory,
		Skills:               g.Skills,
		Characteristics:      g.Characteristics,
		Relationship:         g.Relationship,
		IsTemplate:           g.IsTemplate,
		IsRunning:            g.IsRunning,
		PlaythroughStartTime: postgres.TimeFromTimestamptz(g.PlaythroughStartTime),
		PlaythroughEndTime:   postgres.TimeFromTimestamptz(g.PlaythroughEndTime),
		LastActivityTime:     postgres.TimeFromTimestamptz(g.LastActivityTime),
		CreatedAt:            g.CreatedAt.Time,
		UpdatedAt:            g.UpdatedAt.Time,
	}
}

// ToProto converts a Game to its Proto equivalent
func (g *Game) ToProto() *v1.Game {
	characters := make([]*v1.Character, 0, len(g.Characters))
	for _, c := range g.Characters {
		characters = append(characters, c.ToProto())
	}

	inventory := make([]*v1.InventoryItem, 0, len(g.Inventory))
	for _, i := range g.Inventory {
		inventory = append(inventory, i.ToProto())
	}

	game := &v1.Game{
		Id:              g.ID.String(),
		Name:            g.Name,
		Description:     g.Description,
		StartingMessage: g.StartingMessage,
		Scenario:        g.Scenario,
		Objectives:      g.Objectives,
		Characters:      characters,
		Inventory:       inventory,
		Skills:          g.Skills,
		Characteristics: g.Characteristics,
		Relationship:    g.Relationship,
		IsTemplate:      g.IsTemplate,
		IsRunning:       g.IsRunning,
		CreateTime:      timestamppb.New(g.CreatedAt),
	}

	if g.PlaythroughStartTime != nil {
		game.PlaythroughStartTime = timestamppb.New(*g.PlaythroughStartTime)
	}

	if g.PlaythroughEndTime != nil {
		game.PlaythroughEndTime = timestamppb.New(*g.PlaythroughEndTime)
	}

	if g.LastActivityTime != nil {
		game.LastActivityTime = timestamppb.New(*g.LastActivityTime)
	}

	return game
}

// FromProto populates a Game from its Proto equivalent
func (g *Game) FromProto(in *v1.Game) error {
	var err error
	if in.Id != "" {
		g.ID, err = uuid.Parse(in.Id)
		if err != nil {
			return err
		}
	}

	g.Name = in.Name
	g.Description = in.Description
	g.StartingMessage = in.StartingMessage
	g.Scenario = in.Scenario
	g.Objectives = in.Objectives
	g.Skills = in.Skills
	g.Characteristics = in.Characteristics
	g.Relationship = in.Relationship
	g.IsTemplate = in.IsTemplate
	g.IsRunning = in.IsRunning

	// Populate characters
	g.Characters = make([]Character, 0, len(in.Characters))
	for _, c := range in.Characters {
		var char Character
		if err := char.FromProto(c); err != nil {
			return err
		}
		g.Characters = append(g.Characters, char)
	}

	// Populate inventory
	g.Inventory = make([]InventoryItem, 0, len(in.Inventory))
	for _, i := range in.Inventory {
		var item InventoryItem
		if err := item.FromProto(i); err != nil {
			return err
		}
		g.Inventory = append(g.Inventory, item)
	}

	return nil
}

// InventoryItem represents an item in the game inventory
type InventoryItem struct {
	ID          uuid.UUID
	Name        string
	Description string
	Active      bool
	GameID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// FromDBInventoryItem converts a database inventory item to a model inventory item
func FromDBInventoryItem(i postgres.InventoryItem) InventoryItem {
	return InventoryItem{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Active:      i.Active,
		GameID:      i.GameID,
		CreatedAt:   i.CreatedAt.Time,
		UpdatedAt:   i.UpdatedAt.Time,
	}
}

// ToProto converts an InventoryItem to its Proto equivalent
func (i *InventoryItem) ToProto() *v1.InventoryItem {
	return &v1.InventoryItem{
		Id:          i.ID.String(),
		Name:        i.Name,
		Description: i.Description,
		Active:      i.Active,
	}
}

// FromProto populates an InventoryItem from its Proto equivalent
func (i *InventoryItem) FromProto(in *v1.InventoryItem) error {
	var err error
	if in.Id != "" {
		i.ID, err = uuid.Parse(in.Id)
		if err != nil {
			return err
		}
	}

	i.Name = in.Name
	i.Description = in.Description
	i.Active = in.Active
	return nil
}

// History represents a history entry in the game
type History struct {
	ID        uuid.UUID
	GameID    uuid.UUID
	Text      string
	Choice    string
	Outcome   string
	CreatedAt time.Time
}

// FromDBHistory converts a database history to a model history
func FromDBHistory(h postgres.History) History {
	return History{
		ID:        h.ID,
		GameID:    h.GameID,
		Text:      h.Text,
		Choice:    h.Choice,
		Outcome:   h.Outcome,
		CreatedAt: h.CreatedAt.Time,
	}
}
