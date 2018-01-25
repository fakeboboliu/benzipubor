# benzipubor(本子pubor)

本子pubor是一个为了\<del>满足作者欲♂望\</del>在kindle等电子书设备上看本子而生的项目，可以将文件夹中的本子打包为epub，其打包方式具有不先进、不标准的特色。

## 开始之前

首先请确认你整理本子的目录结构受支持，本子pubor目前支持两种目录结构（根据作者整理习惯设计）：

使用子目录表示多个本子：

```
|-- root
  |-- 本子1
    |-- 本子1，页面1.jpg
    |-- 本子1，页面2.jpg
  |-- 本子2
    |-- 本子2，页面1.jpg
    |-- 本子2，页面2.jpg
```

只有一个本子：

```
|-- root
  |-- 页面1.jpg
  |-- 页面2.jpg
```

默认参数下，根目录标题将作为电子书标题，子目录的标题将作为电子书目录的标题。

## 快速开始

废话少说先上演示：

![bz1](https://user-images.githubusercontent.com/7552030/35371865-a7cd50aa-01d1-11e8-831d-a813c8716225.gif)

只需从[Release页面](https://github.com/popu125/benzipubor/releases)下载编译好的二进制文件（请注意位数，尤其是Windows，错误的版本可能会导致图片压缩出现错误）即可快速开始，演示中是最基本的使用方法，即直接启动将当前目录下图片文件制作为epub，使用默认参数。

除此之外，如果你在执行时添加`-h`参数，将看到一个帮助信息。

## 还来点啥

### 转换为mobi

![bz2](https://user-images.githubusercontent.com/7552030/35371864-a77694ea-01d1-11e8-8a6f-1b6b5c436b02.gif)

上面演示中将epub转换为mobi时使用了`kindlegen`，除此以外，还可选用calibre、Epubor等工具转换为你想要的格式。

### 更多选项

```
λ benzipubor -h
Usage of benzipubor:
  -grey
        将图片处理为灰色（有助压缩到更小） (default true)
  -h    打印帮助信息
  -in string
        输入目录 (default ".")
  -log string
        日志输出 (default "stdout")
  -mode uint
        模式选择: 0: 'aio' 制作为一个电子书文件, 1: 'single' 每个子目录一个独立文件
  -quality int
        图片输出质量（1-100，越高越质量越好，体积越大） (default 50)
  -sizex int
        图片压缩尺寸（横向，500-1500） (default 780)
```

## LICENSE

GPL v3