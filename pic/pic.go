package pic

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/xalanq/go-tracing/vec"
)

// Pic store pixel
type Pic struct {
	W, H int
	C    []vec.Vec
}

// New new one
func New(w, h int) *Pic {
	return &Pic{W: w, H: h, C: make([]vec.Vec, w*h)}
}

// Clamp x -> [0, 1]
func Clamp(x float64) float64 {
	if x < 0.0 {
		return 0.0
	} else if x > 1.0 {
		return 1.0
	}
	return x
}

// ToByte 0~1 pixel value to 0~255 value through Gamma Correction
func ToByte(x float64) byte {
	return byte(math.Pow(Clamp(x), 1.0/2.2)*255 + 0.5)
}

// SavePPM save as .ppm
func (a *Pic) SavePPM(filename string) {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	writer := bufio.NewWriter(file)
	fmt.Printf("Writing to %v\n", filename)
	writer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", a.W, a.H))
	for _, c := range a.C {
		writer.WriteString(fmt.Sprintf("%v %v %v ", ToByte(c.X), ToByte(c.Y), ToByte(c.Z)))
	}
	file.Close()
	fmt.Println("Done.")
}
