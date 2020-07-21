package log

// Instance log interface
type Instance interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Panic(msg string)
	Fatal(msg string)
}
