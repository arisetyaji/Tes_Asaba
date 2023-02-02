package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type produk struct {
    id    string
    name  string
    harga   int
}

func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_produk")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func sqlQuery() {
    db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    var age = 27
    rows, err := db.Query("select id, name, grade from tb_produk where age = ?", age)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()

    var result []produk

    for rows.Next() {
        var each = produk{}
        var err = rows.Scan(&each.id, &each.name, &each.harga)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, each := range result {
        fmt.Println(each.name)
    }
}

func main() {
    sqlQuery()
}
