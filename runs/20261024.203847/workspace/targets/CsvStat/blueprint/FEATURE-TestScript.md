# FEATURE: Test Script

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | bin/test.sh runs go vet and go test from the application root and exits zero only when both pass. |
| Depends On  | FEATURE-ModuleBuild.md |
| Provides    | bin/test.sh |

## Purpose

`bin/test.sh` is the single project test entry point. It is written in POSIX sh.

## Behavior

- The script resolves the application root from its own location (the parent of the directory holding the script), changes into it, and does not depend on the caller's working directory.
- It runs `go vet ./...`, then `go test ./...`.
- It exits zero only when both commands pass; a failure of either exits non-zero. It must not mask the first failure.
- It uses no bash-only syntax.
- The application's own tests are written by the build agent beside the code and cover each type, each statistic, null handling, rounding, each output form, and each exit code using files written by the tests.

## Programmatic Acceptance

=== AC testscript-exists-and-passes ===
Intent: bin/test.sh exists and exits zero on the application root.
Requires: executable=go; scope=test

import subprocess
from pathlib import Path

script = Path("bin/test.sh")
assert script.is_file()
result = subprocess.run(["sh", str(script)], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC testscript-exists-and-passes ===

=== AC testscript-fails-on-failing-test ===
Intent: bin/test.sh exits non-zero when a test in the module fails.
Requires: executable=go; scope=test

import shutil
import subprocess
import tempfile
from pathlib import Path

with tempfile.TemporaryDirectory() as tmp:
    root = Path(tmp)
    (root / "bin").mkdir()
    shutil.copy("bin/test.sh", root / "bin" / "test.sh")
    (root / "go.mod").write_text("module tmpmod\n\ngo 1.22\n")
    (root / "x.go").write_text("package tmpmod\n")
    (root / "x_test.go").write_text(
        'package tmpmod\n\nimport "testing"\n\nfunc TestFails(t *testing.T) {\n\tt.Fatal("expected failure")\n}\n'
    )
    result = subprocess.run(["sh", str(root / "bin" / "test.sh")], capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr)
    assert result.returncode != 0
=== END AC testscript-fails-on-failing-test ===

=== AC testscript-fails-on-vet-finding ===
Intent: bin/test.sh exits non-zero when go vet reports a finding even though tests pass.
Requires: executable=go; scope=test

import shutil
import subprocess
import tempfile
from pathlib import Path

with tempfile.TemporaryDirectory() as tmp:
    root = Path(tmp)
    (root / "bin").mkdir()
    shutil.copy("bin/test.sh", root / "bin" / "test.sh")
    (root / "go.mod").write_text("module tmpmod\n\ngo 1.22\n")
    (root / "x.go").write_text(
        'package tmpmod\n\nimport "fmt"\n\nfunc Bad() string {\n\treturn fmt.Sprintf("%d", "text")\n}\n'
    )
    (root / "x_test.go").write_text(
        'package tmpmod\n\nimport "testing"\n\nfunc TestPasses(t *testing.T) {\n\t_ = Bad()\n}\n'
    )
    result = subprocess.run(["sh", str(root / "bin" / "test.sh")], capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr)
    assert result.returncode != 0
=== END AC testscript-fails-on-vet-finding ===

=== AC testscript-passes-on-clean-module ===
Intent: bin/test.sh exits zero for a module whose vet and tests are clean, from a different working directory.
Requires: executable=go; scope=test

import shutil
import subprocess
import tempfile
from pathlib import Path

with tempfile.TemporaryDirectory() as tmp, tempfile.TemporaryDirectory() as elsewhere:
    root = Path(tmp)
    (root / "bin").mkdir()
    shutil.copy("bin/test.sh", root / "bin" / "test.sh")
    (root / "go.mod").write_text("module tmpmod\n\ngo 1.22\n")
    (root / "x.go").write_text("package tmpmod\n\nfunc One() int { return 1 }\n")
    (root / "x_test.go").write_text(
        'package tmpmod\n\nimport "testing"\n\nfunc TestOne(t *testing.T) {\n\tif One() != 1 {\n\t\tt.Fatal("bad")\n\t}\n}\n'
    )
    result = subprocess.run(
        ["sh", str(root / "bin" / "test.sh")], capture_output=True, text=True, cwd=elsewhere
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
=== END AC testscript-passes-on-clean-module ===

## User Acceptance

- None.

## Guardrails

- POSIX sh only; no bash-only features.
- The script never modifies files under `sources/`.
