# port 2379 is for client communication.
# port 2380 is for server-to-server communication.

export INFRA_PRIVATE_IP_0=172.0.0.0;
export INFRA_PRIVATE_IP_1=172.0.0.1;
export INFRA_PRIVATE_IP_2=172.0.0.2;

# if you need to expose to public
export INFRA_PUBLIC_IP_0=52.0.0.0;
export INFRA_PUBLIC_IP_1=52.0.0.1;
export INFRA_PUBLIC_IP_2=52.0.0.2;


# deploy in each machine with
curl -L  https://github.com/coreos/etcd/releases/download/v2.1.0-rc.0/etcd-v2.1.0-rc.0-linux-amd64.tar.gz -o etcd-v2.1.0-rc.0-linux-amd64.tar.gz
tar xzvf etcd-v2.1.0-rc.0-linux-amd64.tar.gz
cd etcd-v2.1.0-rc.0-linux-amd64
./etcd \
-name infra0 \
-initial-cluster-token my-infra \
-initial-cluster infra0=http://$INFRA_PRIVATE_IP_0:2380,infra1=http://$INFRA_PRIVATE_IP_1:2380,infra2=http://$INFRA_PRIVATE_IP_2:2380 \
-initial-cluster-state new \
-initial-advertise-peer-urls http://$INFRA_PRIVATE_IP_0:2380 \
-listen-peer-urls http://$INFRA_PRIVATE_IP_0:2380 \
-listen-client-urls http://$INFRA_PRIVATE_IP_0:2379,http://127.0.0.1:2379 \
-advertise-client-urls http://$INFRA_PRIVATE_IP_0:2379 \
;

curl -L  https://github.com/coreos/etcd/releases/download/v2.1.0-rc.0/etcd-v2.1.0-rc.0-linux-amd64.tar.gz -o etcd-v2.1.0-rc.0-linux-amd64.tar.gz
tar xzvf etcd-v2.1.0-rc.0-linux-amd64.tar.gz
cd etcd-v2.1.0-rc.0-linux-amd64
./etcd \
-name infra1 \
-initial-cluster-token my-infra \
-initial-cluster infra0=http://$INFRA_PRIVATE_IP_0:2380,infra1=http://$INFRA_PRIVATE_IP_1:2380,infra2=http://$INFRA_PRIVATE_IP_2:2380 \
-initial-cluster-state new \
-initial-advertise-peer-urls http://$INFRA_PRIVATE_IP_1:2380 \
-listen-peer-urls http://$INFRA_PRIVATE_IP_1:2380 \
-listen-client-urls http://$INFRA_PRIVATE_IP_1:2379,http://127.0.0.1:2379 \
-advertise-client-urls http://$INFRA_PRIVATE_IP_1:2379 \
;

curl -L  https://github.com/coreos/etcd/releases/download/v2.1.0-rc.0/etcd-v2.1.0-rc.0-linux-amd64.tar.gz -o etcd-v2.1.0-rc.0-linux-amd64.tar.gz
tar xzvf etcd-v2.1.0-rc.0-linux-amd64.tar.gz
cd etcd-v2.1.0-rc.0-linux-amd64
./etcd \
-name infra2 \
-initial-cluster-token my-infra \
-initial-cluster infra0=http://$INFRA_PRIVATE_IP_0:2380,infra1=http://$INFRA_PRIVATE_IP_1:2380,infra2=http://$INFRA_PRIVATE_IP_2:2380 \
-initial-cluster-state new \
-initial-advertise-peer-urls http://$INFRA_PRIVATE_IP_2:2380 \
-listen-peer-urls http://$INFRA_PRIVATE_IP_2:2380 \
-listen-client-urls http://$INFRA_PRIVATE_IP_2:2379,http://127.0.0.1:2379 \
-advertise-client-urls http://$INFRA_PRIVATE_IP_2:2379 \
;


# outside of VPC, you would do
curl -L http://$INFRA_PUBLIC_IP_0:2379/version;
curl -L http://$INFRA_PUBLIC_IP_0:2379/v2/keys/status;
curl -L http://$INFRA_PUBLIC_IP_0:2379/v2/keys/test;
curl -L http://$INFRA_PUBLIC_IP_0:2379/v2/keys/queue;
curl -L -XPUT http://$INFRA_PUBLIC_IP_0:2379/v2/keys/test/55 -d value="Hello";
curl -L -XDELETE http://$INFRA_PUBLIC_IP_0:2379/v2/keys/queue/350;
