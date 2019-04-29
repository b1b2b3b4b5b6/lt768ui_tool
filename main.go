package main

import (
	"fmt"
	"goc/logface"
	"goc/toolcom/errtool"
	"io/ioutil"
	"os"
	"strings"
)

var log = logface.New(logface.DebugLevel)

var fileC *os.File
var fileH *os.File
var fileB *os.File

func codeStart() {
	errtool.Ignore(os.Remove("lt768ui_resources.c"))
	errtool.Ignore(os.Remove("lt768ui_resources.h"))
	errtool.Ignore(os.Remove("lt768ui.bin"))

	var err error
	fileC, err = os.Create("lt768ui_resources.c")
	errtool.Errpanic(err)
	fileC.WriteString("#include \"lt768ui_resources.h\"\r\n")

	fileH, err = os.Create("lt768ui_resources.h")
	errtool.Errpanic(err)
	fileH.WriteString("#ifndef LT768UI_RESOURCES_H\r\n")
	fileH.WriteString("#define LT768UI_RESOURCES_H\r\n")
	fileH.WriteString("#include \"lt768ui_interface.h\"\r\n")

	fileB, err = os.Create("lt768ui.bin")
	errtool.Errpanic(err)
}

func codeStop() {
	fileH.WriteString("#endif")

	fileC.Close()
	fileH.Close()
	fileB.Close()

	info, err := os.Stat("lt768ui.bin")
	errtool.Errpanic(err)
	log.Info("bin size[%v]", info.Size)
}

func codeCformat(name string, offset int, w int, h int) string {
	str := ""
	str += fmt.Sprintf("dma_info_t %s = {\r\n", name)
	str += fmt.Sprintf("\t.addr = %d,\r\n", offset)
	str += fmt.Sprintf("\t.src_w = %d,\r\n", w)
	str += fmt.Sprintf("\t.x = %d,\r\n", 0)
	str += fmt.Sprintf("\t.y = %d,\r\n", 0)
	str += fmt.Sprintf("\t.w = %d,\r\n", w)
	str += fmt.Sprintf("\t.h = %d,\r\n", h)
	str += fmt.Sprintf("};\r\n")
	return str
}

func codeHformat(name string) string {
	str := ""
	str += fmt.Sprintf("extern dma_info_t %s;\r\n", name)
	return str
}

var offset = 0

func codeAdd(path string) {
	log.Debug("handle file[%s]", path)
	bin, w, h := convert2bin(path)
	name := strings.Replace(path, "/", "_", -1)
	name = strings.TrimSuffix(name, ".jpg")
	name = strings.TrimSuffix(name, ".png")
	name = strings.TrimSuffix(name, ".jpeg")

	log.Debug("generate name[%s]", name)

	fileB.Write(bin)
	fileC.WriteString(codeCformat(name, offset, w, h))
	fileH.WriteString(codeHformat(name))

	offset += len(bin)
	log.Debug("code add resources[%s]", name)
}

func judgeBg(dir string) string {
	bgName := ""
	files, err := ioutil.ReadDir(dir)
	errtool.Errpanic(err)
	for _, file := range files {
		if !file.IsDir() {
			nameList := strings.Split(file.Name(), "/")
			switch nameList[len(nameList)-1] {
			case "bg.jpg":
				bgName = dir + "/" + file.Name()

			case "bg.png":
				bgName = dir + "/" + file.Name()

			case "bg.jpeg":
				bgName = dir + "/" + file.Name()
			}
		}
	}

	return bgName
}

func traversalWidget(dir string) {
	log.Debug("traversalWidget dir[%s]", dir)

	files, err := ioutil.ReadDir(dir)
	errtool.Errpanic(err)
	for _, file := range files {
		if !file.IsDir() {
			codeAdd(dir + "/" + file.Name())
		}
	}
}

func traversalInterface(dir string) {
	log.Debug("traversalInterface dir[%s]", dir)
	bgName := judgeBg(dir)
	if bgName == "" {
		log.Panic("find no bg, dir[%s]", dir)
	}
	files, err := ioutil.ReadDir(dir)
	errtool.Errpanic(err)
	codeAdd(bgName)
	for _, file := range files {
		if file.IsDir() {
			traversalWidget(dir + "/" + file.Name())
		}
	}
}

func main() {
	codeStart()
	files, err := ioutil.ReadDir("./")
	errtool.Errpanic(err)

	for _, file := range files {
		if file.IsDir() && !strings.Contains(file.Name(), ".") {
			traversalInterface(file.Name())
		}
	}
	codeStop()
}
