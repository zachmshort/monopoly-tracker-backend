package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferStatus string

const (
	TransferPending   TransferStatus = "pending"
	TransferCompleted TransferStatus = "completed"
	TransferRejected  TransferStatus = "rejected"
)

type TransferReason string

const (
	TransferReasonRent     TransferReason = "rent"
	TransferReasonTax      TransferReason = "tax"
	TransferReasonChance   TransferReason = "chance"
	TransferReasonProperty TransferReason = "property"
	TransferReasonCustom   TransferReason = "custom"
)

type Transfer struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	RoomID       primitive.ObjectID `bson:"roomId" json:"roomId"`
	FromPlayerID primitive.ObjectID `bson:"fromPlayerId,omitempty" json:"fromPlayerId,omitempty"` // Optional for bank transfers
	ToPlayerID   primitive.ObjectID `bson:"toPlayerId,omitempty" json:"toPlayerId,omitempty"`     // Optional for bank transfers
	Amount       int                `bson:"amount" json:"amount" validate:"required,gt=0"`
	Reason       TransferReason     `bson:"reason" json:"reason" validate:"required"`
	TimeStamp    time.Time          `bson:"timestamp" json:"timestamp"`
	Status       TransferStatus     `bson:"status" json:"status"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
