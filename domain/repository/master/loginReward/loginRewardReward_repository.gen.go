//go:generate mockgen -source=./loginRewardReward_repository.gen.go -destination=./loginRewardReward_repository_mock.gen.go -package=loginReward
package loginReward

import (
	"github.com/game-core/gocrafter/domain/entity/master/loginReward"
	"github.com/jinzhu/gorm"
)

type LoginRewardRewardRepository interface {
	Create(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error)

	Delete(entity *loginReward.LoginRewardReward, tx *gorm.DB) error

	FindByID(ID int64) (*loginReward.LoginRewardReward, error)

	FindByItemName(ItemName string) (*loginReward.LoginRewardReward, error)

	FindByLoginRewardModelName(LoginRewardModelName string) (*loginReward.LoginRewardReward, error)

	FindByLoginRewardModelNameAndItemName(LoginRewardModelName string, ItemName string) (*loginReward.LoginRewardReward, error)

	FindByName(Name string) (*loginReward.LoginRewardReward, error)

	FindOrNilByID(ID int64) (*loginReward.LoginRewardReward, error)

	FindOrNilByItemName(ItemName string) (*loginReward.LoginRewardReward, error)

	FindOrNilByLoginRewardModelName(LoginRewardModelName string) (*loginReward.LoginRewardReward, error)

	FindOrNilByLoginRewardModelNameAndItemName(LoginRewardModelName string, ItemName string) (*loginReward.LoginRewardReward, error)

	FindOrNilByName(Name string) (*loginReward.LoginRewardReward, error)

	List(limit int64) (*loginReward.LoginRewardRewards, error)

	ListByItemName(ItemName string) (*loginReward.LoginRewardRewards, error)

	ListByLoginRewardModelName(LoginRewardModelName string) (*loginReward.LoginRewardRewards, error)

	ListByLoginRewardModelNameAndItemName(LoginRewardModelName string, ItemName string) (*loginReward.LoginRewardRewards, error)

	ListByName(Name string) (*loginReward.LoginRewardRewards, error)

	Update(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error)
}