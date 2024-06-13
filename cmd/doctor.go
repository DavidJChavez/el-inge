package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(doctorCmd)
}

var programas = map[string]bool{
	"go":   false,
	"nvm":  false,
	"java": false,
}

var doctorCmd = &cobra.Command{
	Use: "doctor",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	out, _ := exec.Command("go", "version").Output()
	if out != nil {
		programas["go"] = true
	}
	out, _ = exec.Command("node -v").Output()
	if out != nil {
		programas["node"] = true
	}
	if !isAllTrue(programas) {
		fmt.Println("Faltan cositas xd")
	}
	for k, v := range programas {
		fmt.Printf("%v %s : %v\n", smiley, k, v)
	}
}

func isAllTrue(programas map[string]bool) bool {
	for _, v := range programas {
		if !v {
			return false
		}
	}
	return true
}

var smiley = "\U0001F600"
