# Go Rectangles: A Performance Comparison with C#

## Introduction

After completing my [Udemy course in Golang](https://github.com/jahanalem/LinkedIn2GitHub/blob/main/0030_golang-foundations-and-microservices-plan.md), I wanted to put my new knowledge into practice. To do this, I decided to convert my C# project, `RectanglesCalculator`, into Golang (Go).

My goal was not only to rewrite the project but also to compare the performance of these two popular programming languages — C# (.NET 9) and Go — for the same computational task.

## About the Original Project (C#)

The original project, [RectanglesCalculator](https://github.com/jahanalem/RectanglesCalculator), is a program that finds all possible rectangles formed by a given set of points on a 2D plane. Each rectangle is identified when four points create two horizontal and two vertical lines that match perfectly. The algorithm must efficiently detect and count all valid rectangles, even in large datasets.

## The Performance Comparison

To test both versions, I created JSON files with a varying number of points, from a small set of 16 up to a large set of 50,000. I ran the exact same tests on both the C# and Go applications.

The results were amazing and showed a significant performance advantage for the Go version.

### The Results in Detail

The data below speaks for itself. Go runs much faster than C#, and the performance gap increases dramatically with the size of the dataset.


<img width="1812" height="545" alt="Performance-Comparision_Table" src="https://github.com/user-attachments/assets/0773130d-f271-4bbb-a60a-6824ab1de4bf" />

<img width="1461" height="947" alt="Performance-Comparision" src="https://github.com/user-attachments/assets/000058f3-f79c-4382-bb86-67a8c1d89b63" />

As you can see in the chart, the orange line for C# climbs very steeply, while the blue line for Go remains much flatter. At 50,000 points, the Go version was nearly **20 times faster** than the C# version.

## Why is Go So Much Faster?

The algorithm is identical in both projects, so the difference is not in the logic. Instead, the performance advantage comes from how Go is designed, especially in how it handles **concurrency** and **memory management**.

Here is a simple explanation:

### 1. Lightweight Concurrency (Ants vs. Elephants)

Imagine you need to move thousands of small packages.

-   **Go uses Goroutines**: Think of goroutines as a huge team of tiny, efficient ants. You can have hundreds of thousands of them, and they work together perfectly without getting in each other's way. They require very few resources and are managed with incredible efficiency.

-   **C# uses Threads**: Think of threads as powerful but heavy elephants. They are very strong, but it's slow and expensive to get them started and to manage thousands of them for small tasks. This creates a lot of overhead, which slows down the whole process.

For this project, Go's "ant colony" approach is far more effective.

### 2. Efficient Memory Management (A Shopping List vs. a Scavenger Hunt)

Imagine you need to collect several items from a store.

-   **Go uses [Structs](https://www.w3schools.com/go/go_struct.php)**: Go stores related data together in one continuous block of memory. This is like having a perfectly organized shopping list where all items are in the same aisle. The CPU can grab everything it needs in one quick trip.

-   **C# uses [Classes](https://learn.microsoft.com/en-us/dotnet/standard/design-guidelines/choosing-between-class-and-struct)**: In C#, data can be spread out across different memory locations. To get all the data for one object, the CPU has to jump from one location to another. This is like a scavenger hunt, where each clue leads to the next. All that jumping around takes time.

Go's "shopping list" approach means data is accessed much faster, which makes a huge difference.

## Conclusion

This project was a fantastic opportunity to practice my Go skills and understand how it compares to C#. While both languages are powerful, Go proved to be significantly faster and more memory-efficient for this type of computational and concurrent processing.

The results show that Go is an excellent choice for high-performance applications, especially when dealing with large amounts of data and concurrency.





## ⚙️ Concurrency vs Parallelism

These two words — **concurrency** and **parallelism** — look similar, but they are not the same.  
Let’s understand them in a very simple way.

---

### 🧠 What is Concurrency?

**Concurrency** means *doing many things at the same time in a shared period of time.*

But it doesn’t mean that all things run exactly at the same moment.  
They just **take turns quickly** — so it feels like they happen together.

🟢 Example:  
Imagine you are cooking dinner and also chatting with a friend.  
You cut vegetables, then you check your phone, then you stir the soup.  
You are doing both tasks *in the same time frame*, but not *at the same exact moment.*

That’s **concurrency**.

In Go, **goroutines** make your program concurrent — they share CPU time efficiently.

---

### ⚡ What is Parallelism?

**Parallelism** means *doing many things at the same exact time.*

This happens when you have **multiple CPU cores**, and each core runs a different task simultaneously.

🟣 Example:  
Imagine two cooks in the kitchen.  
One is cutting vegetables while the other is boiling water — both truly at the same time.

That’s **parallelism**.

---

### 🔍 The Difference in One Sentence

| Concept | Meaning | Example |
|----------|----------|----------|
| **Concurrency** | Many tasks taking turns | One cook doing many things quickly |
| **Parallelism** | Many tasks running exactly at the same time | Two cooks working together at once |

---

### 💡 In Go

Go is designed for **concurrency first**.  
It uses **goroutines** to let tasks run together smoothly.  
If your computer has multiple CPU cores, Go can also run them **in parallel** — but that depends on your system and the Go scheduler.

✅ So:  
- **Goroutines → Concurrency**  
- **Multiple CPU cores → Parallelism**

---

## 🧩 Understanding Goroutines, WaitGroups, and Channels (Step by Step)

In Go, these tools — **goroutines**, **WaitGroups**, and **channels** — work together to make your program **concurrent**.  
They help you run many small tasks at the same time safely and efficiently.

Let’s learn each one with very simple examples.

---

### 🟢 1. What is a Goroutine?

A **goroutine** is a small, lightweight thread.  
It allows a function to run **independently** from the rest of the program.

To start a goroutine, you add the word **`go`** before a function call.

Example:
```go
go fmt.Println("Hello from a goroutine!")
```

This line will print the message **while the main program continues running**.
So, the program can do many things “at once.”

💡 **Think of it like this:**
A goroutine is like telling your helper, “Please do this task while I do something else.”

---

### 🟣 2. What is a WaitGroup?

When you start several goroutines, your main program might finish **before** they do.
To prevent that, Go gives us a **WaitGroup** — it helps us *wait* for all goroutines to finish before continuing.

Example:

```go
var wg sync.WaitGroup // create a WaitGroup

wg.Add(1) // tell it: "we have 1 goroutine to wait for"

go func() {
    defer wg.Done() // say "I'm done" when finished
    fmt.Println("Work done!")
}()

wg.Wait() // wait until all goroutines call Done()
```

Explanation:

* `wg.Add(1)` — adds one task to the waiting list.
* `wg.Done()` — marks that task as finished.
* `wg.Wait()` — blocks (pauses) the program until all tasks are done.
* `defer` — means “run this line at the end of the function.”

💡 **Tip:**
Without `WaitGroup`, your program might end before goroutines complete their work.

---

### 🟠 3. What is a Channel?

A **channel** is like a pipe that lets goroutines **send and receive data** between each other safely.

You create a channel like this:

```go
linesChan := make(chan geometry.Line, 1000)
```

Here:

* `chan geometry.Line` means the channel will carry data of type `geometry.Line`.
* The number `1000` means it’s a *buffered* channel — it can hold up to 1000 items at once.

You can:

* **Send** data into the channel using `<-`
* **Receive** data from the channel using `<-`

Example:

```go
// Send a value
linesChan <- geometry.NewLine(p1, p2)

// Receive a value
line := <-linesChan
```

💡 **Think of a channel like a mailbox:**
One goroutine puts letters inside (sending), and another takes them out (receiving).

---

### 🧠 4. What is `close(linesChan)`?

When no more data will be sent to a channel, we **close** it.

```go
close(linesChan)
```

This tells the receiver goroutine:
“No more data is coming. You can stop reading soon.”

If you don’t close the channel, the program may wait forever for new data.

---

### 🔵 5. How These Work Together in `createLinesParallel()`

Let’s look at the code and understand each part:

```go
linesChan := make(chan geometry.Line, 1000)
var wg sync.WaitGroup

for _, group := range pointsByY {
    if len(group) < 2 {
        continue
    }
    wg.Add(1) // one goroutine starts

    go func(g []geometry.Point) {
        defer wg.Done() // mark as done when finished
        for i := 0; i < len(g); i++ {
            for j := i + 1; j < len(g); j++ {
                // send line into the channel
                linesChan <- geometry.NewLine(g[i], g[j])
            }
        }
    }(group)
}

// another goroutine closes the channel when all done
go func() {
    wg.Wait()      // wait for all goroutines
    close(linesChan) // close the channel
}()
```

Explanation in simple:

1. `linesChan` — a shared mailbox where all goroutines put their results.
2. `wg.Add(1)` — says “we’re starting one new worker.”
3. `go func(...)` — starts the worker in the background.
4. `defer wg.Done()` — tells the WaitGroup when that worker finishes.
5. `linesChan <- geometry.NewLine(...)` — sends a line into the mailbox.
6. Another goroutine waits for all workers to finish (`wg.Wait()`), then **closes the mailbox** (`close(linesChan)`).

---

### 🧩 6. What Happens in Simple Steps

1. Many goroutines start working on different groups.
2. Each goroutine sends its results (lines) into the shared channel.
3. The main goroutine collects all lines from that channel.
4. When all workers finish, the channel closes.
5. The program continues safely — no lost data, no waiting forever.

---

### ✅ In One Sentence

> **Goroutines** make Go concurrent,
> **WaitGroups** help control them,
> and **Channels** let them talk safely.

Together, they make Go programs fast, efficient, and easy to manage — like a **team of workers sharing one smart mailbox**.


