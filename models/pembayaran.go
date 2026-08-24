package models

import "time"

type Pembayaran struct {
	ID          string    `json:"id" bson:"_id"`
	TransaksiID string    `json:"transaksi_id" bson:"transaksi_id"`
	KasirID     string    `json:"kasir_id" bson:"kasir_id"`
	Metode      string    `json:"metode" bson:"metode"`
	TotalBayar  float64   `json:"total_bayar" bson:"total_bayar"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
