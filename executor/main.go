package executor

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Execute(str string) {
	executable := "bash"
	cmd := exec.Command(executable, "-c", str)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println()

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

// func confirmExecution() bool {
// 	reply := InputPrompt("\nDo you want to execute this command?[y/n] ")
// 	if strings.Title(reply) == "Y" {
// 		return true
// 	}
// 	return false
// }
