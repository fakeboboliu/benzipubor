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
	"archive/zip"
	"log"
	"os"
	"strconv"
	"sync"
	"bytes"
)

type Gen struct {
	imgList  []string
	tocNodes []toc
	tocNum   int

	bi bookInfo
	l  *log.Logger

	x       int
	Grey    bool
	quality int
}

func (g *Gen) SetX(x int) {
	g.x = inRange(x, 500, 1500)
}

func (g *Gen) SetQuality(quality int) {
	g.quality = inRange(quality, 1, 100)
}

func (g *Gen) AddTocNode(pic int, name string) {
	g.tocNodes = append(g.tocNodes, toc{Pic: pic, Name: name, ID: g.tocNum})
	g.tocNum += 1
}

func (g *Gen) AppendImgList(slice []string) int {
	count := len(g.imgList)
	g.imgList = append(g.imgList, slice...)
	return count
}

func (g *Gen) SetTitle(t string) {
	g.bi.Title = t
}

func (g *Gen) SetLogger(l log.Logger) {
	l.SetPrefix("[EPUB]")
	g.l = &l
}

func (g Gen) Do(dst string) {
	f, err := os.Create(dst)
	if err != nil {
		g.l.Fatalln("Cannot open file to write:", dst, err)
	}
	defer f.Close()

	w := zip.NewWriter(f)
	defer w.Close()

	for k, v := range staticFiles {
		sta, err := w.Create(k)
		if err != nil {
			g.l.Fatalln("Unknown error:", err)
		}
		sta.Write([]byte(v))
	}

	var wg sync.WaitGroup
	for i, fn := range g.imgList {
		wg.Add(1)
		id := i + 1
		g.bi.Objects = append(g.bi.Objects, id)
		go func(id int, fn string) {
			// Pic
			b := new(bytes.Buffer)
			g.doZip(fn, b)
			pic := getZipWriter(w, "image/i_"+strconv.Itoa(id)+".jpg")
			b.WriteTo(pic)

			// Page
			page := getZipWriter(w, "text/p_"+strconv.Itoa(id)+".xhtml")
			tpls["page"].Execute(page, pageInfo{ID: id, Title: g.bi.Title})

			wg.Done()

			g.l.Println("Processed:", fn)
		}(id, fn)
	}
	wg.Wait()

	g.bi.TocNodes = g.tocNodes

	opf := getZipWriter(w, "content.opf")
	tpls["opf"].Execute(opf, g.bi)

	toc := getZipWriter(w, "toc.ncx")
	tpls["toc"].Execute(toc, g.bi)

	nav := getZipWriter(w, "toc.xhtml")
	tpls["nav"].Execute(nav, g.bi)
}
