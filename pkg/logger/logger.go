package logger

import (
	"context"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"tuyoops/go/terraform-plugin/config"
)

var lg *zap.Logger

// InitLogger initializes Logger
func InitLogger(cfg *config.LogConfig) (err error) {
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) // Replace the global logger instance in the zap package
	return
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Log the start of the request
	zap.L().Info(fmt.Sprintf("gRPC method: %s", info.FullMethod), zap.Any("request", req))

	// Call the handler to process the request
	resp, err := handler(ctx, req)

	// Log the end of the request
	if err != nil {
		st, _ := status.FromError(err)
		zap.L().Error(fmt.Sprintf("gRPC method: %s, error: %s", info.FullMethod, st.Message()), zap.Error(err))
		return nil, err
	}
	zap.L().Info(fmt.Sprintf("gRPC method: %s, response: %v", info.FullMethod, resp))

	return resp, nil
}

func StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Log the start of the request
	zap.L().Info(fmt.Sprintf("gRPC method: %s", info.FullMethod))

	// Call the handler to process the request
	err := handler(srv, ss)

	// Log the end of the request
	if err != nil {
		st, _ := status.FromError(err)
		zap.L().Error(fmt.Sprintf("gRPC method: %s, error: %s", info.FullMethod, st.Message()), zap.Error(err))
		return err
	}
	zap.L().Info(fmt.Sprintf("gRPC method: %s, stream completed", info.FullMethod))

	return nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
