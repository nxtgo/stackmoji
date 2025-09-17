package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"

	xdraw "golang.org/x/image/draw"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	inputPath := flag.String("input", "", "path to the input image")
	outputDir := flag.String("output", "emojis", "directory to save emoji slices")
	rows := flag.Int("rows", 3, "number of rows in the grid")
	cols := flag.Int("cols", 3, "number of columns in the grid")
	prefix := flag.String("prefix", "e", "prefix for emoji filenames")
	flag.Parse()

	if *inputPath == "" {
		fmt.Println("provide an input image path using -input, cuck")
		return
	}

	file, err := os.Open(*inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	fmt.Printf("original image size: %dx%d\n", width, height)

	cellWidth := width / *cols
	cellHeight := height / *rows

	maxCell := cellWidth
	maxCell = max(cellHeight, maxCell)

	totalWidth := maxCell * *cols
	totalHeight := maxCell * *rows
	squareImg := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))
	draw.Draw(squareImg, squareImg.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

	offsetX := (totalWidth - width) / 2
	offsetY := (totalHeight - height) / 2
	draw.Draw(squareImg, image.Rect(offsetX, offsetY, offsetX+width, offsetY+height), img, image.Point{}, draw.Over)

	err = os.MkdirAll(*outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	count := 0
	for y := 0; y < totalHeight; y += maxCell {
		for x := 0; x < totalWidth; x += maxCell {
			slice := squareImg.SubImage(image.Rect(x, y, x+maxCell, y+maxCell)).(*image.RGBA)

			final := image.NewRGBA(image.Rect(0, 0, 128, 128))
			xdraw.CatmullRom.Scale(final, final.Bounds(), slice, slice.Bounds(), xdraw.Over, nil)

			outPath := filepath.Join(*outputDir, fmt.Sprintf("%s%d.png", *prefix, count))
			outFile, err := os.Create(outPath)
			if err != nil {
				panic(err)
			}
			png.Encode(outFile, final)
			outFile.Close()

			count++
		}
	}

	fmt.Printf("generated %d emoji slices in %s\ndon't forget to star, cuck. <3\n--> https://github.com/nxtgo/stackmoji <--\n", count, *outputDir)
}
