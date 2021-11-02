# Run Avalanche on local machine

*Last update: November 1, 2021*

*Next: [Run Avalanche on cloud with Kubernetes](./run-avalanche-on-cloud-with-kubernetes.md)*

- [What is Avalanche"?](#what-is-avalanche)
- [What is "Run Avalanche"?](#what-is-run-avalanche)
- [What is "Run Avalanche on local machine"?](#what-is-run-avalanche-on-local-machine)
    - [Install](#install)
    - [Set up a local test network](#set-up-a-local-test-network)
        - [Understanding basic configuration](#understanding-basic-configuration)
        - [Generate TLS certificates](#generate-tls-certificates)
        - [Run nodes](#run-nodes)
    - [Verify nodes are connected](#verify-nodes-are-connected)
- [Test transaction](#test-transaction)
- [Reference](#reference)

### What is "Avalanche"?

[Avalanche](https://www.avax.network/) is an open-source [decentralized application](https://support.avax.network/en/articles/4587146-what-is-a-decentralized-application-dapp) platform, powered by [Avalanche consensus](./nakamoto-bitcoin-vs-snow-avalanche-consensus.md). It achieves global finance scalability with near-instant transaction finality, without compromising the decentralization. It enables enterprise blockchain deployments in one interoperable, highly scalable ecosystem. Ethereum developers can quickly build on Avalanche as Solidity (smart contract platform in Ethereum) works out-of-the-box.

### What is "Run Avalanche"?

"Run Avalanche" here means "install and run Avalanche node". Anybody can participate in the Avalanche network to help secure the ecosystem in a decentralized manner. Each Avalanche node installs [AvalancheGo](https://github.com/ava-labs/avalanchego) software, the Go implementation of an Avalanche node, to verify transaction data in the distributed network.

### What is "Run Avalanche on local machine"?

The below instruction shows how to install [AvalancheGo](https://github.com/ava-labs/avalanchego) and run Avalanche network in your local machine (MacOS/Linux).

#### Install

The below commands assume you have [Go](https://golang.org/dl/) installed.

```bash
if [[ ! -d ${HOME}/go/src/github.com/ava-labs/avalanchego ]]
then
  echo "cloning ava-labs/avalanchego"
  mkdir -p ${HOME}/go/src/github.com/ava-labs
  rm -rf ${HOME}/go/src/github.com/ava-labs/avalanchego
  cd ${HOME}/go/src/github.com/ava-labs
  git clone git@github.com:ava-labs/avalanchego.git
  cd ${HOME}/go/src/github.com/ava-labs/avalanchego
else
  echo "syncing ava-labs/avalanchego"
  cd ${HOME}/go/src/github.com/ava-labs/avalanchego
  git fetch --all
  git checkout master
  git pull origin master
fi
```

To compile [AvalancheGo](https://github.com/ava-labs/avalanchego):

```bash
cd ${HOME}/go/src/github.com/ava-labs/avalanchego
./scripts/build.sh
./build/avalanchego -h
```

See [Download AvalancheGo](https://docs.avax.network/build/tutorials/nodes-and-staking/run-avalanche-node#download-avalanchego) for more.

#### Set up a local test network

##### Understanding basic configuration

First, let's go over some common configuration to the `avalanchego` binary. For example, the [documentation](https://docs.avax.network/build/tutorials/platform/create-a-local-test-network#manually) has the below commands:

```bash
# first node
./build/avalanchego \
--public-ip=127.0.0.1 \
--snow-sample-size=2 \
--snow-quorum-size=2 \
--http-port=9650 \
--staking-port=9651 \
--db-dir=db/node1 \
--staking-enabled=true \
--network-id=local \
--bootstrap-ips= \
--staking-tls-cert-file=$(pwd)/staking/local/staker1.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker1.key

# second node
./build/avalanchego \
--public-ip=127.0.0.1 \
--snow-sample-size=2 \
--snow-quorum-size=2 \
--http-port=9652 \
--staking-port=9653 \
--db-dir=db/node2 \
--staking-enabled=true \
--network-id=local \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-cert-file=$(pwd)/staking/local/staker2.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker2.key

# ...
```

- *`network-id`*: Network ID this node will connect to. The default is to connect to the `mainnet`. Set it to `local` for cluster test network. For example, the network ID is used for checking the version compatibility.
- *`public-ip`*: Public IP of this node for peer-to-peer communication. If empty, the node tries to discover the node IP with [NAT traversal](https://en.wikipedia.org/wiki/NAT_traversal). `public-ip` value is ignored if `dynamic-public-ip` is non-empty (e.g., `opendns`).
- *`http-port`*: Port of the HTTP server. [AvalancheGo API](https://docs.avax.network/build/avalanchego-apis) calls are made to this `[node-ip]:[http-port]`. For instance, `/ext/health` endpoint responds with `200` if the node is healthy.
- *`snow-sample-size`*: As in Snowball, the node samples \\(k\\) peers to query. Once the querying node collects \\(k\\) responses, the node calculates the color ratio to check against a threshold and decide on the agreement.
- *`snow-quorum-size`*: Alpha \\(α\\) in Snowman protocol. The threshold, sufficiently large fraction of the sample (quorum). If more than \\(α\\) (quorum) respond positively to the querying node, the protocol sets the chit value of the querying node for the respective transaction \\(T\\) to 1 -- "strongly preferred".
- *`db-dir`*: Path to database directory.
- *`staking-enabled`*: Every Avalanche node must stake `AVAX` tokens to validate. Such proof-of-stake makes it infeasibly expensive for a malicious actor to gain enough influence over the network and compromise the network (e.g., sybil attack). Set `true` to enable staking. If enabled, TLS network is required. If disabled, sybil control is not enforced. See [Staking](https://docs.avax.network/learn/platform-overview/staking) for more.
- *`staking-port`*: Port of the consensus server.
- *`staking-tls-key-file`*: Path to the TLS certificate for staking.
- *`staking-tls-cert-file`*: Path to the TLS private key for staking.
- *`bootstrap-ips`*: Comma separated list of bootstrap peer IPs to connect to. Empty by default. If empty, the value defaults to sample beacons in the network. If not empty, it overwrites the sample beacon IPs. The default beacon nodes are hardcoded in [`genesis/beacons.go`](https://github.com/ava-labs/avalanchego/blob/v1.6.3/genesis/beacons.go).
- *`bootstrap-ids`*: Comma separated list of bootstrap peer IDs to connect to. Empty by default. If empty, the value defaults to sample beacons in the network. If not empty, it overwrites the sample beacon IDs.

Since each node runs on the local network, we will set *`--network-id=local`* and use the standard IPv4 loopback address *`--public-ip=127.0.0.1`*. And we need to assign each node a unique port for *`--http-port=`* for API endpoints, because otherwise the port will conflict with other nodes on the host. If the cluster size is only 5, the sample and quorum size of 2 will be sufficient for consensus: *`--snow-sample-size=2`* and *`--snow-quorum-size=2`*. We will enable staking with *`--staking-enabled=true`* to simulate proof-of-stake mechanisms as in practice. We will set *`--bootstrap-ips`* to the first node with `[node-ip]:[staking-port]`. In practice, we set it to empty, so that the node can connect to the pre-defined beacon nodes in the network.

##### Generate TLS certificates

[staking/tls.go](https://github.com/ava-labs/avalanchego/blob/v1.6.3/staking/tls.go#L107-L122) auto-generates the certs if the cert and key paths do not exist.

```bash
# https://github.com/gyuho/avax-tester
avax-tester certs create --dir-path "/tmp/avalanchego-certs"

openssl x509 \
-in /tmp/avalanchego-certs/s1.pem \
-text -noout
```

##### Run nodes

```bash
rm -rf /tmp/avalanchego-db \
&& mkdir -p /tmp/avalanchego-db
```

```bash
# to log verbose
# --log-level=verbo

kill -9 $(lsof -t -i:9650)
kill -9 $(lsof -t -i:9651)
cd ${HOME}/go/src/github.com/ava-labs/avalanchego
./build/avalanchego \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9650 \
--snow-sample-size=2 \
--snow-quorum-size=2 \
--db-dir=/tmp/avalanchego-db/s1 \
--staking-port=9651 \
--staking-enabled=true \
--bootstrap-ips= \
--staking-tls-key-file=$(pwd)/staking/local/staker1.key \
--staking-tls-cert-file=$(pwd)/staking/local/staker1.crt

# TODO: not working...
# --staking-tls-key-file=/tmp/avalanchego-certs/s1-key.pem \
# --staking-tls-cert-file=/tmp/avalanchego-certs/s1.pem
```

```bash
kill -9 $(lsof -t -i:9652)
kill -9 $(lsof -t -i:9653)
cd ${HOME}/go/src/github.com/ava-labs/avalanchego
./build/avalanchego \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9652 \
--snow-sample-size=2 \
--snow-quorum-size=2 \
--db-dir=/tmp/avalanchego-db/s2 \
--staking-port=9653 \
--staking-enabled=true \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-key-file=$(pwd)/staking/local/staker2.key \
--staking-tls-cert-file=$(pwd)/staking/local/staker2.crt

# TODO: not working...
# --staking-tls-key-file=/tmp/avalanchego-certs/s2-key.pem \
# --staking-tls-cert-file=/tmp/avalanchego-certs/s2.pem
```

```bash
kill -9 $(lsof -t -i:9654)
kill -9 $(lsof -t -i:9655)
cd ${HOME}/go/src/github.com/ava-labs/avalanchego
./build/avalanchego \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9654 \
--snow-sample-size=2 \
--snow-quorum-size=2 \
--db-dir=/tmp/avalanchego-db/s3 \
--staking-port=9655 \
--staking-enabled=true \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-key-file=$(pwd)/staking/local/staker3.key \
--staking-tls-cert-file=$(pwd)/staking/local/staker3.crt

# TODO: not working
# --staking-tls-key-file=/tmp/avalanchego-certs/s3-key.pem \
# --staking-tls-cert-file=/tmp/avalanchego-certs/s3.pem
```

```bash
kill -9 $(lsof -t -i:9656)
kill -9 $(lsof -t -i:9657)
cd ${HOME}/go/src/github.com/ava-labs/avalanchego
./build/avalanchego \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9656 \
--snow-sample-size=2 \
--snow-quorum-size=2 \
--db-dir=/tmp/avalanchego-db/s4 \
--staking-port=9657 \
--staking-enabled=true \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-key-file=$(pwd)/staking/local/staker4.key \
--staking-tls-cert-file=$(pwd)/staking/local/staker4.crt

# TODO: not working
# --staking-tls-key-file=/tmp/avalanchego-certs/s4-key.pem \
# --staking-tls-cert-file=/tmp/avalanchego-certs/s4.pem
```

```bash
kill -9 $(lsof -t -i:9658)
kill -9 $(lsof -t -i:9659)
cd ${HOME}/go/src/github.com/ava-labs/avalanchego
./build/avalanchego \
--network-id=local \
--public-ip=127.0.0.1 \
--http-port=9658 \
--snow-sample-size=2 \
--snow-quorum-size=2 \
--db-dir=/tmp/avalanchego-db/s5 \
--staking-port=9659 \
--staking-enabled=true \
--bootstrap-ips=127.0.0.1:9651 \
--bootstrap-ids=NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg \
--staking-tls-cert-file=$(pwd)/staking/local/staker5.crt \
--staking-tls-key-file=$(pwd)/staking/local/staker5.key

# TODO: not working...
# --staking-tls-key-file=/tmp/avalanchego-certs/s5-key.pem \
# --staking-tls-cert-file=/tmp/avalanchego-certs/s5.pem
```

#### Verify nodes are connected

To verify nodes are connected via HTTP endpoints:

```bash
curl -X POST --data '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"info.peers"
}' \
-H 'content-type:application/json;' \
127.0.0.1:9650/ext/info
```

```bash
curl -X POST --data '{
    "jsonrpc":"2.0",
    "id"     :1,
    "method" :"info.peers"
}' \
-H 'content-type:application/json;' \
127.0.0.1:9658/ext/info
```

### Test transaction

Now that we created the network, let's ["fund the local test network"](https://docs.avax.network/build/tutorials/platform/fund-a-local-test-network).

### Reference

- [Avalanche documentation](https://docs.avax.network/)
- [Create a Local Avalanche Test Network](https://docs.avax.network/build/tutorials/platform/create-a-local-test-network)
- [Run an Avalanche Node](https://docs.avax.network/build/tutorials/nodes-and-staking/run-avalanche-node)

[↑ top](#run-avalanche-on-local-machine)
