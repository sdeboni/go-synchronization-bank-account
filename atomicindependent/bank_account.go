// Credit to backdevjung's mentoring suggestion
package atomicindependent

import "sync/atomic"

type Account struct {
    balance	atomic.Int64
    closed	atomic.Bool
}

func Open(amount int64) *Account {
	if amount < 0 {
        return nil
    }
    var acc Account
    acc.balance.Store(amount)
    return &acc
}

func (a *Account) Balance() (bal int64, ok bool) {
    if a.closed.Load() {
        return
    }
    bal, ok = a.balance.Load(), true
    return
}

func (a *Account) Deposit(amount int64) (bal int64, ok bool) {
    if a.closed.Load() {
        return
    }
    var current, updated int64
    for {
        current = a.balance.Load()
        updated = current + amount
        if updated < 0 {
            return
        }
        if a.balance.CompareAndSwap(current, updated) {
           if a.closed.Load() {
             a.balance.Store(0)
             bal, ok = 0, false
             return
           } else {
             bal, ok = updated, true
             return
           }
        }
    }
    return
}

func (a *Account) Close() (bal int64, ok bool) {
    if a.closed.CompareAndSwap(false, true) {
        bal, ok = a.balance.Swap(0), true
        return
    }
    return
}
