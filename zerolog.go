package fxlogger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx/fxevent"
)

type logger struct {
	Logger zerolog.Logger
}

func (l *logger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Info().Str("callee", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStart hook executing")
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Warn().Err(e.Err).
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Msg("OnStart hook failed")
		} else {
			l.Logger.Info().Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStart hook executed")
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Info().Str("callee", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStop hook executing")
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Warn().Err(e.Err).
				Str("callee", e.FunctionName).
				Str("callee", e.CallerName).
				Msg("OnStop hook failed")
		} else {
			l.Logger.Info().Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStop hook executed")
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.Logger.Warn().Err(e.Err).Str("type", e.TypeName).Msg("supplied")
		} else {
			l.Logger.Info().Str("type", e.TypeName).Msg("supplied")
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Info().Str("type", rtype).
				Str("constructor", e.ConstructorName).
				Msg("provided")
		}
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("error encountered while applying options")
		}
	case *fxevent.Invoking:
		// Do nothing. Will log on Invoked.

	case *fxevent.Invoked:
		msg := "invoked"
		if e.Err != nil {
			msg = "invoke failed"
			l.Logger.Error().Err(e.Err).Str("stack", e.Trace).
				Str("function", e.FunctionName).Msg(msg)
		} else {
			l.Logger.Info().Str("function", e.FunctionName).Msg(msg)
		}
	case *fxevent.Stopping:
		l.Logger.Info().Str("signal", strings.ToUpper(e.Signal.String())).Msg("received signal")
	case *fxevent.Stopped:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("stop failed")
		}
	case *fxevent.RollingBack:
		l.Logger.Error().Err(e.StartErr).Msg("start failed, rolling back")
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("rollback failed")
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("start failed")
		} else {
			l.Logger.Info().Msg("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("custom logger initialization failed")
		} else {
			l.Logger.Info().Str("function", e.ConstructorName).Msg("initialized custom fxevent.Logger")
		}
	}
}

// Default logger for fxevent.
func Default() func() fxevent.Logger {
	l := log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    true,
	})
	return WithZerolog(l)
}

// WithZerolog customize zerolog instance for fxevent.
func WithZerolog(l zerolog.Logger) func() fxevent.Logger {
	return func() fxevent.Logger {
		return &logger{Logger: l}
	}
}
