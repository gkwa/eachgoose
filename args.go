package eachgoose

import (
	"log/slog"
	"os"
	"strings"
)

type Resource struct {
	Service string
	Regions []string
}

func parseArgs() []Resource {
	args := os.Args[1:] // Skip the program name

	if len(args) == 0 {
		slog.Debug("Please provide a list of resource pairs in the format 'service:region,...'")
		return []Resource{}
	}

	resources := make([]Resource, 0)

	for _, arg := range args {
		slog.Debug("args", "arg", arg)
		parts := strings.Split(arg, ":")
		if len(parts) != 2 {
			slog.Debug("Invalid resource format: %s\n", arg)
			continue
		}

		service := parts[0]
		regionsCommaSep := parts[1]

		regions := strings.Split(regionsCommaSep, ",")

		resources = append(resources, Resource{Service: service, Regions: regions})
	}

	return resources
}
