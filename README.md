# Go Concurrency Race Condition

This repository demonstrates a common race condition in Go programs that involves incorrect usage of mutexes in concurrent scenarios.  The example showcases how a simple counter increment operation within goroutines can produce unexpected results if not properly synchronized.  The solution provides the corrected code with appropriate mutex handling.