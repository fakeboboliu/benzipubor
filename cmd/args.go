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
	"fmt"
	. "github.com/popu125/benzipubor/epub-gen"
	"log"
	"os"
)

const (
	MODE_AIO    = iota
	MODE_SINGLE
)

var (
	autoMode uint
	inputDir string
	logFile  string

	gen *Gen
	l   *log.Logger
)

func init() {
	flag.UintVar(&autoMode, "mode", 0, "模式选择: 0: 'aio' 制作为一个电子书文件, 1: 'single' 每个子目录一个独立文件")
	flag.StringVar(&inputDir, "in", ".", "输入目录")
	flag.StringVar(&logFile, "log", "stdout", "日志输出")

	noGrey := flag.Bool("nogrey", false, "不要将图片处理为灰色")
	sizex := flag.Int("sizex", 780, "图片压缩尺寸（横向）")
	help := flag.Bool("h", false, "打印帮助信息")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if autoMode >= 2 {
		fmt.Println("Unknown mode:", autoMode)
		os.Exit(1)
	}

	gen = NewGen()
	gen.NoGrey = *noGrey
	gen.X = *sizex
}
