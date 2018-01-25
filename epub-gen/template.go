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
	"math/rand"
	"text/template"
	"time"
)

var tplstrs = map[string]string{
	"opf": `<?xml version = "1.0" encoding = "UTF-8"?>
<package xmlns = "http://www.idpf.org/2007/opf" version="3.0" xml:lang="ja" unique-identifier="unique-id" prefix="rendition: http://www.idpf.org/vocab/rendition/#">
<metadata xmlns:dc = "http://purl.org/dc/elements/1.1/">
 <dc:title id = "title">{{.Title}}</dc:title>
 <dc:creator>本子Pubor</dc:creator>
 <dc:publisher>本子Pubor</dc:publisher>
 <dc:language>ja</dc:language>
 <dc:identifier id = "unique-id">urn:uuid:{{.UUID}}</dc:identifier>
 <dc:date>{{.CreateTime}}</dc:date>
 <meta property = "rendition:layout">pre-paginated</meta>
 <meta property = "rendition:spread">landscape</meta>
 <meta property = "ebpaj:guide-version">1.1</meta>
</metadata>

<manifest>
 <item id="ncx" href="toc.ncx" media-type="application/x-dtbncx+xml" />
 <item id="toc" href="toc.xhtml" media-type="application/xhtml+xml" />
 <item media-type = "text/css" id = "bzcss" href = "style/bz.css"></item>
{{range .Objects}}
 <item id="i_{{.}}" href="image/i_{{.}}.jpg" media-type="image/jpeg"></item>
 <item id="p_{{.}}" href="text/p_{{.}}.xhtml" media-type="application/xhtml+xml"></item>
{{end}}
</manifest>

<spine toc="ncx">
{{range .Objects}}
 <itemref idref="p_{{.}}"/>
{{end}}
</spine>
<guide>
 <reference href="toc.xhtml" type="toc" title="Table of Contents" />
</guide>
</package>`,

	"page": `<!DOCTYPE html>
<html xmlns = "http://www.w3.org/1999/xhtml" xmlns:epub = "http://www.idpf.org/2007/ops" xml:lang = "ja">
<head>
<meta charset = "UTF-8" />
<link rel = "stylesheet" type = "text/css" href = "../style/bz.css"/>
<meta name = "viewport" content = "width=1000, height=1500"/>
<title>{{.Title}}</title>
</head>
<body>
<div class="main">
<img width="100%" height="100%" src="../image/i_{{.ID}}.jpg" />
</div>
</body>
</html>`,

	"toc": `<?xml version="1.0" encoding="UTF-8"?>
<ncx version="2005-1" xmlns="http://www.daisy.org/z3986/2005/ncx/">
  <head>
    <meta name="dtb:uid" content="{{.UUID}}"/>
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
</ncx>`,

	"nav": `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml"
      xmlns:epub="http://www.idpf.org/2007/ops" lang="en" xml:lang="en">
  <head>
    <meta charset="utf-8" />
    <title>Table of Contents</title>
  </head>
  <body>
    <nav epub:type="toc">
      <h1>Table of Contents</h1>
      <ol>
      {{range .TocNodes}}<li><a href="text/p_{{.Pic}}.xhtml">{{.Name}}</a></li>{{end}}
      </ol>
      </nav>
   </body>
</html>`,
}

var staticFiles = map[string]string{
	"META-INF/container.xml": `<?xml version="1.0" encoding="UTF-8"?>
<container version="1.0" xmlns="urn:oasis:names:tc:opendocument:xmlns:container">
  <rootfiles>
    <rootfile full-path="content.opf" media-type="application/oebps-package+xml"/>
  </rootfiles>
</container>`,
	"style/bz.css": `@charset "UTF-8"; html,body{margin:0;padding:0;}svg{margin:0;padding:0;}.main{width:1000px;height:1500px;}`,
	"mimetype":     "application/epub+zip",
	"cover.html":   "Welcome",
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
