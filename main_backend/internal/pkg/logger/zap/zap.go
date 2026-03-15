package zap

import "go.uber.org/zap"



type Logger struct{
	Zap *zap.SugaredLogger
}


func InitLogger()*Logger{
	logger := zap.NewExample().Sugar()
	return &Logger{Zap: logger}
}

func(l *Logger) Info(msg string,args ...any){
	l.Zap.Info(msg,args)
}

func(l *Logger ) Warn(msg string,args ...any){
	l.Zap.Warn(msg,args)
}

func(l *Logger ) Error(msg string,args ...any){
	l.Zap.Error(msg,args)
}

