package models

import "gorm.io/gorm"

type Bluetaber struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Categoria string `json:"categoria"`
}
