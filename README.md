# StreamingFast Key/Value Store

[![reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/zkpkg/kvdb)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)


This library contains different implementations for storing blocks and transactions in a key/value database.
It is used by **[StreamingFast](https://github.com/streamingfast/streamingfast)**.

## Usage

```go
db, err := store.New(dsn)
if err != nil {
    return fmt.Errorf("failed setting up db: %w", err)
}
```


## Backends

The following DSNs are provided by this package:
* Google Cloud BigTable: `bigkv://project.instance/namespace-prefix?createTables=true`
  This works very well, is fully managed, and scales horizontally with minimal effort.

* TiKV:   `tikv://pd0,pd1,pd2:2379?prefix=namespace_prefix`
  This is useful for bare metal deployments, is self managed, and scales very well (with the hardware you throw at it)

* Badger: `badger:///home/user/dfuse-data/component/my-badger.db`
  This is useful for local development.  It is a library (similar to RocksDB and LevelDB), and thus creates a database that cannot be shared.

* NetKV: `netkv://localhost:6789?insecure=true`
  This connects to a `netkv` server (which you can install with `go install -v ./store/netkv/server/netkvserver` from this repo), which in turn can serve a `badger://` database.  It allows for simple badger-based backend (single database, no replication, no scaling), but allow decoupling of StreamingFast processes


**Beware** that the TiKV backend does not support 0-length values. If
your application uses 0-length values, use the `WithEmptyValue`
option.


## Contributing

**Issues and PR in this repo related strictly to the kvdb library.**

Report any protocol-specific issues in their
[respective repositories](https://github.com/streamingfast/streamingfast#protocols)

**Please first refer to the general
[StreamingFast contribution guide](https://github.com/streamingfast/streamingfast/blob/master/CONTRIBUTING.md)**,
if you wish to contribute to this code base.


## License

[Apache 2.0](LICENSE)
