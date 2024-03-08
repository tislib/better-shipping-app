package service

import (
	"better-shipping-app/internal/dao"
	"better-shipping-app/internal/model"
)

type PackService interface {
	ListPacks() ([]model.Pack, error)
}

type packService struct {
	packDao dao.PackDao
}

func (p packService) ListPacks() ([]model.Pack, error) {
	return p.packDao.ListPacks()
}

func NewPackService(packDao dao.PackDao) PackService {
	return &packService{packDao: packDao}
}
