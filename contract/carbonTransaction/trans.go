package main

import (
	"errors"
	"fmt"
	"github.com/ssbcV2/contract"
	"time"
)

var trans transaction
var sideofTransaction string //交易发起方身份
var counterid string         //交易对手身份id
var transactionEnd time.Time // 交易发起结束时间
func init() {
	trans = transaction{
		transactionid: "",
		category:      "standardCarbonCredit",
		Value:         0,
		progress:      "ing",
	}
	transactionEnd = time.Now().Add(time.Hour * 24) //申报被交易系统接受后即刻生效，并在当日交易时间内有效
}
func setSide(args map[string]string) (interface{}, error) {
	fmt.Println("Please enter your counterid") //提示输入交易方
	fmt.Scan(&counterid)
	if trans.buyerid == contract.Caller {
		sideofTransaction = "buyer"
		trans.sellerid = counterid
	} else {
		sideofTransaction = "seller"
		trans.buyerid = counterid
	}

}
func generateTrans(args map[string]string) (interface{}, error) {
	if trans.status != true {
		contract.Transfer(contract.Caller(), contract.Value()) //退回转账
		log.Info("交易双方未通过审核，交易无法进行")
		return nil, errors.New("交易双方未通过审核，交易无法进行")
	}
	if transactionEnd.Before(time.Now()) {
		contract.Transfer(contract.Caller(), contract.Value()) // 退回转账
		log.Info("交易申报已结束")
		return nil, errors.New("交易申报已结束")
	}

	trans.transactionid = contract.Call("random", "GetRandom", map[string]string{}) //生成交易代码
	if sideofTransaction == "seller" {
		trans.carbonAmount = contract.Value //交易的碳配额数量为合约价值

	}
	if sideofTransaction == "buyer" {
		trans.Value = contract.Value //交易的总价为合约价值

	}

	trans.progress = "ed"
	trans.time = time.Now() //交易发起时间
	log.Info("交易已生成")
	return nil, nil
}
