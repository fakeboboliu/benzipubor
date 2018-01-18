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
	"log"
	"archive/zip"
	"strconv"
	"os"
	"bufio"
)

type Gen struct {
	imgList []string

	bi bookInfo
	l  *log.Logger

	X      int
	noGrey bool
}

func (g *Gen) SetImgList(imgList []string) {
	g.imgList = imgList
}

func (g *Gen) SetNoGrey(noGrey bool) {
	g.noGrey = noGrey
}

func (g *Gen) SetTitle(t string) {
	g.bi.Title = t
}

func (g *Gen) SetLogger(l log.Logger) {
	l.SetPrefix("[EPUB]")
	g.l = &l
}

func (g Gen) Do(dst string) {
	f, err := os.Open(dst)
	if err != nil {
		g.l.Fatalln("Cannot open file to write:", dst)
	}

	buf := bufio.NewWriter(f)
	w := zip.NewWriter(buf)

	for k, v := range staticFiles {
		sta, err := w.Create(k)
		if err != nil {
			g.l.Fatalln("Unknown error:", err)
		}
		sta.Write([]byte(v))
	}

	for i, fn := range g.imgList {
		// Pic
		pic := getZipWriter(w, "OEBPS/image/i_"+strconv.Itoa(i)+".jpg")
		g.doZip(fn, pic)

		// Pages
		page := getZipWriter(w, "OEBPS/text/p_"+strconv.Itoa(i)+".xhtml")
		tpls["page"].Execute(page, i)
	}

	opf := getZipWriter(w, "OEBPS/standard.opf")
	tpls["opf"].Execute(opf, g)

	buf.Flush()
}
