package engine

import (
    //"os"
    "flag"
)


type Command int

const (
    Build Command = iota
    Tool
    Interactive
    None
)

type Parameters struct {
    Command Command
    Tool *ToolParameters
    Build *BuildParameters

    // General
    RootFolder string
    VerbosityLevel Verbosity // [0; 2]
    Werror bool
}

func ReadParameters(engine *Engine) (Parameters) {
    params := Parameters{}

    // General parameters
    flag.StringVar(&params.RootFolder, "root-folder", "", "the root folder of the FAQ files (should include at least the file 'faq.yaml').")

    var verboseOpt bool
    var quietOpt bool

    flag.BoolVar(&verboseOpt, "verbose", false, "Set the verbosity level to 'Verbose'. In this mode, all information are displayed. This option is exclusive with 'quiet'.")
    flag.BoolVar(&verboseOpt, "quiet", false, "Set the verbosity level to 'Quiet'. In this mode, only important information are displayed. This option is exclusive with 'verbose'.")
    flag.BoolVar(&params.Werror, "Werror", false, "All warning will be treated as errors.")

    flag.Parse()
/*
    // Handle command
    validCmdStr := `valid commands :
 - build: build the FAQ.
 - tool: actions on the FAQ.
 - interactive: interactive mode.
 `

    if(len(os.Args) < 2) {
        engine.Fatal("fatal: no command specified.\n" + validCmdStr)
        return params
    }
    switch os.Args[1] {
    case "build":
        params.Command = Build
        //params.Build = ReadBuildParameters(engine)
    case "tool":
        params.Command = Tool
        engine.Warning("'tool' command not yet available.")
    case "interactive":
        params.Command = Interactive
        engine.Warning("'interactive' command not yet available.")
    default:
        params.Command = None
        engine.Fatal("unknown command '" + os.Args[1] + "'.")
    }
*/
    // General checks
    if(params.RootFolder == "") {
        engine.Fatal("root-folder unspecified.")
        return params
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

    return params
}
