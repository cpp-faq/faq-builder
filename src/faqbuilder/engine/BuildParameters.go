package engine

import (
    "flag"
)


type Verbosity int

const (
    Quiet Verbosity = iota
    Normal
    Verbose
)

type BuildParameters struct {
    // Required
    BuildFolder string
    BaseURL string

    // Options
    CleanOnError bool
    ExitOnError bool
}

func ReadBuildParameters(engine *Engine) (*BuildParameters) {
    params := BuildParameters{}

    // Read args.
    flag.StringVar(&params.BuildFolder, "build-folder", "", "the output folder for the FAQ.")
    flag.StringVar(&params.BaseURL, "base-url", "", "the root URL of the github project (used for absolute linking).")

    flag.BoolVar(&params.CleanOnError, "clean-error", false, "Clean the build folder if an error occures.")
    flag.BoolVar(&params.ExitOnError, "exit-on-error", true, "Stop the process if an error occures.")
    
    flag.Parse()

    if(params.BuildFolder == "") {
        engine.Fatal("build-folder unspecified.")
        return &params
    }
    if(params.BaseURL == "") {
        engine.Fatal("base-url unspecified.")
        return &params
    }

    return &params
}
