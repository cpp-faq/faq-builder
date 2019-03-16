package util

import (
    "fmt"
    "errors"
    "io"
    "io/ioutil"
    "strings"
    "os"
    "path"
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

func ExtractHeader(path string) (string, error) {
    file, e := ioutil.ReadFile(path)
    if e != nil {
        return "", e
    }

    strs := strings.Split(string(file), "---")
    if len(strs) < 2 {
        return "", errors.New("file invalid, cannot split header and content.")
    }

    return strs[0] + strs[1], nil
}

func ExtractContent(path string) (string, error) {
    file, e := ioutil.ReadFile(path)
    if e != nil {
        return "", e
    }

    strs := strings.Split(string(file), "---")
    if len(strs) < 2 {
        return "", errors.New("file invalid, cannot split header and content.")
    }

    ret := ""
    for i := 2; i < len(strs); i++ {
        ret = ret + strs[i]
    }

    return ret, nil
}

func ExistDir(path string) (bool, error) {
    fi, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }

    return fi.Mode().IsDir(), nil
}

func CopyFile(src string, dst string) error {
    sourceFileStat, err := os.Stat(src)
    if err != nil {
          return err
    }

    if !sourceFileStat.Mode().IsRegular() {
          return fmt.Errorf("%s is not a regular file and cannot be copied.", src)
    }

    source, err := os.Open(src)
    if err != nil {
          return err
    }
    defer source.Close()

    destination, err := os.Create(dst)
    if err != nil {
          return err
    }
    defer destination.Close()
    _, ret := io.Copy(destination, source)
    return ret
}

// => https://blog.depado.eu/post/copy-files-and-directories-in-go
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil { return err }
	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil { return err }
	if fds, err = ioutil.ReadDir(src); err != nil { return err }

	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
