# Design Document

This document summarizes my approach to the Teleport Level 1 Product Engineer challenge (per requirements).

## Worker Library

In order to avoid concurrency concerns (concurrent iteration, reading, and writing), I've divided my Worker Queue into three parts. Each part is implemented as a Map:

1. [Worker Queue](src/models/WorkerQueue.go) - table used to process Worker objects based on scheduled time
1. [Worker Table](src/models/WorkerTable.go) - lookup table for actual Worker objects
1. [Status Table](src/models/StatusTable.go) - lookup table for the status of any previous or extant Worker

Advantages of this approach:

1. I can define a unique mutex by dividing the Worker Queue into several individual, singleton, units. If these were combined into a single object (say one with three Maps) that shared a single mutex, concurrency issues would likely arise.
1. Each part is a singleton - a source of truth throughout the app and service. It guarantees that when an object is read, it's accurate.
1. Each part will be backed by CRUD operations (Java Spring Boot repository-style design pattern) that act as Getters and Setters. This simplifies management and allows mutexes to be locked and unlocked so that concurrent operations on the same object don't lead to issues.
1. By dividing the Worker Queue into multiple parts, I can reduce I/O and reads on the same objects. For many operations I only need to see a worker's status. For others, I only need to see the worker's schedule time.

Comparison with other approaches:

1. I don't use buffered channels here since jobs need to persisted (in memory) and scheduled to be run in the future.
1. There isn't a need to use buffered channels here except when the worker task is actually executed.

Worker:

1. A [worker](src/models/WorkerModel.go) is defined as a scheduled time, a bash command to be executed at that time, a status, and an output capturing the result of an executed task.
1. I've hardcoded this to always be `ls` since the requirement doc doesn't specify that commands must be unique - in a real-world scenario, a list of commands would likely be specified as an enum and some underlying Bash commands executed in a select, case, or switch statement.
1. Worker are saved into the [Worker Table](src/models/WorkerTable.go) map in-memory.

[Worker](src/models/WorkerModel.go) receiver functions provide execute task support. Adding and stopping operations involve modification of several tables and have been abstracted to application-wide [job helpers](./src/job/Jobs.go).

## TLS API