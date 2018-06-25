package main

import (
    "fmt"
    "io/ioutil"
    //"encoding/json"
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
    q := &Section{}

    // Read JSON
    file, e := ioutil.ReadFile(path)
    if( e != nil) {
        return nil, e // TODO => chain message.
    }

    fmt.Printf("%s\n", string(file))

    /** TODO **/

    return q, nil
}
