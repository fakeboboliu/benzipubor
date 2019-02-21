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
	"bytes"
)

type zipOp struct {
	w    *zip.Writer
	c    chan task
	done chan bool
}

type zipW struct {
	fn  string
	z   *zipOp
	buf *bytes.Buffer
}

func (zw *zipW) Write(p []byte) (n int, err error) {
	return zw.buf.Write(p)
}

func (zw *zipW) Flush() {
	zw.z.WriteFile(zw.fn, zw.buf.Bytes())
}

func newZipOp(w *zip.Writer) *zipOp {
	op := &zipOp{w: w, c: make(chan task, 64), done: make(chan bool)}
	go op.zipWriter()
	return op
}

func (z *zipOp) WriteFile(name string, data []byte) {
	d := make([]byte, len(data))
	copy(d, data)
	z.c <- task{Name: name, Data: d}
}

func (z *zipOp) zipWriter() {
	for t := range z.c {
		a, err := z.w.Create(t.Name)
		if err != nil {
			panic(err)
		}
		a.Write(t.Data)
	}
	z.done <- true
}

func (z *zipOp) Writer(name string) *zipW {
	return &zipW{fn: name, z: z, buf: new(bytes.Buffer)}
}

func (z *zipOp) Done() {
	close(z.c)
}

func (z *zipOp) Wait() {
	<-z.done
}

type task struct {
	Name string
	Data []byte
}
