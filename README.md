Testing performance differences synchronizing access to struct representing a bank account:

- lock/bank_account.go: sync.Mutex
- rw/bank_account.go  : sync.RWMutex 
- atomicint64/bank_account.go: sync/atomic.Int64 (state is encoded in 64 bits)
- atomicvalue/bank_account.go: sync/atomic.Value 
- atomicindependent/bank_account.go: sync/atomic.Bool + sync/atomic.Int64


Benchmarks:
```
Sequential
----------
Mutex          	    	86723455	        11.57 ns/op
RWMutex        	    	62778496	        16.79 ns/op
AtomicInt64         	287237896	         4.075 ns/op
AtomicValue    	    	23777126	        47.66 ns/op
AtomicIndependent    	293519960	         3.970 ns/op

Parallel
--------
Mutex                	11726322	       102.2 ns/op
RWMutex              	28465296	        44.31 ns/op
AtomicInt64         	18870572	        62.08 ns/op
AtomicValue         	 2381992	       510.4 ns/op
AtomicIndependent    	26383873	        43.57 ns/op
```
