package dao

import "better-shipping-app/internal/model"

type PackDao interface {
	ListPacks() ([]model.Pack, error)
}

type packDao struct {
	dbShell DbShell
}

func (p packDao) ListPacks() ([]model.Pack, error) {
	rows, err := p.dbShell.getDb().Query("SELECT id, item_count FROM pack")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	packs := make([]model.Pack, 0)
	for rows.Next() {
		pack := model.Pack{}
		err := rows.Scan(&pack.Id, &pack.ItemCount)
		if err != nil {
			return nil, err
		}
		packs = append(packs, pack)
	}

	return packs, nil
}

func NewPackDao(dbShell DbShell) PackDao {
	return &packDao{dbShell: dbShell}
}
