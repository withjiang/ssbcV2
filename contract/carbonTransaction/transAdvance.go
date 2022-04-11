package main

import (
	"errors"
	"github.com/ssbcV2/contract"
	"time"
)

var trans transaction
var value int
var caller string    //合约发起方
var counterud string //交易对方
var capitalFlag bool //资金账户划转
var carbonFlag bool  //配额账户划转

func init() {
	value = contract.Value()
	caller = contract.Caller()
	trans = transaction{
		transactionid: "",
		category:      "standardCarbonCredit",
		Value:         value,
		progress:      "ing",
		status:        true,
	}
	capitalFlag = false
	carbonFlag = false
}
func proceed(args map[string]string) (interface{}, error) {
	if trans.status != true {
		contract.Transfer(contract.Caller(), contract.Value()) // 退回转账
		log.Info("交易资质存在问题")
		return nil, errors.New("交易资质存在问题")
	}
	if trans.progress == "ed" {
		contract.Transfer(contract.Caller(), contract.Value()) // 退回转账
		log.Info("交易已结束")
		return nil, errors.New("交易已结束")
	}

	if caller == trans.sellerid { //发起方为卖方
		var callerr seller
		var counter buyer

		callerr.carbonBalance -= value
		callerr.CapitalBalance += value * perValue
		counter.carbonBalance += value
		counter.CapitalBalance -= value * perValue

	}
	if caller == trans.buyeridrid { //发起方为买方
		var callerr buyer
		var counter seller
		var carbon int
		carbon = value / perValue
		callerr.carbonBalance += carbon
		callerr.CapitalBalance -= value
		counter.carbonBalance -= carbon
		counter.CapitalBalance += value

	}
	capitalFlag = true
	carbonFlag = true
	trans.progress = "ed"
	trans.EndTime = time.Now() //交易结束时间
	return nil, nil
}
