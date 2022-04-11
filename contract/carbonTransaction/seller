package main

import (
	"errors"
	"github.com/ssbcV2/contract"
	"time"
)

type seller struct {
	name                  string
	registerAccountStatus bool              //登记账户状态
	carbonBalance         int               //配额库存
	CapitalBalance        int               //资金账户余额
	transactionRecord     map[string]string //交易记录  交易编号/交易对方

}

type transaction struct {
	transactionid string
	category      string //产品类型
	carbonAmount  int    //排放额交易数量
	Value         int    //总价

	statusofBuyer  bool      //买方交易状态
	statusofSeller bool      //卖方交易状态
	status         bool      //交易审核状态
	progress       string    //交易进行状态
	buyerid        string    //买方信息
	sellerid       string    //卖方信息
	buyerCredit    int       //买方评价
	sellerCredit   int       //卖方评价
	time           time.Time //发起时间
	EndTime        time.Time //交易完成时间
}

var entity seller

var carbonAmount int //碳配额库存
var perValue int
var veriMark bool // 有效卖方标记

func init() {
	entity = seller{
		name:                  "",
		registerAccountStatus: true,
		carbonBalance:         100,
		CapitalBalance:        0,
		transactionid:         "",
	}

	perValue = 4
	veriMark = false
}
func setPerValue(args map[string]string) (interface{}, error) {
	entity.name = contract.Caller()
	entity.registerAccountStatus = true
	carbonAmount = contract.Value()

	return nil, nil
}

func sellerPrimary(args map[string]string) (interface{}, error) {
	entity.name = contract.Caller()
	if entity.registerAccountStatus != true {
		contract.Transfer(contract.Caller(), contract.Value()) //退回转账
		log.Info("登记账户存在异常，退出")
		return nil, errors.New("登记账户存在异常，退出")
	}

	if entity.carbonBalance == 0 {
		contract.Transfer(contract.Caller(), contract.Value()) //退回转账
		log.Info("无可用碳排放配额")
		return nil, errors.New("无可用碳排放配额")
	}

	if carbonAmount > entity.carbonBalance {
		contract.Transfer(contract.Caller(), contract.Value()) //退回转账
		log.Info("超过可用库存")
		return nil, errors.New("超过可用库存")
	}

	var trans transaction
	trans.statusofSeller = true
	veriMark = true //通过审核
	if trans.statusofBuyer == true {
		trans.status = true //生成一笔交易
		trans.carbonAmount = carbonAmount
		trans.Value = carbonAmount * perValue

	}

	return nil, nil
}
