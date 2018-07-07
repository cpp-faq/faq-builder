package util

import (
    "fmt"
)

func JoinAll(prefix string, paths []string) ([]string) {
    ret := []string{}

    for _, path := range paths {
        ret = append(ret, prefix + path)
    }

    return ret
}

func PrintAll(strs []string) {
    for _, str := range strs {
        fmt.Println(str)
    }
}
