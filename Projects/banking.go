package accounts

// Account struct
type Account struct {
	owner   string // 소문자는 private
	balance int    // 대문자는 public, 함수도 마찬가지.
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 복사본을 만드는 것이 아니라 실제를 반환 -> 그래서 main에서 수정할 수 없다.
}
