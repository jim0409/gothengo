package main

import (
	"errors"
	"testing"
	"time"
)

// 定義一個滿足 Transferer 接口的 mock 類
type MockedClient struct {
	responseError error // 實例化的時候可以將期望的返回值保存進來
}

func (m *MockedClient) TransferMoney(id int, buyerID int, sellerID int, amount float64) error {
	return m.responseError
}

const (
	buyerID  = 1
	sellerID = 1
	amount   = 100
)

func Test_transaction_Execute(t *testing.T) {
	// 实例化一个可以自由控制结果的 client
	transferer := &MockedClient{
		responseError: errors.New("insufficient balance"),
	}
	tnx := New(buyerID, sellerID, amount, transferer)
	if succeeded := tnx.Execute(); succeeded != false {
		t.Errorf("Execute() = %v, want %v", succeeded, false)
	}
}

func Test_expired_transaction_Execute(t *testing.T) {
	// 用同样的函数签名改写业务中需要的时间函数
	// 这里能改变私有的全局变量是因为测试代码跟业务代码处于同一个包中
	nowFn = func() time.Time {
		return time.Now().Add(-24 * time.Hour)
	}
	// 依旧需要实例化一个假的的 client
	transferer := &MockedClient{
		responseError: nil,
	}
	tnx := New(buyerID, sellerID, amount, transferer)
	//...
	_ = tnx
}
