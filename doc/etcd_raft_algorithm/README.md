[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# etcd, raft algorithm


**Disclaimer**.

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
- [raft algorithm: log replication](#raft-algorithm-log-replication)
- [raft algorithm: log consistency](#raft-algorithm-log-consistency)
- [raft algorithm: safety](#raft-algorithm-safety)
- [raft algorithm: membership changes](#raft-algorithm-membership-changes)
- [raft algorithm: leader changes](#raft-algorithm-leader-changes)
- [raft algorithm: summary](#raft-algorithm-summary)
- [`etcd` internals: RPC between machines](#etcd-internals-rpc-between-machines)
- [`etcd` internals: leader election](#etcd-internals-leader-election)
- [`etcd` internals: log replication](#etcd-internals-log-replication)
- [`etcd` internals: log consistency](#etcd-internals-log-consistency)
- [`etcd` internals: safety](#raft-algorithm-safety)
- [`etcd` internals: membership changes](#etcd-internals-membership-changes)
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
- [Protocol Buffers](https://en.wikipedia.org/wiki/Protocol_Buffers)

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

*Byzantine failures* are far more disruptive because they affect
*agreement*, *consensus* services in distributed computing systems.
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
A *state machine* can be any program or application that takes inputs
and returns outputs. **Replicated state machines** in a distributed system
**compute identical copies** of the same state, so that even if some servers
are down, other **state machines can keep running**. A **replicated state
machine** is usually **implemented by replicating logs identically across
the servers**. The goal of *Raft* algorithm is to **keep the replicated logs
consistent**.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### raft algorithm: terminology

- **`state machine`**: Any program or application that *takes input* and
  *returns output*.
- **`replicated state machines`**: State machines that are distributed on a
  collection of servers and compute identical copies of the same state.
  Even when some of the servers are down, other state machines can keep
  running. Typically **replicated state machines are implemented by
  replicating log entries identically on the collection of servers**.
- **`log`**: A log contains a list of commands, so that *state machines*
  can apply those log entries *when it is safe to do so*. A log entry is the
  primary work unit of *Raft algorithm*.
- **`log commit`**: A leader `commits` a log entry only after the leader has
  replicated the entry on a majority of servers in the cluster. Such log entry
  is considered safe to be applied to state machines. `Commit` also includes
  preceding entries, such as the ones from previous leaders. This is done by
  the leader keeping track of the highest index to commit.
- **`leader`**: *Raft algorithm* first elects a `leader` to handle
  client requests and replicate log entries to other servers(followers).
  Once logs are replicated, the leader tells when to apply log entries to
  their state machines. When a leader fails or gets disconnected from
  other servers, then the algorithm elects a new leader. In normal operation,
  there is **exactly only one leader**. A leader sends periodic heartbeat
  messages to other servers to maintain its authority.
- **`client`**: A client *requests* that **a leader append a new log entry**.
  Then the leader writes and replicates them to its followers. A client does
  **not need to know which machine is the leader**, sending write requests to
  any machine in the cluster. If a client sends a request to a follower, it
  redirects to the current leader (Raft paper §5.1). A leader sends out
  `AppendEntries` RPCs with its `leaderId` to other servers, so that a
  follower knows where to redirect client requests.
- **`follower`**: A follower is completely passive, issuing no RPCs and only
  responds to incoming RPCs from candidates and leaders. All servers start as
  followers. If a follower receives no communication(heartbeat), it becomes a
  candidate to start an election. 
- **`candidate`**: A candidate is used to elect a new leader. It's a server
  state between `follower` and `leader`. If a candidate receives votes from
  the majority of servers, it becomes the new `leader`.
- **`term`**: *Raft* divides time into `terms` of arbitrary duration, indexed
  with consecutive integers. Each term begins with an *election*. And if the
  election ends with no leader (split vote), it creates a new `term`. *Raft*
  ensures that each `term` has at most one leader in the given `term`. `Term`
  index is also used to detect obsolete information. Servers always sync with
  biggest `term` number(index), and any *server with stale `term`* number
  **reverts back to `follower` state**, and any requests from such servers
  are rejected.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### raft algorithm: leader election

*Raft* servers communicate through remote procedure calls (RPCs).
The basic Raft algorithm requires only two types of RPCs:

- `RequestVote` RPCs, issued by candidates during elections.
- `AppendEntries` RPCs, issued by leaders:
  - **to replicate log entries**.
  - **to send out heartbeat messages**.

<br>
This is the summary of
[§5.2 Leader election](http://ramcloud.stanford.edu/raft.pdf):

1. A server begins as a `follower` with a new `term`.
2. A `leader` sends periodic heartbeat messages to its followers in order to
   maintain its authority.
3. `Followers` wait for **randomized** `election timeout` until they receive
   heartbeats from a valid leader of equal or greater `term` number.
4. If **`election times out`** and `followers` receive no such communication
   from a leader, then it assumes there is no current leader in the cluster,
   and it begins a new `election` and the **`follower` becomes the
   `candidate`**, **incrementing its current `term` index(number)**,
   and **resetting its `election timer`**.
5. **`Candidate` first votes for itself** and **sends `RequestVote` RPCs**
   to other servers(followers). `RequestVote` RPC includes `candidate`'s log
   information so that a follower can deny its voting if the follower's log
   is more up-to-date than `candidate`'s. *Raft* compares log's index and
   `term` number to determine the up-to-date log.
6. Then the **`candiate`** either:
	- **_becomes the leader_** by *winning the election* when it gets **majority
	  of votes**. Then it must send out heartbeats to others
	  to establish itself as a leader.
	- **_reverts back to a follower_** when it receives a RPC from a **valid
	  leader**. A valid heartbeat message must have a `term` number that is
	  equal to or greater than `candidate`'s. RPCs with lower `term`
	  numbers are rejected. A leader **only appends to log**. Therefore,
	  future-leader will have the **most complete** log among electing
	  majority: a leader's log is the truth and a leader will eventually
	  make followers' logs identical to the leader's.
	- **_starts a new election and increments its current `term` number_**
	  **when votes are split with no winner** That is, its **`election
	  times out` receiving no heartbeat message from a valid leader, so
	  it retries**. *Raft* randomizes `election timeout` duration to avoid
	  split votes. It remains as a `candidate`.

<br>
And server states in *Raft*:

![raft_server_state](img/raft_server_state.png)


<br>
Here's how election works:

![raft_leader_election_00](img/raft_leader_election_00.png)
![raft_leader_election_01](img/raft_leader_election_01.png)
![raft_leader_election_02](img/raft_leader_election_02.png)
![raft_leader_election_03](img/raft_leader_election_03.png)
![raft_leader_election_04](img/raft_leader_election_04.png)
![raft_leader_election_05](img/raft_leader_election_05.png)
![raft_leader_election_06](img/raft_leader_election_06.png)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### raft algorithm: log replication

*Raft* servers communicate through remote procedure calls (RPCs).
The basic Raft algorithm requires only two types of RPCs:

- `RequestVote` RPCs, issued by candidates during elections.
- `AppendEntries` RPCs, issued by leaders:
  - **to replicate log entries**.
  - **to send out heartbeat messages**.

<br>
This is the summary of
[§5.3 Log replication](http://ramcloud.stanford.edu/raft.pdf).

<br>
Once the cluster has elected a leader, it starts receiving `client` requests.

1. A `client` request contains a command for replicated state machines.
2. The leader **appends** the command to its log as a **new entry**.
3. The leader **replicates** the *log entry* to its `followers`,
   with `AppendEntries` RPC. The leader keeps sending those RPCs until
   all followers eventually store all log entries. Each `AppendEntries` RPC
   contains `term` number of the leader, its log entry index, its `leaderId`.
4. When the log entry has been *safely replicated* on a majority of servers,
   the leader applies the entry to its state machine. What its means by
   `apply the entry to state machine` is *execute the command in the log
   entry*.
5. Once a log entry has been *safely replicated* and applied to the leader's
   state machine, the leader `commits` the log.
6. After the leader applies the log entry to its state machine (`committed`),
   it returns the execution result to the client and tell `followers` that
   the log entry is committed.

Note that **a leader only appends**, never overwriting nor deleting entries.

<br>
Here's how log replication works:

![raft_log_replication_00](img/raft_log_replication_00.png)
![raft_log_replication_01](img/raft_log_replication_01.png)
![raft_log_replication_02](img/raft_log_replication_02.png)
![raft_log_replication_03](img/raft_log_replication_03.png)
![raft_log_replication_04](img/raft_log_replication_04.png)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### raft algorithm: log consistency

![raft_log_consistency_00](img/raft_log_consistency_00.png)
<br>
![raft_log_consistency_01](img/raft_log_consistency_01.png)
<br>
![raft_log_consistency_02](img/raft_log_consistency_02.png)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>


 





#### raft algorithm: safety

This is the summary of
[§5.4 Safety](http://ramcloud.stanford.edu/raft.pdf).

*Raft* algorithm's **_safety_** is ensured when:

1. each state machine executes exactly the same commands in the same order.
2. a leader for any given term contains all of log entries committed
   in previous terms.

<br>
And to guarantee the safety requirement:

- A leader never overwrites nor deletes log entries.
- Only leader log entries can be committed.
- Entries must be committed before applying to a state machine.
- Elect the candidate with most complete log.

In order to commit entries from previous terms, a leader
overwrites followers' log to handle the conflict entries.
Leader first *finds the latest log entry* matching with leader's entry,
and then *deletes follower's extraneous entries after that index*.
This is done by `AppendEntries` RPC.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr> 







#### raft algorithm: membership changes

Not ready yet. I am researching right now.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>







#### raft algorithm: leader changes

Not ready yet. I am researching right now.

<br>
If a `follower` or `candidate` crashes, `RequestVote` and `AppendEntries` RPCs
will fail. *Raft* simply keeps retrying until they succeed. *Raft* RPCs are
*idempotent*, which means calling multiple times has no additional effects.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>







#### raft algorithm: summary

Here's pseudo-code that summarizes *Raft* algorithm:

```go
// ServerState contains persistent, volatile states of
// all servers(follower, candidate, leader).
type ServerState struct {

	// Persistent state on all servers.
	// This should be updated on stable storage
	// before responding to RPCs.
	//
	// currentTerm is the latest term that server
	// has been in. A server begins with currentTerm 0,
	// and it increases monotonically
	currentTerm int
	//
	// votedFor is the candidateId that received vote
	// in current term, from this server.
	votedFor string
	//
	// logs is a list of log entries, of which contains
	// command for state machine, and the term when the
	// entry was received by a leader.
	logs []string

	// Volatile state on all servers.
	//
	// commitIndex is the index of latest(or highest)
	// committed log entry. It starts with 0 and increases
	// monotonically.
	commitIndex int
	//
	// lastApplied is the index of the highest log entry
	// applied to state machine. It is the index of last
	// executed command. It starts with 0 and increases
	// monotonically.
	lastApplied int

	// Volatile state on leaders.
	// This must be reinitialized after election.
	//
	// serverToNextIndex maps serverID to the index of
	// next log entry to send to that server.
	// NextIndex gets initialized with the last leader
	// log index + 1.
	serverToNextIndex map[string]int
	//
	// serverToMatchIndex maps serverID to the index of
	// highest log entry that has been replicated on that server.
	// The MatchIndex begins with 0, increases monotonically.
	serverToMatchIndex map[string]int
}

```


[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>







#### `etcd` internals: RPC between machines

*Raft* servers communicate through remote procedure calls (RPCs).
The basic Raft algorithm requires only two types of RPCs:

- `RequestVote` RPCs, issued by candidates during elections.
- `AppendEntries` RPCs, issued by leaders:
  - **to replicate log entries**.
  - **to send out heartbeat messages**.

<br>
`etcd` uses [`Protocol Buffers`](https://developers.google.com/protocol-buffers/docs/overview?hl=en)
for inter-machine communication of structured data. Below are some of
the related, core packages:

- [`raft`](http://godoc.org/github.com/coreos/etcd/raft): implements of the raft consensus algorithm.
- [`raft/raftpb`](http://godoc.org/github.com/coreos/etcd/raft/raftpb):
- [`rafthttp`](http://godoc.org/github.com/coreos/etcd/rafthttp):
- [`etcdserver`](http://godoc.org/github.com/coreos/etcd/etcdserver): connects servers in the cluster, using `HTTP`.
- [`etcdserver/etcdhttp`](http://godoc.org/github.com/coreos/etcd/etcdserver/etcdhttp):
- [`etcdserver/etcdserverpb`](http://godoc.org/github.com/coreos/etcd/etcdserver/etcdserverpb):
- [`discovery`](http://godoc.org/github.com/coreos/etcd/discovery):
- [`client`](http://godoc.org/github.com/coreos/etcd/client):

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: leader election

Not ready yet. I am researching right now.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: log replication

Not ready yet. I am researching right now.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: safety

Not ready yet. I am researching right now.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: membership changes

Not ready yet. I am researching right now.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### `etcd` internals: leader changes

Not ready yet. I am researching right now.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>
