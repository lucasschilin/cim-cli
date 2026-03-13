package git

import (
	"os/exec"
	"strings"
)

func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	diff := string(out)

	return diff, nil
}

func LimitDiff(diff string, maxLines int) string {
	lines := strings.Split(diff, "\n")

	if len(lines) <= maxLines {
		return diff
	}

	limitedLines := lines[:maxLines]

	return strings.Join(limitedLines, "\n")
}
