package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"time"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procSetCursorPos = user32.NewProc("SetCursorPos")
)

func setCursorPos(x, y int32) {
	procSetCursorPos.Call(uintptr(x), uintptr(y))
}

func main() {
	// Channel to signal when to stop the loop
	stop := make(chan bool)

	// Goroutine to move the mouse cursor in a loop
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				setCursorPos(100, 100)
				time.Sleep(2 * time.Second)
				setCursorPos(200, 200)
				time.Sleep(2 * time.Second)
				setCursorPos(300, 300)
				time.Sleep(2 * time.Second)
			}
		}
	}()

	// Wait for user input to exit
	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Signal the goroutine to stop
	close(stop)
}
