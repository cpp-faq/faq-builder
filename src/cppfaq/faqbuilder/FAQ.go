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
    RootFolder string
}

func NewFAQ(root_folder string) (*FAQ, error) {
    faq := &FAQ{}
    path := root_folder + "/faq.json"
    faq.RootFolder = root_folder

    // Read JSON
    file, e := ioutil.ReadFile(path)
    if( e != nil) {
        return nil, e // TODO => chain message.
    }

    var faqjson FAQJsonObject

    err := json.Unmarshal(file, &faqjson)

    if(err != nil) {
        fmt.Printf("error : cannot parse file [" + path + "]")
        return faq, err
    }

    faq.Version = faqjson.Version
    faq.Name = faqjson.Name
    faq.Sections, err = GetSection(JoinAll(root_folder, faqjson.Sections))

    return faq, nil
}


func (faq *FAQ) Print() {
    fmt.Println("name : " + faq.Name)
    fmt.Println("version : " + faq.Version)
}
