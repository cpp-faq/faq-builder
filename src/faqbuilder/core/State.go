package core

import (
    "faqbuilder/util"
)

/**
 * Represent the state of the build process.
 */
type State struct {
    Version string
    Params *util.Params
    Warnings []string
    Errors []string
}

func NewState(params *util.Params) *State {
    var state State

    state.Version = "1.0"
    state.Params = params

    return &state
}

func Warning(this *State, msg string) {
    this.Warnings = append(this.Warnings, msg)
}

func Error(this *State, msg string) {
    this.Errors = append(this.Errors, msg)
}

func Stop(this *State) bool {
    return len(this.Errors) > 0
}
