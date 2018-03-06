# Inverse
Transaction-like rollback inside your code.

# Use case
 * You have a program flow with many steps
 * Steps can fails
 * Previous successful steps must be cancelled

# Usage examples

```go
// It is a typical example of stream inverse function.
// Just define function which will be called if you will inverse your flow.
// See usage of InverseBeginTransaction in DoSomething function below.
func InverseBeginTransaction(tx *gorm.DB) inverse.Func {
	return func() error {
		return tx.Rollback().Error
	}
}

// It is another stream inverse function.
// For example, that function delete file from FileService if you will inverse your flow.
// See usage of InverseCreateFile in DoSomething function below.
func InverseCreateFile(fileSvc FileService, file File) inverse.Func {
	return func() error {
		return fileSvc.Delete(file)
	}
}

func DoSomething(db *gorm.DB, fileSvc FileService) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create new stream and add InverseBeginTransaction Func.
	// If stream will be reversed, InverseBeginTransaction will be called.
	stream := inverse.NewStream(InverseBeginTransaction(tx))

	// ...

	// Create file in file service.
	file, err := fileSvc.Create("name", data)
	if err != nil {
		// There stream.Inverse() will call InverseCreateFile which rollback tx.
		return util.Err(stream.Inverse(), err)
	}

	// Add InverseCreateFile to stream.
	// If stream will be reversed, InverseBeginTransaction and InverseBeginTransaction will be called.
	stream.Add(InverseCreateFile(fileSvc, file))

	// ...

	// For example, buildSomeErr() will produce error.
	err = buildSomeErr()
	if err != nil {
		// There stream.Inverse() will call InverseBeginTransaction and after that InverseCreateFile.
		return util.Err(stream.Inverse(), err)
	}

	return nil
}
```

# Testing

### Running tests
```
make test
```

# Dependencies

### Install dependencies

Using [dep](https://github.com/golang/dep) you can install dependencies. It is necessary only for testing.
In usual case of usage the library you no need to install any dependency.

```
make dep
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

---

Project is under development and is open for any proposals.