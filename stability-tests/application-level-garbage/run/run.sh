#!/bin/bash
rm -rf /tmp/freed-temp

freed --devnet --appdir=/tmp/freed-temp --profile=6061 --loglevel=debug &
KASPAD_PID=$!
KASPAD_KILLED=0
function killFreedIfNotKilled() {
    if [ $KASPAD_KILLED -eq 0 ]; then
      kill $KASPAD_PID
    fi
}
trap "killFreedIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KASPAD_PID

wait $KASPAD_PID
KASPAD_KILLED=1
KASPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Freed exit code: $KASPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
