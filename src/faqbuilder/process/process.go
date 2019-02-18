package process

import (
    "faqbuilder/model"
    "faqbuilder/util"
    "fmt"
    //"strings"
    //"io/ioutil"
    //"encoding/json"
    //"path/filepath"
)

func Process(faq *model.FAQ, params *util.Params) {
    fmt.Println("Processing FAQ...")
    fmt.Println(" - root_folder:  " + params.RootFolder)
    fmt.Println(" - build_folder: " + params.BuildFolder)
}

func PathToRootFolder(path string, faq *model.FAQ) string {
    // Check if relative.
    return path
}

func TransformLink(link string, faq *model.FAQ) string {
    /*if(strings.HasPrefix("q://")) {

    }
    if(strings.HasPrefix("faq://")) {

    }*/
    return link
}
