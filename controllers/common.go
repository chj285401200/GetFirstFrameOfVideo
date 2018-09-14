/*
@Time : 2018/9/10 16:29 
@Author : CaoHuaijie
@File : common
*/
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/tnextday/goseaweed"
	"os/exec"
	"fmt"
	"bytes"
	"time"
	"strconv"
	"os"
)

var Weed *goseaweed.Seaweed

type CommonController struct {
	beego.Controller
}
func (c *CommonController) URLMapping() {
	c.Mapping("POST", c.GetFirstFrame)

}

func (r *CommonController) GetFirstFrame() {
	result := NewResponse()
	f, h, _ := r.GetFile("myfile") 
	SDPATH := ("F:\\newfile\\")
	logs.Debug("Filename:", h.Filename)
	path := SDPATH + h.Filename 
	logs.Debug("path:", path)
	f.Close() 
	r.SaveToFile("myfile", path)
	
	picPath := getFirstFrame(path)
	result.Code = 200
	url :=  "http://127.0.0.1:8080/"
	logs.Debug("url:", url)
	result.Data = url+"1"
	result.Data2 = url +picPath
	r.Data["json"] = &result
	r.ServeJSON()
}


func getFirstFrame(filename string ) (picPath string){

	width := 1000
	height := 750
	
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	fmt.Println(cmd)
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	logs.Debug(buffer,"---------1")
	if cmd.Run() != nil {
		panic("could not generate frame")
	}

	name := time.Now().UnixNano()
	name2:=strconv.FormatInt(name,10)

	path := "F:/newfile/"+ name2 +".jpg"
	fmt.Println(path)
	f,err := os.Create(path)
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		_,err=f.Write(buffer.Bytes())
	}
	picPath = path
	return picPath
}
