# Advent of Code 2019

### Run tests and benchmark

```
# Run tests
cd d01/a
cargo +nightly test

# Run benchmarks
cd d01/a
cargo +nightly bench -- --nocapture
```

### Benchmarks

Benchmarks done with `cargo bench` on Intel i7-1185G7.

| Day                                          | a                              | b                              |
| -------------------------------------------- | ------------------------------ | ------------------------------ |
| [Day 1](https://adventofcode.com/2022/day/1) | [0.001 ms](./d01/a/src/lib.rs) | [0.002 ms](./d01/b/src/lib.rs) |
| [Day 2](https://adventofcode.com/2022/day/2) | [0 ms](./d02/a/src/lib.rs) | [0 ms](./d02/b/src/lib.rs) |
