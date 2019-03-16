package engine

import (
    "fmt"
)

/**
 * Represents the engine of the build process.
 */
type Engine struct {
    Version string // Version of faq-builder.
    Params Parameters // User parameters.

    // Storing messages instead of count could be useful.
    Warnings []string
    Errors []string
    Fatals []string
}

func NewEngine() (*Engine) {
    var engine Engine

    engine.Version = "1.0"
    engine.Params = ReadParameters(&engine)

    return &engine
}

func (this *Engine) Debug(msg string) bool {
    if(this.Params.VerbosityLevel == Verbose) {
        fmt.Println(msg)
    }

    return this.Abord()
}

func (this *Engine) Info(msg string) bool {
    if(this.Params.VerbosityLevel > Quiet) {
        fmt.Println(msg)
    }

    return this.Abord()
}

func (this *Engine) Warning(msg string) bool {
    if this.Params.Werror {
        this.Error(msg)
    } else {
        this.Warnings = append(this.Warnings, msg)
        fmt.Println("warning: " + msg)
    }

    return this.Abord()
}

func (this *Engine) Error(msg string) bool {
    this.Errors = append(this.Errors, msg)
    fmt.Println("error: " + msg)

    return this.Abord()
}

func (this *Engine) Fatal(msg string) bool {
    this.Fatals = append(this.Fatals, msg)
    fmt.Println("fatal error: " + msg)

    return this.Abord()
}

func (this *Engine) Abord() bool {
    return len(this.Fatals) > 0 || (this.Params.Build != nil && this.Params.Build.ExitOnError && len(this.Errors) > 0)
}
