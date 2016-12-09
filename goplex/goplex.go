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
	
	m := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = 0
			m.Pix[i+2] = 0
			m.Pix[i+3] = 255-v
		}
	}
    // Create img file
    fmt.Println("Creating img file")
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
    fmt.Println("img file created")
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
                    cplane[y][x] = uint8(iteration)
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
    //r.ParseForm()
    //form := r.Form
    //fmt.Println(form)
    params := strings.Split(url, "_")
    html, err := template.ParseFiles("appface/goplex.html")
    if err != nil {
        panic(err)
    }
    // Screen Resolution
    if len(params) > 2 {
        
    //xresolution, _ := strconv.ParseInt(params[2], 10, 16)
    //yresolution, _ := strconv.ParseInt(params[4], 10, 16)
    xresolution := 800
    yresolution := 800
    // Params
    var xmin, xmax, ymin, ymax float64
    fmt.Println(params)
    if len(params) < 10 {
        xmin = -2
        xmax = 0.75
        ymin = -1.5
        ymax = 1.5
    } else {
        xmin, _ = strconv.ParseFloat(params[6], 64)
        xmax, _ = strconv.ParseFloat(params[8], 64)
        ymin, _ = strconv.ParseFloat(params[10], 64)
        ymax, _ = strconv.ParseFloat(params[12], 64)
    }
    
    if len(params) > 13 {
        fmt.Println(params)
        new_xmin, _ := strconv.ParseFloat(params[13], 64)
        new_xmax, _ := strconv.ParseFloat(params[14], 64)
        new_ymin, _ := strconv.ParseFloat(params[15], 64)
        new_ymax, _ := strconv.ParseFloat(params[16], 64)
        xmin = new_xmin // xmin+float64(new_xmin)/float64(xresolution)*(xmax-xmin)
        xmax = new_xmax // xmin+float64(new_xmax)/float64(xresolution)*(xmax-xmin)
        ymin = new_ymin // ymin+float64(new_ymin)/float64(yresolution)*(ymax-ymin)
        ymax = new_ymax // ymin+float64(new_ymax)/float64(yresolution)*(ymax-ymin)
    }
    page_url := fmt.Sprintf("_r_%d_x_%d_Xmin_%f_Xmax_%f_Ymin_%f_Ymax_%f", xresolution, yresolution, xmin, xmax, ymin, ymax)
    fmt.Println(page_url)
    mandelbrot_set := Mandelbrot(int(xresolution), int(yresolution), xmin, xmax, ymin, ymax)
	Show(int(xresolution), int(yresolution), mandelbrot_set)
    set := Set{Xmin:xmin, Xmax:xmax, Ymin:ymin, Ymax:ymax}
    if len(params) < 7 || len(params) > 13 {
        http.Redirect(w, r, "/"+page_url, http.StatusFound)
    }
    html.Execute(w, set)
    } else {
        html.Execute(w, nil)
    }
}


func main() {
	http.HandleFunc("/", viewHandler)
    http.Handle("/set/", http.StripPrefix("/set/", http.FileServer(http.Dir("img"))))
    http.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("appface"))))
	http.ListenAndServe(":8080", nil)
}


