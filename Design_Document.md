# Design Document

This document summarizes my approach to the Teleport Level One Backend Engineer challenge (per [requirements](https://github.com/gravitational/careers/blob/main/challenges/systems/worker.pdf)).

## Worker Library

In order to avoid concurrency concerns (concurrent iteration, reading, and writing), I've divided my "Worker Queue" into three parts. Each part is implemented as a **Map**:

1. [Worker Queue](src/models/WorkerQueue.go) - table used to process **Worker** objects based on scheduled time implemented as a `map[string]time.Time` (`uuid` -> `scheduled time`).
1. [Worker Table](src/models/WorkerTable.go) - lookup table for actual **Worker** objects implemented as a `map[string]WorkerModel` (`uuid` -> `worker object`).
1. [Status Table](src/models/StatusTable.go) - lookup table for the status of any previous or extant **Worker**  implemented as a `map[string]string` (`uuid` -> `worker status`).

The **Worker Queue** object is processed via the [Job Loop](src/job/JobLoop.go) which polls the **Worker Queue** every 5 seconds in [ProcessQueue()](src/job/Run.go). A simple **Time** comparison is performed to verify whether a task should be executed or not.

Advantages of this approach:

1. The "Worker Queue" need not be a **Stack**, **Linked List**, **Dequeue**, or actual **Queue** since the order of arrival doesn't matter. (This is why I put "Worker Queue" in quotes.)
1. I can define a unique **Mutex** by dividing the Worker Queue into several individual, singleton, units. If these were combined into a single object (say one with three **Maps**) that shared a single **Mutex**, concurrency issues would likely arise.
1. Each part is a singleton - a source of truth throughout the app and service. It guarantees that when an object is read, it's accurate.
1. Each part will be backed by CRUD operations (Java Spring Boot repository-style design pattern) that act as getters and setters. This simplifies management and allows **Mutexes** to be locked and unlocked so that concurrent operations on the same object don't lead to issues.
1. By dividing the "Worker Queue" into multiple parts, I can reduce I/O and reads on the same objects. For many operations I only need to see a **Worker's** status. For others, I only need to see the **Worker's** schedule time.
1. Read, delete, update, and insert operations are performed in approximately O(1) time.

Comparison with other approaches:

1. I don't use buffered channels here since jobs need to be persisted (in memory) and scheduled to be run in the future.
1. There isn't a need to use buffered channels here except when the **Worker** task is actually executed.

Worker:

1. A [worker](src/models/WorkerModel.go) is defined as an uuid, scheduled time, bash command to be executed at that time, status, and an output capturing the result of an executed task.
1. I've hardcoded this to always be `ls` since the requirement doc doesn't specify that commands must be unique. In a real-world scenario, a list of commands would likely be specified as an **Enum** with some underlying Bash commands executed within a select, case, or switch statement. To mimic lengthier processes, a timeout of 3 seconds is enforced within each task.
1. **Worker** are saved into the [Worker Table](src/models/WorkerTable.go) map in-memory. They're removed following a `failed`, `completed`, or `stopped` status update.
1. The status workflow goes as follows: [`queued`] -> [`executing`] -> [`failed`, `completed`] or [`queued`] ->[`stopped`].
1. When a **Worker** is executed, the bash command is executed within a go routine. Its output is passed to a buffered channel specific to that worker. This is done to allow the contents of a command operation to be loggable and queryable.

[Worker](src/models/WorkerModel.go) receiver functions provide execute task support. 

Adding and stopping operations involve modifying several tables and have been abstracted to job-specific [helpers](./src/job/Jobs.go).

## TLS API

