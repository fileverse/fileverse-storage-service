package logger

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

var Logger *slog.Logger

func getSlogHandlerOptions() *slog.HandlerOptions {
	level := viper.GetString("logger.level")
	var slogHandlerOpt *slog.HandlerOptions
	switch level {
	case "warn":
		slogHandlerOpt = &slog.HandlerOptions{
			Level: slog.LevelWarn,
		}
	case "debug":
		slogHandlerOpt = &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
	case "error":
		slogHandlerOpt = &slog.HandlerOptions{
			Level: slog.LevelError,
		}
	default:
		slogHandlerOpt = &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
	}

	return slogHandlerOpt
}

func Setup() {
	slogHandlerOpt := getSlogHandlerOptions()

	format := viper.GetString("logger.format")
	switch format {
	case "json":
		Logger = slog.New(slog.NewJSONHandler(os.Stdout, slogHandlerOpt))
	default:
		Logger = slog.New(slog.NewTextHandler(os.Stdout, slogHandlerOpt))
	}
}

func GetContextLogger(ctx context.Context) *slog.Logger {
	logger := ctx.Value("logger")
	if logger == nil {
		return Logger
	}
	return logger.(*slog.Logger)
}
