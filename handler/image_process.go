package handler

import (
	"demo/constant"
	"demo/dto"
	"demo/utils"
	"log"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/fogleman/gg"
	"github.com/gin-gonic/gin"
)


const ResourceRoot = "./media/"

func ImageGenerate(c *gin.Context){
	var req dto.ImageGenerateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil{
		log.Printf("failed to bind json, err=%v\n", err)
		utils.OnFailure(c, constant.StatusCodeJsonFail, "failed to bind json")
		return
	}

	switch strings.ToLower(req.Template){
	case "default":
		url, err := defaultImageGenerate(&req)
		if err != nil{
			log.Printf("failed to process image, err=%v\n", err)
			utils.OnFailure(c, constant.StatusCodeSystemError, "failed to process image")
			return
		}
		log.Printf("success to process image, url=%v\n", url)
		utils.OnSuccess(c, dto.ImageGenerateReponse{URL: url})
		return
	default:
		utils.OnFailure(c, constant.StatusCodeInvalidParam, "unrecognized template")
		return
	}

}

// default image generators, return the name(uri) of this image.
func defaultImageGenerate(req *dto.ImageGenerateRequest) (string, error){
	imgFile, err := gg.LoadImage(path.Join(ResourceRoot, "default.jpeg"))
	if err != nil{
		log.Printf("failed to open file, err=%v\n", err)
		return "", err
	}
	size := imgFile.Bounds().Size()
	canvas := gg.NewContext(size.X, size.Y)
	canvas.SetRGB(0,0,0)
	canvas.DrawImage(imgFile, 0, 0)
	if err := canvas.LoadFontFace(path.Join(ResourceRoot, "Fresh Steak.ttf"), 96); err != nil {
		log.Printf("failed to load font file, err=%v\n", err)
		return "", err
    }

	for _, item := range req.Customization{
		switch strings.ToLower(item.Type){
		case "text":
			canvas.DrawStringWrapped(item.Value, float64(size.X)/1.3, float64(size.Y)/3, 0.5, 0.5, float64(size.X/2), float64(size.Y/2), gg.AlignCenter)
		}
	}
	name := path.Join("/root/media/",strconv.FormatInt(time.Now().UnixNano(), 10)+".png")
	canvas.SavePNG(name)
	return name, nil
}