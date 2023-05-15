package models

type Bluetaber struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Nombre string `json:"nombre"`
	Categoria string `json:"categoria"`
}
