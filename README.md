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

