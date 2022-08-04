package config

type OtoConfig struct {
	cacheDir    string
	rootDir     string
	rootFile    string
	targetRoots []string
	logLevel    string
}

func defaultRootFilenames() []string {
	return []string{"workspace.oto", "root.oto"}
}

func defaultTargetFilenames() []string {
	return []string{"target.oto", "project.oto", "package.oto"}
}
