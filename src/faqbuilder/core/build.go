package core

import (
    "faqbuilder/model"
    "faqbuilder/util"
    "io/ioutil"
    "fmt"
    "os"
    "path/filepath"
    //"strings"
    //"io/ioutil"
    //"encoding/json"
    //"path/filepath"
)

/**
 * Returns the string corresponding to the relative path compared to the specified rootPath.
 * This function assumes that rootPath is a substring of completePath.
 * @param completePath the path of the item (Section / Question), including the rootPath. This
 *      path should have been computed during the JSON parsin phase.
 * @param rootPath the root path of the FAQ folder.
 */
func GetRelativePath(completePath string, rootPath string) (string) {
    return completePath[len(rootPath):]
}

func ProcessQuestion(question *model.Question, state *State) (string, bool) {
    ret := "## " + question.DisplayName + "\n\n"

    content, err := ioutil.ReadFile(filepath.Join(question.RootFolder, question.Name + ".md"))
    if err != nil {
        fmt.Print(err)
        return "", false
    }

    ret += string(content)

    // Handle end-links
    if len(question.EndLinks) > 0 {
        ret += "\n#### Liens et compléments\n"
        for _, link := range question.EndLinks {
            // TODO() if q:// / faq:// process and replace, else display warning if no descr
            ret += " - [" + link.URL + "](" + link.Description + ").\n"
        }
    }

    return ret, true
}



func BuildSection(faq *model.FAQ, section *model.Section, state *State) (bool) {
    // Build header
    content := "# " + section.Name + "\n\n"

    mdcontent, err := ioutil.ReadFile(filepath.Join(section.RootFolder, "index.md"))
    if err == nil {
        content += string(mdcontent) + "\n"
    }

    content += "### Sommaire\n\n"

    questionsContent := ""

    // Build questions
    for _, name := range section.Questions {
        if question, found := faq.Questions[name]; found {
            // Summary
            content += " - [" + question.DisplayName + "](#" + util.EncodeURL(question.DisplayName) + ").\n"

            // Question content
            res, err := ProcessQuestion(&question, state)
            if !err {
                return false
            }
            questionsContent += "\n" + res

        } else {
            fmt.Println("error: question '" + name + "' does not exist.")
        }
    }

    content += questionsContent

    fmt.Println("Section : \n" + content)

    // Save
    os.MkdirAll(filepath.Join(state.Params.BuildFolder, GetRelativePath(section.RootFolder, faq.RootFolder)), os.ModePerm)

    return true
}

func Build(faq *model.FAQ, state *State) (bool) {
    fmt.Println("Building FAQ...")
    fmt.Println(" - root_folder:  " + state.Params.RootFolder)
    fmt.Println(" - build_folder: " + state.Params.BuildFolder)

    for _, section := range faq.Sections {
        if !BuildSection(faq, &section, state) {
            return false
        }
    }

    // Replace
    // Replace all q: with real URL
    // Replace all faq: with real URL
    // Notice all local link (to be copied later)

    return true
}

// TODO()
/*
func PathToRootFolder(path string, faq *model.FAQ) string {
    // Check if relative.
    return path
}

func TransformLink(link string, faq *model.FAQ) string {
    if(strings.HasPrefix("q://")) {

    }
    if(strings.HasPrefix("faq://")) {

    }
    return link
}*/
