package service

import dto "recipe-calculator/data/dto"

type IMaterialService interface {
	CreateMaterial(material dto.IMaterial) bool
	GetMaterial(name string) dto.IMaterial
	DeleteMaterial(name string) bool
	UpdateMaterial(material dto.IMaterial) bool
}
