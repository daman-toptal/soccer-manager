package util

import (
	"errors"
	"fmt"
	grpcRoot "protobuf-v1/golang"
	grpcPlayer "protobuf-v1/golang/player"
	grpcTxn "protobuf-v1/golang/transaction"
	"regexp"
	"soccer-manager/util/config"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var amountRegexp = regexp.MustCompile(`^([+-]?)([0-9]+)\.([0-9]+)$`)
var amountRegWexp = regexp.MustCompile(`^([+-]?)([0-9]+)$`)

func pow10(n int) int64 {
	var result int64 = 10
	for i := 1; i < n; i++ {
		result *= 10
	}
	return result
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.GetInt("BCRYPT_COST"))
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ParseAmountString(amount string, precision ...int) (int64, error) {
	var intPart, fraPart string
	var s int64 = 1
	var pw int = 2
	if len(precision) > 0 {
		pw = precision[0]
	}
	match := amountRegexp.FindStringSubmatch(amount)
	if match != nil {
		if match[1] == "-" {
			s = -1
		}
		intPart = match[2]
		fraPart = AppendZero(match[3], pw)[:pw]
	}
	if match == nil {
		if match = amountRegWexp.FindStringSubmatch(amount); match != nil {
			if match[1] == "-" {
				s = -1
			}
			intPart = match[2]
			fraPart = AppendZero("", pw)[:pw]
		} else {
			return 0, errors.New("invalid format")
		}
	}

	a, err := strconv.ParseInt(intPart, 10, 64)
	if err != nil {
		return 0, errors.Unwrap(err)
	}
	b, err := strconv.ParseInt(fraPart, 10, 64)
	if err != nil {
		return 0, errors.Unwrap(err)
	}

	return (a*pow10(pw) + b) * s, nil
}

func AppendZero(amount string, precision int) string {
	result := amount
	for i := 0; i < precision-len(amount); i++ {
		result += "0"
	}
	return result
}

func ParseAmountToString(amount int64, precision ...int) string {
	var pw int = 2
	if len(precision) > 0 {
		pw = precision[0]
	}
	var format string = "%d.%0" + fmt.Sprintf("%d", pw) + "d"
	if amount < 0 {
		amount = -1 * amount
		format = "-" + format
	}
	return fmt.Sprintf(format, amount/pow10(pw), amount%pow10(pw))
}

type Currency string

const (
	CurrencyUnspecified = Currency("")
	CurrencyUSD         = Currency("USD")
)

var CurrencyFromProto = map[grpcRoot.Currency]Currency{
	grpcRoot.Currency_CURRENCY_UNSPECIFIED: CurrencyUnspecified,
	grpcRoot.Currency_CURRENCY_USD:  CurrencyUSD,
}

type TransactionType string

const (
	TransactionTypeUnspecified = TransactionType("")
	TransactionTypeBuy         = TransactionType("Buy")
	TransactionTypeSell        = TransactionType("Sell")
)

var TransactionTypeFromProto = map[grpcTxn.TransactionType]TransactionType{
	grpcTxn.TransactionType_TT_UNSPECIFIED: TransactionTypeUnspecified,
	grpcTxn.TransactionType_TT_BUY:         TransactionTypeBuy,
	grpcTxn.TransactionType_TT_SELL:        TransactionTypeSell,
}

type PlayerType string

const (
	PlayerTypeUnspecified = PlayerType("")
	PlayerTypeGoalKeeper  = PlayerType("goalKeeper")
	PlayerTypeDefender    = PlayerType("defender")
	PlayerTypeMidFielder  = PlayerType("midFielder")
	PlayerTypeAttacker    = PlayerType("attacker")
)

var PlayerTypeFromProto = map[grpcPlayer.PlayerType]PlayerType{
	grpcPlayer.PlayerType_PT_UNSPECIFIED: PlayerTypeUnspecified,
	grpcPlayer.PlayerType_PT_GOAL_KEEPER: PlayerTypeGoalKeeper,
	grpcPlayer.PlayerType_PT_DEFENDER:    PlayerTypeDefender,
	grpcPlayer.PlayerType_PT_MID_FIELDER: PlayerTypeMidFielder,
	grpcPlayer.PlayerType_PT_ATTACKER:    PlayerTypeAttacker,
}
