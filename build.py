#!/usr/bin/env python3
"""Generate data.js and exercises.json for Cutesy Rogue.

One continuous course: steps/NN-slug/*.go is the program as it stands after
step NN, and course/NN-slug.md explains what that step added, using
{{file x.go}} (whole file) and {{diff x.go}} (exact changes vs the previous
step) so the shown code can never drift from the real code.

Exercises live in exercises/<slug>/ with meta.json + starter.go and either
expected.txt (stdout comparison) or check_test.go (hidden go test). Only the
prompt and starter are shipped to the browser.
"""
import difflib, html, json, os, re, subprocess, sys, tempfile, shutil

ROOT = os.path.dirname(os.path.abspath(__file__))
STEPS = os.path.join(ROOT, "steps")
COURSE = os.path.join(ROOT, "course")
EXERCISES = os.path.join(ROOT, "exercises")


def inline(text):
    """Inline markdown -> HTML: code, bold, italics, links, kbd."""
    text = html.escape(text, quote=False)
    text = re.sub(r"`([^`]+)`", r"<code>\1</code>", text)
    text = re.sub(r"\*\*([^*]+)\*\*", r"<strong>\1</strong>", text)
    text = re.sub(r"(?<![\w*])\*([^*\n]+)\*(?![\w*])", r"<em>\1</em>", text)
    text = re.sub(r"\[\[([^\]]+)\]\]", r"<kbd>\1</kbd>", text)
    text = re.sub(r"\[([^\]]+)\]\(([^)]+)\)", r'<a href="\2" target="_blank" rel="noopener">\1</a>', text)
    return text


def render_file(name, code, status):
    return (f'<div class="snippet {status}"><span class="snippet-file">{html.escape(name)}</span>'
            f'<pre><code class="language-go">{html.escape(code)}</code></pre></div>')


def render_diff(name, old, new, context=3):
    """A line diff between two versions of a file, with line numbers of the
    NEW file, green added lines and red removed lines. Every line is emitted;
    lines far from any change get the class "far" and are hidden unless the
    block is in whole-file mode. Short files, or files where the hunks cover
    most of the file anyway, start in whole-file mode."""
    a, b = old.split("\n"), new.split("\n")
    sm = difflib.SequenceMatcher(None, a, b, autojunk=False)
    rows = []
    for tag, i1, i2, j1, j2 in sm.get_opcodes():
        if tag == "equal":
            for k in range(i2 - i1):
                rows.append(("ctx", j1 + k + 1, b[j1 + k]))
        else:
            for k in range(i1, i2):
                rows.append(("del", None, a[k]))
            for k in range(j1, j2):
                rows.append(("add", k + 1, b[k]))
    keep = [False] * len(rows)
    for idx, (kind, _, _) in enumerate(rows):
        if kind != "ctx":
            for k in range(max(0, idx - context), min(len(rows), idx + context + 1)):
                keep[k] = True
    out, in_gap = [], False
    for idx, (kind, ln, text) in enumerate(rows):
        if not keep[idx] and not in_gap:
            out.append('<div class="diff-gap">· · ·</div>')
            in_gap = True
        if keep[idx]:
            in_gap = False
        n = "" if ln is None else str(ln)
        mark = {"add": "+", "del": "−", "ctx": " "}[kind]
        far = "" if keep[idx] else " far"
        out.append(f'<div class="diff-line {kind}{far}"><span class="ln">{n}</span><span class="mark">{mark}</span>'
                   f'<code>{html.escape(text)}</code></div>')
    shown = sum(keep)
    full = " full" if len(rows) <= 60 or shown > 0.6 * len(rows) else ""
    return (f'<div class="snippet changed diff{full}"><span class="snippet-file">{html.escape(name)}</span>'
            f'<pre>{"".join(out)}</pre></div>')


def expand_directives(body, files, prev):
    """{{file x.go}} / {{diff x.go}} -> one-line placeholders + rendered blocks."""
    blocks = []

    def repl(m):
        kind, name = m.group(1), m.group(2)
        cur = files.get(name)
        if cur is None:
            raise SystemExit(f"directive refers to unknown file {name}")
        if kind == "file" or name not in prev:
            blocks.append(render_file(name, cur, "new" if name not in prev else ""))
        else:
            blocks.append(render_diff(name, prev[name], cur))
        return f"@@BLOCK{len(blocks) - 1}@@"

    return re.sub(r"^\{\{(file|diff) ([\w.]+)\}\}$", repl, body, flags=re.M), blocks


def markdown(src):
    """A small markdown subset -> HTML."""
    out, lines, i = [], src.split("\n"), 0
    para = []

    def flush():
        if para:
            out.append("<p>" + inline(" ".join(para)) + "</p>")
            para.clear()

    while i < len(lines):
        ln = lines[i]
        m = re.match(r"^```(\w+)?\s*(.*)$", ln)
        if m:
            flush()
            lang, label = m.group(1) or "", m.group(2).strip()
            i += 1
            code = []
            while i < len(lines) and not lines[i].startswith("```"):
                code.append(lines[i]); i += 1
            i += 1
            cls = f' class="language-{lang}"' if lang else ""
            cap = f'<span class="snippet-file">{html.escape(label)}</span>' if label else ""
            out.append(f'<div class="snippet">{cap}<pre><code{cls}>{html.escape(chr(10).join(code))}</code></pre></div>')
            continue
        if re.match(r"^@@BLOCK\d+@@$", ln):
            flush()
            out.append(ln)
        elif ln.startswith(">>> "):
            # a "try it first" card: consecutive >>> lines are its paragraphs
            flush()
            paras = []
            while i < len(lines) and lines[i].startswith(">>> "):
                paras.append(inline(lines[i][4:])); i += 1
            out.append('<div class="try"><div class="try-head">🧪 try it first</div>' +
                       "".join(f"<p>{x}</p>" for x in paras) +
                       '<div class="try-actions"><button class="ghost-btn try-load">load the previous step into the editor</button>'
                       '<label class="try-always"><input type="checkbox" class="try-always-box"> always show solutions</label></div></div>')
            continue
        elif ln.strip() == "--- reveal":
            flush()
            out.append('<details class="reveal"><summary>show me how it is done in this step</summary><div class="reveal-body">')
        elif ln.strip() == "--- end":
            flush()
            out.append('</div></details>')
        elif ln.startswith("%%% "):
            flush()
            out.append(f'<div class="experiment">{inline(ln[4:])}</div>')
        elif re.match(r"^### Step (\d+):?\s*(.*)$", ln):
            flush()
            m = re.match(r"^### Step (\d+):?\s*(.*)$", ln)
            out.append(f'<div class="step"><span class="step-n">step {m.group(1)}</span><h3>{inline(m.group(2))}</h3></div>')
        elif ln.startswith("!!! "):
            flush()
            out.append(f'<div class="expect">{inline(ln[4:])}</div>')
        elif ln.startswith("??? "):
            flush()
            out.append(f'<div class="pitfall">{inline(ln[4:])}</div>')
        elif ln.startswith("#"):
            flush()
            level = len(ln) - len(ln.lstrip("#"))
            out.append(f"<h{level + 1}>{inline(ln[level:].strip())}</h{level + 1}>")
        elif re.match(r"^\s*[-*] ", ln):
            flush()
            items = []
            while i < len(lines) and re.match(r"^\s*[-*] ", lines[i]):
                items.append(re.sub(r"^\s*[-*] ", "", lines[i])); i += 1
            out.append("<ul>" + "".join(f"<li>{inline(x)}</li>" for x in items) + "</ul>")
            continue
        elif re.match(r"^\s*\d+\. ", ln):
            flush()
            items = []
            while i < len(lines) and re.match(r"^\s*\d+\. ", lines[i]):
                items.append(re.sub(r"^\s*\d+\. ", "", lines[i])); i += 1
            out.append("<ol>" + "".join(f"<li>{inline(x)}</li>" for x in items) + "</ol>")
            continue
        elif ln.startswith(">"):
            flush()
            quote = []
            while i < len(lines) and lines[i].startswith(">"):
                quote.append(lines[i][1:].strip()); i += 1
            out.append('<blockquote>' + inline(" ".join(quote)) + "</blockquote>")
            continue
        elif ln.strip() == "":
            flush()
        else:
            para.append(ln.strip())
        i += 1
    flush()
    return "\n".join(out)


def go_files(d):
    return [f for f in sorted(os.listdir(d)) if f.endswith(".go") and not f.endswith("_test.go")]



def build_steps():
    """One continuous course: steps/NN-slug/*.go explained by course/NN-slug.md.
    A step's markdown starts with '# Step N · Title'; an optional second line
    '## Chapter: Name' opens a new chapter in the sidebar."""
    parts, prev = [], {}
    for part in sorted(d for d in os.listdir(STEPS) if os.path.isdir(os.path.join(STEPS, d))):
        md_path = os.path.join(COURSE, part + ".md")
        if not os.path.exists(md_path):
            raise SystemExit(f"missing lesson text {md_path}")
        with open(md_path) as fh:
            src = fh.read()
        title_m = re.match(r"^# (.+)$", src, re.M)
        title = title_m.group(1).strip() if title_m else part
        body = src[title_m.end():] if title_m else src
        chapter = ""
        ch_m = re.match(r"^\s*## Chapter: (.+)$", body, re.M)
        if ch_m:
            chapter = ch_m.group(1).strip()
            body = body[:ch_m.start()] + body[ch_m.end():]
        files = []
        for f in go_files(os.path.join(STEPS, part)):
            with open(os.path.join(STEPS, part, f)) as fh:
                code = fh.read()
            old = prev.get(f)
            files.append({"name": f, "code": code,
                          "status": "new" if old is None else ("changed" if old != code else "same")})
        body, blocks = expand_directives(body, {f["name"]: f["code"] for f in files}, prev)
        prev = {f["name"]: f["code"] for f in files}
        prose = re.sub(r"@@BLOCK(\d+)@@", lambda m: blocks[int(m.group(1))], markdown(body))
        parts.append({"slug": part, "track": "course", "chapter": chapter, "kind": "tui", "title": title,
                      "prose": prose, "files": files, "output": "", "base": f"steps/{part}"})
    return parts


def build_exercises():
    exercises = {}
    for slug in sorted(os.listdir(EXERCISES)):
        d = os.path.join(EXERCISES, slug)
        meta_path = os.path.join(d, "meta.json")
        if not os.path.exists(meta_path):
            continue
        with open(meta_path) as fh:
            meta = json.load(fh)
        with open(os.path.join(d, "starter.go")) as fh:
            starter = fh.read()
        exercises[slug] = {"prompt": meta["prompt"], "base": meta["base"], "file": meta["file"], "starter": starter}
    return exercises


def main():
    lessons = build_steps()
    exercises = build_exercises()
    with open(os.path.join(ROOT, "data.js"), "w") as fh:
        fh.write("const LESSONS = ")
        json.dump(lessons, fh, ensure_ascii=False, separators=(",", ":"))
        fh.write(";\n")
    with open(os.path.join(ROOT, "exercises.json"), "w") as fh:
        json.dump(exercises, fh, ensure_ascii=False, indent=1)
    print(f"{len(lessons)} steps, {len(exercises)} exercises -> data.js, exercises.json")


if __name__ == "__main__":
    main()
