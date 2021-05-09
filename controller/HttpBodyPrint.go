package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func PrintAllData(c *gin.Context)  {
	//log.Print("handle log")
	body,_ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("---body/--- \r\n "+string(body))

	fmt.Println("---header/--- \r\n")
	for k,v :=range c.Request.Header {
		fmt.Println(k,v)
	}
	//fmt.Println("header \r\n",c.Request.Header)

	c.JSON(200,gin.H{
		"receive":"1024",
	})
}