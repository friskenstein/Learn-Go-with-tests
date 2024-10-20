package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // Flag for picking which exercise to run
    ex := flag.String("ex", "", "Specify which exercise to run (e.g., hello)")
    flag.Parse()

    if *ex == "" {
        fmt.Println("Please specify an exercise to run with the -ex flag")
        return
    }

    // Build the path to the example's main file
    cmd := exec.Command("go", "run", fmt.Sprintf("./%s/%s.go", *ex, *ex))

    // Set the output of the command to Stdout and Stderr
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Execute the command
    if err := cmd.Run(); err != nil {
        fmt.Printf("Failed to run %s: %v\n", *ex, err)
    }
}
