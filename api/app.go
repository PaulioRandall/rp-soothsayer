package api

import (
	"context"

	"path/filepath"
)

type App struct {
	ctx context.Context
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
