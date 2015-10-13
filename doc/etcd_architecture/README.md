[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# etcd architecture

<br>

- [Reference](#reference)
- [distributed systems, consensus algorithm](#distributed-systems-consensus-algorithm)
- [raft algorithm: introduction](#raft-algorithm-introduction)

[↑ top](#etcd-architecture)
<br><br><br><br>
<hr>









#### Reference

- [The Raft Consensus Algorithm](https://raft.github.io/)
- [Raft lecture (Raft user study)](https://www.youtube.com/watch?v=YbZ3zDzDnrw)
- [`coreos/etcd`](https://github.com/coreos/etcd)

[↑ top](#etcd-architecture)
<br><br><br><br>
<hr>







#### distributed systems, consensus algorithm

> In distributed computing, a problem is divided into many tasks, each of which
> is solved by one or more computers, which communicate with each other by
> message passing.
>
> A distributed system may have a common goal, such as solving a large
> computational problem. Alternatively, each computer may have its own user
> with individual needs, and the **purpose of the distributed system** is to
> **coordinate the use of shared resources** or **provide communication
> services** to the users.
> 
> [*Distributed computing*](https://en.wikipedia.org/wiki/Distributed_computing)
> *by Wikipedia*

- *In parallel computing*, multiple processors may have access to a globally 
  **shared memory to exchange data between processors**.
- *In distributed computing*, each processor has its **own private memory**
  exchanging data by **passing messages between processors**.

<br>
> A fundamental problem in **distributed computing** is to achieve overall **system
> reliability** in the presence of a number of *faulty processes*. This often
> requires processes to agree on some data value that is needed during
> computation. Examples of applications of **consensus** include **whether to commit
> a transaction to a database, agreeing on the identity of a leader, state
> machine replication, and atomic broadcasts.**
>
> [*Consensus*](https://en.wikipedia.org/wiki/Consensus_(computer_science))
> *by Wikipedia*

A process can fail either from a *crash* or a *loss of the process presenting
different symptoms to different observers*
([*Byzantine fault*](https://en.wikipedia.org/wiki/Byzantine_failure)).
The consensus algorithm handles these kinds of *faulty processes* in
distributed computing systems, and keeps data consistent even when it loses one
of its communications.

<br>
One of the most important properties of distributed computing is
*linearizability*:

> In concurrent programming, an operation (or set of operations) is atomic,
> **linearizable**, indivisible or uninterruptible if it appears to the rest
> of the system to occur instantaneously. **Atomicity is a guarantee of
> isolation from concurrent processes**. Additionally, **atomic operations**
> commonly have a **succeed-or-fail** definition—they either successfully
> change the state of the system, or have no apparent effect.
>
> [*Linearizability*](https://en.wikipedia.org/wiki/Linearizability)
> *by Wikipedia*

<br>
> **Linearizability** provides **the illusion that each operation applied by
> concurrent processes takes effect instantaneously at some point between
> its invocation and its response**, implying that the meaning of a concurrent
> object's operations can be given by pre- and post-conditions.
>
> [*Linearizability: A Correct Condition for Concurrent
> Objects*](https://cs.brown.edu/~mph/HerlihyW90/p463-herlihy.pdf)

<br>
In other words, once an operation finishes, every other machine in the cluster
must see it. While operations are concurrent in distributed system, every
machine sees each operation in the same linear order. Think of
*linearizability* as *atomic consistency* with an atomic operation where a set
of operations occur atomically with respect to other parts of the system.

*Linearizability* is *local*. Since an operation on each object is linearized,
all operations in the system are linearizable. *Linearizability* is
**non-blocking**, since a pending invocation does not need to wait for other
pending invocation to complete, or an object with a pending operation does
not block the total operation, which makes it suitable for concurrent and
real-time systems.

**Sequential consistency** is another consistency model, *weaker than
linearizability*. Each operation can take effect **before its invocation**
or **after its response** (*not necessarily between its invocation and its
response as in linearizability*). And it is still considered *consistent*.

> Many caches also behave like sequentially consistent systems. If I write a
> tweet on Twitter, or post to Facebook, it **takes time to percolate through
> layers of caching systems**. **Different users will see my message at
> different times–but each user will see my operations in order**. Once seen,
> a post shouldn’t disappear. **If I write multiple comments, they’ll become
> visible sequentially, not out of order**.
>
> [*Sequential consistency*](https://aphyr.com/posts/313-strong-consistency-models)
> *by aphyr*

<br>
The goal of `etcd` as a **distributed consistent key-value store**
is **sequential consistency**:

> `etcd` tries to ensure **sequential consistency**, which means each replica
> have the same command execution ordering.
>
> [*Xiang Li*](https://github.com/coreos/etcd/issues/741)

[↑ top](#etcd-architecture)
<br><br><br><br>
<hr>








#### raft algorithm: introduction

To make your program reliable, you would:
- execute program in a collection of machines (distributed system).
- ensure that they all run exactly the same way (consistency).

This is the definition of **replicated state machine**. And a *state machine*
can be any program or application with inputs and outputs. *Each replicated
state machine* computes identical copy with a same state, which means when
some servers are down, other state machines can keep running. A distributed
system usually implements *replicated state machines* by **replicating logs
identically across cluster**. And the goal of *Raft algorithm* is to **keep
those replicated logs consistent**.

> **Raft is a consensus algorithm for managing a replicated
> log.** It produces a result equivalent to (multi-)Paxos, and
> it is as efficient as Paxos, but its structure is different
> from Paxos; this makes Raft more understandable than
> Paxos and also provides a better foundation for building
> practical systems.
>
> [*In Search of an Understandable Consensus Algorithm*](http://ramcloud.stanford.edu/raft.pdf)
> *by Diego Ongaro and John Ousterhout*

<br>
Raft nodes(servers) must be one of three states: `follower`, `candidate`, or
`leader`. A `leader` sends periodic heartbeat messages to its `followers`
to maintain its authority. In normal operation, there is **exactly only one
`leader`** for each term. All servers start as a `follower`, and the
`follower` becomes a `candidate` when there is no current `leader` and starts
an election. If a `candidate` receives the majority of votes, it becomes a
`leader`. The `leader` then accepts new log entries from clients and replicates
those log entries to its `followers`.

*Raft* inter-server communication is done by remote procedure calls
(RPCs). The basic Raft algorithm requires only two types of RPCs
(later `InstallSnapshot` RPC added):

- `RequestVote` RPCs, issued by `candidates` during elections.
- `AppendEntries` RPCs, issued by `leaders`:
  - **to replicate log entries**.
  - **to send out heartbeat messages**.

**Servers retry RPCs** *when they do not receive a response in time*,
and **send RPCs in parallel** *for best performance*.

A `log entry` is considered *safely replicated* when the leader has replicated
it on the **quorum of its followers**. Once `log entry` has been *safely
replicated* on a majority of servers, it is considered **safe to be applied**
to its state machine. And such `log entry` is *called* **committed**. Then
**`leader`** **applies committed entry to its state machine**. `Applying
committed entry to state machine` means *executing the command in the log
entry*. Again, `leader` attempts to **replicate a log entry on the quorum
of its followers**. Once they are replicated on the majority of its followers,
it is **safely replicated**. Therefore it is **safe to be applied**. Then the
`leader` **commits that log entry**. *Raft* guarantees that such entries are
committed in a durable storage, and that they will eventually be
applied *(executed)* by other available state machines. When a `log entry` is
committed, it is safe to be applied. And for a `log entry` to be committed,
it only needs to be stored on the quorum of cluster. This means each `command`
can complete as soon as the majority of cluster has responded to a single
round of `AppendEntries` RPCs. In other words, the `leader` does not need to
wait for responses from every node.

Most critical case for performance is when a leader replicates log entries.
*Raft* algorithm minimizes the number of messages by requiring a single
round-trip request only to half of the cluster. *Raft* also has mechanism
to discard obsolete information accumulated in the log. Since system
cannot handle infinitely growing logs, *Raft* uses `snapshot` to save the
state of the entire system on a stable storage, so that logs stored
up to the `snapshot` point can be discarded.

[↑ top](#etcd-architecture)
<br><br><br><br>
<hr>

