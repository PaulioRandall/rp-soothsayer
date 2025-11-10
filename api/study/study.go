package study

type Study struct {
	Name     string  `json:"name"`
	Filepath string  `json:"filepath"`
	Media    []Media `json:"media"`
	Codes    []Code  `json:"code"`
	Tags     []Tag   `json:"tag"`
}

type StudyEntity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TaggableStudyEntity struct {
	StudyEntity
	Tags string `json:"tags"`
}

type Media struct {
	TaggableStudyEntity
	Filepath     string         `json:"filepath"`
	Sections     []MediaSection `json:"sections"`
	Observations []Observation  `json:"observations"`
}

type MediaSection struct {
	TaggableStudyEntity
	Start int `json:"start"`
	End   int `json:"end"`
}

type Observation struct {
	TaggableStudyEntity
	Start       int      `json:"start"`
	Quote       string   `json:"quote"`
	Description string   `json:"description"`
	Codes       []string `json:"codes"`
}

type Code struct {
	TaggableStudyEntity
	Description string `json:"description"`
}

type Tag struct {
	StudyEntity
	Description string `json:"description"`
}
