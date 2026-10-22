#!/bin/sh
# run.sh — CsvStat grader. Do not filter, skip, or reinterpret.
#
# Builds ./cmd/csvstat with `go build` and compares every command's exit code and standard output
# over the supplied data files with the known answer below. Standard error is never compared.
#
#   sh sources/run.sh                    from the application root (how `drydock uat` runs it)
#   sh uat/CsvStat/bin/run.sh APP        from anywhere, grading the application root APP
#
# Exit 0 only when the build succeeds and every check passes.
set -u

HERE=$(cd "$(dirname "$0")" && pwd)
APP=$(cd "${1:-.}" && pwd)
if [ -f "$HERE/employees.csv" ]; then DATA=$HERE; else DATA=$(cd "$HERE/../sources" && pwd); fi
WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT INT TERM
BIN="$WORK/csvstat"
MISSING="$WORK/no-such-file.csv"

if ! (cd "$APP" && go build -o "$BIN" ./cmd/csvstat); then
    echo "error: go build ./cmd/csvstat failed at the application root $APP." >&2
    echo "The deliverable is a Go module whose main package is ./cmd/csvstat." >&2
    exit 1
fi

PASSED=0
FAILED=0

# check <id> <expected exit> <csvstat args...>  — expected standard output on stdin.
check() {
    id=$1
    want_rc=$2
    shift 2
    cat > "$WORK/want"
    (cd "$APP" && "$BIN" "$@") > "$WORK/got" 2> "$WORK/err"
    rc=$?
    if [ "$rc" -eq "$want_rc" ] && cmp -s "$WORK/want" "$WORK/got"; then
        PASSED=$((PASSED + 1))
        echo "PASS $id: csvstat $*"
    else
        FAILED=$((FAILED + 1))
        echo "FAIL $id: csvstat $*"
        echo "  exit: expected $want_rc, got $rc"
        diff "$WORK/want" "$WORK/got" | sed 's/^/  /'
        sed 's/^/  stderr: /' "$WORK/err"
    fi
}

# ── A. Every column of each data file ──────────────────────────────────────────────────
check A1 0 "$DATA/employees.csv" <<'EOF'
column,type,rows,nulls,distinct,min,max,sum,mean,median
id,integer,24,0,24,1001,1024,24300,1012.50,1012.50
name,text,24,0,23,Ada,Tony,,,
department,text,24,0,4,Engineering,Support,,,
hire_date,date,24,0,24,2015-08-27,2025-06-02,,,
salary,number,24,0,24,43699.25,147150.99,2085635.10,86901.46,79496.68
remote,boolean,24,0,4,false,true,,,
manager_id,integer,24,15,4,2,5,30,3.33,4.00
notes,text,24,10,4,mentor,"remote, part-time",,,
EOF
check A2 0 "$DATA/sensors.csv" <<'EOF'
column,type,rows,nulls,distinct,min,max,sum,mean,median
reading_id,integer,36,0,36,1,36,666,18.50,18.50
recorded_at,text,36,0,36,2025-03-01T08:00,2025-03-03T19:00,,,
sensor,text,36,0,4,east,west,,,
temp_c,number,36,3,32,-12.21,29.81,207.06,6.27,2.41
humidity,integer,36,7,25,18,92,1698,58.55,60.00
ok,boolean,36,0,2,false,true,,,
battery,empty,36,36,0,,,,,
EOF
check A3 0 "$DATA/orders.csv" <<'EOF'
column,type,rows,nulls,distinct,min,max,sum,mean,median
order_id,integer,30,0,30,5000,5029,150435,5014.50,5014.50
order_date,date,30,0,19,2025-02-01,2025-02-27,,,
customer,text,30,0,5,Acme Corp,Umbrella,,,
qty,integer,30,0,9,1,12,206,6.87,6.00
unit_price,number,30,0,30,15.27,231.18,4177.97,139.27,131.31
coupon,text,30,18,2,SAVE10,WELCOME,,,
ref,text,30,0,5,-3,A7,,,
shipped,text,30,0,20,2025-02-30,2025-03-27,,,
EOF
check A4 0 "$DATA/header_only.csv" <<'EOF'
column,type,rows,nulls,distinct,min,max,sum,mean,median
sku,empty,0,0,0,,,,,
price,empty,0,0,0,,,,,
in_stock,empty,0,0,0,,,,,
EOF

# ── B. One column ────────────────────────────────────────────────────────────────────────
check B1 0 --column salary "$DATA/employees.csv" <<'EOF'
column,type,rows,nulls,distinct,min,max,sum,mean,median
salary,number,24,0,24,43699.25,147150.99,2085635.10,86901.46,79496.68
EOF
check B2 0 --column notes "$DATA/employees.csv" <<'EOF'
column,type,rows,nulls,distinct,min,max,sum,mean,median
notes,text,24,10,4,mentor,"remote, part-time",,,
EOF
check B3 0 --column temp_c "$DATA/sensors.csv" <<'EOF'
column,type,rows,nulls,distinct,min,max,sum,mean,median
temp_c,number,36,3,32,-12.21,29.81,207.06,6.27,2.41
EOF

# ── C. Frequencies ───────────────────────────────────────────────────────────────────────
check C1 0 --freq department "$DATA/employees.csv" <<'EOF'
value,count
Sales,8
Engineering,7
Support,7
Finance,2
EOF
check C2 0 --freq sensor --top 2 "$DATA/sensors.csv" <<'EOF'
value,count
east,9
north,9
EOF
check C3 0 --freq customer "$DATA/orders.csv" <<'EOF'
value,count
Globex,10
Acme Corp,9
Umbrella,7
Initech,3
Hooli,1
EOF
check C4 0 --freq manager_id --top 3 "$DATA/employees.csv" <<'EOF'
value,count
4,4
2,3
3,1
EOF
check C5 0 --freq battery "$DATA/sensors.csv" <<'EOF'
value,count
EOF

# ── E. Rejected input and operational failures ───────────────────────────────────────────
check E1 3 "$DATA/ragged.csv" < /dev/null
check E2 3 --column nosuch "$DATA/orders.csv" < /dev/null
check E3 3 --freq nosuch "$DATA/orders.csv" < /dev/null
check E4 1 "$MISSING" < /dev/null

# ── U. Usage errors ──────────────────────────────────────────────────────────────────────
check U1 2 < /dev/null
check U2 2 "$DATA/orders.csv" "$DATA/sensors.csv" < /dev/null
check U3 2 --freq sensor --top 0 "$DATA/sensors.csv" < /dev/null
check U4 2 --top 3 "$DATA/sensors.csv" < /dev/null
check U5 2 --column qty --freq qty "$DATA/orders.csv" < /dev/null
check U6 2 --frobnicate "$DATA/orders.csv" < /dev/null

TOTAL=$((PASSED + FAILED))
echo "CsvStat: $PASSED passed, $FAILED failed (of $TOTAL)"
[ "$FAILED" -eq 0 ]
