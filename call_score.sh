#!/usr/bin/env bash
# Usage: call_score.sh AGE INCOME CARDS CAR_LOANS [BASE_URL]
set -euo pipefail
if [ "$#" -lt 4 ]; then
  echo "usage: $0 AGE INCOME CARDS CAR_LOANS [BASE_URL]" >&2; exit 2
fi
BASE_URL="${5:-http://127.0.0.1:5000}"
curl --fail --silent --show-error --max-time 10 --get "$BASE_URL/score" \
  --data-urlencode "age=$1" --data-urlencode "income=$2" \
  --data-urlencode "numberOfCard=$3" --data-urlencode "numberOfCarLoan=$4"
echo
