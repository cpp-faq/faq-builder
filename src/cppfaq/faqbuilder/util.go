package main

func JoinAll(prefix string, paths []string) ([]string) {
    ret := []string{}

    for _, path := range paths {
        ret = append(ret, prefix + path)
    }

    return ret
}
