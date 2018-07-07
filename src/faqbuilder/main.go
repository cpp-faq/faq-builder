package main

import (
    "fmt"
    "os"
    "flag"
    //"faqbuilder/process"
    "faqbuilder/util"
    "faqbuilder/model"
)

func main() {
    // Read args.
    folder := flag.String("root-folder", "", "the root folder of the FAQ files (should include at least the faq.json file).")
    flag.Parse()

    if(*folder == "") {
        fmt.Println("error: root-folder unspecified.")
        os.Exit(1)
    }

    fmt.Printf("Faq Builder v1.0.\n");

    // Read FAQ.
    faq, e := model.NewFAQ(*folder)
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(1)
    }
    fmt.Println("FAQ parsed.")

    util.PrintAll(faq.ToStrings())

    // Process
    // faq.Process()



}
