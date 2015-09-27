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
- [consensus algorithm vs. replication](#consensus-algorithm-vs-replication)
- [raft algorithm: introduction](#raft-algorithm-introduction)
- [raft algorithm: terminology](#raft-algorithm-terminology)
- [raft algorithm: leader election](#raft-algorithm-leader-election)
- [raft algorithm: log replication](#raft-algorithm-log-replication)
- [raft algorithm: log consistency](#raft-algorithm-log-consistency)
- [raft algorithm: safety](#raft-algorithm-safety)
- [raft algorithm: follower and candidate crashes](#raft-algorithm-follower-and-candidate-crashes)
- [raft algorithm: client interaction](#raft-algorithm-client-interaction)
- [raft algorithm: log compaction](#raft-algorithm-log-compaction)
- [**raft algorithm: summary**](#raft-algorithm-summary)
- [`etcd` internals: RPC between machines](#etcd-internals-rpc-between-machines)
  - [**`raft`**](#raft)
  - [**`etcdserver`**](#etcdserver)

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
- [`coreos/etcd`](https://github.com/coreos/etcd)
- [Raft Protocol Overview by Consul](https://www.consul.io/docs/internals/consensus.html)
- [Protocol Buffers](https://en.wikipedia.org/wiki/Protocol_Buffers)
- [Protocol Buffers](https://en.wikipedia.org/wiki/Protocol_Buffers)
- [`gyuho/go-fuzz-etcd`](https://github.com/gyuho/go-fuzz-etcd)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### consensus algorithm vs. replication

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
A consensus algorithm must be resilient to these failures and prevent
data inconsistency even when it loses one of its communications.

<br>
An ultimate **consensus algorithm** should achieve:
- **_consistency_**.
- **_availability_**.
- **_partition tolerance_**.

[CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem) states that
it is impossible that a distributed computer system simultaneously satisfies
them all. 

<br>
> In concurrent programming, an operation (or set of operations) is atomic,
> **linearizable**, indivisible or uninterruptible if it appears to the rest
> of the system to occur instantaneously. Atomicity is a guarantee of isolation
> from concurrent processes. Additionally, atomic operations commonly have a
> succeed-or-fail definition — they either successfully change the state of
> the system, or have no apparent effect.
>
> [*Linearizability*](https://en.wikipedia.org/wiki/Linearizability)
> *by Wikipedia*


database replication, elb, redis cluster

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### raft algorithm: introduction

To make your program reliable, you would:
- execute program in a collection of machines (distributed system).
- ensure that they all run exactly the same way (consistency).

This is the definition of **replicated state machine**.
And a *state machine* can be any program or application with inputs and
outputs. *Each replicated state machines* computes identical copy with a
same state, which means when some servers are down, other state machines 
can keep running. A distributed system usually implements *replicated state
machines* by **replicating logs identically across cluster**. And the goal
of *Raft* algorithm is to **keep those replicated logs consistent**.

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
  can apply those commands *when it is safe to do so*. A log entry is the
  primary work unit of *Raft algorithm*. A `command` completes when only a
  majority of cluster has responded to a single round of remote procedure
  calls, so that the minority of slow servers do not affect the overall
  performance.
- **`log commit`**: A leader `commits` a log entry only after the leader has
  replicated the entry on majority of servers in a cluster. Then such entry
  is safe to be applied to state machines. `commit` also includes preceding
  entries, such as the ones from previous leaders. This is done by the leader
  keeping track of the highest index to commit.
- **`leader`**: *Raft algorithm* first elects a `leader` that handles
  client requests and replicates log entries to followers.
  Once logs are replicated, `leader` tells followers when to apply log
  entries to their state machines. When a leader fails, *Raft* elects a
  new leader. In normal operation, there is **exactly only one leader**
  for each term. A leader sends periodic heartbeat messages to its followers
  to maintain its authority.
- **`client`**: A `client` requests a `leader` to append its new log entry.
  Then `leader` writes and replicates them to its followers. A client does
  **not need to know which machine is the leader**, sending write requests to
  any machine in the cluster. If a client sends a request to a follower, it
  redirects to the current leader (Raft paper §5.1). A leader sends out
  `AppendEntries` RPCs with its `leaderId` to other servers, so that a
  follower knows where to redirect client requests.
- **`follower`**: A `follower` is completely passive, issuing no RPCs and only
  responding to incoming RPCs from candidates or leaders. All servers start as
  followers. If a follower receives no communication or no heartbeat from a
  valid `leader`, it becomes a `candidate` and then starts an election.
- **`candidate`**: A server becomes a `candidate` from a `follower` when there
  is no current `leader`, so electing a new `leader`: it's a state between
  `follower` and `leader`. If a candidate receives votes from the majority
  of cluster, it becomes the new `leader`.
- **`term`**: *Raft* divides time into `terms` of arbitrary duration, indexed
  with consecutive integers. Each term begins with an *election*. And if the
  election ends with no leader *(split vote)*, it creates a new `term`. *Raft*
  ensures that each `term` has at most one leader for the given `term`. `term
  index` is used to detect obsolete data, only allowing the sync with servers
  with biggest `term number` (`index`). Any server with stale `term` must stay
  as or *revert back to* `follower` state. And requests from such servers are
  rejected.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### raft algorithm: leader election

*Raft* inter-server communication is done by remote procedure calls
(RPCs). The basic Raft algorithm requires only two types of RPCs:

- `RequestVote` RPCs, issued by `candidates` during elections.
- `AppendEntries` RPCs, issued by `leaders`:
  - **to replicate log entries**.
  - **to send out heartbeat messages**.

**Servers retry RPCs** *when they do not receive a response in time*,
and **send RPCs in parallel** *for best performance*.

<br>
Summary of
[§5.2 Leader election](http://ramcloud.stanford.edu/raft.pdf):

1. A server begins as a `follower` with a new `term`.
2. A `leader` sends periodic heartbeat messages to its followers in order to
   maintain its authority.
3. Each `follower` waits for heartbeats from a valid leader. The heartbeat
   message from `leader` is **valid** only when the `leader`'s `term number`
   is equal to or greater than `follower`'s `term number`.
4. Each `follower` has its own **randomized** `election timeout`, in order to
   avoid *split vote*. So in most cases, only a single `follower` server
   times out. If the follower has not received such heartbeats within timeout,
   it assumes that there is no current `leader` in cluster.
5. Then the `follower` starts a new `election` and becomes `candidate`.
6. When a `follower` becomes a `candidate`, it:
	- increments its `term number`.
	- resets its `election timeout`.
7. `candidate` first votes for itself.
8. `candidate` then **sends `RequestVote` RPCs** to other servers.
9. `RequestVote` RPC includes the information of `candidate`'s log
   (`index`, `term number`, etc).
10. `follower` denies voting if its log is more complete
   log than `candidate`.
11. Then **`candiate`** either:
	- **_becomes the leader_** by *winning the election* when it gets **majority
	  of votes**. Then it must send out heartbeats to others
	  to establish itself as a leader.
	- **_reverts back to a follower_** when it receives a RPC from a **valid
	  leader**. A valid `leader` must have `term number` that is
	  equal to or greater than `candidate`'s. RPCs with lower `term`
	  numbers are rejected. A leader **only appends to log**. Therefore,
	  future-leader will have **most complete** log among electing
	  majority: a leader's log is the truth and `leader` will eventually
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

*Raft* inter-server communication is done by remote procedure calls
(RPCs). The basic Raft algorithm requires only two types of RPCs:

- `RequestVote` RPCs, issued by `candidates` during elections.
- `AppendEntries` RPCs, issued by `leaders`:
  - **to replicate log entries**.
  - **to send out heartbeat messages**.

**Servers retry RPCs** *when they do not receive a response in time*,
and **send RPCs in parallel** *for best performance*.

<br>
Summary of
[§5.3 Log replication](http://ramcloud.stanford.edu/raft.pdf):

0. Once cluster has elected a leader, it starts receiving `client` requests.
1. A `client` request contains `command` for replicated state machines.
2. The leader **only appends** `command` to its log, never overwriting nor
   deleting its log entries.
3. The leader **replicates** the *log entry* to its `followers` with
   `AppendEntries` RPCs. The leader keeps sending those RPCs until
   all followers eventually store all log entries. Each `AppendEntries` RPC
   contains leader's `term number`, its log entry index, its `leaderId`
4. Again `command` can complete as soon as a majority of cluster has
   responded to a single round of `AppendEntries` RPCs. The leader does
   not need to wait for all servers' responses..
5. When `log entry` has been *safely replicated* on a majority of servers,
   the **`leader`** applies the entry to its state machine. What its means by
   `apply the entry to state machine` is *execute the command in the log
   entry*.
6. Then the `leader` returns the execution result to the client.
7. The log entry that has been *safely replicated* and *applied to `leader`'s
   state machine* is *called* **_committed_**.
8. Future `AppendEntries` RPCs from the `leader` has the highest index of
   `committed` log entry, so that `followers` could learn that a log entry is
   `committed`, and they can apply the entry to their local state machines as
   well. *Raft* ensures all committed entries are durable and eventually
   executed by all of available state machines.

<br>
Here's how log replication works:

![raft_log_replication_00](img/raft_log_replication_00.png)
![raft_log_replication_01](img/raft_log_replication_01.png)
![raft_log_replication_02](img/raft_log_replication_02.png)
![raft_log_replication_03](img/raft_log_replication_03.png)
![raft_log_replication_04](img/raft_log_replication_04.png)
![raft_log_replication_05](img/raft_log_replication_05.png)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>






#### raft algorithm: log consistency

![raft_log_consistency_00](img/raft_log_consistency_00.png)
<br>
![raft_log_consistency_01](img/raft_log_consistency_01.png)
<br>
![raft_log_consistency_02](img/raft_log_consistency_02.png)

<br>
**Log Matching Property**:
- If two entries in different logs have the same `index` and `term`, then they
  store the same `command`. This is because `leader` creates maximum one entry
  per `index` for the given `term`, and log entries never change their position
  in the log.
- If two entries in different logs have the same `index` and `term`, then the
  `logs` are identical in all preceding entries. This is guaranteed by
  `AppendEntries` RPC that does consistency check: the RPC contains `index`
  and `term` that immediately precede new entries, and if `follower` does not
  find an entry in its log with the same `index` and `term`, then it refuses
  the new entries. As a result, whenever `AppendEntries` returns successfully,
  `leader` learns that the `follower`'s log is identical to `leader`'s.

<br>
Then what if `AppendEntries` RPC fails? It fails when **`follower` does not
have the entry at immediate-preceding `index` and `term`**. And the `leader`
keeps sending `AppendEntries` RPCs until all `followers` eventually contain all
log entries in order to maintain the log consistency.

<br>
Then how does `leader` achieve this by keep sending `AppendEntries` RPCs?
<br>

A `follower` may have extraneous entries. The `leader` checks the log with its
latest log entry that two logs agree. And it deletes logs after that point.
A `follower` may be missing some entries. In this case, `leader` keeps sending
`AppendEntries` RPCs until it finds the matching entry.

![raft_log_matching_00](img/raft_log_matching_00.png)
![raft_log_matching_01](img/raft_log_matching_01.png)
![raft_log_matching_02](img/raft_log_matching_02.png)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>


 





#### raft algorithm: safety

Summary of
[§5.4 Safety](http://ramcloud.stanford.edu/raft.pdf):

<br>
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

When committing entries from previous terms, `leader` overwrites
`followers` with logs to handle the conflict entries. `leader`
first *finds the latest `follower` log entry* that matches with
leader's entry. And *deletes any extraneous entries after that
index*, in `follower`'s log. This is done by `AppendEntries` RPC.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>







#### raft algorithm: follower and candidate crashes

Summary of
[§5.5 Follower and candidate crashes](http://ramcloud.stanford.edu/raft.pdf):

> If a *follower or candidate crashes*, then future `RequestVote` and
> `AppendEntries` RPCs sent to it will fail. Raft handles these failures
> by **retrying indefinitely**; if the *crashed server restarts*, then
> the *RPC will complete successfully*. If a server crashes after
> completing an RPC but before responding, then it will receive the same RPC
> again after it restarts. **Raft RPCs are idempotent**, so this causes no
> harm. For example, if a follower receives an `AppendEntries` request
> that includes log entries already present in its log, it ignores those
> entries in the new request.
>
> [*§5.5 Follower and candidate crashes*](http://ramcloud.stanford.edu/raft.pdf)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### raft algorithm: client interaction

Summary of
[§8 Client interaction](http://ramcloud.stanford.edu/raft.pdf):

> If a client sends a request to a follower, it
> redirects to the current leader
>
> [*Raft paper §5.1*](http://ramcloud.stanford.edu/raft.pdf)

<br>
> In concurrent programming, an operation (or set of operations) is atomic,
> **linearizable**, indivisible or uninterruptible if it appears to the rest
> of the system to occur instantaneously. Atomicity is a guarantee of isolation
> from concurrent processes. Additionally, atomic operations commonly have a
> succeed-or-fail definition — they either successfully change the state of
> the system, or have no apparent effect.
>
> [*Linearizability*](https://en.wikipedia.org/wiki/Linearizability)
> *by Wikipedia*

<br>
Again, *clients send all requests to the Raft leader*. A client first connects
to a randomly chosen server. And if the server is not the leader, the server
will reject the requests from the client, and tell the client which server is
the most recent leader. `AppendEntries` RPCs to `followers` include `leader`'s
network address.

<br>
In *Raft*, each operation should appear to execute instantaneously, only once,
at some point between the call and response: **_linearizable semantics_**.
And changes in the cluster should appear in the same order to all of the
machines in the cluster. However, **if a leader crashes**, *client requests*
will **time out** and try again with randomly-chosen servers. If it were
after the leader had committed the log entry but before responding to
the client, the client tries the same command with the new leader:
*the command would get executed a second time*.

To prevent this, clients assign unique serial number to each command. Then the
state machine in a server stores the most recent serial number processes for
each client with the associated response. If the state machine finds a command
whose serial number has already been executed, it immediately returns the
response without re-executing the request the second time.

<br>
But still, *Raft* client does not need to know which machine is the `leader`.
Then how do we redirect clients to `leader`, which send requests to
`followers`? 

> Yes, a client can submit a write request to any machine in the cluster. What
> happens is the etcd machine you contact **creates a raft RPC** and **tags it
> with a unique id**. When the machine sees that unique id, tag come back as a
> raft log commit, then it returns a 200 OK to the client.
>
> [*Brandon
> Phillips*](https://groups.google.com/d/msg/coreos-user/et7-Lm0gQxo/jkeZPKo0uaEJ)

<br>
`TODO: find related code in etcd`

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### raft algorithm: log compaction

Summary of
[§7 Log compaction](http://ramcloud.stanford.edu/raft.pdf):

Most critical case for performance is when a leader replicates log entries.
*Raft* algorithm minimizes the number of messages by sending a single
round-trip request to half of the cluster. For stored logs, *Raft* has
mechanism to discard obsolete information accumulated in the log.

<br>
*Raft* uses `snapshot` to save the state of the entire system on a stable
storage, so that the log up to that `snapshot` point can be discarded.
Here's how `snapshot` works in *Raft* log:

![raft_log_compaction_00](img/raft_log_compaction_00.png)
![raft_log_compaction_01](img/raft_log_compaction_01.png)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>










#### **raft algorithm: summary**

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
	// commitIndex is the index of the latest(or highest)
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


// AppendEntries is invoked by a leader to:
//	- replicate log entries.
//	- send heartbeat messages.
//
//
// leaderTerm is the term number of the leader.
//
// leaderID is the serverID of the leader so that
// followers can redirect clients to the leader.
//
// prevLogIndex is the index of immediate preceding
// log entry (log entry before the new ones).
//
// prevLogTerm is the term number of prevLogIndex entry.
//
// entries contains log entries to store. It's empty for
// heartbeats. And it is a list so that a leader can be
// more efficient replicating log entries.
//
// leaderCommit is the leader's commitIndex.
//
// It returns the currentTerm and its success.
// Current terms are exchanged whenever servers
// communicate. It always updates its current term
// with a larger value, so that if a candidate or
// leader discovers that its term is out of date,
// it immediately reverts to follower state (§5.1).
//
// It returns true if the follower has the matching
// entry with the prevLogIndex and prevLogTerm.
//
func AppendEntries(
	targetServer      string,

	leaderTerm        int,
	leaderID          string,
	prevLogIndex      int,
	prevLogTerm       int,
	entries           []string,
	leaderCommitIndex int,
) (int, bool) {
	
	currentTerm := getCurrentTerm(targetServer)
	if leaderTerm < currentTerm {
		// so that the leader can update itself.
		// This means the leader is out of date.
		// The leader will revert back to a follower state.
		return currentTerm, false
	}

	logs := getLogs(targetServer)
	if v, exist := logs[prevLogIndex]; exist {
		if v['term'] != prevLogTerm {
			// An existing entry conflicts with a new entry.
			// The entry at prevLogIndex has a different term.
			// We need to delete the existing entry and
			// all the following entries.
			delete(logs, prevLogTerm)
			delete(logs, prevLogTerm + 1 ...)
		}
	} else {
		// the log entry with prevLogIndex
		// and prevLogTerm does not exist
		return currentTerm, false
	}

	// append any new entries that are not already in the log.
	for _, entry := range entries {
		if !findEntry(logs, entry) {
			logs = append(logs, entry)
		}
	}

	commitIndex := getCommitIndex(targetServer)
	if leaderCommitIndex > commitIndex {
		setCommitIndex(targetServer, min(leaderCommit, indexOfLastNewEntry))
	}

	return currentTerm, true
}

// RequestVote is invoked by candidates to gather votes.
//
// lastLogIndex is the index of candiate's last log entry.
// lastLogTerm is the term number of lastLogIndex.
//
// It returns currentTerm, and boolean value if the candidate
// has received vote or not.
//
func RequestVote(
	targetServer  string,

	candidateTerm int,
	candidateID   string,
	lastLogIndex  int,
	lastLogTerm   int,
) (int, bool) {

	currentTerm := getCurrentTerm(targetServer)
	if candidateTerm < currentTerm {
		return currentTerm, false
	}

	// votedFor is the candidateID that received vote
	// in the current term.
	votedFor := targetServer.vote()
	if votedFor == nil || votedFor == candidateID {
		// candidate's log is at least up-to-date
		// as receiver's log, so grant vote.
		return currentTerm, true
	}

	return currentTerm, false
}

func doServer(server *ServerState) {
	// All servers.
	if server.commitIndex > server.lastApplied {
		server.lastApplied++
		execute(server.logs[server.lastApplied])
	}
}

func candidate(server *ServerState) {
	// All servers.
	if server.commitIndex > server.lastApplied {
		server.lastApplied++
		execute(server.logs[server.lastApplied])
	}

	// on conversion to candidate
	server.currentTerm++
	
	// vote for itself
	vote(server)

	// reset election timer
	server.electionTimer.init()

	if SendRequestVote(allServers) > majority {
		becameLeader(server)
	}

	if electionTimeOut() {
		startNewElection()
	}
}

func becameLeader(server *ServerState) {
	// this needs to be sent periodically
	// while idle.
	SendHeartBeats(allServers)

	select {
	case entry := <-command:
		server.logs = append(server.logs, entry)
		execute(entry) // apply to state machine
		respond(entry.client)
	}

	for follower, nextIndex := range server.serverToNextIndex {
		if server.lastLogIndex >= nextIndex {
			AppendEntries(server.logs[nextIndex:])
		}
	}

	for index := range server.logs {
		if index > server.commitIndex {
			if server.logs[index].term == server.currentTerm {
				if majority of matchIndex[followers] >= index {
					server.commitIndex = index
				}					
			}
		}
	}
}

```


[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>







#### `etcd` internals: RPC between machines

*Raft* inter-server communication is done by remote procedure calls
(RPCs). The basic Raft algorithm requires only two types of RPCs:

- `RequestVote` RPCs, issued by candidates during elections.
- `AppendEntries` RPCs, issued by leaders:
  - **to replicate log entries**.
  - **to send out heartbeat messages**.

<br>
`etcd` uses [`Protocol Buffers`](https://developers.google.com/protocol-buffers/docs/overview?hl=en)
for inter-machine communication of structured data. Below are some of
core packages:

- [**`raft`**](http://godoc.org/github.com/coreos/etcd/raft):
  implements the raft consensus algorithm.
- [`raft/raftpb`](http://godoc.org/github.com/coreos/etcd/raft/raftpb):
  [auto-generated](https://github.com/coreos/etcd/blob/master/raft/raftpb/raft.pb.go#L1-L3)
  by protocol buffer compiler. It defines `MessageType`, `Entry`,
  `Message`, `State`, and other structured data required for *Raft* algorithm.
- [`rafthttp`](http://godoc.org/github.com/coreos/etcd/rafthttp):
  implements `http` operations in *Raft*.
- [**`etcdserver`**](http://godoc.org/github.com/coreos/etcd/etcdserver):
  connects servers in the cluster, using `HTTP`. It defines `Cluster`
  interface with methods: `ID` to return the cluster ID, `ClientURLs` to
  return the list of all clients URLs, `Members` to return the slice of
  members, etc. It also defines `Server` interface: `Start` to start a `etcd`
  server(*cluster*), `Stop` to stop the server, `ID` to return the ID of the
  server, `Leader` to return the server ID of leader, `Do` to handle the
  server requests, `Process` to take the raft message and apply it to the
  server's state machine (execute the command in log entry), `AddMember` to add
  a member into the cluster, `RemoveMember` to remove a member from the
  cluster, `UpdateMember` to update an existing member in the cluster.
- [`etcdserver/etcdhttp`](http://godoc.org/github.com/coreos/etcd/etcdserver/etcdhttp):
  implements `etcdserver` endpoints with muxed handlers.
- [`etcdserver/etcdserverpb`](http://godoc.org/github.com/coreos/etcd/etcdserver/etcdserverpb):
  auto-generated by protocol buffer compiler. It defines `Request` types used
  in `etcdserver` package.
- [`discovery`](http://godoc.org/github.com/coreos/etcd/discovery):
  implements cluster discovery.
- [**`client`**](http://godoc.org/github.com/coreos/etcd/client): is the official
  *Go* `etcd` client.


[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>




##### **`raft`**

Package `raft` implements the raft consensus algorithm.
And package `etcdserver` imports `raft` and `raftpb` package
to create and run `etcd` clusters:

- https://github.com/coreos/etcd/blob/master/etcdserver/server.go
- https://github.com/coreos/etcd/blob/master/etcdserver/raft.go
- https://github.com/coreos/etcd/blob/master/etcdserver/storage.go


<br>
Here's how `raft` and `raftpb` are **used** in `etcdserver`:

- [`Server`](https://godoc.org/github.com/coreos/etcd/etcdserver#Server)
  interface requires `Process` method to take `raftpb.Message` and
  to apply(execute) the message to its state machine. `Server` interface in
  `etcdserver` package is satisfied by type
  [`EtcdServer`](https://godoc.org/github.com/coreos/etcd/etcdserver#EtcdServer).
- `raftNode` struct embeds `raft.Node` interface that represents a node in a
  cluster. `raftNode` also embeds `raft.MemoryStorage` to store data
  in-memory.
- `raftNode` has `applyc chan apply` as a channel to do `apply` (execute
  commands in log entry). And `apply` is defined as a struct that contains
  a slice of `raftpb.Entry` which contains `Type`, `Term`, `Index`, `Data`.
  `apply` struct also contains `snapshot raftpb.SnapShot`, which represents
  current state of the system, with member `conf_state`, `index`, `term`.

<br>
First, it's helpful to look at
[`raft/raftpb/raft.proto`](https://github.com/coreos/etcd/blob/master/raft/raftpb/raft.proto)
because it defines structured data format used in *Raft* RPCs.

<br>
Here's how `raft` gets implemented. First, to define the server state:

```go
// Possible values for StateType.
const (
	StateFollower StateType = iota
	StateCandidate
	StateLeader
)

// StateType represents the role of a node in a cluster.
type StateType uint64

```

<br>
`Config` struct contains the configuration of a *Raft* node.

```go
type Config struct {
	ID             uint64     // non-zero ID of a raft node.
	peers          []uint64   // slice of all node IDs, including the node itself.

    ElectionTick   int        // election timeout, must be greater than
	                          // HeartbeatTick

	HearbeatTick   int        // heartbeat interval

	Storage Storage  // storage for a raft node.

	Applied         uint64   // the index of last applied entry.
	MaxSizePerMsg   uint64   // maximum size of each appending message.
	MaxInflightMsgs int      // maximum number of in-flight append-waiting
	                         // messages. TCP/UDP has its own buffer but useful
	                         // for avoid overflowing.

    Logger Logger
}

```

Please look at https://godoc.org/github.com/coreos/etcd/raft#Config for full
comments. And here's the definition of `Storage` in the `Config`:

```go
// Storage is an interface that may be implemented by the application
// to retrieve log entries from storage.
//
// If any Storage method returns an error, the raft instance will
// become inoperable and refuse to participate in elections; the
// application is responsible for cleanup and recovery in this case.
type Storage interface {
	// InitialState returns the saved HardState and ConfState information.
	InitialState() (pb.HardState, pb.ConfState, error)
	
	// Entries returns a slice of log entries in the range [lo,hi).
	// MaxSize limits the total size of the log entries returned, but
	// Entries returns at least one entry if any.
	Entries(lo, hi, maxSize uint64) ([]pb.Entry, error)
	
	// Term returns the term of entry i, which must be in the range
	// [FirstIndex()-1, LastIndex()]. The term of the entry before
	// FirstIndex is retained for matching purposes even though the
	// rest of that entry may not be available.
	Term(i uint64) (uint64, error)
	
	// LastIndex returns the index of the last entry in the log.
	LastIndex() (uint64, error)
	
	// FirstIndex returns the index of the first log entry that is
	// possibly available via Entries (older entries have been incorporated
	// into the latest Snapshot; if storage only contains the dummy entry the
	// first log entry is not available).
	FirstIndex() (uint64, error)

	// Snapshot returns the most recent snapshot.
	Snapshot() (pb.Snapshot, error)
}

```

So
[`raft/storage.go`](https://github.com/coreos/etcd/blob/master/raft/storage.go)
defines `Storage` interface, and package `raft` defines `MemoryStorage` type
that satisfies `Storage` interface by implementing methods in `Storage`
method.

<br>
[`raft/node.go`](https://github.com/coreos/etcd/blob/master/raft/node.go)
defines [`Node`](https://godoc.org/github.com/coreos/etcd/raft#Node)
interface, as below:

```go
type Node interface {
    // Tick increments the internal logical clock for the Node by a single tick. Election
    // timeouts and heartbeat timeouts are in units of ticks.
    Tick()

    // Campaign causes the Node to transition to candidate state and start campaigning to become leader.
    Campaign(ctx context.Context) error

    // Propose proposes that data be appended to the log.
    Propose(ctx context.Context, data []byte) error

    // ProposeConfChange proposes config change.
    // At most one ConfChange can be in the process of going through consensus.
    // Application needs to call ApplyConfChange when applying EntryConfChange type entry.
    ProposeConfChange(ctx context.Context, cc pb.ConfChange) error

    // Step advances the state machine using the given message. ctx.Err() will be returned, if any.
    Step(ctx context.Context, msg pb.Message) error

    // Ready returns a channel that returns the current point-in-time state
    // Users of the Node must call Advance after applying the state returned by Ready
    Ready() <-chan Ready

    // Advance notifies the Node that the application has applied and saved progress up to the last Ready.
    // It prepares the node to return the next available Ready.
    Advance()

    // ApplyConfChange applies config change to the local node.
    // Returns an opaque ConfState protobuf which must be recorded
    // in snapshots. Will never return nil; it returns a pointer only
    // to match MemoryStorage.Compact.
    ApplyConfChange(cc pb.ConfChange) *pb.ConfState

    // Status returns the current status of the raft state machine.
    Status() Status

    // Report reports the given node is not reachable for the last send.
    ReportUnreachable(id uint64)

    // ReportSnapshot reports the stutus of the sent snapshot.
    ReportSnapshot(id uint64, status SnapshotStatus)

    // Stop performs any necessary termination of the Node
    Stop()
}

```

Then what types satisfy this `Node` interface?
[`raft/node.go`](https://github.com/coreos/etcd/blob/master/raft/node.go)
has internal type `node` that satisfies `Node` interface.

```go
// node is the canonical implementation of the Node interface
type node struct {
	propc      chan pb.Message
	recvc      chan pb.Message
	confc      chan pb.ConfChange
	confstatec chan pb.ConfState
	readyc     chan Ready
	advancec   chan struct{}
	tickc      chan struct{}
	done       chan struct{}
	stop       chan struct{}
	status     chan chan Status
}

func newNode() node {
	return node{
		propc:      make(chan pb.Message),
		recvc:      make(chan pb.Message),
		confc:      make(chan pb.ConfChange),
		confstatec: make(chan pb.ConfState),
		readyc:     make(chan Ready),
		advancec:   make(chan struct{}),
		tickc:      make(chan struct{}),
		done:       make(chan struct{}),
		stop:       make(chan struct{}),
		status:     make(chan chan Status),
	}
}

```

By doing this, the type `node` is implicitly exported
as an interface type `Node`. And other packages use `Node`
with methods implemented in `node` type. For example:

```go
func StartNode(c *Config, peers []Peer) Node {
	...
	n := newNode()
	go n.run(r)
	return &n
}

```

Please check out
[this](https://github.com/gyuho/learn/tree/master/doc/go_interface#implicitly-exporting-interface)
for more detailed explanation of Go `interface` behavior.

Most operations in *Raft* is based on `raft.Node`, and states and data are
stored through `raft.MemoryStorage`. For the full source code and
documentation, please go to
[godoc.org/github.com/coreos/etcd/raft](http://godoc.org/github.com/coreos/etcd/raft).

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>




##### **`etcdserver`**

Package `etcdserver` defines interfaces for `etcd` cluster and servers.
Let's look at the actual code.

<br>
[`etcdserver/raft.go`](https://github.com/coreos/etcd/blob/master/etcdserver/raft.go)
imports package
[`raft`](http://godoc.org/github.com/coreos/etcd/raft)
and [`raft/raftpb`](http://godoc.org/github.com/coreos/etcd/raft/raftpb).

<br>
[`etcdserver/member.go`](https://github.com/coreos/etcd/blob/master/etcdserver/member.go)
implements `Member` as a member in a cluster. `Member` contains attributes of
peers and clients.

```go
// RaftAttributes represents the raft related attributes of an etcd member.
type RaftAttributes struct {
	PeerURLs []string `json:"peerURLs"`
}

// Attributes represents all the non-raft related attributes of an etcd member.
type Attributes struct {
	Name       string   `json:"name,omitempty"`
	ClientURLs []string `json:"clientURLs,omitempty"`
}

type Member struct {
	ID types.ID `json:"id"`
	RaftAttributes
	Attributes
}

```


<br>
[`etcdserver/cluster.go`](https://github.com/coreos/etcd/blob/master/etcdserver/cluster.go)
defines `Cluster` interface:

```go
type Cluster interface {
	// ID returns the cluster ID
	ID() types.ID

	// ClientURLs returns an aggregate set of all URLs on which this
	// cluster is listening for client requests
	ClientURLs() []string

	// Members returns a slice of members sorted by their ID
	Members() []*Member

	// Member retrieves a particular member based on ID, or nil if the
	// member does not exist in the cluster
	Member(id types.ID) *Member

	// IsIDRemoved checks whether the given ID has been removed from this
	// cluster at some point in the past
	IsIDRemoved(id types.ID) bool

	// ClusterVersion is the cluster-wide minimum major.minor version.
	Version() *semver.Version
}

```

And it has internal type `cluster` to satisfy this interface:

```go
type cluster struct {
	id    types.ID
	token string
	store store.Store

	sync.Mutex // guards the fields below
	version    *semver.Version
	members    map[types.ID]*Member
	// removed contains the ids of removed members in the cluster.
	// removed id cannot be reused.
	removed map[types.ID]bool
}

```

You can add, remove, update members in a cluster.


<br>
[`etcdserver/server.go`](https://github.com/coreos/etcd/blob/master/etcdserver/server.go)
defines `Server` interface:

```go
type Server interface {
	// Start performs any initialization of the Server necessary for it to
	// begin serving requests. It must be called before Do or Process.
	// Start must be non-blocking; any long-running server functionality
	// should be implemented in goroutines.
	Start()

	// Stop terminates the Server and performs any necessary finalization.
	// Do and Process cannot be called after Stop has been invoked.
	Stop()

	// ID returns the ID of the Server.
	ID() types.ID

	// Leader returns the ID of the leader Server.
	Leader() types.ID

	// Do takes a request and attempts to fulfill it, returning a Response.
	Do(ctx context.Context, r pb.Request) (Response, error)

	// Process takes a raft message and applies it to the server's raft state
	// machine, respecting any timeout of the given context.
	Process(ctx context.Context, m raftpb.Message) error

	// AddMember attempts to add a member into the cluster. It will return
	// ErrIDRemoved if member ID is removed from the cluster, or return
	// ErrIDExists if member ID exists in the cluster.
	AddMember(ctx context.Context, memb Member) error

	// RemoveMember attempts to remove a member from the cluster. It will
	// return ErrIDRemoved if member ID is removed from the cluster, or return
	// ErrIDNotFound if member ID is not in the cluster.
	RemoveMember(ctx context.Context, id uint64) error

	// UpdateMember attempts to update a existing member in the cluster. It will
	// return ErrIDNotFound if the member ID does not exist.
	UpdateMember(ctx context.Context, updateMemb Member) error

	ClusterVersion() *semver.Version
}

```

And `EtcdServer` type satisfies the `Server` interface:

```go
// EtcdServer is the production implementation of the Server interface
type EtcdServer struct {
	// r must be the first element to keep 64-bit alignment for atomic
	// access to fields
	r raftNode

	cfg       *ServerConfig
	snapCount uint64

	w          wait.Wait
	stop       chan struct{}
	done       chan struct{}
	errorc     chan error
	id         types.ID
	attributes Attributes

	cluster *cluster

	store store.Store
	kv    dstorage.KV

	stats  *stats.ServerStats
	lstats *stats.LeaderStats

	SyncTicker <-chan time.Time

	reqIDGen *idutil.Generator

	// forceVersionC is used to force the version monitor loop
	// to detect the cluster version immediately.
	forceVersionC chan struct{}
}

```

And please check out [`gyuho/go-fuzz-etcd`](https://github.com/gyuho/go-fuzz-etcd)
for more fuzz testing on `etcd`.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>

