package main

import "fmt"

import "github.com/gin-gonic/gin"
//import "github.com/gin-gonic/contrib/gzip"


func main() {
    r := gin.Default()
    r.Use(mwAll)

    r.GET("/", func(c *gin.Context) {
        c.String(200, "This is a blank page\n")
    })
    r.GET("/ping1", mwKarl, ping) // use middleware "mw"
    r.GET("/ping2", ping) // no middleware "mw"
    r.Run(":8000")

}

func mwKarl(c *gin.Context){
    c.Header("KARL-ONLY", "Hello")
    c.Next()
}
func mwAll(c *gin.Context){
    c.Header("TO-ALL", "Good Morning")
    c.Next()
}

func ping(c *gin.Context){
    fmt.Println("hello")
    c.String(200, "pongpongpongpongpongpongpong\n")
}


