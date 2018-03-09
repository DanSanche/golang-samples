package main

import (
    "fmt"
    "os"
    "testing"
)

func main() {
    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
    fmt.Println(projectID)
}

func TestTesting(t *testing.T){
    t.Errorf("failed")
}
