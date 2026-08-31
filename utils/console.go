package utils

import (
	"fmt"
	"time"
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

func Debug(a ...any) {
	fmt.Printf(
		Purple+"[DEBUG] "+White+"%s\t%v"+Reset+"\n",
		time.Now().Format("2006-01-02 15:04:05"),
		fmt.Sprint(a...),
	)
}

func PublicErrors(a ...any) { // This function will appear when something goes wrong (permissions, code?, something, e.g)
	fmt.Printf(
		Red+"[DEBUG] "+White+"%s\t%v"+Reset+"\n",
		time.Now().Format("2006-01-02 15:04:05"),
		fmt.Sprint(a...),
	)
}
