package utils

import (
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func Success(msg string) {
	fmt.Println(Blue + "[✔️]\t" + Green + msg + Reset)
}

func Information(msg string) {
	fmt.Println(Blue + "[*]\t" + Green + msg + Reset)
}

func Error(a ...any) (n int, err error) {
	return fmt.Println(Blue + "[!]\t" + Red + fmt.Sprint(a...) + Reset)
}
