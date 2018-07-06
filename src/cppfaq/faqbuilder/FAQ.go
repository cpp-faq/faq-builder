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
    Question []Question
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
        return nil, err
    }

    faq.Version = faqjson.Version
    faq.Name = faqjson.Name
    faq.Sections, err = GetSection(JoinAll(root_folder, faqjson.Sections))

    if(err != nil) {
        return nil, err
    }

    return faq, nil
}


func (faq *FAQ) ToStrings() []string {
    ret := []string{}

    ret = append(ret, "FAQ : " + faq.Name)
    ret = append(ret, "version : " + faq.Version)
    ret = append(ret, "root_folder : " + faq.RootFolder)

    // Sections
    for _, section := range faq.Sections {
        ret = append(ret, JoinAll("   - ", section.ToStrings())...)
    }

    return ret
}
