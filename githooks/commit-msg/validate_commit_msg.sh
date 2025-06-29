#!/bin/bash

commit_msg_file="$1"
commit_msg=$(cat "$commit_msg_file")
pattern="^(feat|fix|chore|docs|style|refactor|perf|test|build|ci|revert|merge)(\(.+\))?: .+"

if ! [[ "$commit_msg" =~ $pattern ]]; then
  echo "Invalid commit message."
  echo "Expect: <type>(scope): message"
  exit 1
fi

exit 0