package main

import (
    "fmt"
    "net/http"
    "html/template"
    "strings"
    "log"
    "mime/multipart"
    "crypto/md5"
    "time"
    "io"
    "strconv"
    "os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello Kasheem Lew")
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        // set token to void duplicate submit
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))
        // same as token : fmt.Sprintf("%x", h.Sum(strconv.FormatInt(crutime,
        // 10)))

        t, _ := template.ParseFiles("./templates/login.gtpl")
        t.Execute(w, token)
    } else {
        r.ParseForm()
        token := r.Form.Get("token")

        // check token
        if token != "" {
        } else {
        }

        // r.Form.Getl() for nil
        if len(r.FormValue("username"))==0 {
            fmt.Println("no username!")
        }
        // fmt.Println("username:", r.Form["username"])
        // use r.FormValue("username") instead of
        // r.ParseForm()+r.Form["username"]
        // fmt.Println("password:", r.Form["password"])
        fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
        fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
        template.HTMLEscape(w, []byte(r.Form.Get("username")))
    }
}

func fruitLegal(fruit string) bool {
    slice := []string{"apple", "pear", "banane"}
    for _, item := range slice {
        if item == fruit {
            return true
        }
    }
    return false
}

func genderLegal(gender string) bool {
    slice := []string{"1", "2"}
    for _, v := range slice {
        if gender == v {
            return true
        }
    }
    return false
}

func info(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./templates/info.gtpl")
        t.Execute(w, nil)
    } else {
        v := r.FormValue("fruit")
        g := r.FormValue("gender")
        i := r.Form["interest"]
        if fruitLegal(v) {
            fmt.Println("legal fruit:", v)
        } else {
            fmt.Println("illegal fruit!")
        }
        if genderLegal(g) {
            fmt.Println("legal gender:", g)
        } else {
            fmt.Println("illegal gender!", g)
        }
        fmt.Println("interest: ", i)
    }
}

func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method: ", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("./templates/upload.gtpl")
        t.Execute(w, token)
    } else {
        r.ParseMultipartForm(32 << 20) // max file size
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v",handler.Header)
        f, err := os.OpenFile("./uploadfiles/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    http.HandleFunc("/info", info)
    http.HandleFunc("/upload", upload)
    server := "localhost"
    port := "9090"
    fmt.Println("Server", server, " Listening on", port, "...")
    err := http.ListenAndServe(server+":"+port, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
