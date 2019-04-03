package utils

import (
	"fmt"
	"os/exec"
	"sync"
)

func ExecCmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("\n\nexecuting... ",cmd)
	// splitting head => g++ parts => rest of the command
	//parts := strings.Fields(cmd)
	//head := parts[0]
	//parts = parts[1:]

	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("here is erorr \n\n%s\n", err)
	}
	fmt.Printf("here is output \n\n%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
	//return out, err
}