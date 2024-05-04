package dao

import (
	"github.com/justary/gin-blog/internal/model"
	"github.com/justary/gin-blog/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	return tag.Count(d.engine)
}
func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}
func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Common: &model.Common{
			CreatedBy: createBy,
		},
	}
	return tag.Create(d.engine)
}
func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Common: &model.Common{
			ID:         id,
			ModifiedBy: modifiedBy,
		},
	}
	return tag.Update(d.engine)
}
func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Common: &model.Common{ID: id}}
	return tag.Delete(d.engine)
}
