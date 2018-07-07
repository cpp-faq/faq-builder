package model

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "errors"
    "strconv"
    "faqbuilder/util"
)

type SectionJsonObject struct {
    DisplayName string`json:"display-name"`
    Questions []string`json:"questions"`
    SubSections []string`json:"sections"`
}

type Section struct {
    Questions []string
    SubSections []Section
    RootFolder string
    Name string
}

func GetSections(paths []string, faq *FAQ) ([]Section, error) {
    sections := []Section{};

    for _, path := range paths {
        sec, err := NewSection(path, faq)
        if(err != nil) {
            return sections, err
        }
        sections = append(sections, *sec)
    }

    return sections, nil
}

func NewSection(path string, faq *FAQ) (*Section, error) {
    s := &Section{}

    s.RootFolder = path
    path = path + "/index.json"

    // Read JSON
    file, e := ioutil.ReadFile(path)
    if( e != nil) {
        fmt.Printf("error : cannot find section [" + path + "]\n")
        return nil, e // TODO => chain message.
    }

    var secjson SectionJsonObject

    err := json.Unmarshal(file, &secjson)

    if(err != nil) {
        return s, err
    }

    s.SubSections, err = GetSections(util.JoinAll(s.RootFolder + "/", secjson.SubSections), faq)
    s.Questions = secjson.Questions

    if(err != nil) {
        return nil, err
    }

    for _, qstr := range s.Questions {
        var q Question

        q, err = ParseQuestion(s.RootFolder + "/" + qstr)

        if(err != nil) {
            return nil, err
        }
        if b := faq.AddQuestion(q); !b {
            return nil, errors.New("error [Section '" + s.Name + "']: question '" + qstr + "' already defined.")
        }
    }

    return s, nil
}


func (sec *Section) ToStrings(faq *FAQ) []string {
    ret := []string{}

    ret = append(ret, "Section : " + sec.Name)
    ret = append(ret, "  - root_folder : " + sec.RootFolder)
    ret = append(ret, "  - subsections : (" + strconv.Itoa(len(sec.SubSections)) + ")")
    for _, subsec := range sec.SubSections {
        ret = append(ret, util.JoinAll("      ", subsec.ToStrings(faq))...)
    }
    ret = append(ret, "  - questions : (" + strconv.Itoa(len(sec.Questions)) + ")")
    for _, qstr := range sec.Questions {
        //ret = append(ret, "      " + qstr)
        q, _ := faq.Questions[qstr]
        ret = append(ret, util.JoinAll("      ", q.ToStrings())...)
    }
    return ret
}
