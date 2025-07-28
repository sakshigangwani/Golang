# 📦 Variables in Go

This folder demonstrates how to **declare and use variables** in Go using both the standard (`var`) and shorthand (`:=`) syntax.

Go is a statically typed language, which means **every variable has a fixed type at compile time**, but Go can also **infer the type** based on the assigned value.

---

## 🧠 What Are Variables?

Variables are **containers that hold data**. In Go, you must declare a variable before using it. You can declare variables in two main ways:

---

## ✍️ Declaring Variables the Traditional Way

```go
var quantity int = 10
var username string = "John"
var isPermission bool = true
var costPerMsg float32 = 2.89
```

## ✍️ Declaring Variables in the Easy Way (Using := Short Assignment Operator)

```go
quantity := 10
username := "John"
isPermission := true
```