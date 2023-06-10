package main

import (
	"fmt"

	"github.com/sigurn/crc16"
)

const (
	InitPayload         = "000201"
	PointOfMethod       = "010212"
	MerchantAccountCode = "2937"
	AID                 = "0016A000000677010111"
	CountryCode         = "5802TH"
	CurrencyCode        = "5303764"
	CheckSumCode        = "6304"

	ThaiPhoneTypeCode    = "01130066"
	IDCardNumberTypeCode = "02"
	AmountCode           = "54"
)

func generateQrPromptPayWithAmount(promptPayID string, amount float64) string {
	promptPayAmount := fmt.Sprintf("%.2f", amount)
	resultQR := InitPayload + PointOfMethod + MerchantAccountCode + AID + ThaiPhoneTypeCode
	resultQR += promptPayID[1:] + CountryCode + CurrencyCode + AmountCode
	resultQR += "0" + fmt.Sprintf("%d", len(promptPayAmount)) + promptPayAmount + CheckSumCode
	h := crc16.New(crc16.MakeTable(crc16.CRC16_CCITT_FALSE))
	h.Write([]byte(resultQR))
	return fmt.Sprintf("%s%X", resultQR, h.Sum16())
}

func generateQrPromptPay(promptPayID string) string {
	resultQR := InitPayload + PointOfMethod + MerchantAccountCode + AID + ThaiPhoneTypeCode
	resultQR += promptPayID[1:] + CountryCode + CurrencyCode + CheckSumCode
	h := crc16.New(crc16.MakeTable(crc16.CRC16_CCITT_FALSE))
	h.Write([]byte(resultQR))
	return fmt.Sprintf("%s%X", resultQR, h.Sum16())
}

func main() {
	fmt.Println("QR string -", generateQrPromptPay("0882314328"))
	fmt.Println("QR string with amount -", generateQrPromptPayWithAmount("0882314328", 1))
}
