package api

type DirInfo struct {
	DirName      string
	IsProjectDir bool
}

func ReadDir(path string) []DirInfo {
	// TODO

	// MOCK
	return MockDirInfoData
}
