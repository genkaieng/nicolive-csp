package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/genkaieng/nicolive-csp/pkg/nicolive"
)

func main() {
	var lvid string
	if len(os.Args) > 1 {
		lvid = os.Args[1]
	}
	if lvid == "" {
		fmt.Println("Usage: go run main.go <lvid>")
		return
	}
	if !strings.HasPrefix(lvid, "lv") {
		fmt.Println("Please specify the <lvid> with 'lv' prefix (e.g., lv123456).")
		return
	}
	fmt.Println("Nicolive:", lvid)

	nicolive.Connect(lvid)
}
