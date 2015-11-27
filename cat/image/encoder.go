package image

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"

	"github.com/securityartwork/cat/binanal"
)

// Encodes data to monochromatic-color-scale PNG file
func EncodeColor(file []byte, secDat []binanal.SectionData) (png.Encoder, *image.RGBA) {
	fSize := len(file)
	min := 0
	max := int(math.Sqrt(float64(fSize)))
	binIndex := 0

	// fmt.Printf("[+] File size: %d\n", fSize)
	// fmt.Printf("[+] Max vector size: %d\n", max)
	// fmt.Printf("[+] Total vectors: %d\n", fSize/max)

	binImage := image.NewRGBA(
		image.Rect(min, min, max, max),
	)

	var c color.Color
	sectionNumber := 0

	// Fill the image with the file bytes
	for y := min; y < max; y++ {
		for x := min; x < max; x++ {
			// Set section color delimiters
			idxA := binIndex > secDat[sectionNumber].Offset
			idxB := binIndex < secDat[sectionNumber].End
			lim := sectionNumber < len(secDat)-1

			// Increase section number
			if binIndex > secDat[sectionNumber].End && lim {
				sectionNumber++
			}

			// If the same section map with color
			if idxA && idxB && lim {
				// Get a color for every different section
				// NOTE: "sectionNumber%12" is for color table security, only 12 colors
				// available at the moment.
				c = GetColor(sectionNumber%12, file[binIndex])

			} else {
				c = color.RGBA{
					uint8(file[binIndex]),
					uint8(file[binIndex]),
					uint8(file[binIndex]),
					uint8(255),
				}
			}

			binIndex++
			binImage.Set(x, y, c)
		}
	}

	var enc png.Encoder
	enc.CompressionLevel = 0

	return enc, binImage
}

// Encodes data to black and white PNG file
func EncodeBW(file []byte) (png.Encoder, *image.RGBA) {
	fSize := len(file)
	min := 0
	max := int(math.Sqrt(float64(fSize)))
	binIndex := 0

	fmt.Printf("[+] File size: %d\n", fSize)
	fmt.Printf("[+] Max vector size: %d\n", max)
	fmt.Printf("[+] Total vectors: %d\n", fSize/max)

	binImage := image.NewRGBA(
		image.Rect(min, min, max, max),
	)

	var c color.Color

	// Fill the image with the file bytes
	for y := min; y < max; y++ {
		for x := min; x < max; x++ {
			c = color.RGBA{
				uint8(file[binIndex]),
				uint8(file[binIndex]),
				uint8(file[binIndex]),
				uint8(255),
			}

			binIndex++
			binImage.Set(x, y, c)
		}
	}

	var enc png.Encoder
	enc.CompressionLevel = 0

	return enc, binImage
}
