package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type FAQJsonObject struct {
    Name string
    Version string
    Sections []string
}

type FAQ struct {
    Name string
    Version string
    Header string
    Sections []Section
}

func NewFAQ(root_folder string) (*FAQ, error) {
    faq := &FAQ{}

    // Read JSON
    file, e := ioutil.ReadFile(root_folder + "/faq.json")
    if( e != nil) {
        return nil, e // TODO => chain message.
    }

    var faqjson FAQJsonObject

    err := json.Unmarshal(file, &faqjson)

    if(err != nil) {
        return faq, err
    }

    fmt.Printf("Results: %v\n", faqjson)

    faq.Version = faqjson.Version
    faq.Name = faqjson.Name



    return faq, nil
}

func (faq *FAQ) Print() {
    fmt.Println("name : " + faq.Name)
    fmt.Println("version : " + faq.Version)
}
