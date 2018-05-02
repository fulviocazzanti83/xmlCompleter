package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(`USE xmlCompleter --file myfile.xml`)
		os.Exit(0)
	}

	file := args[2]
	fillFile(file, file+"_BASE64")

}

func fillFile(originFile string, destinationFile string) {

	content, err := ioutil.ReadFile(originFile)
	if err != nil {
		log.Fatal(err)
	}

	contentString := string(content)
	fmt.Println(contentString)

	re := regexp.MustCompile(`FILENAME=\"(.)+\"\sNUMBER`)
	imageFileName := re.FindString(contentString)
	imageFileName = strings.Replace(imageFileName, `FILENAME="`, "", -1)
	imageFileName = strings.Replace(imageFileName, `" NUMBER`, "", -1)

	fmt.Println(imageFileName)
	base64image := ConvertFile(imageFileName)

	fmt.Println(base64image)

	xmlBase64 := fmt.Sprintf(`<data encoding="base64"><![CDATA[%s]]></data>`, base64image)
	newXml := strings.Replace(contentString, `<data encoding="base64"></data>`, xmlBase64, -1)
	fmt.Println(newXml)

	f, err := os.Create("./sampleData/exportFederalMugol_NEW.xml")
	f.WriteString(newXml)
	f.Sync()
	f.Close()

}

func ConvertFile(imageName string) string {
	imgFile, err := os.Open(imageName) // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	return imgBase64Str

}
