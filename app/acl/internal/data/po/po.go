package po

import (
	"github.com/yrcs/nicehouse/pkg/repo"
)

type Role struct {
	repo.Base
	Name        string `gorm:"not null;comment:名称"`
	Description string `gorm:"comment:描述"`
	IsSystem    bool   `gorm:"type:tinyint(1) unsigned not null;default:0;comment:是否内置"`
}
