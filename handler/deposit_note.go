package handler

import (
	"eps-backend/model"
	"eps-backend/structs"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
)

// id
// created_at
// updated_at
// deleted_at
// name
// supplier
// amount
// origin_account
// destination_account
// image_upload
// reply
// status

func (h *Handler) CreateDeposit(c echo.Context) error {
	c.Logger().Info("::CreateDeposit::")
	name := c.FormValue("name")
	supplier := c.FormValue("supplier")
	amount := c.FormValue("amount")
	originAccount := c.FormValue("origin_account")
	destinationAccount := c.FormValue("destination_account")

	i, _ := strconv.Atoi(amount)
	notes := model.DepositNote{
		Name:               name,
		Supplier:           supplier,
		Amount:             float64(i),
		OriginAccount:      originAccount,
		DestinationAccount: destinationAccount,
	}

	err := h.depositNoteStore.Create(notes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       "result",
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}

func (h *Handler) GetDeposit(c echo.Context) error {
	c.Logger().Info("::GetDeposit::")
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	note, err := h.depositNoteStore.GetById(i)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       note,
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}

func (h *Handler) GetImage(c echo.Context) error {
	c.Logger().Info("::GetImage::")
	id := c.Param("id")

	// Assuming images are stored in a directory named "uploads"
	imagePath := "uploads/" + id + ".jpg" // Adjust the file extension as needed

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusNotFound,
			Message:    "image not found",
		})
	}
	defer file.Close()

	// Set the appropriate content type header
	c.Response().Header().Set("Content-Type", "image/jpeg") // Adjust content type based on image format

	// Copy the image file to the response
	_, err = io.Copy(c.Response(), file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to send image",
		})
	}
	return nil
}

func (h *Handler) UpdateDeposit(c echo.Context) error {
	c.Logger().Info("::UpdateDeposit::")
	//logic form value
	id := c.Param("id")
	name := c.FormValue("name")
	supplier := c.FormValue("supplier")
	amount := c.FormValue("amount")
	originAccount := c.FormValue("origin_account")
	destinationAccount := c.FormValue("destination_account")

	i, _ := strconv.Atoi(amount)
	intID, _ := strconv.Atoi(id)

	// Get the image file from the form
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create a destination file
	imagePath := "uploads/" + id + ".jpg"
	dst, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// Create a notes object
	notes := model.DepositNote{
		ID:                 intID,
		Name:               name,
		Supplier:           supplier,
		Amount:             float64(i),
		OriginAccount:      originAccount,
		DestinationAccount: destinationAccount,
		ImageUpload:        imagePath,
	}

	//end of logic
	err = h.depositNoteStore.Update(notes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       "note",
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}
