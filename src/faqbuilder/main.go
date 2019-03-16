package main

import (
    "os"
    "strconv"
    "faqbuilder/util"
    "faqbuilder/model"
    "faqbuilder/core"
    "faqbuilder/engine"
)

func main() {

    engine := engine.NewEngine()

    if engine.Abord() {
        os.Exit(1)
    }

    engine.Info("Faq Builder v" + engine.Version + ".\n");

    // Read FAQ.
    faq := model.NewFAQ(engine.Params.RootFolder, engine)

    if engine.Abord() {
        os.Exit(2)
    }

    engine.Info("FAQ parsed.")

    util.PrintAll(faq.ToStrings())

    if !core.Build(faq, engine) {

        engine.Error(strconv.Itoa(len(engine.Errors)) + " error(s) occured, abording.")
        os.Exit(3)
    } else {
        engine.Info("the FAQ has been successfully builded.")
    }
}
