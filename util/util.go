package util

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb"
	"github.com/shopspring/decimal"
)

func ToDate(dateString string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateString)
	return parsedDate
}

func RemoveNewLine(str string) string {
	return strings.Trim(str, "\n")
}

func GenerateRandomString() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func IsTopUpAmountValid(amount decimal.Decimal) bool {
	minAmount := decimal.NewFromInt(50000)
	maxAmount := decimal.NewFromInt(10000000)

	if amount.LessThan(minAmount) || amount.GreaterThan(maxAmount) {
		return false
	}
	return true
}

func IsTransferAmountValid(amount decimal.Decimal) bool {
	minAmount := decimal.NewFromInt(1000)
	maxAmount := decimal.NewFromInt(50000000)

	if amount.LessThan(minAmount) || amount.GreaterThan(maxAmount) {
		return false
	}
	return true
}

func ToTransactionTypeEnum(input model.TransactionTypes) (res pb.TransactionRes_TRANSACTION_TYPES) {
	if input == model.Transfer {
		return pb.TransactionRes_TRANSFER
	} else if input == model.TopUp {
		return pb.TransactionRes_TOP_UP
	} else if input == model.GameReward {
		return pb.TransactionRes_GAME_REWARD
	}
	return
}

func ToSourceOfFundEnum(input model.SourceOfFunds) (res pb.TransactionRes_SOURCE_OF_FUNDS) {
	if input == model.BankTransfer {
		return pb.TransactionRes_BANK_TRANSFER
	} else if input == model.CreditCard {
		return pb.TransactionRes_CREDIT_CARD
	} else if input == model.Cash {
		return pb.TransactionRes_CASH
	} else if input == model.Reward {
		return pb.TransactionRes_REWARD
	}
	return
}
