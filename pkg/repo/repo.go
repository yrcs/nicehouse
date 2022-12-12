package repo

import (
	"context"
	"time"
)

import (
	"github.com/jinzhu/copier"

	"gorm.io/gorm"

	"gorm.io/plugin/optimisticlock"
)

import (
	"github.com/yrcs/nicehouse/pkg/util"
)

type BasePO struct {
	Id        string                 `gorm:"type:varchar(26);primaryKey;comment:分布式全局唯一 ULID"`
	CreatedAt time.Time              `gorm:"type:datetime not null;comment:创建时间"`
	UpdatedAt time.Time              `gorm:"type:datetime not null;comment:更新时间"`
	DeletedAt gorm.DeletedAt         `gorm:"type:datetime;index;comment:删除时间"`
	Version   optimisticlock.Version `gorm:"not null;default:0;comment:版本号（乐观锁专用）"`
}

type Repo[E, T any] interface {
	Create(ctx context.Context, value any) error
	FindOne(ctx context.Context, conds ...any) (E, error)
	Find(ctx context.Context, conds ...any) ([]E, error)
	FindByPage(ctx context.Context, offset int, limit int, conds map[string]any, orderBy map[string]string) ([]E, int, error)
	Update(ctx context.Context, column string, value any, conds ...any) (E, error)
	Updates(ctx context.Context, values map[string]any, query any, conds ...any) (E, error)
	Delete(ctx context.Context, ids []string, query any, conds ...any) error
}

type BaseRepo[E, T any] struct {
	DB *gorm.DB
}

var _ Repo[*any, *any] = (*BaseRepo[*any, *any])(nil)

func (r *BaseRepo[E, T]) Create(ctx context.Context, value any) error {
	return r.DB.WithContext(ctx).Create(value).Error
}

func (r *BaseRepo[E, T]) FindOne(ctx context.Context, conds ...any) (E, error) {
	var (
		po T
		do E
	)

	if err := r.DB.WithContext(ctx).First(&po, conds...).Error; err != nil {
		return do, err
	}

	do = util.InstantiateStruct(do)
	copier.Copy(do, po)
	return do, nil
}

func (r *BaseRepo[E, T]) Find(ctx context.Context, conds ...any) ([]E, error) {
	var pos []T
	if err := r.DB.WithContext(ctx).Find(&pos, conds...).Error; err != nil {
		return nil, err
	}

	var dos []E
	copier.Copy(&dos, pos)
	return dos, nil
}

func (r *BaseRepo[E, T]) FindByPage(ctx context.Context, offset int, limit int, conds map[string]any, orderBy map[string]string) ([]E, int, error) {
	var pos []T
	tx := r.DB.WithContext(ctx).Where(conds)
	total := int(tx.Find(&pos).RowsAffected)

	for k, v := range orderBy {
		tx = tx.Order(k + " " + v)
	}
	err := tx.Offset(offset).Limit(limit).Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}

	var dos []E
	copier.Copy(&dos, pos)
	return dos, total, nil
}

func (r *BaseRepo[E, T]) Update(ctx context.Context, column string, value any, conds ...any) (E, error) {
	var (
		po T
		do E
	)

	tx := r.DB.WithContext(ctx)
	if err := tx.Take(&po, conds...).Error; err != nil {
		return do, err
	}

	if err := tx.Model(&po).Update(column, value).Error; err != nil {
		return do, err
	}

	do = util.InstantiateStruct(do)
	copier.Copy(do, po)
	return do, nil
}

func (r *BaseRepo[E, T]) Updates(ctx context.Context, values map[string]any, query any, conds ...any) (E, error) {
	var (
		po T
		do E
	)
	po = util.InstantiateStruct(po)
	if err := r.DB.WithContext(ctx).Model(&po).Where(query, conds...).Updates(values).Error; err != nil {
		return do, err
	}

	do = util.InstantiateStruct(do)
	copier.Copy(do, po)
	return do, nil
}

func (r *BaseRepo[E, T]) Delete(ctx context.Context, ids []string, query any, conds ...any) error {
	var po T
	po = util.InstantiateStruct(po)
	return r.DB.WithContext(ctx).Where(query, conds...).Delete(po, ids).Error
}
