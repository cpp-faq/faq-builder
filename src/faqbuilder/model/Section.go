package model

import (
    "gopkg.in/yaml.v2"
    "strconv"
    "faqbuilder/util"
    "faqbuilder/engine"
)

type SectionOptions struct {
}

type SectionSerializer struct {
    DisplayName string`yaml:"display-name"`
    Questions []string`yaml:"questions"`
    SubSections []string`yaml:"sections"`

    Options SectionOptions`yaml:"options"`
}

type Section struct {
    Questions []string
    SubSections []Section
    RootFolder string
    Name string

    Options SectionOptions
}

func GetSections(paths []string, faq *FAQ, engine *engine.Engine) ([]Section) {
    sections := []Section{};

    for _, path := range paths {
        sec := NewSection(path, faq, engine)
        if engine.Abord() {
            return sections
        }
        sections = append(sections, *sec)
    }

    return sections
}

func NewSection(path string, faq *FAQ, engine *engine.Engine) (*Section) {
    s := &Section{}

    s.RootFolder = path

    // Read file
    header, e := util.ExtractHeader(path + "/index.md")
    if e != nil && engine.Error("cannot parse section: " + e.Error() + ".\n") {
        return nil
    }

    var secser SectionSerializer

    err := yaml.Unmarshal([]byte(header), &secser)

    if err != nil && engine.Error("cannot parse section : " + err.Error() + ".") {
        return s
    }

    s.SubSections = GetSections(util.JoinAll(s.RootFolder + "/", secser.SubSections), faq, engine)
    s.Questions = secser.Questions
    s.Name = secser.DisplayName
    s.Options = secser.Options

    if engine.Abord() {
        return s
    }

    for _, qstr := range s.Questions {
        q := ParseQuestion(s.RootFolder + "/" + qstr, engine)

        if engine.Abord() {
            return s
        }
        if b := faq.AddQuestion(q); !b && engine.Error("duplicate question : " + q.Name + " in section '" + s.Name + "'.") {
            return s
        }
    }

    return s
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
