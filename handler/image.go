package handler

import (
	"net/http"
	"os"
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//var filePath string

type ImageHandler struct {
	svc repo.ImageRepository
	log *logger.Logger
}

func NewimageHandler(svc repo.ImageRepository, log *logger.Logger) *ImageHandler {
	return &ImageHandler{
		svc,
		log,
	}
}

type uploadimage struct {
	Filename        string //`json:"filename"`
	Size            int    //`json:"size"`
	Mimetype        string //`json:"mimetype"`
	Pickuprequestid int    `validate:"required"`
}

// Hanlder function to upload the image of open article

// CaptureImage godoc
//
//	@Summary        Capture the image of the open pickuprequest
//	@Description    Capture the image of the open pickuprequest
//	@Tags           Image
//	@Accept         json
//	@Produce        json
//	@Param          uploadimage   body        uploadimage   true    "uploadimage"
//	@Success        200                     {object}    ImageResponse        		"image captured"
//	@Failure        400                     {object}    errorValidResponse          "Validation error"
//	@Failure        401                     {object}    errorValidResponse          "Unauthorized error"
//	@Failure        403                     {object}    errorValidResponse          "Forbidden error"
//	@Failure        404                     {object}    errorValidResponse          "Data not found error"
//	@Failure        409                     {object}    errorValidResponse          "Data conflict error"
//	@Failure        500                     {object}    errorValidResponse          "Internal server error"
//	@Router         /image/ [post]
func (ih *ImageHandler) uploadimage(ctx *gin.Context) {
	ih.log.Debug("Entered upload image handler method")

	var upl uploadimage

	if err := ctx.ShouldBindWith(&upl, binding.Form); err != nil {
		handleError(ctx, err)
		ih.log.Debug("error occured while binding:", err)
		return
	}
	if !handleValidation(ctx, upl) {
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ih.log.Debug("error creating file:", err)
		handleError(ctx, err)
		return
	}

	//filePath := "D:/uploadedimage/" + file.Filename //in drive image will upload
	//filePath := "OpenArticleImages/" + file.Filename //in project folder image will uploaded in OpenArticleImages folder

	subfolder := "OpenArticleImages/" + strconv.Itoa(upl.Pickuprequestid)
	filePath := subfolder + "/" + file.Filename

	// Ensure the subfolder exists, create it if not
	if err := os.MkdirAll(subfolder, 0755); err != nil {
		ih.log.Debug("error creating subfolder:", err)
		handleError(ctx, err)
		return
	}

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ih.log.Debug("Failed to save file:", err)
		handleError(ctx, err)
		return
	}
	ih.log.Debug("filepath is :", filePath)

	image := domain.Image{
		Filename:        upl.Filename,
		Size:            upl.Size,
		Mimetype:        upl.Mimetype,
		Pickuprequestid: upl.Pickuprequestid,
	}

	_, err = ih.svc.SaveImage(ctx, &image, filePath)
	if err != nil {
		ih.log.Debug("error occured while calling repo:", err)
		handledbError(ctx, err)
		return
	}

	rsp := newImageResponse(&image, filePath)

	handleSuccess(ctx, rsp)

}

type getimagebyid struct {
	Imageid int `uri:"imageid" validate:"required"`
}

// Hanlder function to get the image of the pickuprequest based on imageid

// GetImagebyimageid godoc
//
//	@Summary		Get a image based on imageid
//	@Description	Get a image based on imageid
//	@Tags			Image
//	@Accept			json
//	@Produce		json
//	@Param			imageid	path		int				true	"Imageid"
//	@Success		200	{object}	ImageResponse	"Category retrieved"
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/image/{imageid} [get]
func (ih *ImageHandler) getimagebyimageid(ctx *gin.Context) {

	var image getimagebyid

	if err := ctx.ShouldBindUri(&image); err != nil {
		ih.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, image) {
		return
	}

	imageBytes, imagedetails, err := ih.svc.GetImagedetailsByID(ctx, image.Imageid)
	if err != nil {
		ih.log.Debug("error occured while calling repo function:", err)
		handledbError(ctx, err)
		return
	}

	ctx.Data(http.StatusOK, "image/png", imageBytes)

	rsp := newImageDetailsResponse(imagedetails)

	handleSuccess(ctx, rsp)

}
