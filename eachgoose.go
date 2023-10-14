package eachgoose

import (
	"log/slog"
)

func Execute() int {
	resources := ParseArgs()

	for _, res := range resources {
		slog.Debug("args", "service", res.Service, "region", res.Regions)
	}

	return 0
}
