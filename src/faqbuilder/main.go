package main

import (
    "fmt"
    "os"
    //"faqbuilder/process"
    "faqbuilder/util"
    "faqbuilder/model"
    "faqbuilder/process"
)

func main() {

    params, e := util.ReadParams()

    if(e != nil) {
        fmt.Println(e.Error())
        os.Exit(1)
    }

    fmt.Printf("Faq Builder v1.0.\n");

    // Read FAQ.
    faq, e := model.NewFAQ(params.RootFolder)
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(2)
    }
    fmt.Println("FAQ parsed.")

    util.PrintAll(faq.ToStrings())

    process.Process(faq)
    //build.Build(faq, params)
}
