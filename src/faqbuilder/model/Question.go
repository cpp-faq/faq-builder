package model

import (
    "gopkg.in/yaml.v2"
    //"encoding/json"
    "path/filepath"
    "faqbuilder/util"
    "faqbuilder/engine"
)

type Link struct {
    Description string`yaml:"descr"` //`json:"descr"`
    URL string`yaml:"url"` //`json:"url"`
}

type QuestionSerializer struct {
    DisplayName string`yaml:"display-name"` //`json:"display-name"`
    EndLinks []Link`yaml:"end-links"` //`json:"end-links"`
}

type Question struct {
    Name string // Duplicate with map key.
    DisplayName string
    RootFolder string
    EndLinks []Link

    // After processing
    LocalPath string
    LocalFiles []string // List of local files to be copied
}

func ParseQuestion(path string, engine *engine.Engine) (Question) {
    q := Question{}

    q.RootFolder = filepath.Dir(path)
    q.Name = filepath.Base(path)

    // Read JSON
    header, e := util.ExtractHeader(path + ".md")
    if( e != nil && engine.Error("cannot parse question: " + e.Error() + ".")) {
        return q
    }

    var quser QuestionSerializer

    err := yaml.Unmarshal([]byte(header), &quser)

    if(err != nil && engine.Error("cannot parse question: " + e.Error() + ".")) {
        return q
    }

    q.DisplayName = quser.DisplayName
    q.EndLinks = quser.EndLinks

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
