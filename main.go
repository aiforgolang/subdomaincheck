package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/weppos/publicsuffix-go/publicsuffix"
)

func checkDomain(input string) {
    domainName := strings.TrimSpace(strings.ToLower(input))
    if domainName == "" {
        return
    }

    mainDomain, err := publicsuffix.Domain(domainName)
    if err != nil {
        fmt.Printf("%s\tunknown\n", input)
        return
    }

    if domainName == mainDomain {
        fmt.Printf("%s\tmain\n", input)
    } else {
        fmt.Printf("%s\tsubdomain\n", input)
    }
}

func main() {
    if len(os.Args) > 1 {
        for _, arg := range os.Args[1:] {
            checkDomain(arg)
        }
        return
    }

    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            checkDomain(scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "error:", err)
        }
        return
    }

    fmt.Println("eg 1: ./main google.com")
    fmt.Println("eg 2: cat domains.txt | ./main")
}
