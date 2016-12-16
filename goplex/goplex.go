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
    //"math/rand"
    //"time"
)

/*
func RandStringRunes(n int) string {
    rand.Seed(time.Now().UnixNano())
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
*/

func Show(dx, dy int, data [][]uint8, background, color, pic_name string) {
	
	m := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = 255-v
			m.Pix[i+1] = 0
			m.Pix[i+2] = 0
			m.Pix[i+3] = 255-v
		}
	}
    // Create img file
    fmt.Println("Creating img file")
    filename := fmt.Sprintf("img/%s.png", pic_name)
    img_file, err :=  os.Create(filename)
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
    PicPath string
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
    // Session number
    //session := RandStringRunes(10)
    params := strings.Split(url, "_")
    html, err := template.ParseFiles("appface/goplex.html")
    if err != nil {
        panic(err)
    }
    // Screen Resolution
    if len(params) > 3 {
        
        var xresolution, yresolution int64        
        xscreen, _ := strconv.ParseInt(params[2], 10, 16)
        yscreen, _ := strconv.ParseInt(params[4], 10, 16)
        xresolution = 800
        yresolution = 800
        if xscreen < xresolution {
            xresolution = xscreen
            yresolution = xscreen
        }
        if yscreen < yresolution {
            xresolution = yscreen
            yresolution = yscreen
        }
        // Params
        var xmin, xmax, ymin, ymax float64
        // Colors
        var bw, rgb string
        //fmt.Println(url)
        if len(params) < 10 {
            xmin = -2
            xmax = 0.75
            ymin = -1.5
            ymax = 1.5
            bw = "b"
            rgb = "r"
        } else {
            xmin, _ = strconv.ParseFloat(params[6], 64)
            xmax, _ = strconv.ParseFloat(params[8], 64)
            ymin, _ = strconv.ParseFloat(params[10], 64)
            ymax, _ = strconv.ParseFloat(params[12], 64)
            bw = params[13]
            rgb = params[14]
        }
        // Create the new url
        page_url := fmt.Sprintf("_r_%d_x_%d_Xmin_%f_Xmax_%f_Ymin_%f_Ymax_%f_%s_%s", xresolution, yresolution, xmin, xmax, ymin, ymax, bw, rgb)
        fmt.Println(page_url)
        pic_path := "set/"+page_url+".png"
        // Create the pic if it does not already exist
        if _, err := os.Stat("img/"+page_url+".png"); os.IsNotExist(err) {
          mandelbrot_set := Mandelbrot(int(xresolution), int(yresolution), xmin, xmax, ymin, ymax)
	      Show(int(xresolution), int(yresolution), mandelbrot_set, bw, rgb, page_url)
        }
        set := Set{Xmin:xmin, Xmax:xmax, Ymin:ymin, Ymax:ymax, PicPath:pic_path}
        if len(params) < 9 || len(params) > 15 {
            fmt.Println(params)
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


