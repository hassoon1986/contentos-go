package test

import (
	"fmt"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/prototype"
	"time"
)

func bpRegistrationOP(name string, pubKey *prototype.PublicKeyType) *prototype.BpRegisterOperation {
	bpRegisterOp := &prototype.BpRegisterOperation{
		Owner:           &prototype.AccountName{Value: name},
		Url:             "",
		Desc:            "",
		BlockSigningKey: pubKey,
		Props: &prototype.ChainProperties{
			AccountCreationFee: prototype.NewCoin(1),
			StaminaFree:        constants.DefaultStaminaFree,
			TpsExpected:        constants.DefaultTPSExpected,
			EpochDuration:      constants.InitEpochDuration,
			TopNAcquireFreeToken: constants.InitTopN,
			PerTicketPrice:     prototype.NewCoin(1000000),
			PerTicketWeight:    constants.PerTicketWeight,
		},
	}

	return bpRegisterOp
}

func tx(css iservices.IConsensus, op interface{}, privKey *prototype.PrivateKeyType) *prototype.SignedTransaction {
	head := css.GetHeadBlockId()
	chainID := prototype.ChainId{ Value:common.GetChainIdByName("main") }
	refBlockPrefix := common.TaposRefBlockPrefix(head.Data[:])
	// occupant implement
	refBlockNum := common.TaposRefBlockNum(head.BlockNum())
	tx := &prototype.Transaction{
		RefBlockNum: refBlockNum,
		RefBlockPrefix: refBlockPrefix,
		Expiration: &prototype.TimePointSec{UtcSeconds: uint32(time.Now().Unix()) + 30},
	}
	tx.AddOperation(op)
	signTx := &prototype.SignedTransaction{Trx: tx}

	res := signTx.Sign(privKey, chainID)
	signTx.Signature = &prototype.SignatureType{Sig: res}

	if err := signTx.Validate(); err != nil {
		fmt.Println("tx validate ", err)
		return nil
	}

	return signTx
}

func RegisterBP(name, sk string, css iservices.IConsensus) error {
	privKey, err := prototype.PrivateKeyFromWIF(sk)
	if err != nil {
		fmt.Println("registerBP get priv key ", err)
		return nil
	}
	pubKey, err := privKey.PubKey()
	if err != nil {
		fmt.Println("registerBP get pub key ", err)
		return nil
	}

	op := bpRegistrationOP(name, pubKey)
	tx := tx(css, op, privKey)
	err = css.PushTransactionToPending(tx)
	if err != nil {
		return err
	}
	return nil
}

func createAccountOP(name, creator string, pubKey *prototype.PublicKeyType) *prototype.AccountCreateOperation {
	acop := &prototype.AccountCreateOperation{
		Fee:            prototype.NewCoin(constants.MinBpRegisterVest),
		Creator:        &prototype.AccountName{Value: creator},
		NewAccountName: &prototype.AccountName{Value: name},
		PubKey:          pubKey,
	}
	return acop
}

func CreateAcc(accName, accKey, creatorKey string, css iservices.IConsensus) error {
	accPrivKey, err := prototype.PrivateKeyFromWIF(accKey)
	if err != nil {
		fmt.Println("createAcc get acc priv key ", err)
		return nil
	}
	accPubKey, err := accPrivKey.PubKey()
	if err != nil {
		fmt.Println("createAcc get acc pub key ", err)
		return nil
	}

	creatorPrivKey, err := prototype.PrivateKeyFromWIF(creatorKey)
	if err != nil {
		fmt.Println("createAcc get creator priv key ", err)
		return nil
	}

	op := createAccountOP(accName, "initminer", accPubKey)
	tx := tx(css, op, creatorPrivKey)
	err = css.PushTransactionToPending(tx)
	if err != nil {
		return err
	}
	return nil
}

func bpDisableOP(name string) *prototype.BpEnableOperation {
	bpDisableOp := &prototype.BpEnableOperation{
		Owner:           &prototype.AccountName{Value: name},
		Cancel:          true,
	}

	return bpDisableOp
}

func DisableBP(name, sk string, css iservices.IConsensus) error {
	privKey, err := prototype.PrivateKeyFromWIF(sk)
	if err != nil {
		fmt.Println("disableBP get priv key ", err)
		return nil
	}

	op := bpDisableOP(name)
	tx := tx(css, op, privKey)
	err = css.PushTransactionToPending(tx)
	if err != nil {
		return err
	}
	return nil
}

func bpEnableOP(name string) *prototype.BpEnableOperation {
	bpEnableOp := &prototype.BpEnableOperation{
		Owner:           &prototype.AccountName{Value: name},
	}

	return bpEnableOp
}

func EnableBP(name, sk string, css iservices.IConsensus) error {
	privKey, err := prototype.PrivateKeyFromWIF(sk)
	if err != nil {
		fmt.Println("enableBP get priv key ", err)
		return nil
	}

	op := bpEnableOP(name)
	tx := tx(css, op, privKey)
	err = css.PushTransactionToPending(tx)
	if err != nil {
		return err
	}
	return nil
}