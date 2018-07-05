package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type SectionJsonObject struct {
    Name string
    Questions []string
    Sections []string
}

type Section struct {
    Questions []*Question
    SubSections []Section
    RootFolder string
}

func GetSection(paths []string) ([]Section, error) {
    sections := []Section{};

    fmt.Printf("Searching sections.\n")
    for _, path := range paths {
        sec, err := NewSection(path)
        if(err != nil) {
            return sections, err
        }
        sections = append(sections, *sec)
        fmt.Printf("path = " + path + "\n")
    }

    return sections, nil
}

func NewSection(path string) (*Section, error) {
    s := &Section{}

    path = path + "/index.json"

    // Read JSON
    file, e := ioutil.ReadFile(path)
    if( e != nil) {
        fmt.Printf("error : cannot find section [" + path + "]\n")
        return nil, e // TODO => chain message.
    }

    fmt.Printf("%s\n", string(file))

    var secjson SectionJsonObject

    err := json.Unmarshal(file, &secjson)

    if(err != nil) {
        return s, err
    }

    fmt.Printf("Results: %v\n", secjson)

    return s, nil
}
