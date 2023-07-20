package log

var (
	maxSize    = 10
	maxBackups = 5
	maxAge     = 7
)

func SetMaxSize(n int) {
	maxSize = n
}

func SetMaxBackups(n int) {
	maxBackups = n
}

func SetMaxAge(n int) {
	maxAge = n
}
