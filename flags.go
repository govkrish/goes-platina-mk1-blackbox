// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"flag"
	"os"
)

const (
	DefaultGoes   = "./goes-platina-mk1"
	InstalledGoes = "/usr/bin/goes"
)

var (
	IsAlpha = flag.Bool("test.alpha", false, "zero based ports")
	DynDbg  = flag.String("test.dyndbg", "",
		"platform-driver load flag (use +pmf for verbose)")
	Goes = flag.String("test.goes", DefaultGoes,
		"GO Embedded System for Platina's Mk1 TOR Switch")
	NoVnet = flag.Bool("test.novnet", false,
		"manual vnet start (debugger)")
	PlatformDriver = flag.String("test.platform-driver", "platina-mk1",
		"Linux Kernel Platform Driver")
	SingleStep = flag.Bool("test.step", false,
		"single step (manual testing)")
	Repeat   = flag.Uint("test.repeat", 1, "test suite iterations")
	XethStat = flag.Bool("test.xeth-stat", false,
		"show /sys/kernel/platina-mk1/xeth stats")
)

func assertFlags() {
	flag.Parse()
	if _, err := os.Stat(*Goes); err != nil {
		if *Goes != DefaultGoes {
			panic(err)
		}
		if _, err = os.Stat(InstalledGoes); err != nil {
			panic("can't find goes")
		} else {
			*Goes = InstalledGoes
		}
	}
}
