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
	"html/template"
	"math/rand"
	"time"
)

var tplstrs = map[string]string{
	"opf": `<?xml version = "1.0" encoding = "UTF-8"?>
<package xmlns = "http://www.idpf.org/2007/opf" version = "3.0" xml:lang = "ja" unique-identifier = "unique-id" prefix = "rendition: http://www.idpf.org/vocab/rendition/#         epub-bundle-tool: https://wing-kai.github.io/epub-manga-creator/         ebpaj: http://www.ebpaj.jp/         fixed-layout-jp: http://www.digital-comic.jp/">

<metadata xmlns:dc = "http://purl.org/dc/elements/1.1/">

<dc:title id = "title">{{.Title}}</dc:title>
<meta refines = "#title" property = "file-as"></meta>

<dc:creator id="creator01">本子Pubor</dc:creator>

<dc:subject></dc:subject>

<dc:publisher id = "publisher">本子Pubor</dc:publisher>
<meta refines = "#publisher" property = "file-as"></meta>

<dc:language>ja</dc:language>

<dc:identifier id = "unique-id">urn:uuid:{{.UUID}}</dc:identifier>

<meta property = "dcterms:modified">{{.CreateTime}}</meta>

<meta property = "rendition:layout">pre-paginated</meta>
<meta property = "rendition:spread">landscape</meta>

<meta property = "ebpaj:guide-version">1.1</meta>
<meta name = "SpineColor" content = "#FFFFFF"></meta>
<meta name = "cover" content = "cover"></meta>

</metadata>

<manifest>

<item id="ncx" href="toc.ncx" media-type="application/x-dtbncx+xml" />
<!-- style -->
<item media-type = "text/css" id = "fixed-layout-jp" href = "style/fixed-layout-jp.css"></item>

{{range .Objects}}
<item id="i_{{.}}" href="image/i_{{.}}.jpg" media-type="image/jpeg"></item>
<item id="p_{{.}}" href="text/p_{{.}}.xhtml" media-type="application/xhtml+xml" properties="svg" fallback="i_{{.}}"></item>
{{end}}

</manifest>

<spine toc="ncx">

{{range .Objects}}
<itemref idref="p_{{.}}"/>
{{end}}

</spine>

</package >`,

	"page": `<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE html>
	<html
	xmlns = "http://www.w3.org/1999/xhtml"
	xmlns:epub = "http://www.idpf.org/2007/ops"
	xml:lang = "ja"
	>
	<head>
	<meta charset = "UTF-8" />
	<title>Yens</title>
	<link rel = "stylesheet" type = "text/css" href = "../style/fixed-layout-jp.css"/>
	<meta name = "viewport" content = "width=1000, height=1500"/>
	</head>
	<body>
	<div class = "main">
	<svg xmlns = "http://www.w3.org/2000/svg" version = "1.1"
	xmlns:xlink = "http://www.w3.org/1999/xlink"
	width = "100%" height = "100%" viewBox = "0 0 1000 1500">
	<image width = "100%" height = "100%" preserveAspectRatio = "none" xlink:href = "../image/i_{{.}}.jpeg" />
	</svg>
	</div>
	</body>
	</html>`, // page.xhtml need: ID
	"toc": `<?xml version="1.0"?>
<ncx version="2005-1" xmlns="http://www.daisy.org/z3986/2005/ncx/">
  <head>
    <meta name="dtb:uid" content="unknown"/>
    <meta name="dtb:depth" content="1"/>
    <meta name="dtb:totalPageCount" content="0"/>
    <meta name="dtb:maxPageNumber" content="0"/>
  </head>
  <docTitle>
    <text>{{.Title}}</text>
  </docTitle>
  <navMap>
  {{range .TocNodes}}
    <navPoint id="navPoint{{.ID}}" playOrder="{{.ID}}">
      <navLabel>
        <text>{{.Name}}</text>
      </navLabel>
      <content src="text/p_{{.Pic}}.xhtml" />  
    </navPoint>
   {{end}}
   </navMap>
</ncx>
`,
}

var staticFiles = map[string]string{
	"META-INF/container.xml": `<?xml version="1.0" encoding="UTF-8"?>
<container version="1.0" xmlns="urn:oasis:names:tc:opendocument:xmlns:container">
  <rootfiles>
    <rootfile full-path="content.opf" media-type="application/oebps-package+xml"/>
  </rootfiles>
</container>`,
	"style/fixed-layout-jp.css": `@charset "UTF-8"; html,body{margin:0;padding:0;font-size:0;}svg{margin:0;padding:0;}`,
	"mimetype":                  "application/epub+zip",
}

var tpls = map[string]*template.Template{}

func init() {
	for k, v := range tplstrs {
		tpls[k] = template.Must(template.New("tpl").Parse(v))
	}
	rand.Seed(time.Now().UnixNano())
}

func randUUID() string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	uuid := []byte("")
	for i := 36; i > 0; i-- {
		switch i {
		case 27, 22, 17, 12:
			uuid = append(uuid, 45) // 45 is "-"
		default:
			uuid = append(uuid, chars[rand.Intn(36)])
		}
	}
	return string(uuid)
}
