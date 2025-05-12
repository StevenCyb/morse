package display

import (
	"strings"

	"github.com/fatih/color"
)

func ColorizeMorseString(morse string) string {
	morse = strings.Replace(morse, ".", color.CyanString("."), -1)
	morse = strings.Replace(morse, "-", color.RedString("-"), -1)
	return morse
}
