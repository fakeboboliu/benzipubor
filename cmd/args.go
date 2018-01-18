/*
 *     benzipubor
 *     Copyright (C) 2018 bobo liu
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"flag"
	. "github.com/popu125/benzipubor/epub-gen"
	"log"
	"os"
)

var (
	autoMode string
	inputDir string
	logFile  string

	gen *Gen
	l   *log.Logger
)

func init() {
	flag.StringVar(&autoMode, "mode", "off", "Auto mode: 'off' to turn off, 'aio' write all to one file, 'single' make a single file for each dictionary")
	flag.StringVar(&inputDir, "in", ".", "Input dir")
	flag.StringVar(&logFile, "log", "stdout", "Log file")

	noGrey := flag.Bool("nogrey", false, "Don't make pictures grey")
	help := flag.Bool("h", false, "Show this message")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	gen = NewGen()
	gen.SetNoGrey(*noGrey)
}
