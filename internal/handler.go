package internal

import (
	"USERLOGIN/model"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service actions
}

type actions interface {
	UserLogin(username, pass string) (model.User, error)
}

func NewHandler(service actions) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/login/username/:username/pass/:pass", h.UserLogin) //I know that is not safe, i am working on this !!!!
}

func (h *Handler) UserLogin(c *fiber.Ctx) error {

	username := c.Params("username")
	pass := c.Params("pass")

	_, err := h.service.UserLogin(username, pass)
	if err != nil {
		return err
	}

	return nil
}
