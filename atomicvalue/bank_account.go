package atomicvalue

import "sync/atomic"

type Account struct {
  state atomic.Value
}

type state struct {
  isClosed bool
  balance int64
}

func Open(amount int64) *Account {
  if amount < 0 {
    return nil
  } 
  var init atomic.Value
  init.Store(state{isClosed: false, balance: amount})
  return &Account { state: init }  
}

func (a *Account) getState() state {
  return a.state.Load().(state)
}

func (a *Account) Balance() (bal int64, ok bool) {
  current := a.getState()
  if current.isClosed {
    return
  }

  bal = current.balance
  ok = true
  return
}

func (a *Account) Deposit(amount int64) (bal int64, ok bool) {
  for {
    current := a.getState()
    if current.isClosed {
      return
    }
    if current.balance + amount < 0 {
      bal = current.balance
      ok = false
      return
    }

    updated := current
    updated.balance += amount
    
    if a.state.CompareAndSwap(current, updated) {
      bal = updated.balance 
      ok = true
      return
    }
  }
}

func (a *Account) Close() (bal int64, ok bool) {
  for {
    current := a.getState()
    if current.isClosed {
      return
    }
    updated := current
    updated.isClosed = true
    updated.balance = 0
    if a.state.CompareAndSwap(current, updated) {
      bal = current.balance
      ok = true
      return
    }
  }
}
