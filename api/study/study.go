package study

import (
	"encoding/json"
	"os"
)

func CreateAndSaveNewStudy(name string, path string) (*Study, error) {
	study := Study{
		Name:     name,
		Filepath: path,
		Media:    []Media{},
		Codes:    []Code{},
		Tags:     []Tag{},
	}

	data, e := json.Marshal(study)
	if e != nil {
		return nil, e
	}

	e = os.WriteFile(name+".json", data, 0666)
	if e != nil {
		return nil, e
	}

	return &study, nil
}

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
