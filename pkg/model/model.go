package model

type Workspace struct {
	Name        string
	Description string

	Landscape    Landscape
	Perspectives []Perspective
}

type Landscape struct {
	Name        string
	Description string
}

type Perspective struct {
	Name        string
	Description string
}

type Entity struct {
	Type        string
	Name        string
	Description string
	Tech        string
	Shape       string
	Relations   []struct {
		Source      string
		Target      string
		Description string
		Tech        string
	}
}
