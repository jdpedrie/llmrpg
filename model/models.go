package model

import (
	"github.com/geldata/gel-go/geltypes"
	v1 "github.com/jdpedrie/llmrpg/genproto/jdpedrie/llmrpg/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CharacterSet []Character

func (c CharacterSet) ToProto() []*v1.Character {
	out := make([]*v1.Character, 0, len(c))
	for _, ch := range c {
		out = append(out, ch.ToProto())
	}

	return out
}

type Character struct {
	ID              geltypes.UUID         `gel:"id"`
	Name            geltypes.OptionalStr  `gel:"name"`
	Description     geltypes.OptionalStr  `gel:"description"`
	Context         []string              `gel:"context"`
	Active          bool                  `gel:"active"`
	MainCharacter   bool                  `gel:"main_character"`
	Skills          CharacterAttributeSet `gel:"skills"`
	Characteristics CharacterAttributeSet `gel:"characteristics"`
	Relationship    CharacterAttributeSet `gel:"relationship"`
}

func (Character) DBType() string {
	return "Character"
}

func (c *Character) FromProto(in *v1.Character) error {
	if err := setUUID(&c.ID, in.Id); err != nil {
		return err
	}

	setOptionalString(&c.Name, in.Name)
	setOptionalString(&c.Description, in.Description)

	c.Context = in.Context
	c.Active = in.Active
	c.MainCharacter = in.MainCharacter

	for _, a := range in.Skills {
		var ca CharacterAttribute
		if err := ca.FromProto(a); err != nil {
			return err
		}

		c.Skills = append(c.Skills, ca)
	}

	for _, a := range in.Characteristics {
		var ca CharacterAttribute
		if err := ca.FromProto(a); err != nil {
			return err
		}

		c.Characteristics = append(c.Characteristics, ca)
	}

	for _, a := range in.Relationship {
		var ca CharacterAttribute
		if err := ca.FromProto(a); err != nil {
			return err
		}

		c.Relationship = append(c.Relationship, ca)
	}

	return nil
}

func (c *Character) ToProto() *v1.Character {
	return &v1.Character{
		Id:              c.ID.String(),
		Name:            singleValue(c.Name.Get),
		Description:     singleValue(c.Description.Get),
		Context:         c.Context,
		Active:          c.Active,
		MainCharacter:   c.MainCharacter,
		Skills:          c.Skills.ToProto(),
		Characteristics: c.Characteristics.ToProto(),
		Relationship:    c.Relationship.ToProto(),
	}
}

type CharacterAttributeSet []CharacterAttribute

func (c CharacterAttributeSet) ToProto() []*v1.CharacterAttribute {
	out := make([]*v1.CharacterAttribute, 0, len(c))
	for _, cas := range c {
		out = append(out, cas.ToProto())
	}

	return out
}

type CharacterAttribute struct {
	ID    geltypes.UUID `gel:"id"`
	Name  string        `gel:"name"`
	Value int16         `gel:"value"`
}

func (CharacterAttribute) DBType() string {
	return "CharacterAttribute"
}

func (a *CharacterAttribute) FromProto(in *v1.CharacterAttribute) error {
	if err := setUUID(&a.ID, in.Id); err != nil {
		return err
	}

	a.Name = in.Name
	a.Value = int16(in.Value)
	return nil
}

func (a *CharacterAttribute) ToProto() *v1.CharacterAttribute {
	return &v1.CharacterAttribute{
		Id:    a.ID.String(),
		Name:  a.Name,
		Value: int32(a.Value),
	}
}

type Game struct {
	ID              geltypes.UUID        `gel:"id"`
	Name            string               `gel:"name"`
	Description     geltypes.OptionalStr `gel:"description"`
	StartingMessage geltypes.OptionalStr `gel:"starting_message"`
	Scenario        geltypes.OptionalStr `gel:"scenario"`
	Objectives      geltypes.OptionalStr `gel:"objectives"`
	Characters      CharacterSet         `gel:"characters"`

	Skills          []string `gel:"skills"`
	Characteristics []string `gel:"characteristics"`
	Relationship    []string `gel:"relationship"`

	IsTemplate bool `gel:"is_template"`
	IsRunning  bool `gel:"is_running"`

	PlaythroughStartTime geltypes.OptionalDateTime `gel:"playthrough_start_time"`
	PlaythroughEndTime   geltypes.OptionalDateTime `gel:"playthrough_end_time"`
	LastActivityTime     geltypes.OptionalDateTime `gel:"last_activity_time"`
	CreateTime           geltypes.OptionalDateTime `gel:"create_time"`
}

func (Game) DBType() string {
	return "Game"
}

func (g *Game) FromProto(in *v1.Game) (err error) {
	if err := setUUID(&g.ID, in.Id); err != nil {
		return err
	}

	setOptionalString(&g.Description, in.Description)
	setOptionalString(&g.StartingMessage, in.StartingMessage)
	setOptionalString(&g.Scenario, in.Scenario)
	setOptionalString(&g.Objectives, in.Objectives)

	for _, item := range in.Characters {
		var c Character
		if err := c.FromProto(item); err != nil {
			return err
		}

		g.Characters = append(g.Characters, c)
	}

	g.Skills = in.Skills
	g.Characteristics = in.Characteristics
	g.Relationship = in.Relationship

	g.IsTemplate = in.IsTemplate
	g.IsRunning = in.IsRunning

	return nil
}

func (g *Game) ToProto() *v1.Game {
	return &v1.Game{
		Id:                   g.ID.String(),
		Name:                 g.Name,
		Description:          singleValue(g.Description.Get),
		StartingMessage:      singleValue(g.StartingMessage.Get),
		Scenario:             singleValue(g.Scenario.Get),
		Objectives:           singleValue(g.Objectives.Get),
		Characters:           g.Characters.ToProto(),
		Skills:               g.Skills,
		Characteristics:      g.Characteristics,
		Relationship:         g.Relationship,
		IsTemplate:           g.IsTemplate,
		IsRunning:            g.IsRunning,
		PlaythroughStartTime: timestamppb.New(singleValue(g.PlaythroughStartTime.Get)),
		PlaythroughEndTime:   timestamppb.New(singleValue(g.PlaythroughEndTime.Get)),
		LastActivityTime:     timestamppb.New(singleValue(g.LastActivityTime.Get)),
		CreateTime:           timestamppb.New(singleValue(g.CreateTime.Get)),
	}
}

type History struct {
	ID      geltypes.UUID `gel:"id"`
	GameID  geltypes.UUID
	Text    string `gel:"text"`
	Choice  string `gel:"choice"`
	Outcome string `gel:"outcome"`
}

func (History) DBType() string {
	return "History"
}

func (h *History) FromProto(in *v1.History) error {
	if err := setUUID(&h.ID, in.Id); err != nil {
		return err
	}

	if err := setUUID(&h.GameID, in.GameId); err != nil {
		return err
	}

	h.Text = in.Text
	h.Choice = in.Choice
	h.Outcome = in.Outcome

	return nil
}

func setUUID(dst *geltypes.UUID, v string) error {
	if v == "" {
		return nil
	}

	out, err := geltypes.ParseUUID(v)
	if err != nil {
		return err
	}

	*dst = out
	return nil
}

func setOptionalString(dst *geltypes.OptionalStr, v string) {
	if v == "" {
		return
	}

	dst.Set(v)
}

func singleValue[T any](in func() (T, bool)) T {
	v, _ := in()
	return v
}
