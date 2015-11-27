package image

import (
	"image/color"
)

// Generate color with the chromatic scale value setted
func GetColor(selector int, value byte) color.Color {
	palette := []color.Color{
		// Red
		color.RGBA{
			uint8(value),
			uint8(16),
			uint8(16),
			uint8(255),
		},
		// Green
		color.RGBA{
			uint8(16),
			uint8(value),
			uint8(16),
			uint8(255),
		},
		// Blue
		color.RGBA{
			uint8(16),
			uint8(16),
			uint8(value),
			uint8(255),
		},
		// Yellow
		color.RGBA{
			uint8(value),
			uint8(value),
			uint8(16),
			uint8(255),
		},
		// Turquoise
		color.RGBA{
			uint8(16),
			uint8(value),
			uint8(value),
			uint8(255),
		},
		// Pink
		color.RGBA{
			uint8(value),
			uint8(16),
			uint8(value),
			uint8(255),
		},
		// == … ==
		color.RGBA{
			uint8(value),
			uint8(32),
			uint8(32),
			uint8(255),
		},
		// …
		color.RGBA{
			uint8(32),
			uint8(value),
			uint8(32),
			uint8(255),
		},
		// …
		color.RGBA{
			uint8(32),
			uint8(32),
			uint8(value),
			uint8(255),
		},
		// …
		color.RGBA{
			uint8(value),
			uint8(value),
			uint8(32),
			uint8(255),
		},
		// …
		color.RGBA{
			uint8(32),
			uint8(value),
			uint8(value),
			uint8(255),
		},
		// …
		color.RGBA{
			uint8(value),
			uint8(32),
			uint8(value),
			uint8(255),
		},
	}

	return palette[selector]
}
