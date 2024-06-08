package types

import "strings"

type Plot struct {
	Id                 string
	Name               string
	StartX             int
	StartZ             int
	EndX               int
	EndZ               int
	OwnerDiscordId     string
	OwnerMinecraftUuid string
	OwnerMinecraftName string
	Software           string
	NotesRaw           string
}

func (p *Plot) GetNotes() []Note {
	var notes []Note

	linesRaw := strings.Split(p.NotesRaw, "¦")
	for _, line := range linesRaw {
		parts := strings.Split(line, "§")
		notes = append(notes, Note{
			Message:   parts[0],
			CreatedAt: parts[1],
			CreatedBy: parts[2],
		})
	}

	return notes
}

func (p *Plot) SetNotes(notes []Note) {
	var out []string

	for _, note := range notes {
		out = append(out, strings.Join([]string{note.Message, note.CreatedAt, note.CreatedBy}, "§"))
	}

	p.NotesRaw = strings.Join(out, "¦")
}

type Note struct {
	Message   string
	CreatedAt string
	CreatedBy string
}

type User struct {
	Uuid string `json:"uuid"`
}
