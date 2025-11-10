package api

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type DirInfo struct {
	DirName      string
	DirPath      string
	IsProjectDir bool
}

func ReadDir(path string) ([]DirInfo, error) {
	absPath, e := filepath.Abs(path)
	if e != nil {
		return nil, e
	}

	entries, e := os.ReadDir(absPath)
	if e != nil {
		return nil, e
	}

	dirs := filterDirs(entries)
	dirInfoData := toDirInfoData(dirs, absPath)

	return dirInfoData, nil
}

func filterDirs(entries []fs.DirEntry) []fs.DirEntry {
	result := []fs.DirEntry{}

	for _, entry := range entries {
		if isPresentableDir(entry) {
			result = append(result, entry)
		}
	}

	return result
}

func isPresentableDir(entry fs.DirEntry) bool {
	return entry.IsDir() && !isHiddenFile(entry.Name())
}

func isHiddenFile(name string) bool {
	return strings.HasPrefix(name, ".")
}

func toDirInfoData(dirs []fs.DirEntry, parentPath string) []DirInfo {
	result := []DirInfo{}

	for _, dir := range dirs {
		dirInfo := dirEntrytoDirInfo(dir, parentPath)
		result = append(result, dirInfo)
	}

	return result
}

func dirEntrytoDirInfo(dir fs.DirEntry, parentPath string) DirInfo {
	return DirInfo{
		DirName:      dir.Name(),
		DirPath:      filepath.Join(parentPath, dir.Name()),
		IsProjectDir: false,
	}
}
