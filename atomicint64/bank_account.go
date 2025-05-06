package atomicint64

import "sync/atomic"

type Account struct {
  state atomic.Int64
}

func Open(amount int64) *Account {
  if amount < 0 {
    return nil
  }
  var state atomic.Int64
  state.Store(amount)
  return &Account { state }
}

func (a *Account) Balance() (bal int64, ok bool) {
  state := a.state.Load()

  if state < 0 {
    return
  }

  bal = state
  ok = true
  return
}

func (a *Account) Deposit(amount int64) (bal int64, ok bool) {
  for {
    state := a.state.Load()
    if state < 0 {
      bal = 0
      ok = false
      return
    } 

    bal = state

    var updated = bal + amount
    if updated < 0 {
      ok = false
      return
    }

    if a.state.CompareAndSwap(bal, updated) {
      bal = updated
      ok = true
      return
    }
  }
}

func (a *Account) Close() (bal int64, ok bool) {
  for {
    state := a.state.Load() 
    if state < 0 {
      return
    }
    
    if a.state.CompareAndSwap(state, -1) {
      bal = state
      ok = true
      return
    }
  }
}
