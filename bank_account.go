package account

import "sync"

type Account struct {
  balance int64
  isClosed bool
  mu sync.Mutex
}

func Open(amount int64) *Account {
  if amount < 0 {
    return nil
  }
  return &Account { balance: amount, isClosed: false }
}

func (a *Account) Balance() (int64, bool) {
  if a.isClosed {
    return 0, false
  }
  a.mu.Lock()
  defer a.mu.Unlock()

  if a.isClosed {
    return 0, false
  }

  return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
  if a.isClosed {
    return 0, false
  }

  a.mu.Lock() 
  defer a.mu.Unlock()
  // syncyhronized section 

  if a.isClosed {
    return 0, false
  }

  if a.balance + amount < 0 {
    return 0, false
  }

  a.balance += amount
  return a.balance, true
}

func (a *Account) Close() (int64, bool) {
  if a.isClosed {
    return 0, false
  }

  a.mu.Lock()
  defer a.mu.Unlock()
  // syncyhronized section 

  if a.isClosed {
    return 0, false
  }

  a.isClosed = true
  balance := a.balance
  a.balance = 0

  return balance, true
}
