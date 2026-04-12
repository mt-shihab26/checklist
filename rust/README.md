# Rust

## 1. Basics

* [ ] Install Rust via rustup
* [ ] Verify installation (`rustc --version`)
* [ ] Use Cargo:
  * [ ] `cargo new`
  * [ ] `cargo run`
* [ ] Hello World program
* [ ] Variables (`let`, `mut`)
* [ ] Constants (`const`)
* [ ] Shadowing
* [ ] Data types:
  * [ ] Integers
  * [ ] Floats
  * [ ] Boolean
  * [ ] Char
* [ ] Compound types:
  * [ ] Tuple
  * [ ] Array
* [ ] Operators
* [ ] Input/Output (`println!`)
* [ ] Control flow:
  * [ ] `if / else`
  * [ ] `loop`
  * [ ] `while`
  * [ ] `for`

## 2. Core Language

* [ ] Functions
  * [ ] Parameters
  * [ ] Return values
  * [ ] Expressions vs statements
* [ ] Ownership (critical concept)
  * [ ] Move semantics
  * [ ] Copy trait
* [ ] Borrowing
  * [ ] Immutable references
  * [ ] Mutable references
* [ ] Slices (`&str`, array slices)
* [ ] Strings:
  * [ ] `String`
  * [ ] `&str`

## 3. Structs, Enums, Pattern Matching

* [ ] Structs
  * [ ] Named structs
  * [ ] Tuple structs
* [ ] Methods (`impl`)
* [ ] Enums
* [ ] Pattern matching (`match`)
* [ ] `if let`
* [ ] `while let`
* [ ] Standard enums:

  * [ ] `Option`
  * [ ] `Result`

---

## 4. Collections & Iteration

* [ ] Vectors (`Vec<T>`)
* [ ] HashMaps
* [ ] Iterators:

  * [ ] `.iter()`
  * [ ] `.into_iter()`
  * [ ] `.map()`
  * [ ] `.filter()`
  * [ ] `.collect()`

---

## 5. Error Handling

* [ ] `Result<T, E>`
* [ ] `unwrap()` vs `expect()`
* [ ] Error propagation (`?`)
* [ ] Custom error types
* [ ] `thiserror` / `anyhow` basics

---

## 6. Modules & Packages

* [ ] Modules (`mod`)
* [ ] Visibility (`pub`)
* [ ] File structure
* [ ] Crates (binary vs library)
* [ ] Workspaces
* [ ] Dependencies in Cargo.toml

---

## 7. Traits & Generics

* [ ] Define traits
* [ ] Implement traits
* [ ] Derive traits (`Debug`, `Clone`)
* [ ] Generics (`<T>`)
* [ ] Trait bounds
* [ ] `where` clauses

---

## 8. Lifetimes

* [ ] Lifetime annotations
* [ ] Borrow checker rules
* [ ] Lifetime elision
* [ ] Struct lifetimes
* [ ] `'static`

---

## 9. Smart Pointers & Memory

* [ ] `Box`
* [ ] `Rc`
* [ ] `Arc`
* [ ] `RefCell`
* [ ] Interior mutability
* [ ] Ownership vs shared state

---

## 10. Concurrency

* [ ] Threads (`std::thread`)
* [ ] Message passing (channels)
* [ ] Shared state:

  * [ ] `Mutex`
  * [ ] `Arc`
* [ ] Send & Sync traits

---

## 11. Async Rust

* [ ] Async/await
* [ ] Futures
* [ ] Executors
* [ ] Use Tokio
* [ ] Async file/network handling

---

## 12. Macros & Advanced Features

* [ ] `macro_rules!`
* [ ] Procedural macros (intro)
* [ ] Unsafe Rust
* [ ] FFI (C interop)

---

## 13. File & Data Handling

* [ ] File read/write
* [ ] JSON (`serde`)
* [ ] Environment variables
* [ ] CLI args (`std::env`)

---

## 14. Testing & Debugging

* [ ] Unit tests (`#[test]`)
* [ ] Integration tests
* [ ] Benchmarks
* [ ] Logging
* [ ] Debugging tools

---

## 15. Web Development

* [ ] HTTP basics
* [ ] Build API using:

  * [ ] Axum or
  * [ ] Actix Web
* [ ] Routing
* [ ] Middleware
* [ ] REST API design

---

## 16. Database

* [ ] SQL basics
* [ ] Use:

  * [ ] SQLx or
  * [ ] Diesel
* [ ] CRUD operations
* [ ] Migrations
* [ ] Connection pooling

---

## 17. CLI Tools

* [ ] Build CLI apps
* [ ] Argument parsing:

  * [ ] Clap

---

## 18. Tooling

* [ ] rust-analyzer
* [ ] `clippy`
* [ ] `rustfmt`

---

## 19. Advanced Topics

* [ ] Generics deep dive
* [ ] Zero-cost abstractions
* [ ] Memory layout
* [ ] Performance optimization
* [ ] Profiling

---

## 20. Architecture & Design

* [ ] Project structure
* [ ] Clean architecture
* [ ] Error design patterns
* [ ] API design

---

## 21. DevOps & Deployment

* [ ] Build binaries
* [ ] Cross-compilation
* [ ] Docker
* [ ] CI/CD basics

---

## 22. Real-World Projects

* [ ] CLI tool
* [ ] REST API
* [ ] Concurrent system
* [ ] Full backend (auth + DB)
* [ ] File processor

---

## Final Mastery Checklist

* [ ] Understand ownership deeply
* [ ] Use lifetimes confidently
* [ ] Write idiomatic Rust
* [ ] Build safe concurrent systems
* [ ] Read complex Rust codebases
* [ ] Contribute to crates


