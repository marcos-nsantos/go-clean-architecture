package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/errors"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Execute os handlers primeiro

		// Se houver erros, trata o último erro
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			handleError(c, err, logger)
		}
	}
}

func handleError(c *gin.Context, err error, logger *slog.Logger) {
	switch e := err.(type) {
	case *errors.ValidationError:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

	case *errors.NotFoundError:
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: e.Error(),
		})

	case *errors.InvalidInputError:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: e.Error(),
		})

	case *errors.InternalError:
		logger.Error("internal server error",
			"error", e.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Erro interno do servidor",
		})

	default:
		logger.Error("unknown error",
			"error", err.Error(),
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Erro interno do servidor",
		})
	}
}
