package util

import (
    "errors"
    "flag"
)

type Params struct {
    RootFolder string
    BuildFolder string
}

func ReadParams() (Params, error) {
    params := Params{}

    // Read args.
    /*RootFolder := flag.String("root-folder", "", "the root folder of the FAQ files (should include at least the faq.json file).")*/
    flag.StringVar(&params.RootFolder, "root-folder", "", "the root folder of the FAQ files (should include at least the faq.json file).")
    flag.StringVar(&params.BuildFolder, "out-folder", "", "the output folder for the FAQ.")
    flag.Parse()

    if(params.RootFolder == "") {
        return params, errors.New("error: root-folder unspecified.")
    }

    return params, nil
}
