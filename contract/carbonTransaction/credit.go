package main

import (
	"errors"
	"fmt"
	"github.com/ssbcV2/contract"
	"time"
)

var creditTime time.Time //交易评价完成时间
var creditStatusOfB bool //买方评价进度
var creditStatusOfS bool //卖方评价进度
var creditStatus bool    //评价进度
var side string          //评价方

func init() {
	trans = transaction{
		transactionid: "",
		category:      "standardCarbonCredit",
		progress:      "ed",
	}
	creditStatus = false

}

func setValue(args map[string]string) (interface{}, error) {

	fmt.Println("Please enter transactionid")
	fmt.Scan(&trans.transactionid)

	fmt.Println("Please enter your side")
	fmt.Scan(&side)
	return nil, nil
}

func evaluate(args map[string]string) (interface{}, error) {
	side = contract.Caller()
	tranId = trans.transactionid

	if trans.progress == "ing" {
		log.Info("交易未完成。不能评价")
		return nil, errors.New("交易未完成。不能评价")
	}
	if side != trans.buyerid && side != trans.sellerid {
		log.Info("非法交易身份，不能评价")
		return nil, errors.New("非法交易身份，不能评价")
	}
	if side == trans.buyerid {
		fmt.Println("Please enter your credit")
		fmt.Scan(&trans.buyerCredit)
		creditStatusOfB = true
	}
	if side == trans.selleridid {
		fmt.Println("Please enter your credit")
		fmt.Scan(&trans.sellerCredit)
		creditStatusOfS = true
	}
	if creditStatusOfB && creditStatusOfS {
		creditStatus = true
		creditTime = time.Now()
	}

	return nil, nil
}
