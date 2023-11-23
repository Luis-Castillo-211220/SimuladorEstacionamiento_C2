package controllers

import (
	"Prueba1/models"
	"Prueba1/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type EntranceController struct {
	model *models.Entrada 
	view  *views.ViewEntrada
	mu    *sync.Mutex
}

func NewEntranceController(win *pixelgl.Window, mu *sync.Mutex) *EntranceController {
	return &EntranceController{
		model: models.NewEntrada(),
		view:  views.NewViewEntrada(win),
		mu:    mu,
	}
}



