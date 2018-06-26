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
    Questions []Question
    SubSections []Section
}

func NewSection(path string) (*Section, error) {
    s := &Section{}

    // Read JSON
    file, e := ioutil.ReadFile(path)
    if( e != nil) {
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
