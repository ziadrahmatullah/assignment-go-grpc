package util

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"

	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/pb/emergency-funds"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
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

func ToTransactionTypeEnum(input model.TransactionTypes) (res string) {
	if input == model.Transfer {
		return string(model.Transfer)
	} else if input == model.TopUp {
		return string(model.TopUp)
	} else if input == model.GameReward {
		return string(model.GameReward)
	}
	return
}

func ToSourceOfFundEnum(input model.SourceOfFunds) (res string) {
	if input == model.BankTransfer {
		return string(model.BankTransfer)
	} else if input == model.CreditCard {
		return string(model.CreditCard)
	} else if input == model.Cash {
		return string(model.Cash)
	} else if input == model.Reward {
		return string(model.Reward)
	}
	return
}

func ToMaritalStatusEnum(input constant.MaritalStatuses) (res pb.MaritalStatus) {
	if input == constant.MARRIED {
		return pb.MaritalStatus_MARITAL_STATUS_MARRIED
	} else if input == constant.SINGLE {
		return pb.MaritalStatus_MARITAL_STATUS_SINGLE
	}
	return
}

func IsValidNumChildren(noc uint, input constant.MaritalStatuses) bool {
	if input == constant.SINGLE {
		if noc == 0 {
			return true
		} else {
			return false
		}
	}
	return true
}

func IsValidDecimal(input decimal.Decimal) bool {
	return !input.LessThan(decimal.Zero) 
}
