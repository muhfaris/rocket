package liblog

import (
	"log/slog"
	"os"
)

func init() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)
}
