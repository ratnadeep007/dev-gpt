package executor

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Execute(str string) {
	// goodToExecute := confirmExecution()

	// if !goodToExecute {
	// 	return
	// }

	// splitted := strings.Split(str, " ")
	// splitted = append([]string{"-c"}, splitted...)
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

func confirmExecution() bool {
	reply := InputPrompt("\nDo you want to execute this command?[y/n] ")
	if strings.Title(reply) == "Y" {
		return true
	}
	return false
}

func InputPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
