package util

import (
    "errors"
    "flag"
)

type Params struct {
    RootFolder string
    BuildFolder string
    BaseURL string

    // Options
    Verbose bool
    WError bool
}

func ReadParams() (Params, error) {
    params := Params{}

    // Read args.
    /*RootFolder := flag.String("root-folder", "", "the root folder of the FAQ files (should include at least the faq.json file).")*/
    flag.StringVar(&params.RootFolder, "root-folder", "", "the root folder of the FAQ files (should include at least the faq.json file).")
    flag.StringVar(&params.BuildFolder, "build-folder", "", "the output folder for the FAQ.")
    flag.StringVar(&params.BaseURL, "base-url", "", "the root URL of the github project (used for absolute linking).")
    flag.Parse()

    if(params.RootFolder == "") {
        return params, errors.New("error: root-folder unspecified.")
    }
    if(params.RootFolder == "") {
        return params, errors.New("error: build-folder unspecified.")
    }
    if(params.BaseURL == "") {
        return params, errors.New("error: base-url unspecified.")
    }

    return params, nil
}
