package main


import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
    "math/cmplx"
    "bufio"
    "os"
    "net/http"
    "html/template"
    "strings"
    "strconv"
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
    img_file, err :=  os.Create("img/mandelbrot.png")
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


func Mandelbrot(dx, dy int, xs, xe, ys, ye float64) [][]uint8 {
    // Initialize the complex plane
    cplane := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        cplane[y] = make([]uint8, dx)
        for x := 0; x < dx; x++ {
            re := xs+float64(x)/float64(dx)*(xe-xs)
            im := ys+float64(y)/float64(dy)*(ye-ys)
            c := complex(re, im)
            z := complex(0, 0)
            f := z*z + c
            iteration := 0
            for iteration < 255 {
                if cmplx.Abs(f) > 2 {
                    cplane[y][x] = uint8(255-iteration)
                    break
                }
                f = f*f + c
                iteration++
            }
            if iteration == 255 {
                cplane[y][x] = uint8(0)
            }
        }
    }
    return cplane
}


type Set struct {
    Xmin float64
    Xmax float64
    Ymin float64
    Ymax float64
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
    params := strings.Split(url, "_")
    html, err := template.ParseFiles("appface/index.html")
    if err != nil {
        panic(err)
    }
    // Params
    var xmin, xmax, ymin, ymax float64
    fmt.Println(params)
    if len(params) < 2 {
        xmin = -2
        xmax = 0.75
        ymin = -1.5
        ymax = 1.5
    } else {
        xmin, _ = strconv.ParseFloat(params[1], 64)
        xmax, _ = strconv.ParseFloat(params[3], 64)
        ymin, _ = strconv.ParseFloat(params[5], 64)
        ymax, _ = strconv.ParseFloat(params[7], 64)
    }
    xresolution := 900
    yresolution := 900
    if len(params) > 8 {
        new_xmin, _ := strconv.ParseFloat(params[8], 64)
        new_xmax, _ := strconv.ParseFloat(params[9], 64)
        new_ymin, _ := strconv.ParseFloat(params[10], 64)
        new_ymax, _ := strconv.ParseFloat(params[11], 64)
        xmin = xmin+float64(new_xmin)/float64(xresolution)*(xmax-xmin)
        xmax = xmin+float64(new_xmax)/float64(xresolution)*(xmax-xmin)
        ymin = ymin+float64(new_ymin)/float64(yresolution)*(ymax-ymin)
        ymax = ymin+float64(new_ymax)/float64(yresolution)*(ymax-ymin)
    }
    page_url := fmt.Sprintf("Xmin_%f_Xmax_%f_Ymin_%f_Ymax_%f", xmin, xmax, ymin, ymax)
    fmt.Println(page_url)
    mandelbrot_set := Mandelbrot(xresolution, yresolution, xmin, xmax, ymin, ymax)
	Show(xresolution, yresolution, mandelbrot_set)
    set := Set{Xmin:xmin, Xmax:xmax, Ymin:ymin, Ymax:ymax}
    html.Execute(w, set)
}


func main() {
	http.HandleFunc("/", viewHandler)
    http.Handle("/set/", http.StripPrefix("/set/", http.FileServer(http.Dir("img"))))
	http.ListenAndServe(":8080", nil)
}


