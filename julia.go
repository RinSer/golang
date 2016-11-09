package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
    "math"
    "math/cmplx"
    "bufio"
    "os"
)

func Show(dx, dy int, data [][]uint8) {
	
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
    // Create img file
    img_file, err :=  os.Create("julia.png")
    if err != nil {
        panic(err)
    }
    img_drawer := bufio.NewWriter(img_file)
    er := png.Encode(img_drawer, m)
    if er != nil {
        panic(err)
    }
    img_drawer.Flush()
    img_file.Close()
	//ShowImage(m)
}


func ShowImage(m image.Image) {
    
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}


func JuliaSet(dx, dy int, a, b, c, d float64, C complex128) [][]uint8 {
    // Compute the R value
    r := (1+math.Sqrt(1+4*cmplx.Abs(C)))/2
    // Initialize the complex plane
    cplane := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        cplane[y] = make([]uint8, dx)
        for x := 0; x < dx; x++ {
            var z complex128
            re := a+float64(x)/float64(dx)*(b-a)
            im := c+float64(y)/float64(dy)*(d-c)
            z = complex(re, im)
            f := z*z + C
            iteration := 0
            for iteration < 255 {
                if cmplx.Abs(f) > r {
                    cplane[y][x] = uint8(255-iteration*3)
                    break
                }
                f = f*f + C
                iteration++
            }
            if iteration == 255 {
                cplane[y][x] = uint8(0)
            }
        }
    }
    return cplane
}


func main() {
    set := JuliaSet(640, 640, -2, 2, -2, 2, complex(0, -1/float64(2)))
	Show(640, 640, set)
}

