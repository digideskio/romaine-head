package leaf

import (
	"strings"

	"github.com/danopia/romaine-head/common"
)

func getVersion() map[string]string {
	output, _ := common.RunCmd("croutonversion")
	lines := strings.Split(output, "\n")
	fields := make(map[string]string)

	for _, line := range lines {
		pieces := strings.Split(line, ": ")
		fields[pieces[0]] = pieces[1]
	}

	return fields
}

func runCommand(path string, args []string) string {
	output, _ := common.RunCmd(path, args...)
	return output
}
