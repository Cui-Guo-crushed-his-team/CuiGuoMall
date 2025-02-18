package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	PaymentStatusUnpaid = 0
	PaymentStatusPaying = 1
	PaymentStatusPaid   = 2
	PaymentStatusFailed = 3
)

type PaymentRecord struct {
	ID                int64     `gorm:"primaryKey"`
	OrderID           string    `gorm:"column:order_id;uniqueIndex"`
	Status            int8      `gorm:"column:status"`
	Amount            float64   `gorm:"column:amount"`
	PaymentOutTradeNo string    `gorm:"uniqueIndex"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
}

func (p *PaymentRecord) TableName() string {
	return "payment_record"
}

// todo userid没了
// GetPaymentRecordByOrderID 根据订单号查询支付记录
func GetPaymentRecordByOrderID(db *gorm.DB, ctx context.Context, orderID string) (*PaymentRecord, error) {
	var record PaymentRecord
	err := db.WithContext(ctx).Where("order_id = ?", orderID).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// todo userid没了
// GetPaymentRecordByPaymentOutTradeNo 根据支付单号查询支付记录
func GetPaymentRecordByPaymentOutTradeNo(ctx context.Context, db *gorm.DB, paymentOutTradeNo string) (*PaymentRecord, error) {
	var record PaymentRecord
	err := db.WithContext(ctx).Where("payment_out_trade_no = ?", paymentOutTradeNo).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// CreatePaymentRecord 创建支付记录
func CreatePaymentRecord(db *gorm.DB, ctx context.Context, record *PaymentRecord) error {
	return db.WithContext(ctx).Create(record).Error
}

// UpdatePaymentStatus 更新支付状态
func UpdatePaymentStatus(db *gorm.DB, ctx context.Context, orderID string, status int8) error {
	return db.WithContext(ctx).Model(&PaymentRecord{}).
		Where("order_id = ?", orderID).
		Update("status", status).Error
}
