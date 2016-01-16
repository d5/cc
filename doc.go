/*
Colorize terminal output.

	package main

	import (
		"log"

		"github.com/d5/cc"
	)

	func main() {
		log.Printf("Enabled: %v", cc.Enabled())
		log.Printf(cc.Red("red"))
		log.Printf(cc.BgWhite(cc.Blue("blue")))
		log.Printf(cc.Underline(cc.Bold("bold-underline")))
	}

Features:
	- Clean and simple API
	- Windows support
	- Auto-detect coloring support
*/
package cc
