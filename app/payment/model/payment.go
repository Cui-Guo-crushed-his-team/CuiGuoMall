package model

import (
	"context"
	"database/sql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/biz/dal/mysql"
	"gorm.io/gorm/clause"
)

type PaymentStatus int64

const (
	PaymentStatusUnknown = iota
	PaymentStatusInit
	PaymentStatusSuccess
	PaymentStatusFailed
	PaymentStatusRefund
)

type Payment struct {
	Model
	Amount      string
	Description string `gorm:"description"`

	OutTradeNO string `gorm:"type:varchar(256);unique"`

	TradeNO sql.NullString `gorm:"type:varchar(128);unique"`

	Status PaymentStatus
}

func Create(ctx context.Context, desc string, outTradeNO, amount string) error {
	p := Payment{
		Description: desc,
		OutTradeNO:  outTradeNO,
		Amount:      amount,
		Status:      PaymentStatusInit,
	}
	return mysql.DB.WithContext(ctx).Create(&p).Error
}

// CreateOnce 只会创建一次
func CreateOnce(ctx context.Context, desc string, outTradeNO, amount string) error {
	p := Payment{
		Description: desc,
		OutTradeNO:  outTradeNO,
		Amount:      amount,
		Status:      PaymentStatusInit,
	}
	return mysql.DB.Clauses(clause.OnConflict{
		DoNothing: true,
	}).WithContext(ctx).Create(&p).Error
}
func Finish(ctx context.Context, outTradeNO, tradeNO string) error {
	result := mysql.DB.WithContext(ctx).
		Where("out_trade_no = ?", outTradeNO).
		Updates(map[string]any{
			"trade_no": tradeNO,
			"status":   PaymentStatusSuccess,
		})
	return result.Error
}

func GetByOutTradeNO(ctx context.Context, outTradeNO string) (payment *Payment, err error) {
	payment = &Payment{}
	err = mysql.DB.WithContext(ctx).Where("out_trade_no = ?", outTradeNO).First(payment).Error
	if err != nil {
		return nil, err
	}
	return payment, nil
}
