package api

type DirInfo struct {
	ParentPath   string
	DirName      string
	IsProjectDir bool
}

func ReadDir(path string) []DirInfo {
	// TODO

	// MOCK
	return MockDirInfoData
}
