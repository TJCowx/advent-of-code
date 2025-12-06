# advent-of-code

This repository contains multiple years of Advent of Code solutions.
Older years (2023 and 2024) are included as submodules.
All years from 2025 onward live directly in this repository.

## Structure

```text
advent-of-code/
├── go_utils/ (Go – shared utilities for 2025+)
├── 2025/ (Go)
├── 2024/ (Go – submodule)
└── 2023/ (Rust – submodule)
```

## Go setup

I set it up so they all link locally (2025 beyond). To link the `go_utils` run:
```bash
go work init ./go_utils ./2025
```
You can replace 2025 with anything beyond 2026 (if and when becomes available)

## Running the Code

`cd` into the directory you want to run. (For 2023, 2024 follow their readme).

### Go Projects (2025)

```bash
# go run main.go day-part
go run main.go 1-2 # Runs day 1 part 2
go run main.go 1 # Runs just day 1
```

To run the tests
```bash
go test -v ./day01.go
```


## Execution

I put the 🌟 on what I have completed and executions times

### 2025

```
Runtimes based on 
- Memory: 16GB
- Processor: i9-1200H
```

| Day | Completed | P1 Runtime | P2 Runtime    |
|-----|:---------:|-----------:|--------------:|
| 01  |    🌟     | 92.617µs   |               |
| 02  |    🌟🌟   | 57.13114ms | 265.989988ms  |
| 03  |    🌟🌟   | 35.20985ms | 1.443598065s  |
| 04  |    🌟🌟   | 513.269µs  | 15.96112ms    |
| 05  |    🌟🌟   | 3.343751ms | 27.729µs      |
| 06  |    🌟🌟   | 309.548µs  | 940.218µs     |
| 07  |           |            |               |
| 08  |           |            |               |
| 09  |           |            |               |
| 10  |           |            |               |
| 11  |           |            |               |
| 12  |           |            |               |
| 13  |           |            |               |
| 14  |           |            |               |
