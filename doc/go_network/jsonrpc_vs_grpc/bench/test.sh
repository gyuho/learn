#!/bin/bash
test_size=300000
client_size=100

current_dir=$(pwd)
dot="."
test_function_name="XXX"
benchmark_function_name="BenchmarkStress"

repeat_size=1

printf "\n"
for i in `seq 1 $repeat_size`;
do
	echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh && cd $HOME/go/pkg && rm -rf *;
	cd $current_dir && fpath="$(echo $current_dir)/jsonrpc_$i.txt";
	if [ "$benchmark_function_name" == "$dot" ]; then
		echo "running all benchmarks...";
		go test -opt "jsonrpc" -size $test_size -run $test_function_name -bench . -benchmem -cpu 4 > $fpath;
	else
		echo "running only $benchmark_function_name";
		go test -opt "jsonrpc" -size $test_size -run $test_function_name -bench $benchmark_function_name -benchmem -cpu 4 > $fpath;
	fi
done

printf "\n"
for i in `seq 1 $repeat_size`;
do
	echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh && cd $HOME/go/pkg && rm -rf *;
	cd $current_dir && fpath="$(echo $current_dir)/grpc_$i.txt";
	if [ "$benchmark_function_name" == "$dot" ]; then
		echo "running all benchmarks...";
		go test -opt "grpc" -size $test_size -run $test_function_name -bench . -benchmem -cpu 4 > $fpath;
	else
		echo "running only $benchmark_function_name";
		go test -opt "grpc" -size $test_size -run $test_function_name -bench $benchmark_function_name -benchmem -cpu 4 > $fpath;
	fi
done

printf "\n"
for i in `seq 1 $repeat_size`;
do
	echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh && cd $HOME/go/pkg && rm -rf *;
	cd $current_dir && fpath="$(echo $current_dir)/grpc_multiclients_$i.txt";
	if [ "$benchmark_function_name" == "$dot" ]; then
		echo "running all benchmarks...";
		go test -opt "grpc" -size $test_size -numc $client_size -run $test_function_name -bench . -benchmem -cpu 4 > $fpath;
	else
		echo "running only $benchmark_function_name";
		go test -opt "grpc" -size $test_size -numc $client_size -run $test_function_name -bench $benchmark_function_name -benchmem -cpu 4 > $fpath;
	fi
done

benchmark_name="results_jsonrpc_vs_grpc.txt"
echo "$(echo $benchmark_function_name):" > $current_dir/$benchmark_name;
for i in `seq 1 $repeat_size`;
do
	echo "[$(echo $i)]:" >> $current_dir/results_jsonrpc_vs_grpc.txt;
	jsonrpc_path="$(echo $current_dir)/jsonrpc_$i.txt" && \
	grpc_path="$(echo $current_dir)/grpc_$i.txt" && \
	benchcmp $jsonrpc_path $grpc_path >> $current_dir/$benchmark_name && \
	echo "" >> $current_dir/$benchmark_name && \
	echo "" >> $current_dir/$benchmark_name;
done

benchmark_name="results_jsonrpc_vs_grpc_multiclients.txt"
echo "$(echo $benchmark_function_name):" > $current_dir/$benchmark_name;
for i in `seq 1 $repeat_size`;
do
	echo "[$(echo $i)]:" >>	$current_dir/results_jsonrpc_vs_grpc_multiclients.txt;
	jsonrpc_path="$(echo $current_dir)/jsonrpc_$i.txt" && \
	grpc_multiclients_path="$(echo $current_dir)/grpc_multiclients_$i.txt" && \
	benchcmp $jsonrpc_path $grpc_multiclients_path >> $current_dir/$benchmark_name && \
	echo "" >> $current_dir/$benchmark_name && \
	echo "" >> $current_dir/$benchmark_name;
done

benchmark_name="results_grpc_vs_grpc_multiclients.txt"
echo "$(echo $benchmark_function_name):" > $current_dir/$benchmark_name;
for i in `seq 1 $repeat_size`;
do
	echo "[$(echo $i)]:" >> $current_dir/results_jsonrpc_vs_grpc.txt;
	grpc_path="$(echo $current_dir)/grpc_$i.txt" && \
	grpc_multiclients_path="$(echo $current_dir)/grpc_multiclients_$i.txt" && \
	benchcmp $grpc_path $grpc_multiclients_path >> $current_dir/$benchmark_name && \
	echo "" >> $current_dir/$benchmark_name && \
	echo "" >> $current_dir/$benchmark_name;
done

