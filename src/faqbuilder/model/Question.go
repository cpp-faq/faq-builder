package model

import (
    "gopkg.in/yaml.v2"
    //"encoding/json"
    "path/filepath"
    "faqbuilder/util"
    "faqbuilder/engine"
)

type QuestionOptions struct {
}

type LinkOptions struct {
    Prefix string`yaml:"prefix" json:"prefix"`
}

type Link struct {
    Description string`yaml:"descr"`
    URL string`yaml:"url"`

    Options LinkOptions`yaml:"options" json:"options"`
}

type QuestionSerializer struct {
    DisplayName string`yaml:"display-name"`
    EndLinks []Link`yaml:"end-links"`

    Options QuestionOptions`yaml:"options" json:"options"`
}

type Question struct {
    Name string // Duplicate with map key.
    DisplayName string
    RootFolder string
    EndLinks []Link

    Options QuestionOptions
}

func ParseQuestion(path string, engine *engine.Engine) (Question) {
    q := Question{}

    q.RootFolder = filepath.Dir(path)
    q.Name = filepath.Base(path)

    // Read JSON
    header, err := util.ExtractHeader(path + ".md")
    if( err != nil && engine.Error("cannot parse question '" + q.Name  + "': " + err.Error() + ".")) {
        return q
    }

    var quser QuestionSerializer

    err = yaml.Unmarshal([]byte(header), &quser)

    if(err != nil && engine.Error("cannot parse question '" + q.Name  + "': " + err.Error() + ".")) {
        return q
    }

    q.DisplayName = quser.DisplayName
    q.EndLinks = quser.EndLinks
    q.Options = quser.Options

    return q
}

func (q *Question) ToStrings() []string {
    ret := []string{}

    ret = append(ret, "Question : " + q.Name)
    ret = append(ret, "  - display-name : " + q.DisplayName)
    ret = append(ret, "  - root-folder : " + q.RootFolder)
    ret = append(ret, "  - end-links : ")
    for _, link := range q.EndLinks {
        ret = append(ret, "      - " + link.URL + "(" + link.Description + ")")
    }
    return ret
}
