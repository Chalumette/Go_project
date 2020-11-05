/*

import (
	"image/png"
	_ "image/png"
	"log"
)

func main() {
	img, err := png.Decode(gopherPNG())
	if err != nil {
		log.Fatal(err)
	}
}
*/
package main

import (
	"bufio"
	"fmt"
	"os" //to read the file
)

func main() {
	// Open the image (here test.png) and make it a channel called ImageFile
	ImageFile, err := os.Open("test.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// We close the channel when we finish the func
	defer ImageFile.Close()

	imInfo, _ := ImageFile.Stat()
	var size int64 = imInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(ImageFile)
	_, err = buffer.Read(bytes)

	fmt.Println(bytes)

	/* Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, imageType, err := image.Decode(ImageFile)
	if err != nil {
		// Handle error
	}
	fmt.Println(imageData)

	fmt.Println(imageType)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	ImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(ImageFile)
	if err != nil {
		// Handle error
	}
	fmt.Println(loadedImage)*/

	fmt.Println("aaaaaaaaaaaaaaaaaaaa")
}
