package gin_handlers

import (
	"archetype-golang-ddd/internal/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewCreateOrderHttpHandler(useCase *application.CreateOrderCommandHandler) *CreateOrderHttpHandler {
	return &CreateOrderHttpHandler{
		useCase: useCase,
	}
}

type CreateOrderHttpHandler struct {
	useCase *application.CreateOrderCommandHandler
}

func (h *CreateOrderHttpHandler) Handle(c *gin.Context) {

	command := application.CreateOrderCommand{
		ProductIds: []string{"1", "2", "3"},
		ClientName: "test Name",
		ClientRut:  "111111-1",
	}

	order, err := h.useCase.Exec(command)

	if err != nil {
		response := GenericResponse{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		}

		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	response := GenericResponse{
		Code:    http.StatusCreated,
		Success: true,
		Data: DataWrapper{
			Data: order,
		},
		Message: "",
	}

	c.IndentedJSON(http.StatusCreated, response)
}
