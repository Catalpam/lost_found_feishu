package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"lost_found/api"
	"lost_found/api/core/request"
	"lost_found/api/core/response"
	"lost_found/core"
	"lost_found/core/tools"
)

type UploadImage struct {
	ImageKey string `json:"image_key"`
}

// upload image
func Uploadimage2Feishu(imageName string) (string, error) {
	coreCtx := core.WrapContext(context.Background())
	// coreCtx.Set(constants.HTTPHeaderKeyRequestID, "2020122212081301001702714534518-xxxxx")
	var formData = request.NewFormData()
	formData.AddParam("image_type", "message")
	bs, err := ioutil.ReadFile("./images/" + imageName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	formData.AddFile("image", request.NewFile().SetContent(bs))
	/*
		// stream upload, file implement io.Reader
		file, err := os.Open("test.png")
		if err != nil {
			fmt.Println(err)
			return
		}
		formData.AddFile("image", request.NewFile().SetContentStream(file))
	*/
	ret := &UploadImage{}
	err = api.Send(coreCtx, conf, request.NewRequestWithNative("image/v4/put", "POST",
		request.AccessTokenTypeTenant, formData, ret))
	fmt.Println(coreCtx.GetRequestID())
	fmt.Println(coreCtx.GetHTTPStatusCode())
	if err != nil {
		fmt.Println(tools.Prettify(err))
		e := err.(*response.Error)
		fmt.Println(e.Code)
		fmt.Println(e.Msg)
		return "", err
	}
	fmt.Println("图片上传至飞书服务器成功，Imageey: " + ret.ImageKey)
	return ret.ImageKey, nil
}
