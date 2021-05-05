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
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"os"

	"github.com/nfnt/resize"
)

func (g *Gen) doZip(path string, w io.Writer) error {
	f, err := os.Open(path)
	if err != nil {
		g.l.Println("Can not open image file:", path)
		g.l.Println("[DEBUG]", err)
		return err
	}

	var src image.Image
	src, _, err = image.Decode(f)
	if err != nil {
		g.l.Println("Image file can't load:", path)
		g.l.Println("[DEBUG]", err)
		return err
	}

	bound := src.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	dst := src
	if dx > g.x {
		dst = resize.Resize(uint(g.x), uint(dy*(g.x/dx)), src, resize.Bicubic)
	}

	if g.Grey {
		bound = dst.Bounds()
		grey := image.NewRGBA(bound)
		dx = bound.Dx()
		dy = bound.Dy()
		for i := 0; i < dx; i++ {
			for j := 0; j < dy; j++ {
				colorRgb := dst.At(i, j)
				_, green, _, a := colorRgb.RGBA()
				g_uint8 := uint8(green >> 8)
				a_uint8 := uint8(a >> 8)
				grey.SetRGBA(i, j, color.RGBA{g_uint8, g_uint8, g_uint8, a_uint8})
			}
		}
		dst = grey.SubImage(bound)
	}

	jpeg.Encode(w, dst, &jpeg.Options{Quality: g.quality})
	return nil
}
