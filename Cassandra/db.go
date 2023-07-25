package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func Run(command string) execute.ExecResult {
	words := strings.Fields(command)

	cmd := execute.ExecTask{
		Command:     words[0],
		Args:        words[1:],
		StreamStdio: false,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	return res
}

func RunUntilValid(command string) {
	color.Green(command)
	success := false

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "Trying command: "
	s.Start()
	for !success {
		res := Run(command)
		if res.ExitCode == 0 {
			success = true
		} else {
			time.Sleep(5 * time.Second)
		}
	}
	s.Stop()
}

func RunPrint(command string) {
	res := Run(command)
	color.Green(command)

	if res.ExitCode != 0 {
		color.Red(fmt.Sprintf("%s\n", res.Stderr))
	}
}

func main() {
	success := color.New(color.Bold, color.FgWhite, color.BgBlue).PrintFunc()
	fail := color.New(color.Bold, color.FgWhite, color.BgRed).PrintFunc()

	if len(os.Args) <= 1 {
		fail("No Argument Provided (start or stop)")
		return
	}

	mode := strings.ToLower(os.Args[1])

	if mode == "start" {
		RunPrint("docker network create cassandra")
		RunPrint("docker run --rm -d -p 127.0.0.1:9042:9042 --name cassandra --hostname cassandra --network cassandra cassandra")
		RunPrint("docker cp data.cql cassandra:/data.cql")
		RunUntilValid("docker exec cassandra cqlsh -f data.cql")
		success("Server successfully started...")
	} else if mode == "stop" {
		RunPrint("docker kill cassandra")
		RunPrint("docker network rm cassandra")
		success("Server successfully stopped...")
	}
}
