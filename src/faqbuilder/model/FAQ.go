package model

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "strconv"
    "faqbuilder/util"
    "faqbuilder/engine"
)

type FAQOptions struct {

}

type FAQSerializer struct {
    Name string
    Version string
    Language string
    Sections []string

    Options FAQOptions`yaml:"options"`
}

type FAQ struct {
    Name string
    Version string
    Header string
    RootFolder string

    Options FAQOptions

    Sections []Section
    Questions map[string]Question
}

func NewFAQ(root_folder string, engine *engine.Engine) (*FAQ) {
    faq := &FAQ{}

    path := root_folder + "/faq.yaml"

    faq.RootFolder = root_folder
    faq.Questions = make(map[string]Question)

    // Read file
    file, e := ioutil.ReadFile(path)
    if e != nil && engine.Error("cannot find index file: " + e.Error() + ".") {
        return nil
    }

    var faqser FAQSerializer

    err := yaml.Unmarshal(file, &faqser)


    if err != nil &&  engine.Error("cannot parse index file: " + e.Error() + ".") {
        return nil
    }

    faq.Version = faqser.Version
    faq.Name = faqser.Name
    faq.Sections = GetSections(util.JoinAll(root_folder, faqser.Sections), faq, engine)
    faq.Options = faqser.Options

    if engine.Abord() {
        return nil
    }

    return faq
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
