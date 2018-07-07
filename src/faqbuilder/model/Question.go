package model

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "path/filepath"
    "faqbuilder/util"
)

type QuestionJsonObject struct {
    DisplayName string`json:"display-name"`
    EndLinks []string`json:"end-links"`
}

type Question struct {
    Name string // Duplicate with map key.
    DisplayName string
    RootFolder string
    EndLinks []string
}

func ParseQuestion(path string) (Question, error) {
    q := Question{}
    fmt.Println("test")
    q.RootFolder = filepath.Dir(path)
    q.Name = filepath.Base(path)

    // Read JSON
    file, e := ioutil.ReadFile(path + ".json")
    if( e != nil) {
        fmt.Printf("error : cannot find question [" + path + "]\n")
        return q, e // TODO => chain message.
    }

    var qjson QuestionJsonObject

    fmt.Println("file : " + string(file))

    err := json.Unmarshal(file, &qjson)

    if(err != nil) {
        return q, err
    }

    fmt.Printf("Results: %v\n", qjson)

    q.DisplayName = qjson.DisplayName
    q.EndLinks = qjson.EndLinks

    /*s.SubSections, err = GetSections(JoinAll(s.RootFolder + "/", secjson.SubSections), faq)
    s.Questions = secjson.Questions

    if(err != nil) {
        return nil, err
    }

    for _, qstr := range s.Questions {
        var q *Question

        q, err = NewQuestion(s.RootFolder + "/" + qstr)

        if(err != nil) {
            return nil, err
        }
        faq.AddQuestion(q)
    }*/

    return q, nil
}

func (q *Question) ToStrings() []string {
    ret := []string{}

    ret = append(ret, "Question : " + q.Name)
    ret = append(ret, "  - display-name : " + q.DisplayName)
    ret = append(ret, "  - root-folder : " + q.RootFolder)
    ret = append(ret, "  - end-links : ")
        ret = append(ret, util.JoinAll("      - ", q.EndLinks)...)
    return ret
}
