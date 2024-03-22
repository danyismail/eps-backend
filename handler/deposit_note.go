package handler

import (
	"eps-backend/model"
	"eps-backend/structs"
	"eps-backend/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func (h *Handler) CreateDeposit(c echo.Context) error {
	h.e.Logger.Info("::CreateDeposit::")
	name := c.FormValue("name")
	supplier := c.FormValue("supplier")
	amount := c.FormValue("amount")
	originAccount := c.FormValue("origin_account")
	destinationAccount := c.FormValue("destination_account")

	i, _ := strconv.Atoi(amount)
	notes := model.DepositNote{
		CreatedAt:          time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:          time.Now().Format("2006-01-02 15:04:05"),
		Name:               name,
		Supplier:           supplier,
		Amount:             float64(i),
		OriginAccount:      originAccount,
		DestinationAccount: destinationAccount,
		Status:             "pending",
	}
	err := h.depositNoteStore.Create(notes, c.Param("e"))
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
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
	h.e.Logger.Info("::GetDeposit::")
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	note, err := h.depositNoteStore.GetById(i, c.Param("e"))
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	if note == nil {
		h.e.Logger.Error("note is nil")
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    "record not found",
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       note,
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}

func (h *Handler) CancelDeposit(c echo.Context) error {
	h.e.Logger.Info("::CancelDeposit::")
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	note, err := h.depositNoteStore.GetById(i, c.Param("e"))
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	if note == nil {
		h.e.Logger.Error("note is nil")
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    "record not found",
		})
	}

	//delete image
	err = deleteImage(note.ImageUpload)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	err = h.depositNoteStore.Delete(i, c.Param("e"))
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
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

func (h *Handler) GetAllDeposit(c echo.Context) error {
	h.e.Logger.Info("::GetAllDeposit::")
	dt := c.QueryParam("dt")
	note, err := h.depositNoteStore.GetAllStatus(c.Param("e"), dt)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
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

func (h *Handler) GetDepositCreated(c echo.Context) error {
	h.e.Logger.Info("::GetDepositCreated::")
	note, err := h.depositNoteStore.GetStatusCreated(c.Param("e"))
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
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

func (h *Handler) GetDepositUploaded(c echo.Context) error {
	h.e.Logger.Info("::GetDepositCreated::")
	note, err := h.depositNoteStore.GetStatusUploaded(c.Param("e"))
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
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

func (h *Handler) GetDepositDone(c echo.Context) error {
	h.e.Logger.Info("::GetDepositCreated::")
	startDt := c.QueryParam("startDt")
	endDt := c.QueryParam("endDt")
	note, err := h.depositNoteStore.GetStatusDone(c.Param("e"), startDt, endDt)
	if err != nil {
		h.errorBot.SendMessage(err)
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
	h.e.Logger.Info("::GetImage::")
	id := c.Param("id")

	// Assuming images are stored in a directory named "uploads"
	envr := c.Param("e")
	imagePath := "uploads/dev/" + id + ".jpg" // Adjust the file extension as needed
	if envr == utils.DIGI_EPS {
		imagePath = "uploads/prod/" + id + ".jpg"
	}

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
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
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to send image",
		})
	}
	return nil
}

func (h *Handler) DeleteImage(c echo.Context) error {
	h.e.Logger.Info("::GetImage::")
	id := c.Param("id")

	// Assuming images are stored in a directory named "uploads"
	envr := c.Param("e")
	imagePath := "uploads/dev/" + id + ".jpg" // Adjust the file extension as needed
	if envr == utils.DIGI_AMAZONE {
		imagePath = "uploads/prod/" + id + ".jpg"
	}

	err := os.Remove(imagePath)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusNotFound,
			Message:    "image not found",
		})
	}
	h.e.Logger.Info("delete image from " + imagePath + " successfully")
	return nil
}

func (h *Handler) UpdateDeposit(c echo.Context) error {
	h.e.Logger.Info("::UpdateDeposit::")
	//logic form value
	envr := c.Param("e")
	id := c.Param("id")
	name := c.FormValue("name")
	supplier := c.FormValue("supplier")
	amount := c.FormValue("amount")
	originAccount := c.FormValue("origin_account")
	destinationAccount := c.FormValue("destination_account")
	reply := c.FormValue("reply")

	i, _ := strconv.Atoi(amount)
	intID, _ := strconv.Atoi(id)

	note, err := h.depositNoteStore.GetById(intID, c.Param("e"))
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return err
	}
	if note == nil {
		h.e.Logger.Error("note is nil")
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    "record not found",
		})
	}

	// Get the image file from the form
	var imagePath string
	file, err := c.FormFile("image")
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
	} else {
		// Open the uploaded file
		src, err := file.Open()
		if err != nil {
			h.e.Logger.Error(err)
			h.errorBot.SendMessage(err)
			return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
				Data:       nil,
				StatusCode: http.StatusInternalServerError,
				Message:    "error open file",
			})
		}
		defer src.Close()

		// Create a destination file
		imagePath = "uploads/dev/" + id + ".jpg" // Adjust the file extension as needed
		if envr == utils.DIGI_AMAZONE {
			imagePath = "uploads/prod/" + id + ".jpg"
		}
		dst, err := os.Create(imagePath)
		if err != nil {
			h.e.Logger.Error(err)
			h.errorBot.SendMessage(err)
			return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
				Data:       nil,
				StatusCode: http.StatusInternalServerError,
				Message:    "error create file",
			})
		}
		defer dst.Close()

		// Copy the uploaded file to the destination file
		if _, err = io.Copy(dst, src); err != nil {
			h.e.Logger.Error(err)
			h.errorBot.SendMessage(err)
			return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
				Data:       nil,
				StatusCode: http.StatusInternalServerError,
				Message:    "error copy file",
			})
		}
	}

	// Create a notes object
	notes := model.DepositNote{
		ID:                 note.ID,
		UpdatedAt:          time.Now().Format("2006-01-02 15:04:05"),
		Name:               name,
		Supplier:           supplier,
		Amount:             float64(i),
		OriginAccount:      originAccount,
		DestinationAccount: destinationAccount,
		ImageUpload:        imagePath,
		Reply:              reply,
	}
	if notes.ImageUpload != "" {
		notes.Status = "process"
	}
	if notes.Reply != "" {
		notes.Status = "success"
	}
	//end of logic
	err = h.depositNoteStore.Update(notes, envr)
	if err != nil {
		h.e.Logger.Error(err)
		h.errorBot.SendMessage(err)
		return c.JSON(http.StatusInternalServerError, structs.CommonResponse{
			Data:       nil,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, structs.CommonResponse{
		Data:       nil,
		StatusCode: http.StatusOK,
		Message:    "success",
	})
}

func deleteImage(imagePath string) error {
	fmt.Println("deleteImage " + imagePath)
	err := os.Remove(imagePath)
	if err != nil {
		return err
	}
	fmt.Println("deleteImage " + imagePath + " successfully")
	return nil
}
