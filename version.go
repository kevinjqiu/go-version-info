package version

import "fmt"

type BuildInfo struct {
	Version string
	Date    string
	Commit  string
}

func (bi BuildInfo) String() string {
	return fmt.Sprintf("%s (Commit=%s, BuildDate=%s)", bi.Version, bi.Commit, bi.Date)
}
