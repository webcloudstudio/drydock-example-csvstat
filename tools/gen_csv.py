#!/usr/bin/env python3
# Generate the CsvStat fixtures into <outdir> from a fixed seed.
import csv, random, sys, os
out = sys.argv[1]; rng = random.Random(int(sys.argv[2]) if len(sys.argv) > 2 else 7)
def write(name, header, rows):
    with open(os.path.join(out, name), "w", newline="") as f:
        w = csv.writer(f, lineterminator="\n"); w.writerow(header); w.writerows(rows)
first = ["Ada", "Grace", "Alan", "Edsger", "Barbara", "Donald", "Ken", "Dennis", "Margaret", "Linus",
         "Radia", "Frances", "John", "Niklaus", "Leslie", "Tony", "Barbara", "Guido", "Bjarne", "Anita",
         "Sophie", "Lynn", "Jean", "Hedy"]
depts = ["Engineering", "Finance", "Sales", "Support"]
rows = []
for i in range(24):
    mgr = rng.choice(["", "NA", str(rng.randint(1, 6))]) if i > 0 else ""
    note = rng.choice(["", "", "", "remote, part-time", "on leave", "relocating to Austin, TX", "mentor"])
    rows.append([1001 + i, first[i], rng.choice(depts),
                 f"20{rng.randint(15, 25)}-{rng.randint(1, 12):02d}-{rng.randint(1, 28):02d}",
                 f"{rng.randint(4200000, 15800000) / 100:.2f}",
                 rng.choice(["true", "false", "TRUE", "False"]), mgr, note])
write("employees.csv", ["id", "name", "department", "hire_date", "salary", "remote", "manager_id", "notes"], rows)
rows = []
sensors = ["north", "south", "east", "west"]
for i in range(36):
    temp = rng.choice(["N/A"] + [f"{rng.randint(-1250, 3150) / 100:.2f}"] * 9)
    hum = rng.choice([""] + [str(rng.randint(18, 97))] * 7)
    rows.append([i + 1, f"2025-03-{1 + i // 12:02d}T{8 + i % 12:02d}:00", sensors[i % 4], temp, hum,
                 rng.choice(["true", "false"]), rng.choice(["", "NULL", "null", " "])])
write("sensors.csv", ["reading_id", "recorded_at", "sensor", "temp_c", "humidity", "ok", "battery"], rows)
rows = []
customers = ["Acme Corp", "  Globex  ", "Initech", "Umbrella", "Acme Corp ", "Hooli"]
for i in range(30):
    rows.append([5000 + i, f"2025-02-{rng.randint(1, 28):02d}", rng.choice(customers), rng.randint(1, 12),
                 f"{rng.randint(199, 25000) / 100:.2f}", rng.choice(["", "", "", "SAVE10", "WELCOME"]),
                 rng.choice(["12", "12.5", "7", "A7", "-3"]),
                 "2025-02-30" if i == 17 else f"2025-03-{rng.randint(1, 28):02d}"])
write("orders.csv", ["order_id", "order_date", "customer", "qty", "unit_price", "coupon", "ref", "shipped"], rows)
write("header_only.csv", ["sku", "price", "in_stock"], [])
with open(os.path.join(out, "ragged.csv"), "w") as f:
    f.write("sku,price,in_stock\nA1,9.99,true\nA2,4.50,false\nA3,1.25,true,extra\nA4,3.00,false\n")
