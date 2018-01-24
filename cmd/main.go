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
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("======== Welcome to benzipubor ========\n")

	switch logFile {
	case "stdout":
		l = log.New(os.Stdout, "[MAIN]", log.LstdFlags)
	case "stderr":
		l = log.New(os.Stdin, "[MAIN]", log.LstdFlags)
	default:
		f, err := os.Open(logFile)
		if err != nil {
			fmt.Println("Cannot open file for write:", logFile)
		}
		l = log.New(f, "[MAIN]", log.LstdFlags)
		defer f.Close()
	}

	walkRootAndGen(inputDir)
	l.Println("Done.")
}
