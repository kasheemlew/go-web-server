package main

import (
    "fmt"
    "net/http"
    "html/template"
    "strings"
    "log"
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
        t, _ := template.ParseFiles("./templates/login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        // r.Form.Getl() for nil
        if len(r.FormValue("username"))==0 {
            fmt.Println("no username!")
        }
        fmt.Println("username:", r.Form["username"])
        // use r.FormValue("username") instead of
        // r.ParseForm()+r.Form["username"]
        fmt.Println("password:", r.Form["password"])
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

func main() {
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    http.HandleFunc("/info", info)
    server := "localhost"
    port := "9090"
    fmt.Println("Server", server, " Listening on", port, "...")
    err := http.ListenAndServe(server+":"+port, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
