package main

import (
    "fmt"
    "os"
    "faqbuilder/util"
    "faqbuilder/model"
    "faqbuilder/core"
)

func main() {

    params, e := util.ReadParams()
    state := core.NewState(&params)

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

    if !core.Build(faq, state) {
        fmt.Println("error: error(s) occured, abording.")
    }
}
