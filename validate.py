#!/usr/bin/env python3
"""Sanity-check the content.

  python3 validate.py        vet everything; every exercise solution must
                             reproduce expected.txt (or pass check_test.go)
                             and every starter must NOT.
  python3 validate.py --gen  (re)generate expected.txt from each solution
                             first, then check as above.

Exercises live in exercises/<slug>/: meta.json (base, file, prompt),
starter.go, solution.go and expected.txt or check_test.go.
"""
import json, os, shutil, subprocess, sys, tempfile

ROOT = os.path.dirname(os.path.abspath(__file__))
ENV = {**os.environ, "GOFLAGS": "-mod=mod -buildvcs=false"}


def stage(base, overlay_name, overlay_path, extra=None):
    d = tempfile.mkdtemp(prefix="rogue-validate-")
    src = os.path.join(ROOT, base)
    for f in os.listdir(src):
        if f.endswith(".go") and not f.endswith("_test.go"):
            shutil.copy(os.path.join(src, f), d)
    shutil.copy(overlay_path, os.path.join(d, overlay_name))
    if extra:
        shutil.copy(extra, os.path.join(d, os.path.basename(extra)))
    for f in ("go.mod", "go.sum"):
        shutil.copy(os.path.join(ROOT, f), d)
    return d


def go(d, *args, timeout=120):
    r = subprocess.run(["go", *args], cwd=d, capture_output=True, text=True, env=ENV, timeout=timeout)
    return r.returncode == 0, r.stdout, r.stderr


def norm(s):
    return "\n".join(l.rstrip() for l in s.strip().split("\n"))


def run_exercise(meta, slug, kind):
    """Returns (passed, output)."""
    d = os.path.join(ROOT, "exercises", slug)
    test = os.path.join(d, "check_test.go")
    expected = os.path.join(d, "expected.txt")
    if os.path.exists(test):
        tmp = stage(meta["base"], meta["file"], os.path.join(d, kind), test)
        ok, out, err = go(tmp, "test", "-count=1", "-run", ".", ".")
        shutil.rmtree(tmp)
        return ok, out + err
    tmp = stage(meta["base"], meta["file"], os.path.join(d, kind))
    ok, out, err = go(tmp, "run", ".", timeout=60)
    shutil.rmtree(tmp)
    if not ok:
        return False, err
    if kind == "solution.go" and "--gen" in sys.argv:
        with open(expected, "w") as fh:
            fh.write(out.rstrip("\n") + "\n")
    if not os.path.exists(expected):
        return False, "no expected.txt (run with --gen)"
    with open(expected) as fh:
        return norm(out) == norm(fh.read()), out


def main():
    ok = True
    vet, _, err = go(ROOT, "vet", "./...")
    print("go vet ./...:", "ok" if vet else "FAIL\n" + err)
    ok &= vet

    exdir = os.path.join(ROOT, "exercises")
    for slug in sorted(os.listdir(exdir)):
        meta_path = os.path.join(exdir, slug, "meta.json")
        if not os.path.exists(meta_path):
            continue
        meta = json.load(open(meta_path))
        for kind, expect in (("solution.go", True), ("starter.go", False)):
            passed, out = run_exercise(meta, slug, kind)
            good = passed == expect
            ok &= good
            print(f"{'ok  ' if good else 'FAIL'} {slug:32} {kind:12} -> {'pass' if passed else 'fail'} (expected {'pass' if expect else 'fail'})")
            if not good:
                print(out[:1500])
    print("ALL GOOD" if ok else "PROBLEMS FOUND")
    sys.exit(0 if ok else 1)


if __name__ == "__main__":
    main()
