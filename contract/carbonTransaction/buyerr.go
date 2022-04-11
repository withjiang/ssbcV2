package main

import (
	"errors"
	"github.com/ssbcV2/contract"
)

type buyer struct {
	name                  string
	registerAccountStatus bool              //登记账户状态
	carbonBalance         int               //配额库存
	CapitalBalance        int               //资金账户余额
	transactionRecord     map[string]string //交易记录 交易编号 交易对方

}

var entityB buyer

var CapitalBalance int //资金账户余额
var veriMarkB bool     // 有效买方标记

func init() {
	entityB = buyer{
		name:                  "",
		registerAccountStatus: true,
		carbonBalance:         0,
		CapitalBalance:        1000,
		transactionid:         "",
	}

	veriMark = false
}
func setBuyerValue(args map[string]string) (interface{}, error) {
	entityB.name = contract.Caller()
	entityB.registerAccountStatus = true
	money = contract.Value()

	return nil, nil
}

func BuyerPrimary(args map[string]string) (interface{}, error) {
	entityB.name = contract.Caller()
	if entityB.registerAccountStatus != true {
		contract.Transfer(contract.Caller(), contract.Value()) // 退回转账
		log.Info("登记账户存在异常，退出")
		return nil, errors.New("登记账户存在异常，退出")
	}

	if entityB.CapitalBalance == 0 {
		contract.Transfer(contract.Caller(), contract.Value()) // 退回转账
		log.Info("资金不足")
		return nil, errors.New("资金不足")
	}
	value = contract.Value

	if value > entityB.CapitalBalance {
		return nil, errors.New("超过可用库存")
	}

	var trans transaction

	trans.carbonAmount = carbonAmount
	trans.Value = carbonAmount * perValue

	trans.statusofBuyer = true //生成一笔交易
	if trans.statusofSeller == true {
		trans.status = true //生成一笔交易
		trans.carbonAmount = carbonAmount
		trans.Value = carbonAmount * perValue

	}
	contract.Transfer(entityB.name, trans.Value)
	return nil, nil
}
