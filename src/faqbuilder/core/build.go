package core

import (
    "faqbuilder/model"
    "faqbuilder/util"
    "faqbuilder/engine"
    "os"
    "path/filepath"
    "net/url"
    "io/ioutil"
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

func EncodeURL(str string, engine *engine.Engine) (string) {
    Url, err := url.Parse(str)
    if err != nil {
        engine.Error("the string '" + str + "' cannot be encoded to HTTP URL.")
        return str
    }
    return Url.String()
}

func ProcessEndLink(link model.Link) (string) {
    ret := ""

    if link.Options.Prefix != "" {
        ret += "**[" + link.Options.Prefix + "]** "
    }
    return ret + "[" + link.Description + "](" + link.URL + ")"
}

func ProcessQuestion(question *model.Question, engine *engine.Engine) (string) {
    engine.Debug("    generating question : " + question.Name + "'.")

    ret := "## " + question.DisplayName + "\n"

    path := filepath.Join(question.RootFolder, question.Name + ".md")
    content, err := util.ExtractContent(path)
    if err != nil {
        engine.Warning("unable to find '" + path + "', the question '" + question.DisplayName + "' is empty.")
        return ""
    }

    ret += string(content)

    // Handle end-links
    if len(question.EndLinks) > 0 {
        ret += "\n#### Liens et compléments\n"
        for _, link := range question.EndLinks {
            // TODO() if q:// / faq:// process and replace, else display warning if no descr
            ret += " - " + ProcessEndLink(link) + ".\n"
        }
    }

    return ret
}

func BuildSection(faq *model.FAQ, section *model.Section, engine *engine.Engine) {
    engine.Debug("  generating section : " + section.Name + "'.")

    // Build header
    content := "# " + section.Name + "\n"

    mdcontent, err := util.ExtractContent(filepath.Join(section.RootFolder, "index.md"))
    if err != nil && engine.Error("cannot read section content: " + err.Error() + ".") {
        return
    }

    content += mdcontent

    content += "\n### Sommaire\n\n"

    questionsContent := ""

    // Build sub-sections.

    for _, subsection := range section.SubSections {
        // Summary
        content += " - **[Section] [" + subsection.Name + "](" + EncodeURL(subsection.Name + "/README.md", engine) + ")**.\n"

        BuildSection(faq, &subsection, engine)
        if engine.Abord() {
            return
        }
    }

    // Build questions
    for _, name := range section.Questions {
        if question, found := faq.Questions[name]; found {
            // Summary
            content += " - [" + question.DisplayName + "](#" + EncodeURL(question.DisplayName, engine) + ").\n"

            // Question content
            res := ProcessQuestion(&question, engine)
            if engine.Abord() {
                return
            }
            questionsContent += "\n" + res

        } else {
            engine.Error("question '" + name + "' does not exist.")
            continue;
        }
    }

    content += questionsContent

    // Save
    dir := filepath.Join(engine.Params.BuildFolder, GetRelativePath(section.RootFolder, faq.RootFolder))
    os.MkdirAll(dir, os.ModePerm)

    err = ioutil.WriteFile(filepath.Join(dir + "/README.md"), []byte(content), 0644)
    if err != nil && engine.Error("cannot save section '" + section.Name + "': " + err.Error() + ".") {
        return
    }

    // Copy resources folder, if exist

    resdir := filepath.Join(section.RootFolder, "res/")

    if b, _ := util.ExistDir(resdir); b {
        err := util.CopyDir(resdir, filepath.Join(dir, "/res/"))
        if err != nil && engine.Error("cannot copy ressources directory for section '" + section.Name + "': " + err.Error() + ".") {
            return
        }
    }
}

func Build(faq *model.FAQ, engine *engine.Engine) (bool) {
    engine.Info("building FAQ...")

    for _, section := range faq.Sections {
        BuildSection(faq, &section, engine)
        if engine.Abord() {
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
