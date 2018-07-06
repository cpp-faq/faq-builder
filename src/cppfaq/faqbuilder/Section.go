package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type SectionJsonObject struct {
    DisplayName string`json:"display-name"`
    Questions []string`json:"questions"`
    SubSections []string`json:"sections"`
}

type Section struct {
    Questions []*Question
    SubSections []Section
    RootFolder string
    Name string
}

func GetSection(paths []string) ([]Section, error) {
    sections := []Section{};

    //fmt.Printf("Searching sections.\n")
    for _, path := range paths {
        sec, err := NewSection(path)
        if(err != nil) {
            return sections, err
        }
        sections = append(sections, *sec)
        //fmt.Printf("path = " + path + "\n")
    }

    return sections, nil
}

func NewSection(path string) (*Section, error) {
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

    fmt.Println("file : " + string(file))

    err := json.Unmarshal(file, &secjson)

    if(err != nil) {
        return s, err
    }

    fmt.Printf("Results: %v\n", secjson)

    s.SubSections, err = GetSection(JoinAll(s.RootFolder + "/", secjson.SubSections))

    if(err != nil) {
        return nil, err
    }

    return s, nil
}


func (sec *Section) ToStrings() []string {
    ret := []string{}

    ret = append(ret, "Section : " + sec.Name)
    ret = append(ret, "root_folder : " + sec.RootFolder)
    ret = append(ret, "subsections : ")
    for _, subsec := range sec.SubSections {
        ret = append(ret, JoinAll("   - ", subsec.ToStrings())...)
    }
    ret = append(ret, "questions : (TODO)")
    /*for _, question := range sec.SubSections {
        append(ret, " - " + subsec.Name)
    }*/
    return ret
}
