package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
)

func main() {
    db_uri := ""
    db, err := sql.Open("mysql", db_uri)
    checkErr(err)

    // Insert Data
    stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?, created=?")
    checkErr(err)

    res, err := stmt.Exec("kasheemlew", "backend", "2017-03-12")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)
    // Update Data
    stmt, err = db.Prepare("update userinfo set username=?, where uid=?")
    checkErr(err)

    res, err = stmt.Exec("kasheemlewupdate", id)
    checkErr(err)

    affect, err := res.RowAffected()
    checkErr(err)

    fmt.Println(affect)

    // Data Query
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    // Delete Data
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)
    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
