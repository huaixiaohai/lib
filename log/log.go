package log

import (
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

const (
	OutputFile = "file"
	OutputStd  = "stdout"
)

//const FileSuffix = ".log"

type Cfg struct {
	Level    uint32
	Output   string
	FilePath string
}

var instance = logrus.New()
var cfg *Cfg

func Init(c *Cfg) {
	cfg = c
	instance.SetLevel(logrus.Level(cfg.Level))
	if cfg.Output == OutputFile {
		instance.SetOutput(&lumberjack.Logger{
			Filename:   path.Join(cfg.FilePath, "log"),
			MaxSize:    50, // 50M
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   false, // disabled by default
		})
	} else if cfg.Output == OutputStd {
		instance.SetOutput(os.Stdout)
	} else {
		panic("log output error")
	}
}

var (
	Printf  = instance.Printf
	Println = instance.Println
	Info    = instance.Info
	Error   = instance.Error
)
