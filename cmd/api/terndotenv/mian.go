package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command("tern", "migrate", "--migrations", "./internal/storee/pgstore/migrations/tern.conf")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Comand execut failed", err)
		fmt.Println("Output", string(output))
	}

	fmt.Println("Comand Output sucess", string(output))

}
