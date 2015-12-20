#!/bin/bash

current_dir=$(pwd)

################################################################
printf "\n"
echo "TEST #1 $ go test -v ./..."
go test -v ./...;

################################################################
printf "\n"
echo "TEST #2 $ go test -v -race ./.."
go test -v -race ./...;

################################################################
cd $HOME && go get -u golang.org/x/tools/cmd/benchcmp && \
cat /proc/cpuinfo | grep "model name" | head -1 && \
go version && \
printf "linux kernel: %s\n" $(uname -r) && \

################################################################

printf "\n"
echo "TEST #3 Benchmarks..."

dot="."
test_function_name="XXX"
benchmark_function_name="."

repeat_size=2

for i in `seq 1 $repeat_size`;
do
	echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh && cd $HOME/go/pkg && rm -rf *;
	cd $current_dir && fpath="$(echo $current_dir)/old_$i.txt";
	if [ "$benchmark_function_name" == "$dot" ]; then
		echo "running all benchmarks...";
		go test -opt "slice" -run $test_function_name -bench . -benchmem -cpu 1,2,4,8 > $fpath;
	else
		echo "running only $benchmark_function_name";
		go test -opt "slice" -run $test_function_name -bench $benchmark_function_name -benchmem -cpu 1,2,4,8 > $fpath;
	fi
done

for i in `seq 1 $repeat_size`;
do
	echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh && cd $HOME/go/pkg && rm -rf *;
	cd $current_dir && fpath="$(echo $current_dir)/new_$i.txt";
	if [ "$benchmark_function_name" == "$dot" ]; then
		echo "running all benchmarks...";
		go test -opt "map" -run $test_function_name -bench . -benchmem -cpu 1,2,4,8 > $fpath;
	else
		echo "running only $benchmark_function_name";
		go test -opt "map" -run $test_function_name -bench $benchmark_function_name -benchmem -cpu 1,2,4,8 > $fpath;
	fi
done

echo "$(echo $benchmark_function_name):" > $current_dir/benchmark_results.txt;
for i in `seq 1 $repeat_size`;
do
	echo "[$(echo $i)]:" >> $current_dir/benchmark_results.txt;
	old_path="$(echo $current_dir)/old_$i.txt" && new_path="$(echo $current_dir)/new_$i.txt" && \
	benchcmp $old_path $new_path >> $current_dir/benchmark_results.txt && \
	echo "" >> $current_dir/benchmark_results.txt && \
	echo "" >> $current_dir/benchmark_results.txt;
done

################################################################

printf "\n"
echo "Done"

