package main

import (
        "github.com/gin-gonic/gin"
        cors "github.com/rs/cors/wrapper/gin"
        "log"
        //"net/http"
        //"encoding/json"
        //"fmt"
        //"os"
        "io/ioutil"
        //"strings"
        //"errors"
        //client "github.com/influxdata/influxdb1-client/v2"

        //"time"
        //"errors"
        //"strconv"
        //"github.com/gin-contrib/cors"
        //client "github.com/influxdata/influxdb/client/v2"
)


func main() {

        log.Println("main(): Entering main()")
        log.Println("main(): Starting web server")

        StartWebServer()
}

func StartWebServer() {

        // ROUTER WITH CUSTOM SETTINGS 2 (with "github.com/rs/cors/wrapper/gin")
        router := gin.Default()
        router.Use(cors.AllowAll())

        // homepage
        router.GET("/", func(c *gin.Context) {
                c.JSON(200, gin.H{
                     "message": "Welcome!",
                })
        })

        router.GET("/message", func(c *gin.Context) {
	        message, err := ReadMessageFromFile();

                if (err != nil) {
                	c.JSON(200, gin.H{"message": "Error reading message file",})
		} else {
                	c.JSON(200, gin.H{"message": message,})
                }
        })

        // start the web server
        router.Run(":8081")

}

func ReadMessageFromFile() (message string, err error){

	 var messageFile string = "/tmp/messages/message1.txt"	

        b, err := ioutil.ReadFile(messageFile) // just pass the file name
        if err != nil {
	    return "", err
        }

        str := string(b) // convert content to a 'string'

        //fmt.Println(b) // print the content as 'bytes'
        //fmt.Println(str) // print the content as a 'string'

        return str, nil
}
