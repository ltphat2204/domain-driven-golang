package utils

import "math/rand"

func GetRandomColor(palette []string) string {
    if len(palette) == 0 {
        return "#000000" // Default to black if palette is empty
    }
    return palette[rand.Intn(len(palette))]
}

func IsValidColor(color string, palette []string) bool {
    for _, c := range palette {
        if c == color {
            return true
        }
    }
    return false
}