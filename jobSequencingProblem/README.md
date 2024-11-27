# LEARNING ABOUT GREEDY ALGORITHMS

## PROBLEM 1

### Imagine This Scenario:

You have a bunch of jobs, and you need to do them one by one. Each job takes **1 unit of time** (let's say 1 hour) to finish. But there's a twist: each job has a **deadline** — which means you must finish the job by a certain time, or you won’t earn any money for that job.

Now, you want to **maximize your earnings** by finishing as many jobs as possible, but you can’t do two jobs at the same time. You need to choose the best way to schedule them, so you can do the most profitable ones first!

### The Problem:
1. You have **N jobs**, and each job has:
    - **Profit** (how much money you get for completing it).
    - **Deadline** (the latest time by which you must finish the job).

2. Your goal is to **schedule jobs** so you finish as many of them as possible before their deadlines, and **maximize the profit** you make.

### Rules (or Constraints):
1. **Each job takes exactly 1 unit of time**.
2. **No two jobs can be done at the same time**. You can only do **one job at a time**.
3. **A job must be finished by its deadline**. If you don’t finish a job before its deadline, you get **no profit** for that job.
4. You want to **maximize your profit**, meaning you should choose jobs that give you the most money.

### Example with Simple Jobs:
Let’s look at a few jobs, and I’ll show you how we can solve the problem.

Imagine you have 5 jobs like this:

| Job | Profit | Deadline |
|-----|--------|----------|
| A   | 100    | 2        |
| B   | 19     | 1        |
| C   | 27     | 2        |
| D   | 25     | 1        |
| E   | 15     | 3        |

**The goal** is to schedule jobs so you make the most money by finishing as many jobs as possible before their deadlines.

---
## STEPS TAKEN.
### Step 1: **Sort the Jobs by Profit**
You first want to **pick the jobs that pay the most**. So, you sort the jobs by their profit in **descending order** (from highest to lowest):

Sorted jobs by profit:
- Job **A**: Profit = 100, Deadline = 2
- Job **C**: Profit = 27, Deadline = 2
- Job **D**: Profit = 25, Deadline = 1
- Job **B**: Profit = 19, Deadline = 1
- Job **E**: Profit = 15, Deadline = 3

---

### Step 2: **Schedule the Jobs One by One**

#### Job **A** (Profit = 100, Deadline = 2):
- You can only do this job before time 2.
- The latest time you can do it is time slot **2**.
- **You choose time slot 2** for Job A. It’s a good choice because it gives you the most profit.

#### Job **C** (Profit = 27, Deadline = 2):
- Job C also needs to be finished by time 2.
- But **time 2** is already taken by Job A.
- So, you try to schedule Job C at time **1**.
- **Time slot 1** is free, so you schedule Job C there.

#### Job **D** (Profit = 25, Deadline = 1):
- Job D needs to be done by time 1, but **time 1** is already taken by Job C.
- **You can't do Job D**, so you skip it.

#### Job **B** (Profit = 19, Deadline = 1):
- Job B also needs to be done by time 1.
- But **time 1** is already taken by Job C.
- **You can't do Job B**, so you skip it too.

#### Job **E** (Profit = 15, Deadline = 3):
- Job E needs to be done by time 3.
- **Time slot 3** is free, so you schedule Job E at time 3.

---

### Final Schedule:
You were able to schedule the following jobs:
- **Job A** at time 2 (Profit = 100)
- **Job C** at time 1 (Profit = 27)
- **Job E** at time 3 (Profit = 15)

The **total profit** you earned is **100 + 27 + 15 = 142**.
