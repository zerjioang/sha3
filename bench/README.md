# Performance results

This directory contains performance test results files. We recommend running same tests in your device for fresh results

**Using standard library sha3 implementation**

```bash
BenchmarkGenerateMethodSignatures/generate-std-12         	 1000000	      1641 ns/op	   0.61 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	  824888	      1392 ns/op	   0.72 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	 1283482	      1083 ns/op	   0.92 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	 1378556	      1714 ns/op	   0.58 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	  926504	      1542 ns/op	   0.65 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	 1277808	      1293 ns/op	   0.77 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	  976179	      1174 ns/op	   0.85 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	 1000000	      1328 ns/op	   0.75 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	 1238808	      1495 ns/op	   0.67 MB/s	     976 B/op	       6 allocs/op
BenchmarkGenerateMethodSignatures/generate-std-12         	  902724	      1271 ns/op	   0.79 MB/s	     976 B/op	       6 allocs/op
```

**Using this package**

```
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1985337	       677.4 ns/op	   1.48 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1255321	       801.8 ns/op	   1.25 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1963074	       705.9 ns/op	   1.42 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1882951	       750.5 ns/op	   1.33 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1425180	       878.3 ns/op	   1.14 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1473018	       875.1 ns/op	   1.14 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	  999141	      1039 ns/op	   0.96 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1543413	       831.2 ns/op	   1.20 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1531819	       723.1 ns/op	   1.38 MB/s	     456 B/op	       2 allocs/op
BenchmarkGenerateMethodSignatures/generate-custom-12         	 1483425	       793.4 ns/op	   1.26 MB/s	     456 B/op	       2 allocs/op
```

## Test script

```go
package main

import (
	"encoding/hex"
	sha32 "golang.org/x/crypto/sha3"
	"testing"
)

func GenerateMethodSignatureStd(fname []byte) string {
	var h = sha32.NewLegacyKeccak256()
	_, _ = h.Write(fname)
	hash := h.Sum(nil)
	return hex.EncodeToString(hash[0:4])
}

func BenchmarkGenerateMethodSignatures(b *testing.B) {
	b.Run("generate-std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		var f2 string
		msg := []byte("f09140466846285922(address,bytes)")
		for i := 0; i < b.N; i++ {
			f2 = GenerateMethodSignatureStd(msg)
		}
		if f2 == "" {
			panic("")
		}
	})
}
```