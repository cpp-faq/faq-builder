package main

import (
    //"fmt"
    "io/ioutil"
    //"encoding/json"
)

type QuestionJsonObject struct {
    DisplayName string
    EndLinks []string
}

type Question struct {
    Name string
    DisplayName string
    Version string
    Header string
}

func NewQuestion(path string) (*Question, error) {
    q := &Question{}

    // Read JSON
    _, e := ioutil.ReadFile(path)
    if( e != nil) {
        return nil, e // TODO => chain message.
    }


    return q, nil
}
