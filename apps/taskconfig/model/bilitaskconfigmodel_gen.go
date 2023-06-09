// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	biliTaskConfigFieldNames          = builder.RawFieldNames(&BiliTaskConfig{})
	biliTaskConfigRows                = strings.Join(biliTaskConfigFieldNames, ",")
	biliTaskConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(biliTaskConfigFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	biliTaskConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(biliTaskConfigFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	biliTaskConfigModel interface {
		Insert(ctx context.Context, data *BiliTaskConfig) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BiliTaskConfig, error)
		Update(ctx context.Context, data *BiliTaskConfig) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBiliTaskConfigModel struct {
		conn  sqlx.SqlConn
		table string
	}

	BiliTaskConfig struct {
		Id                 int64     `db:"id"`
		Dedeuserid         string    `db:"dedeuserid"`
		Sessdata           string    `db:"sessdata"`
		BiliJct            string    `db:"bili_jct"`
		DonateCoins        int64     `db:"donate_coins"`
		ReserveCoins       int64     `db:"reserve_coins"`
		AutoCharge         bool      `db:"auto_charge"`
		DonateGift         bool      `db:"donate_gift"`
		DonateGiftTarget   string    `db:"donate_gift_target"`
		AutoChargeTarget   string    `db:"auto_charge_target"`
		DevicePlatform     string    `db:"device_platform"`
		DonateCoinStrategy int64     `db:"donate_coin_strategy"`
		UserAgent          string    `db:"user_agent"`
		SkipTask           bool      `db:"skip_task"`
		FollowDeveloper    bool      `db:"follow_developer"`
		CreateTime         time.Time `db:"create_time"` // 数据创建时间[禁止在代码中赋值]
		UpdateTime         time.Time `db:"update_time"` // 数据更新时间[禁止在代码中赋值]
	}
)

func newBiliTaskConfigModel(conn sqlx.SqlConn) *defaultBiliTaskConfigModel {
	return &defaultBiliTaskConfigModel{
		conn:  conn,
		table: "`bili_task_config`",
	}
}

func (m *defaultBiliTaskConfigModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultBiliTaskConfigModel) FindOne(ctx context.Context, id int64) (*BiliTaskConfig, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", biliTaskConfigRows, m.table)
	var resp BiliTaskConfig
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBiliTaskConfigModel) Insert(ctx context.Context, data *BiliTaskConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, biliTaskConfigRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Dedeuserid, data.Sessdata, data.BiliJct, data.DonateCoins, data.ReserveCoins, data.AutoCharge, data.DonateGift, data.DonateGiftTarget, data.AutoChargeTarget, data.DevicePlatform, data.DonateCoinStrategy, data.UserAgent, data.SkipTask, data.FollowDeveloper)
	return ret, err
}

func (m *defaultBiliTaskConfigModel) Update(ctx context.Context, data *BiliTaskConfig) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, biliTaskConfigRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Dedeuserid, data.Sessdata, data.BiliJct, data.DonateCoins, data.ReserveCoins, data.AutoCharge, data.DonateGift, data.DonateGiftTarget, data.AutoChargeTarget, data.DevicePlatform, data.DonateCoinStrategy, data.UserAgent, data.SkipTask, data.FollowDeveloper, data.Id)
	return err
}

func (m *defaultBiliTaskConfigModel) tableName() string {
	return m.table
}
