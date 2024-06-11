package output

import (
	"errors"
	"flag"
	"strings"
)

type Options struct {
	ColorFlag       string
	OutputFlag      string
	ColorizeLetters string
	InputText       string
	BannerFile      string
}

// ColorFlag function parses command-line arguments to extract color options, text to be colored, and banner file name.
func ParseOptions() (Options, error) {
	// Define flag for color option
	var options Options
	flag.StringVar(&options.ColorFlag, "color", "", "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
	flag.StringVar(&options.OutputFlag, "output", "", "Usage:...")
	flag.Parse()

	// Determine the number of arguments and parse accordingly
	switch len(flag.Args()) {
	case 1: // One argument: input text
		options.InputText = flag.Arg(0)
	case 2:
		// Two arguments: colorize letters and input text or bannerfile
		if flag.Arg(1) == "thinkertoy" || flag.Arg(1) == "standard" || flag.Arg(1) == "shadow" {
			options.BannerFile = flag.Arg(1)
			options.InputText = flag.Arg(0)
		} else {
			if options.ColorFlag != "" {
				options.ColorizeLetters = flag.Arg(0)
				options.InputText = flag.Arg(1)
			} else {
				return Options{}, errors.New("usage: go run . [OPTION] [STRING]\nex: go run . --color=<color> <letters to be colored> \"something\"")
			}
		}
		// colorizeLetters = flag.Arg(0)
	case 3:
		// Three arguments: colirize letters, input text and bannerfile
		if options.ColorFlag != "" {
			options.InputText = flag.Arg(1)
			options.ColorizeLetters = flag.Arg(0)
			options.BannerFile = flag.Arg(2)
		} else {
			return Options{}, errors.New("usage: go run . [OPTION] [STRING]\nex: go run . --color=<color> <letters to be colored> \"something\"")
		}
	default:
		// Invalid number of arguments
		return Options{}, errors.New("usage: go run . [OPTION] [STRING]\nex: go run . --color=<color> <letters to be colored> \"something\"")
	}
	options.ColorFlag = strings.ToLower(options.ColorFlag)
	options.OutputFlag = strings.ToLower(options.OutputFlag)
	options.BannerFile = strings.ToLower(options.BannerFile)

	// Convert color flag and banner file name to lowercase for consistency
	return options, nil
}
