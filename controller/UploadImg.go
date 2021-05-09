package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	kits "github.com/taoshihan1991/imaptool/tools"
	"os"
	"path"
	"strings"
	"time"
)

func UploadImg(c *gin.Context){
	f, err := c.FormFile("image")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "上传失败!",
		})
		return
	} else {

		fileExt:=strings.ToLower(path.Ext(f.Filename))
		if fileExt!=".png"&&fileExt!=".jpg"&&fileExt!=".gif"&&fileExt!=".jpeg"{
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
			})
			return
		}

		fileName:=kits.Md5(fmt.Sprintf("%s%s",f.Filename,time.Now().String()))
		fildDir:=fmt.Sprintf("Image/%d%s/",time.Now().Year(),time.Now().Month().String())
		isExist,_:=kits.IsFileExist(fildDir)
		if !isExist{
			os.Mkdir(fildDir,os.ModePerm)
		}
		filepath:=fmt.Sprintf("%s%s%s",fildDir,fileName,fileExt)
		c.SaveUploadedFile(f, filepath)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传成功!",
			"result":gin.H{
				"path":filepath,
			},
		})
	}
}
