package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	// Внешняя зависимость.
	"golang.org/x/image/draw"
)

func scale(img image.Image, w int, h int) image.Image {
	dstImg := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.NearestNeighbor.Scale(dstImg, dstImg.Bounds(), img, img.Bounds(), draw.Over, nil)
	return dstImg
}

func decodeImageFile(imgName string) (image.Image, error) {
	imgFile, err := os.Open(imgName)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(imgFile)

	return img, err
}

func processPixel(c color.Color) rune {
	gc := color.GrayModel.Convert(c)
	r, _, _, _ := gc.RGBA()
	r = r >> 8

	 symbols := []rune("@80GCLft1i;:,. ")
	 if r>=0 && r<=17 {
		 return symbols[0]
	 }else if r>17 && r<=34{
		 return symbols[1]
	 }else if r>34 && r<=51{
		 return symbols[2]
	 }else if r>51 && r<=68{
		 return symbols[3]
	 }else if r>68 && r<=85{
		 return symbols[4]
	 }else if r>85 && r<=102{
		 return symbols[5]
	 }else if r>102 && r<=119{
		 return symbols[6]
	 }else if r>119 && r<=136{
		 return symbols[7]
	 }else if r>136 && r<=153{
		 return symbols[8]
	 }else if r>153 && r<=170{
		 return symbols[9]
	 }else if r>170 && r<=187{
		 return symbols[10]
	 }else if r>187 && r<=204{
		 return symbols[11]
	 }else if r>204 && r<=221{
		 return symbols[12]
	 }else if r>221 && r<=238{
		 return symbols[13]
	 }else if r>238 && r<=255{
		 return symbols[14]
	 }
	return '0'
}

func convertToAscii(img image.Image) [][]rune {
	textImg := make([][]rune, img.Bounds().Dy())
	for i := range textImg {
		textImg[i] = make([]rune, img.Bounds().Dx())
	}

	for i := range textImg {
		for j := range textImg[i] {
			textImg[i][j] = processPixel(img.At(j, i))
		}
	}
	return textImg
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: asciimg <image.jpg>")
		os.Exit(0)
	}
	imgName := flag.Arg(0)

	img, err := decodeImageFile(imgName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	textImg := convertToAscii(img)
	for i := range textImg {
		for j := range textImg[i] {
			fmt.Printf("%c", textImg[i][j])
		}
		fmt.Println()
	}
}
