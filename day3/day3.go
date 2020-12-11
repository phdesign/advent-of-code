package main

import (
    "strings"
    "fmt"
    "io/ioutil"
    "flag"
    "os"
)

func CountTrees(grid string, down, right int) (count int) {
    rows := strings.Split(strings.Trim(grid, "\n"), "\n")
    width := len(rows[0])
    for x, y := 0, 0; y < len(rows); x, y = x+right, y+down {
        square := rows[y][x % width]
        if square == '#' {
            count++
        }
    }
    return
}

func main() {
    flag.Parse()
    filename := flag.Arg(0)
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    result := CountTrees(string(content), 1, 1)
    result *= CountTrees(string(content), 1, 3)
    result *= CountTrees(string(content), 1, 5)
    result *= CountTrees(string(content), 1, 7)
    result *= CountTrees(string(content), 2, 1)
    fmt.Println(result)
}
