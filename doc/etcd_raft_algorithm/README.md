[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# etcd, raft algorithm

- [Reference](#reference)
- [consensus algorithm](#consensus-algorithm)
- [raft algorithm: introduction](#raft-algorithm-introduction)
- [raft algorithm: terminology](#raft-algorithm-terminology)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### Reference

**_DISCLAIMER_:** This is **my** understanding of the **raft algorithm**.<br>
I may say things out of ignorance. Please refer to the readings below.

- [Consensus (computer science)](https://en.wikipedia.org/wiki/Consensus_(computer_science))
- [CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem)
- [*Raft paper by Diego Ongaro and John Ousterhout*](http://ramcloud.stanford.edu/raft.pdf)
- [The Raft Consensus Algorithm](https://raft.github.io/)
- [Raft (computer science)](https://en.wikipedia.org/wiki/Raft_(computer_science))
- [Raft lecture (Raft user study)](https://www.youtube.com/watch?v=YbZ3zDzDnrw)
- [coreos/etcd](https://github.com/coreos/etcd)
- [Raft Protocol Overview by Consul](https://www.consul.io/docs/internals/consensus.html)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### consensus algorithm

> A fundamental problem in **distributed computing** is to achieve overall **system
> reliability** in the presence of a number of *faulty processes*. This often
> requires processes to agree on some data value that is needed during
> computation. Examples of applications of **consensus** include **whether to commit
> a transaction to a database, agreeing on the identity of a leader, state
> machine replication, and atomic broadcasts.**
>
> [*Consensus*](https://en.wikipedia.org/wiki/Consensus_(computer_science))
> *by Wikipedia*

A process can fail either from a *crash failure* or [*Byzantine
failure*](https://en.wikipedia.org/wiki/Byzantine_failure):
- *crash failure* occurs when the process abruptly stops.
- *Byzantine failure* is the loss of the process presenting different symptoms
  to different observers (*Byzantine fault*).

*Byzantine failures* are far more disruptive because they affect the
*agreement* or *consensus* services in distributed computing systems.
Ideally every process must agree on the same value. If a distributed system
loses one of its communications, it can result in data inconsistency.
A consensus algorithm must be resilient to these failures in order to
guarantee the correctness.

<br>
An ultimate **consensus algorithm** would achieve:
- **_consistency_**.
- **_availability_**.
- **_partition tolerance_**.

[CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem) states that
it is impossible that a distributed computer system simultaneously satisfies
them all.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### raft algorithm: introduction

> **Raft is a consensus algorithm for managing a replicated
> log.** It produces a result equivalent to (multi-)Paxos, and
> it is as efficient as Paxos, but its structure is different
> from Paxos; this makes Raft more understandable than
> Paxos and also provides a better foundation for building
> practical systems.
>
> [*In Search of an Understandable Consensus
> Algorithm*](http://ramcloud.stanford.edu/raft.pdf)
> *by Diego Ongaro and John Ousterhout*

One way to make your program reliable is:
- execute the program in a collection of machines (distributed system).
- ensure that they all get executed exactly the same way (consistency).

This is the definition of **replicated state machine** in the paper.
*State machine* can be any program or application that takes inputs
and returns outputs. **Replicated state machines** in a distributed system
compute identical copies of the same state, so that even if some servers are
down, other **state machines** can keep running. **Replicated state machines**
is usually implemented by replicating logs **identically** across the servers.
And **keeping the replicated logs consistent** is the overall goal of **raft
algorithm**.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### raft algorithm: terminology

- **_state machine_**: any program or application with input and output.
- **_replicated state machines_**: state machines that are distributed on a
  collection of servers and compute identical copies of the same state: in
  doing so, when some of the servers are down, other state machines can keep
  running. Typically, replicated state machines are implemented **replicating
  logs of commands** identically on the collection of servers.
- **_log_**: Logs contains the list of commands, so that *state machines*
  can apply those log entries *when it is safe to do so*. A log entry is the
  primary work unit of *Raft algorithm*.
- **_leader_**: *Raft algorithm* achieves *consensus* **by first electing a
  leader** that accepts log entries from clients, and replicates them on other
  servers(followers) telling them when it is safe to apply log entries to their
  state machines. A leader can fail or get disconnected from other servers.
  Then the algorithm elects a new leader.
- **_client_**: A client requests that **a leader append a new log entry**.
  Then the leader writes and replicates them to its followers. A client does
  **not need to know which machine is the leader**, sending write requests to
  any machine in the cluster. If a client requests to a follower, the request
  will be returned and the client is notified what the current leader is.
- **_follower_**: 
- **_candidate_**:


[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### 

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### 

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### 

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### 

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>
