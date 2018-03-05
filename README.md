# InverseFlow
With InverseFlow you can elegant rollback your flow.

# Testing

### Running tests
```
make test
```

### Running benchmark
```
make bench
```

# Performance

### Benchmark results

```
BenchmarkInverseStream-4        1000000000               2.99 ns/op
BenchmarkInverseNative-4        1000000000               2.33 ns/op
```

# Usage examples

```go
func InverseBeginTransaction(tx *gorm.DB) inverseflow.Func {
	return func() error {
		return tx.Rollback().Error
	}
}

func InverseCreateFile(fileSvc FileService, file File) inverseflow.Func {
	return func() error {
		return fileSvc.Delete(file)
	}
}

func DoSomething(db *gorm.DB, fileSvc FileService) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// create new stream and add InverseBeginTransaction Func
	stream := inverseflow.NewStream(InverseBeginTransaction(tx))

	// ...

	file, err := fileSvc.Create("name", data)
	if err != nil {
		// there stream.Inverse() will rollback transaction
		return util.Err(stream.Inverse(), err)
	}

	stream.Add(InverseCreateFile(fileSvc, file))

	// ...

	err = buildSomeErr()
	if err != nil {
		// there stream.Inverse() will rollback transaction and delete file
		return util.Err(stream.Inverse(), err)
	}
}
```