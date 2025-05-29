package main

import "git.cryptic.systems/volker.raschek/dcmerge/cmd"

var version string

func main() {
	_ = cmd.Execute(version)
}
