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

type Parameters struct {
    // Required
    RootFolder string
    BuildFolder string
    BaseURL string

    // Options
    VerbosityLevel Verbosity // [0; 2]
    Werror bool
    CleanOnError bool
    ExitOnError bool
}

func ReadParameters(engine *Engine) (*Parameters) {
    params := Parameters{}

    // Read args.
    /*RootFolder := flag.String("root-folder", "", "the root folder of the FAQ files (should include at least the faq.json file).")*/
    flag.StringVar(&params.RootFolder, "root-folder", "", "the root folder of the FAQ files (should include at least the faq.json file).")
    flag.StringVar(&params.BuildFolder, "build-folder", "", "the output folder for the FAQ.")
    flag.StringVar(&params.BaseURL, "base-url", "", "the root URL of the github project (used for absolute linking).")

    var verboseOpt bool
    var quietOpt bool

    flag.BoolVar(&verboseOpt, "verbose", false, "Set the verbosity level to 'Verbose'. In this mode, all information are displayed. This option is exclusive with 'quiet'.")
    flag.BoolVar(&verboseOpt, "quiet", false, "Set the verbosity level to 'Quiet'. In this mode, only important information are displayed. This option is exclusive with 'verbose'.")
    flag.BoolVar(&params.Werror, "Werror", false, "All warning about the FAQ building process will be treated as errors.")
    flag.BoolVar(&params.CleanOnError, "clean-error", false, "Clean the build folder if an error occures.")
    flag.BoolVar(&params.ExitOnError, "exit-on-error", true, "Stop the process if an error occures.")
    flag.Parse()

    if(params.RootFolder == "") {
        engine.Fatal("root-folder unspecified.")
        return &params
    }
    if(params.RootFolder == "") {
        engine.Fatal("build-folder unspecified.")
        return &params
    }
    if(params.BaseURL == "") {
        engine.Fatal("base-url unspecified.")
        return &params
    }

    if verboseOpt {
        if quietOpt {
            engine.Warning("both 'quiet' and 'verbose' option specified. 'verbose' will be used.")
        }
        params.VerbosityLevel = Verbose
    } else if quietOpt {
        params.VerbosityLevel = Quiet
    } else {
        params.VerbosityLevel = Normal
    }

    return &params
}
