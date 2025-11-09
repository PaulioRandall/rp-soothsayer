package api

var MockDirInfoData []DirInfo = []DirInfo{
	DirInfo{
		DirName:      "Documents",
		IsProjectDir: false,
	},
	DirInfo{
		DirName:      "Downloads",
		IsProjectDir: false,
	},
	DirInfo{
		DirName:      "Pictures",
		IsProjectDir: false,
	},
	DirInfo{
		DirName:      "Videos",
		IsProjectDir: false,
	},
	DirInfo{
		DirName:      "Books",
		IsProjectDir: false,
	},
	DirInfo{
		DirName:      "jojo-analysis",
		IsProjectDir: true,
	},
	DirInfo{
		DirName:      "one-punch-man-analysis",
		IsProjectDir: true,
	},
}
