[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# etcd, raft algorithm


*Disclaimer*.

This is high-level overview of *raft algorithm* to understand the internals of
[`coreos/etcd`](https://github.com/coreos/etcd). You don't need know these
details to use `etcd`. And I may say things out of ignorance.
Please refer to [Reference](#reference) below.

<br>

- [Reference](#reference)
- [consensus algorithm](#consensus-algorithm)
- [raft algorithm: introduction](#raft-algorithm-introduction)
- [raft algorithm: terminology](#raft-algorithm-terminology)
- [raft algorithm: leader election](#raft-algorithm-leader-election)
- [`etcd` internals: RPC](#etcd-internals-rpc)
- [`etcd` internals: leader election](#etcd-internals-leader-election)
- [raft algorithm: log replication](#raft-algorithm-log-replication)
- [`etcd` internals: log replication](#etcd-internals-log-replication)
- [raft algorithm: safety](#raft-algorithm-safety)
- [`etcd` internals: safety](#raft-algorithm-safety)
- [raft algorithm: leader changes](#raft-algorithm-leader-changes)
- [`etcd` internals: leader changes](#etcd-internals-leader-changes)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### Reference

- [The Raft Consensus Algorithm](https://raft.github.io/)
- [*Raft paper by Diego Ongaro and John Ousterhout*](http://ramcloud.stanford.edu/raft.pdf)
- [Consensus (computer science)](https://en.wikipedia.org/wiki/Consensus_(computer_science))
- [CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem)
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
is usually **implemented by replicating logs identically across the servers**.
And **keeping the replicated logs consistent** is the overall goal of **raft
algorithm**.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### raft algorithm: terminology

- **`state machine`**: Any program or application that *takes input* and
  *returns output*.
- **`replicated state machines`**: State machines are distributed on a
  collection of servers and compute identical copies of the same state:
  those state machines are *replicated state machines*. In doing so, even
  when some of the servers are down, other state machines can keep running.
  Typically, **replicated state machines are implemented replicating
  logs of commands identically on the collection of servers**.
- **`log`**: A log contains the list of commands, so that *state machines*
  can apply those log entries *when it is safe to do so*. A log entry is the
  primary work unit of *Raft algorithm*.
- **`log commit`**: A leader `commit`s a log entry only after the leader has
  replicated the entry on a majority of servers in a cluster. Such log entry
  is considered safe to be applied to state machines. `Commit` includes the
  preceding entries, such as the ones from previous leaders. This is done by
  the leader keeping track of the highest index to commit.
- **`leader`**: *Raft algorithm* achieves *consensus* **by first electing a
  leader** that accepts log entries from clients, and replicates them on other
  servers(followers) telling them when it is safe to apply log entries to their
  state machines. When a leader fails or gets disconnected from other servers,
  then the algorithm elects a new leader. In normal operation, there is
  **exactly only one leader** and all of the other servers are followers.
  A leader must keep sending heartbeats to maintain its authority.
  A leader handles all requests from clients.
- **`client`**: A client requests that **a leader append a new log entry**.
  Then the leader writes and replicates them to its followers. A client does
  **not need to know which machine is the leader**, sending write requests to
  any machine in the cluster. If a client sends request to followers, the
  followers redirects to the current leader (Raft paper §5.1).
- **`follower`**: A follower is completely passive, issuing no RPCs and only
  responds to incoming RPCs from candidates and leaders. All servers start as
  followers. If a follower receives no communication(heartbeat), it becomes a
  candidate to start an election. 
- **`candidate`**: A candidate is used to elect a new leader. It's a state
  between `follower` and `leader`. If a candidate receives votes from the
  majority of a cluster, it becomes the new leader.
- **`term`**: *Raft* divides time into `term`s of arbitrary duration, indexed
  with consecutive integers. Each term begins with an *election*. And if the
  election ends with no leader (split vote), it creates a new `term`. *Raft*
  ensures that each `term` has at most one leader in the given `term`. `Term`
  index is also used to detect obsolete information. Servers always sync with
  biggest `term` number(index), and any server with stale `term` number reverts
  back to `follower` state, and any requests from such servers are rejected.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### raft algorithm: leader election

1. *Raft* starts a server as a `follower`, with the new `term`.
2. A `leader` must send periodic heartbeat messages to its followers in order to
   maintain its authority.
3. `Followers` wait for **randomized** `election timeout` until they receive
   heartbeat messages from a valid leader, with equal or greater `term` number.
4. If **`election times out`** and `followers` receive no such communication
   from a leader, then it assumes there is no current leader in the cluster,
   and it begins a new `election` and the **`follower` becomes the
   `candidate`**, **incrementing its current `term` index(number)**,
   and **resetting its `election timer`**.
5. `Candidate` first votes for itself and sends `RequestVote` RPCs to other
   servers (followers). A follower as a voter deny its vote if its own log
   is more up-to-date than `candidate`'s. 
6. Then the `candiate` either:
	- *becomes the leader* by *winning the election* when it gets **majority
	  of votes**. Then it must send out the heartbeat messages to others
	  to establish itself as a leader.
	- *reverts back to a follower* when it receives a RPC from a **valid
	  leader**. A valid heartbeat message must have a `term` number that is
	  equal to or greater than `candidate`'s. The RPCs with lower `term`
	  numbers are rejected. A leader **only appends to log**. Therefore,
	  future-leader will have the **most complete** log among electing
	  majority: a leader's log is the truth and a leader will eventually
	  make followers' logs identical to the leader's.
	- *starts a new election and increments its current `term` number* **when
	  votes are split with no winner** That is, its **`election times out`
	  receiving no heartbeat message from a valid leader, so it retries**.
	  *Raft* randomizes `election timeout` duration to avoid split votes.
	  It remains as a `candidate`.

<br>
Here's how it works:

![leader_election_00](img/leader_election_00.png)
![leader_election_01](img/leader_election_01.png)
![leader_election_02](img/leader_election_02.png)
![leader_election_03](img/leader_election_03.png)
![leader_election_04](img/leader_election_04.png)
![leader_election_05](img/leader_election_05.png)
![leader_election_06](img/leader_election_06.png)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: RPC

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: leader election

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### raft algorithm: log replication


[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: log replication

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### raft algorithm: safety


[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: safety


[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### raft algorithm: leader changes


[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: leader changes

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>
