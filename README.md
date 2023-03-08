```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/rfjakob/secretbox_benchmark
cpu: Intel(R) Core(TM) i5-3470 CPU @ 3.20GHz
BenchmarkSecretbox/blockSize1024-4         	  568634	      2099 ns/op	 487.93 MB/s
BenchmarkSecretbox/blockSize2048-4         	  335469	      3563 ns/op	 574.81 MB/s
BenchmarkSecretbox/blockSize4096-4         	  184939	      6470 ns/op	 633.09 MB/s
BenchmarkSecretbox/blockSize8192-4         	   97590	     12284 ns/op	 666.86 MB/s
BenchmarkSecretbox/blockSize16384-4        	   50010	     24000 ns/op	 682.67 MB/s
BenchmarkSecretbox/blockSize32768-4        	   25194	     47597 ns/op	 688.44 MB/s
BenchmarkSecretbox/blockSize65536-4        	   12690	     94577 ns/op	 692.94 MB/s
BenchmarkSecretbox/blockSize131072-4       	    6324	    188895 ns/op	 693.89 MB/s
BenchmarkSecretbox/blockSize262144-4       	    3174	    377696 ns/op	 694.06 MB/s
BenchmarkSecretbox/blockSize524288-4       	    1587	    754909 ns/op	 694.50 MB/s
BenchmarkSecretbox/blockSize1048576-4      	     792	   1510909 ns/op	 694.00 MB/s
PASS
ok  	github.com/rfjakob/secretbox_benchmark	15.402s
```
