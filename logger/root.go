package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/spf13/viper"
)

func GetLogger(name, section string) *zap.Logger {
	return zap.New(initCoreWithConfig(name, section), zap.AddCallerSkip(1), zap.AddCaller())
}

func initCoreWithConfig(name, section string) zapcore.Core {
	logPath := "/home/logs/"
	env := strings.ToLower(os.Getenv("envType"))
	if env != "product" && env != "test" && env != "gray" {
		logPath = "runtime/"
	}

	logPath = fmt.Sprintf("%s%s/%s.log", logPath, name, name)

	if closed := viper.GetBool(fmt.Sprintf("log.%s.closed", section)); closed {
		logPath = "/dev/null"
	}

	maxSize := viper.GetInt(fmt.Sprintf("log.%s.max_size", section))
	if maxSize == 0 {
		maxSize = 5120
	}

	maxAge := viper.GetInt(fmt.Sprintf("log.%s.max_age", section))
	if maxSize == 0 {
		maxAge = 7
	}

	maxBackup := viper.GetInt(fmt.Sprintf("log.%s.max_backup", section))
	if maxBackup == 0 { // 如果未设置则默认保留5个
		maxBackup = 5
	}

	if maxBackup < 0 { // 如果设置为 -1 代表保留全部
		maxBackup = 0
	}

	opts := []zapcore.WriteSyncer{
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath, // ⽇志⽂件路径
			MaxSize:    maxSize, // 单位为MB,默认为100MB
			MaxAge:     maxAge,  // 文件最多保存多少天
			LocalTime:  true,    // 采用本地时间
			Compress:   false,   // 是否压缩日志
			MaxBackups: maxBackup,
		}),
	}

	// if l.stdout {
	// 	opts = append(opts, zapcore.AddSync(os.Stdout))
	// }

	syncWriter := zapcore.NewMultiWriteSyncer(opts...)

	// 自定义时间输出格式
	// customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// 	enc.AppendString(t.Format(logTmFmtWithMS))
	// 	// enc.AppendString(strconv.Itoa(int(t.UnixNano() / 1e6)))
	// }

	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("" + level.CapitalString() + "")
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		// enc.AppendString("[" + l.traceId + "]")
		enc.AppendString("" + caller.TrimmedPath() + "")
	}

	encoderConf := zapcore.EncoderConfig{
		// CallerKey:      "caller_line", // 打印文件名和行数
		// LevelKey:       "level",
		// MessageKey:     "msg",
		// TimeKey:        "@timestamp",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.RFC3339TimeEncoder, // 自定义时间格式
		EncodeLevel:    customLevelEncoder,         // 小写编码器
		EncodeCaller:   customCallerEncoder,        // 全路径编码器
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// // level大写染色编码器
	// if l.enableColor {
	// 	encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// }

	// // json 格式化处理
	// if l.jsonFormat {
	// 	return zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf),
	// 		syncWriter, zap.NewAtomicLevelAt(l.logMinLevel))
	// }

	return zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf),
		syncWriter, zap.NewAtomicLevelAt(zapcore.DebugLevel))
}
