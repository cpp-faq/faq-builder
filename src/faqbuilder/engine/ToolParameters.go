package engine

import (
    "flag"
)

type ToolParameters struct {
    ExitOnError bool
}

func ReadToolParameters() (ToolParameters) {
    params := ToolParameters{}

    // Read args.

    flag.BoolVar(&params.ExitOnError, "print-structure", true, "")
    flag.Parse()

    return params
}
