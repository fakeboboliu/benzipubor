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

package epub_gen

import (
	"time"
)

func NewGen() *Gen {
	return &Gen{
		// Book Info
		bi: bookInfo{UUID: randUUID(), CreateTime: time.Now().Format(time.RFC3339)},
		// Table of Contents
		tocNum: 1, imgList: make([]string, 0), tocNodes: make([]toc, 0),
		// Image settings
		x: 720, Grey: true, quality: 80,
	}
}
func inRange(in, min, max int) int {
	if in <= min {
		return min
	} else if in >= max {
		return max
	} else {
		return in
	}
}

type toc struct {
	ID   int
	Pic  int
	Name string
}
