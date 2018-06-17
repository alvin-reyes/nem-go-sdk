package model

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

type TransferTransactionV1 struct {
	TimeStamp string `json:"timeStamp"`
	Amount    int    `json:"amount"`
	Signature string `json:"signature"`
	Fee       string `json:"fee"`
	Recipient string `json:"recipient"`
	Type      int    `json:"type"`
	Deadline  string `json:"deadline"`
	Message   struct {
		Payload string `json:"payload"`
		Type    int    `json:"type"`
	} `json:"message"`
	Version string `json:"version"`
	Signer  string `json:"signer"`
}

type RequestPrepareAnnounceTransferTransactionV1 struct {
	TransferTransactionV1 TransferTransactionV1 `json:"transaction"`
	PrivateKey            string                `json:"privateKey"`
}

//	Support for V2.
type TransferTransactionV2 struct {
	TimeStamp string `json:"timeStamp"`
	Amount    int    `json:"amount"`
	Signature string `json:"signature"`
	Fee       string `json:"fee"`
	Recipient string `json:"recipient"`
	Type      int    `json:"type"`
	Deadline  string `json:"deadline"`
	Message   struct {
		Payload string `json:"payload"`
		Type    int    `json:"type"`
	} `json:"message"`
	Version string             `json:"version"`
	Signer  string             `json:"signer"`
	Mosaics []MosaicIdQuantity `json:"mosaics"`
}

type MosaicIdQuantity struct {
	MosaicId MosaicId `json:"mosaicId"`
	Quantity int            `json:"quantity"`
}

type RequestPrepareAnnounceTransferTransactionV2 struct {
	TransferTransactionV2 TransferTransactionV2 `json:"transaction"`
	PrivateKey            string                `json:"privateKey"`
}

func NewRequestPrepareAnnounceTransferTransactionV2(txn TransferTransactionV2, privateKey string) *RequestPrepareAnnounceTransferTransactionV2 {
	return &RequestPrepareAnnounceTransferTransactionV2{txn, privateKey}
}

func (s RequestPrepareAnnounceTransferTransactionV1) SerializeV1() ([]byte, []byte) {

	//	Common Transaction
	transactionType := make([]byte, 4)
	version := make([]byte, 4)
	timestamp := make([]byte, 4)
	lengthOfPublicKeyByte := make([]byte, 4)
	publicKeySigner := make([]byte, 32)
	fee := make([]byte, 4)
	deadline := make([]byte, 4)

	//	Start of Common
	//	Transaction type
	binary.LittleEndian.PutUint32(transactionType, 0x0101)
	//	Version
	versionByteArr, _ := strconv.Atoi(s.TransferTransactionV1.Version)
	binary.LittleEndian.PutUint32(version, uint32(versionByteArr))

	//	Timestamp
	binary.LittleEndian.PutUint32(timestamp, uint32(time.Now().Unix()))

	//	Length of Public Key Byte
	binary.LittleEndian.PutUint32(lengthOfPublicKeyByte, 0x0020)

	//	PublicKey Signer
	psig, _ := strconv.Atoi(s.TransferTransactionV1.Recipient)
	binary.LittleEndian.PutUint32(publicKeySigner, uint32(psig))

	//	Fee
	binary.LittleEndian.PutUint32(fee, 12)

	//	Deadline
	binary.LittleEndian.PutUint32(deadline, uint32(time.Now().Unix()))

	commonTxnByte := append(transactionType, version...)
	commonTxnByte = append(commonTxnByte, timestamp...)
	commonTxnByte = append(commonTxnByte, lengthOfPublicKeyByte...)
	commonTxnByte = append(commonTxnByte, publicKeySigner...)
	commonTxnByte = append(commonTxnByte, fee...)
	commonTxnByte = append(commonTxnByte, deadline...)
	//	End of Common

	//[]byte{0x01, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x98}
	fmt.Println(commonTxnByte)
	fmt.Println(">>>>>> tt")

	// specific to tt v1
	lengthRecAddress := make([]byte, 4)
	recAddress := make([]byte, 40)
	amount := make([]byte, 8)
	lenghOfMessage := make([]byte, 4)
	messageType := make([]byte, 4)
	lengthOfPayload := make([]byte, 4)

	binary.LittleEndian.PutUint32(lengthRecAddress, 0x0101)
	binary.LittleEndian.PutUint32(recAddress, 0x98)
	binary.LittleEndian.PutUint32(amount, uint32(time.Now().Unix()))
	binary.LittleEndian.PutUint32(lenghOfMessage, 0x0020)
	binary.LittleEndian.PutUint32(messageType, uint32(psig))
	binary.LittleEndian.PutUint32(lengthOfPayload, 12)
	payload := []byte(s.TransferTransactionV1.Message.Payload)

	tTxnByte := append(lengthRecAddress, recAddress...)
	tTxnByte = append(tTxnByte, amount...)
	tTxnByte = append(tTxnByte, lenghOfMessage...)
	tTxnByte = append(tTxnByte, messageType...)
	tTxnByte = append(tTxnByte, lengthOfPayload...)
	tTxnByte = append(tTxnByte, payload...)
	fmt.Println(tTxnByte)

	//dataByte := append(commonTxnByte, tTxnByte...)

	return commonTxnByte,tTxnByte
}

func (s RequestPrepareAnnounceTransferTransactionV2) SerializeV2() ([]byte, []byte) {

	//	Common Transaction
	transactionType := make([]byte, 4)
	version := make([]byte, 4)
	timestamp := make([]byte, 4)
	lengthOfPublicKeyByte := make([]byte, 4)
	publicKeySigner := make([]byte, 32)
	fee := make([]byte, 4)
	deadline := make([]byte, 4)

	//	Start of Common
	//	Transaction type
	binary.LittleEndian.PutUint32(transactionType, 0x0101)
	//	Version
	versionByteArr, _ := strconv.Atoi(s.TransferTransactionV2.Version)
	binary.LittleEndian.PutUint32(version, uint32(versionByteArr))

	//	Timestamp
	binary.LittleEndian.PutUint32(timestamp, uint32(time.Now().Unix()))

	//	Length of Public Key Byte
	binary.LittleEndian.PutUint32(lengthOfPublicKeyByte, 0x0020)

	//	PublicKey Signer
	psig, _ := strconv.Atoi(s.TransferTransactionV2.Recipient)
	binary.LittleEndian.PutUint32(publicKeySigner, uint32(psig))

	//	Fee
	binary.LittleEndian.PutUint32(fee, 12)

	//	Deadline
	binary.LittleEndian.PutUint32(deadline, uint32(time.Now().Unix()))

	commonTxnByte := append(transactionType, version...)
	commonTxnByte = append(commonTxnByte, timestamp...)
	commonTxnByte = append(commonTxnByte, lengthOfPublicKeyByte...)
	commonTxnByte = append(commonTxnByte, publicKeySigner...)
	commonTxnByte = append(commonTxnByte, fee...)
	commonTxnByte = append(commonTxnByte, deadline...)
	//	End of Common

	//[]byte{0x01, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x98}
	fmt.Println(commonTxnByte)
	fmt.Println(">>>>>> tt")

	// specific to tt v2
	lengthRecAddress := make([]byte, 4)
	recAddress := make([]byte, 40)
	amount := make([]byte, 8)
	lenghOfMessage := make([]byte, 4)
	messageType := make([]byte, 4)
	lengthOfPayload := make([]byte, 4)

	lengthOfMosaicStructure := make([]byte, 4)
	lengthOfNamespaceID := make([]byte, 4)
	namespaceID := make([]byte, 4)
	lengthOfMosaicName := make([]byte, 4)
	mosaicName := make([]byte, 4)
	quantity := make([]byte, 4)

	binary.LittleEndian.PutUint32(lengthRecAddress, 0x0101)
	binary.LittleEndian.PutUint32(recAddress, 0x98)
	binary.LittleEndian.PutUint32(amount, uint32(time.Now().Unix()))
	binary.LittleEndian.PutUint32(lenghOfMessage, 0x0020)
	binary.LittleEndian.PutUint32(messageType, uint32(psig))
	binary.LittleEndian.PutUint32(lengthOfPayload, 12)
	payload := []byte(s.TransferTransactionV2.Message.Payload)

	binary.LittleEndian.PutUint32(lengthOfMosaicStructure, 4)
	binary.LittleEndian.PutUint32(lengthOfNamespaceID, 4)
	binary.LittleEndian.PutUint32(namespaceID, 4)
	binary.LittleEndian.PutUint32(lengthOfMosaicName, 4)
	binary.LittleEndian.PutUint32(mosaicName, 4)
	binary.LittleEndian.PutUint32(quantity, 8)

	tTxnByte := append(lengthRecAddress, recAddress...)
	tTxnByte = append(tTxnByte, amount...)
	tTxnByte = append(tTxnByte, lenghOfMessage...)
	tTxnByte = append(tTxnByte, messageType...)
	tTxnByte = append(tTxnByte, lengthOfPayload...)
	tTxnByte = append(tTxnByte, payload...)
	fmt.Println(tTxnByte)

	//dataByte := append(commonTxnByte, tTxnByte...)

	return commonTxnByte,tTxnByte
}
