#!/usr/bin/env python3
# Reference implementation of the CsvStat contract, used only to validate the grader.
# Exact arithmetic; raises on any half-cent tie so the exam never depends on a rounding mode.
import csv, datetime, re, sys
from fractions import Fraction
def die(code, m): print(m, file=sys.stderr); sys.exit(code)
a = sys.argv[1:]; column = freq = None; top = None
while a and a[0].startswith("--"):
    o = a.pop(0)
    if not a: die(2, "value")
    v = a.pop(0)
    if o == "--column": column = v
    elif o == "--freq": freq = v
    elif o == "--top":
        if not re.fullmatch(r"[0-9]+", v) or int(v) < 1: die(2, "top")
        top = int(v)
    else: die(2, "option")
if len(a) != 1 or (column and freq) or (top is not None and not freq): die(2, "usage")
try: text = open(a[0], encoding="utf-8", newline="").read()
except OSError as e: die(1, str(e))
rows = list(csv.reader(text.splitlines(keepends=True)))
if not rows: die(3, "no header")
header = rows[0]; data = rows[1:]
if any(h.strip() == "" for h in header) or len({h.strip() for h in header}) != len(header): die(3, "bad header")
header = [h.strip() for h in header]
for i, r in enumerate(data, 2):
    if len(r) != len(header): die(3, f"line {i}: {len(r)} fields")
NULLS = {"", "na", "n/a", "null"}
def vals(ci): return [r[ci].strip() for r in data]
def nonnull(vs): return [v for v in vs if v.lower() not in NULLS]
def isdate(v):
    if not re.fullmatch(r"\d{4}-\d{2}-\d{2}", v): return False
    try: datetime.date.fromisoformat(v); return True
    except ValueError: return False
def kind(vs):
    if not vs: return "empty"
    if all(re.fullmatch(r"-?[0-9]+", v) for v in vs): return "integer"
    if all(re.fullmatch(r"-?[0-9]+(\.[0-9]+)?", v) for v in vs): return "number"
    if all(isdate(v) for v in vs): return "date"
    if all(v.lower() in ("true", "false") for v in vs): return "boolean"
    return "text"
def r2(x):
    c = x * 100
    if c.denominator == 2: raise SystemExit(f"TIE at {x}")
    n = abs(c); q = int(n + Fraction(1, 2))
    s = f"{q // 100}.{q % 100:02d}"
    return "-" + s if x < 0 and q else s
def stats(name, ci):
    vs = nonnull(vals(ci)); t = kind(vs)
    mn = mx = sm = mean = med = ""
    if t in ("integer", "number"):
        xs = sorted(Fraction(v) for v in vs); total = sum(xs)
        mid = len(xs) // 2
        m = xs[mid] if len(xs) % 2 else (xs[mid - 1] + xs[mid]) / 2
        if t == "integer": mn, mx, sm = str(int(xs[0])), str(int(xs[-1])), str(int(total))
        else: mn, mx, sm = r2(xs[0]), r2(xs[-1]), r2(total)
        mean, med = r2(total / len(xs)), r2(m)
    elif t == "date": mn, mx = min(vs), max(vs)
    elif t == "boolean":
        ls = {v.lower() for v in vs}; mn = "false" if "false" in ls else "true"; mx = "true" if "true" in ls else "false"
    elif t == "text":
        mn, mx = min(vs, key=lambda s: s.encode()), max(vs, key=lambda s: s.encode())
    return [name, t, len(data), len(data) - len(vs), len(set(vs)), mn, mx, sm, mean, med]
w = csv.writer(sys.stdout, lineterminator="\n")
if freq is not None:
    if freq not in header: die(3, "unknown column")
    counts = {}
    for v in nonnull(vals(header.index(freq))): counts[v] = counts.get(v, 0) + 1
    w.writerow(["value", "count"])
    for v, c in sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode()))[: top or 5]: w.writerow([v, c])
else:
    if column is not None and column not in header: die(3, "unknown column")
    w.writerow(["column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"])
    for ci, name in enumerate(header):
        if column is None or name == column: w.writerow(stats(name, ci))
