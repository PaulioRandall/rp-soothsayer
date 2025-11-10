package api

import (
	"context"
	"path/filepath"

	//"soothsayer/api/study/validate"
	"soothsayer/api/study"
)

type App struct {
	ctx   context.Context
	study *study.Study
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// File browser functions

func (a *App) ParentPath(path string) string {
	return filepath.Dir(path)
}

func (a *App) AbsPath(path string) (string, error) {
	return filepath.Abs(path)
}

func (a *App) ReadDir(path string) ([]DirInfo, error) {
	return ReadDir(path)
}

func (a *App) LoadStudy(path string) (*study.Study, error) {
	// TODO
	// 1: read study as map
	// 2: validate study sturcture and content
	// 3: unmarshall map into Study struct

	return nil, nil
}
