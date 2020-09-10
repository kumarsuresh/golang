package main

import (
    "fmt"
)


func readConfig(path string) (*Config, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer file.Close()
    cfg := &Config{}
    return cfg, nil
}

func main() {
    cfg, err := readConfig("/path/to/file.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error %s\n", err)
        os.Exit(1)
    }
    fmt.Println(cfg)
}