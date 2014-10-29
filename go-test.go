package main

import (
	"os/exec"
	"fmt"
	//"bufio"
	"github.com/lrajlich/go-test/remote"
	//"remote"
)

func run() {
	cmd := exec.Command("/Users/lrajlich/Projects/anviltop/anviltop-agent/cron_job.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic (err)
	}

	// Start command
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	// Read stdout
	/*scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Print("Out: ",string(line),"\n")
		channel <- line
	}*/

	go remote.SendLoop("http://localhost:8081/")
	remote.ReadLoop(stdout)

	cmd.Wait()
	fmt.Printf("DONE")
}

func main() {
	run()
}
