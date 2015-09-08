[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# etcd, raft algorithm

- [Reference](#reference)
- [consensus algorithm](#consensus-algorithm)
- [raft algorithm: introduction](#raft-algorithm-introduction)
- [raft algorithm](#raft-algorithm)

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### Reference

**_DISCLAIMER_:** This is **my** understanding of the **raft algorithm**.<br>
I may say things out of ignorance. Please refer to reading below.

- [Consensus (computer science)](https://en.wikipedia.org/wiki/Consensus_(computer_science))
- [CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem)
- [*Raft paper by Diego Ongaro and John Ousterhout*](http://ramcloud.stanford.edu/raft.pdf)
- [The Raft Consensus Algorithm](https://raft.github.io/)
- [Raft (computer science)](https://en.wikipedia.org/wiki/Raft_(computer_science))
- [coreos/etcd](https://github.com/coreos/etcd)

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
A consensus algorithm must be resilient to these kind of failures in order
to guarantee the correctness.

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>









#### raft algorithm: introduction

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










#### 

[↑ top](#etcd-raft-algorithm)
<br><br><br><br>
<hr>
