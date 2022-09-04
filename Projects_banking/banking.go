package accounts

import "errors"

// Account struct
type Account struct {
	owner   string // 소문자는 private
	balance int    // 대문자는 public, 함수도 마찬가지.
}

var errNoMoney = errors.New("can't withdraw") // error를 만들 떄, 이름을 err를 붙이는 것이 좋다.

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 복사본을 만드는 것이 아니라 실제를 반환 -> 그래서 main에서 수정할 수 없다.
}

// Deporsit x amount on your account (Method 작성하는 법)
func (a *Account) Deposit(amount int) { // receiver이라고 부른다. a의 type은 Account이다. 이름은 자유롭게 지을 수 있다.
	a.balance += amount
}

// 단순히 a Account로 작성하면 복사본을 받기 때문에 원본의 값에는 영향을 미치지 않는다. -> a *Account로 수정.

// receiver을 작성하는 데 주의해야할 점
// struct의 첫 글자를 따서 소문자로 지어야 한다는 것.

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil // Null or None 과 같은 것.
}

// Go에는 exception이 없다. try~except / try~catch 등.
// error을 return해줘야 하고 error를 직접 체크해줘야한다.
