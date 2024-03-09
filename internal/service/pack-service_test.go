package service

import (
	"better-shipping-app/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type packDaoMock struct {
	mock.Mock
}

func (m *packDaoMock) ListPacks() ([]model.Pack, error) {
	args := m.Called()
	return args.Get(0).([]model.Pack), args.Error(1)
}

func TestPackListSuccess(t *testing.T) {
	packDao := &packDaoMock{}

	packs := []model.Pack{
		{ItemCount: 250},
		{ItemCount: 500},
		{ItemCount: 1000},
		{ItemCount: 2000},
		{ItemCount: 5000},
	}

	packDao.On("ListPacks").Return(packs, nil)

	service := NewPackService(packDao)

	result, err := service.ListPacks()

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, packs, result)
}

func TestPackListFailOnDaoError(t *testing.T) {
	packDao := &packDaoMock{}

	packDao.On("ListPacks").Return([]model.Pack{}, assert.AnError)

	service := NewPackService(packDao)

	_, err := service.ListPacks()

	if err == nil {
		t.Error("Expected error")
	}
}
