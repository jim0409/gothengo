package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
比如有一個電商系統中的交易類`transaction`，用來記錄每筆訂單的交易情況

其中的`Execute()`函數負責執行轉帳操作，將錢從買假的帳戶轉移到賣家的帳戶中

而真正的轉帳操作則視同過調用銀行(`支付寶、微信`)的 SDK 完成的
*/

type TransactionStatus int

const (
	Executed = iota
	Expired
	Failed
)

type Config struct {
	token string
}

var config = Config{
	token: "123",
}

/* Deprecated */
// type transaction struct {
// 	ID        string
// 	BuyerID   int
// 	SellerID  int
// 	Amount    float64
// 	createdAt time.Time
// 	Status    TransactionStatus
// }

/* Deprecated */
// func (t *transaction) Execute() bool {
// 	if t.Status == Executed {
// 		return true
// 	}
// 	if time.Now().Sub(t.createdAt) > 24*time.Hour { // 交易有有效期
// 		t.Status = Expired
// 		return false
// 	}

// 	client := BankClient.New(config.token) // 调用银行的 SDK 执行转账
// 	if err := client.TransferMoney(t.ID, t.BuyerID, t.SellerID, t.Amount); err != nil {
// 		t.Status = Failed
// 		return false
// 	}
// 	t.Status = Executed
// 	return true
// }

/*
	這個類最重要的功能集中在 Execute() 函數中，但它卻不好測試，因為它有兩個外部依賴
	1. 行為不確定的 time.Now()，他的每一次調用都會產生不同的結果
	2. 銀行提供的轉帳 SDK 我們不可能每次測試都去真的調用一下，那測試成本也很高

	> 解決方法，就是把這兩個依賴 mock 掉，即使用一個 "假的" 服務來替換真的服務，
	  這裏我們先拿測試成本較高的銀行 SDK 測試
*/

// Mock SDK Dependency
/*
	按照 DI 的理論，先將代碼裡使用到的方法抽象成一個接口
	目前這個接口只包含一個方法，實際的場景下抽象出來的接口可能更複雜
*/

type Transferer interface {
	TransferMoney(id int, buyerID int, sellerId int, amount float64) error
}

/*
將創建 BankClient 的行為上移到調用者那邊去，相當於調用者創建一個滿足`Transferer`接口的實例，再注入我們的代碼
所以`transaction`這邊就需要有個地方來接受這個實例

一個方法是通過`Execute()`函數的參數，但如果依賴過多的話，會造成函數爆炸
另一個則是放到`transaction`的成員屬性中
*/

type transaction struct {
	ID        string
	BuyerID   int
	SellerID  int
	Amount    float64
	createdAt time.Time
	Status    TransactionStatus
	// 增加一個存放接口的屬性
	transfer Transferer
}

// New generat an transaction properity to make sure
func New(buyerId, sellerId int, amount float64, transferer Transferer) *transaction {
	return &transaction{
		ID:        fmt.Sprintf("%v", buyerId+1000), // id may generated via ...
		BuyerID:   buyerId,
		SellerID:  sellerId,
		Amount:    amount,
		createdAt: time.Now(),
		Status:    1, // taken from ...
		transfer:  transferer,
	}
}

var nowFn = time.Now

func (t *transaction) Execute() bool {
	// ..
	if t.Status == Executed {
		return true
	}
	if nowFn().Sub(t.createdAt) > 24*time.Hour {
		t.Status = Expired
		return false
	}
	// 不直接創建，而是使用別人注入的接口實例
	id, _ := strconv.Atoi(t.ID)
	t.transfer.TransferMoney(id, t.BuyerID, t.SellerID, t.Amount)
	// ..
	return true
}

/*
Mock `time.Now`

事實上，mock當前時間是一個很常見的問題，類似的函數還有`rand.Intn`
他們的共同點就是輸出是不確定的，這就讓我們的測試無法覆蓋所有的情況

面對這些函數，通常的做法就是套用一個全域變數(函數)，當需要時才呼叫
測試時，可以直接覆蓋全域變數(函數)來做假設

var nowFn = time.Now
*/
