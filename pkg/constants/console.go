package constants

const (
	DEBUG int8 = 0
	INFO int8 = 1
	WARNING int8 = 2
	ERROR int8 = 3
	FATAL int8 = 4

	LABEL int8 = 5
	DESC int8 = 6
	QUERY int8 = 7
	INPUT int8 = 8
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