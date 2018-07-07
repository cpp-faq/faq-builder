package model

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "strconv"
    "faqbuilder/util"
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
    Questions map[string]Question
    RootFolder string
}

func NewFAQ(root_folder string) (*FAQ, error) {
    faq := &FAQ{}

    path := root_folder + "/faq.json"

    faq.RootFolder = root_folder
    faq.Questions = make(map[string]Question)

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
    faq.Sections, err = GetSections(util.JoinAll(root_folder, faqjson.Sections), faq)

    if(err != nil) {
        return nil, err
    }

    return faq, nil
}

func (faq *FAQ) AddQuestion(q Question) bool {
    if _, exist := faq.Questions[q.Name]; exist {
        return false
    }
    faq.Questions[q.Name] = q
    return true
}

func (faq *FAQ) ToStrings() []string {
    ret := []string{}

    ret = append(ret, "FAQ : " + faq.Name)
    ret = append(ret, "  - version : " + faq.Version)
    ret = append(ret, "  - root_folder : " + faq.RootFolder)

    // Sections
    ret = append(ret, "  - sections : (" + strconv.Itoa(len(faq.Sections)) + ")")
    for _, section := range faq.Sections {
        ret = append(ret, util.JoinAll("      ", section.ToStrings(faq))...)
    }

    return ret
}
