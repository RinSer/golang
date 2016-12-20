package main


import (
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
    "math"
)


func MakePic(dx, dy int, data [][]uint8, background, color, pic_name string) {
    // Color settings
    var red, green, blue uint8
    if strings.IndexRune(color, 'r') == -1 {
        red = 0
    } else {
        red = 255
    }
    if strings.IndexRune(color, 'g') == -1 {
        green = 0
    } else {
        green = 255
    }
    if strings.IndexRune(color, 'b') == -1 {
        blue = 0
    } else {
        blue = 255
    }
	// Image creation
	m := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
            i := y*m.Stride + x*4
            if background == "b" {
                if red == 255 {
			        m.Pix[i] = 255-v
                } else {
                    m.Pix[i] = v
                }
                if green == 255 {
                    m.Pix[i+1] = 255-v
                } else {
			        m.Pix[i+1] = v
                }
                if blue == 255 {
                    m.Pix[i+2] = 255-v
                } else {
			        m.Pix[i+2] = v
                }
			    m.Pix[i+3] = 255-v
            } else {
                if red == 255 {
			        m.Pix[i] = 255-v
                } else {
                    m.Pix[i] = v
                }
                if green == 255 {
                    m.Pix[i+1] = 255-v
                } else {
			        m.Pix[i+1] = v
                }
                if blue == 255 {
                    m.Pix[i+2] = 255-v
                } else {
			        m.Pix[i+2] = v
                }
			    m.Pix[i+3] = 255
            }
		}
	}
    // Create img file
    //fmt.Println("Creating img file")
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
    //fmt.Println("img file created")
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
                    cplane[y][x] = uint8(255-iteration*7)
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



type Set struct {
    Xmin float64
    Xmax float64
    Ymin float64
    Ymax float64
    PicPath string
    Cvalue complex128
    ReC float64
    ImC float64
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
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
        var xmin, xmax, ymin, ymax, rec, imc float64
        var c complex128
        // Colors
        var bw, rgb string
        //fmt.Println(url)
        if len(params) < 19 {
            rec = 0
            imc = 1
            xmin = -2
            xmax = 2
            ymin = -2
            ymax = 2
            bw = "b"
            rgb = "r"
        } else {
            rec, _ = strconv.ParseFloat(params[6], 64)
            imc, _ = strconv.ParseFloat(params[8], 64)
            xmin, _ = strconv.ParseFloat(params[10], 64)
            xmax, _ = strconv.ParseFloat(params[12], 64)
            ymin, _ = strconv.ParseFloat(params[14], 64)
            ymax, _ = strconv.ParseFloat(params[16], 64)
            bw = params[17]
            rgb = params[18]
        }
        c = complex(rec, imc)
        // Create the new url
        page_url := fmt.Sprintf("_r_%d_x_%d_ReC_%f_ImC_%f_Xmin_%f_Xmax_%f_Ymin_%f_Ymax_%f_%s_%s", xresolution, yresolution, rec, imc, xmin, xmax, ymin, ymax, bw, rgb)
        //fmt.Println(page_url)
        pic_path := "set/"+page_url+".png"
        // Create the pic if it does not already exist
        if _, err := os.Stat("img/"+page_url+".png"); os.IsNotExist(err) {
          julia_set := JuliaSet(int(xresolution), int(yresolution), xmin, xmax, ymin, ymax, c)
	      MakePic(int(xresolution), int(yresolution), julia_set, bw, rgb, page_url)
        }
        set := Set{Xmin:xmin, Xmax:xmax, Ymin:ymin, Ymax:ymax, PicPath:pic_path, Cvalue:c, ReC:rec, ImC:imc}
        if len(params) != 19 {
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


