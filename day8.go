package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
The image you received is 25 pixels wide and 6 pixels tall.

To make sure the image wasn't corrupted during transmission,
the Elves would like you to find the layer that contains the fewest 0 digits.
On that layer, what is the number of 1 digits multiplied by the number of 2 digits?
*/

func parseImage(data string, width int, height int) [][]int {
	layers := make([][]int, len(data)/(width*height))

	for i, value := range data {
		currLayers := i / (width * height)
		val, _ := strconv.ParseInt(string(value), 10, 8)
		if layers[currLayers] == nil {
			layers[currLayers] = []int{}
		}
		layers[currLayers] = append(layers[currLayers], int(val))
	}

	return layers
}

func findIt(layerList [][]int) int {
	youknowIt := 0
	leastZeros := 1000000000000000000
	for _, layer := range layerList {
		count0 := 0
		count1 := 0
		count2 := 0
		for _, pixel := range layer {
			switch pixel {
			case 0:
				count0++
			case 1:
				count1++
			case 2:
				count2++
			}
		}
		if count0 < leastZeros {
			leastZeros = count0
			youknowIt = count1 * count2
		}
	}
	return youknowIt
}

func renderImage(image [][]int, width int, height int) [][]int {
	// fmt.Println(image)
	finalImage := make([][]int, height)
	for index, _ := range finalImage {
		newRow := make([]int, width)
		finalImage[index] = newRow
		for j, _ := range finalImage[index] {
			finalImage[index][j] = -1
		}
	}
	for i, layer := range image {
		if i == 0 {
			// fmt.Printf("The length of the pixels are: %d\n", len(layer))
		}
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {

				//			fmt.Printf("The pixel index is: %d, %d: %d \n", x, y, x*height+y)
				pixelIndex := width*y + x
				pixel := layer[pixelIndex]
				currentPixel := finalImage[y][x]
				if (finalImage[y][x] == -1) || (currentPixel == 2 && (pixel == 1 || pixel == 0)) {
					// fmt.Printf("Setting pixels [%d][%d]: %d at (%d, %d)\n", i, pixelIndex, pixel, x, y)
					finalImage[y][x] = pixel
				}
			}
		}
	}

	return finalImage
}

func printImage(image [][]int, width int, height int) string {
	pic := make([]string, len(image))
	for i, _ := range image {
		row := make([]string, len(image[i]))
		for j, _ := range row {
			if image[i][j] == 0 {
				row[j] = " "
			} else {
				row[j] = strconv.Itoa(image[i][j])
			}
		}
		pic[i] = strings.Join(row, "")
	}
	return strings.Join(pic, "\n")
}

func day81() {
	fmt.Println("Day 8.1")
	programData := loadFile("./day8_input.txt")
	flatLayers := parseImage(programData, 25, 6)
	value := findIt(flatLayers)
	fmt.Printf("The row is: %d\n", value)
}

func day82() {
	fmt.Println("Day 8.2")
	programData := loadFile("./day8_input.txt")
	message := parseImage(programData, 25, 6)
	imageLayers := renderImage(message, 25, 6)
	image := printImage(imageLayers, 25, 6)
	fmt.Printf("%s\n", image)
}

func day8() {
	fmt.Println("Day 8")
	fmt.Println("---------")
	day81()
	day82()
	fmt.Println()
}
