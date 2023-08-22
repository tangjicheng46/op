package version

import "strconv"

const (
	MajorVersion int = 0
	MinorVersion int = 0
	PatchVersion int = 1
)

func GetVersion() string {
	versionString := strconv.Itoa(MajorVersion) + "." + strconv.Itoa(MinorVersion) + "." + strconv.Itoa(PatchVersion)
	return versionString
}
