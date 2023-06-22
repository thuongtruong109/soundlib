package constants

const (
	DEBUG int8 = iota
	INFO
	WARNING
	ERROR
	FATAL

	LABEL
	DESC
	QUERY
	INPUT
)

const (
	Reset string = "\033[0m"
	Red   string = "\033[31m"
	Green string = "\033[32m"
	Yellow string = "\033[33m"
	Purple string = "\033[34m"
	Pink string = "\033[35m"
	Cyan string = "\033[36m"
	Gray string = "\033[37m"
	White string = "\033[97m"
	Orange string = "\033[38;5;208m"
)