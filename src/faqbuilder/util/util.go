package util

import (
    "fmt"
    "net/url"
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

func EncodeURL(str string) (string) {
    Url, err := url.Parse(str)
    if err != nil {
        fmt.Println("waning: the string '" + str + "' cannot be encoded to HTTP URL.")
        return str
    }
    return Url.String()
}
