package logging 

type LogLevel int

const (
  Trace LogLevel = iota
  Debug
  Information
  Warning
  Fatal
  None
)

type Logger interface {
  
  Trace(string)
  Tracef(string, ...interface{})

  Debug(string)
  Debugf(string, ...interface{})

  Info(string)
  Infof(string, ...interface{})

  Warning(string)
  Warningf(string, ...interface{})

  Panic(string)
  Panicf(string, ...interface{})

}
