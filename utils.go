package utils

import (
	"bitbucket.org/kardianos/osext"
	"path"
)

// Returns the name of the excecutable
func GetExecutableName() string {
	executablePath, err := osext.Executable()
	if err != nil {
		return "(UNKNOWN)"
	} else {
		return path.Base(executablePath)
	}
}
