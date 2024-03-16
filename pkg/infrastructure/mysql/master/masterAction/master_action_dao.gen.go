// Package masterAction アクション
package masterAction

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
)

type masterActionDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterActionDao(conn *database.SqlHandler) masterAction.MasterActionRepository {
	return &masterActionDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionDao) Find(ctx context.Context, id int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindOrNil(ctx context.Context, id int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindByName(ctx context.Context, name string) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindByName", fmt.Sprintf("%s_", name)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindByAnyId(ctx context.Context, anyId int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindByAnyId", fmt.Sprintf("%d_", anyId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadConn.WithContext(ctx).Where("any_id = ?", anyId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindByAnyId", fmt.Sprintf("%d_", anyId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FinOrNilByName(ctx context.Context, name string) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindOrNilByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindOrNilByName", fmt.Sprintf("%s_", name)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FinOrNilByAnyId(ctx context.Context, anyId int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindOrNilByAnyId", fmt.Sprintf("%d_", anyId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadConn.WithContext(ctx).Where("any_id = ?", anyId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindOrNilByAnyId", fmt.Sprintf("%d_", anyId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindList(ctx context.Context) (masterAction.MasterActions, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterAction.MasterActions); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActions()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterAction.NewMasterActions()
	for _, t := range ts {
		ms = append(ms, masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionDao) FindListByName(ctx context.Context, name string) (masterAction.MasterActions, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindListByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(masterAction.MasterActions); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActions()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterAction.NewMasterActions()
	for _, t := range ts {
		ms = append(ms, masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindListByName", fmt.Sprintf("%s_", name)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionDao) FindListByAnyId(ctx context.Context, anyId int64) (masterAction.MasterActions, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindListByAnyId", fmt.Sprintf("%d_", anyId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterAction.MasterActions); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActions()
	res := s.ReadConn.WithContext(ctx).Where("any_id = ?", anyId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterAction.NewMasterActions()
	for _, t := range ts {
		ms = append(ms, masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindListByAnyId", fmt.Sprintf("%d_", anyId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionDao) Create(ctx context.Context, tx *gorm.DB, m *masterAction.MasterAction) (*masterAction.MasterAction, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterAction{
		Id:              m.Id,
		Name:            m.Name,
		ActionType:      m.ActionType,
		AnyId:           m.AnyId,
		TriggerActionId: m.TriggerActionId,
		NextActionId:    m.NextActionId,
	}
	res := conn.Model(NewMasterAction()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId), nil
}

func (s *masterActionDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterAction.MasterActions) (masterAction.MasterActions, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterActions()
	for _, m := range ms {
		t := &MasterAction{
			Id:              m.Id,
			Name:            m.Name,
			ActionType:      m.ActionType,
			AnyId:           m.AnyId,
			TriggerActionId: m.TriggerActionId,
			NextActionId:    m.NextActionId,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterAction()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionDao) Update(ctx context.Context, tx *gorm.DB, m *masterAction.MasterAction) (*masterAction.MasterAction, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterAction{
		Id:              m.Id,
		Name:            m.Name,
		ActionType:      m.ActionType,
		AnyId:           m.AnyId,
		TriggerActionId: m.TriggerActionId,
		NextActionId:    m.NextActionId,
	}
	res := conn.Model(NewMasterAction()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterAction.SetMasterAction(t.Id, t.Name, t.ActionType, t.AnyId, t.TriggerActionId, t.NextActionId), nil
}

func (s *masterActionDao) Delete(ctx context.Context, tx *gorm.DB, m *masterAction.MasterAction) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterAction()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterAction())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}