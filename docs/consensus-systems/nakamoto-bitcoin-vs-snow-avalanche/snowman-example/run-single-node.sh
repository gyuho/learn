#!/usr/bin/env bash
set -xue

if ! [[ "$0" =~ run-single-node.sh ]]; then
    echo "must be run from example root"
    exit 255
fi

targets=(
    single-node/single-block-no-conflict
    single-node/simple-conflict
    single-node/transitive-vote
)
for target in "${targets[@]}"
do
    set +x
    echo ""
    echo ""
    echo ""
    set -x

    pushd $target
    go run main.go
    popd
done

set +x
echo ""
echo ""
echo "SUCCESSFULLY ran all examples"
echo ""
echo ""
set -x
