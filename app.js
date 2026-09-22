/* Cutesy Rogue — app logic */
const $ = (id) => document.getElementById(id);

const state = {
  idx: -1,
  file: 0, // index of the file tab currently shown
  done: new Set(JSON.parse(localStorage.getItem("rogueCuteDone") || "[]")),
  drafts: JSON.parse(localStorage.getItem("rogueCuteDrafts") || "{}"),   // "slug/file.go" -> code
  exDone: new Set(JSON.parse(localStorage.getItem("rogueCuteExDone") || "[]")),
  exDrafts: JSON.parse(localStorage.getItem("rogueCuteExDrafts") || "{}"),
  reveal: !!localStorage.getItem("rogueCuteReveal"),
};

// loadPreviousStep puts the previous step's version of every file into this
// step's editor, so the reader can attempt the step themselves and build/play
// it. The reader's own edits of the previous step win over the site's code,
// so their style carries forward. Files new in this step start as an empty package.
function loadPreviousStep(i) {
  const l = LESSONS[i], prev = LESSONS[i - 1];
  const before = Object.fromEntries(prev.files.map((f) => [f.name, state.drafts[draftKey(prev, f)] ?? f.code]));
  const mine = prev.files.some((f) => state.drafts[draftKey(prev, f)] !== undefined);
  l.files.forEach((f) => {
    const code = before[f.name] ?? "package main\n";
    if (code === f.code) delete state.drafts[draftKey(l, f)];
    else state.drafts[draftKey(l, f)] = code;
  });
  save();
  buildTabs(l);
  showFile(l, state.file);
  squeak(mine ? "loaded your version of the previous step." : "loaded the previous step.");
  $("codeHeading").scrollIntoView({ behavior: "smooth", block: "start" });
}

let EXERCISES = {};

const save = () => {
  localStorage.setItem("rogueCuteDone", JSON.stringify([...state.done]));
  localStorage.setItem("rogueCuteDrafts", JSON.stringify(state.drafts));
  localStorage.setItem("rogueCuteExDone", JSON.stringify([...state.exDone]));
  localStorage.setItem("rogueCuteExDrafts", JSON.stringify(state.exDrafts));
};

/* ---------- sidebar ---------- */
function buildNav(filter = "") {
  const list = $("navList");
  list.innerHTML = "";
  const f = filter.trim().toLowerCase();
  LESSONS.forEach((l, i) => {
    if (f && !l.title.toLowerCase().includes(f) && !l.slug.includes(f)) return;
    if (l.chapter) {
      const h = document.createElement("li");
      h.className = "nav-group";
      h.textContent = l.chapter;
      list.appendChild(h);
    }
    const li = document.createElement("li");
    const btn = document.createElement("button");
    btn.textContent = l.title;
    btn.dataset.idx = i;
    if (i === state.idx) btn.classList.add("active");
    if (state.done.has(l.slug)) {
      const star = document.createElement("span");
      star.className = "done-star";
      star.textContent = "⚔️";
      btn.appendChild(star);
    }
    btn.addEventListener("click", () => openLesson(i));
    li.appendChild(btn);
    list.appendChild(li);
  });
}

function updateProgress() {
  const n = state.done.size, total = LESSONS.length;
  $("progressFill").style.width = (n / total * 100) + "%";
  $("progressLabel").textContent = `${n} / ${total}`;
}

/* ---------- mascot chatter ---------- */
const CHATTER = [
  "kobold says: pick a step.", "kobold says: the @ is you. I am the k.",
  "kobold says: edit the code; the build button tells you what broke.",
  "kobold says: hjkl work in the game, so do the arrows.", "kobold says: gofmt everything.",
  "kobold says: a corpse is just an entity with a different glyph.",
  "kobold says: I get 10 XP in most games. Be kind.",
];
let chatterTimer;
function squeak(msg) {
  $("speechText").textContent = msg;
  $("speech").classList.remove("hidden");
  clearTimeout(chatterTimer);
  chatterTimer = setTimeout(() => $("speech").classList.add("hidden"), 2600);
}
$("mascot").addEventListener("click", () => squeak(CHATTER[Math.floor(Math.random() * CHATTER.length)]));

/* ---------- confetti ---------- */
function confetti() {
  const bits = ["@", "%", "!", "~", ">", "#"];
  for (let i = 0; i < 28; i++) {
    const s = document.createElement("span");
    s.className = "confetti";
    s.textContent = bits[Math.floor(Math.random() * bits.length)];
    s.style.left = Math.random() * 100 + "vw";
    s.style.animationDuration = (2 + Math.random() * 2) + "s";
    s.style.animationDelay = Math.random() * 0.6 + "s";
    s.style.fontSize = 14 + Math.random() * 14 + "px";
    document.body.appendChild(s);
    setTimeout(() => s.remove(), 4800);
  }
}

/* ---------- lesson view ---------- */
const draftKey = (l, file) => `${l.slug}/${file.name}`;
const codeFor = (l, file) => state.drafts[draftKey(l, file)] ?? file.code;

function openLesson(i, push = true) {
  stopGame();
  $("termWrap").classList.add("hidden");
  state.idx = i;
  const l = LESSONS[i];
  $("welcome").classList.add("hidden");
  $("lessonView").classList.remove("hidden");

  $("lessonNum").textContent = `${i + 1} / ${LESSONS.length}`;
  document.body.dataset.kind = l.kind;
  $("sampleOut").classList.toggle("hidden", !l.output);
  $("sampleOut").textContent = "";
  if (l.output) {
    const label = document.createElement("span");
    label.className = "out-label";
    label.textContent = "sample output 🐾";
    $("sampleOut").appendChild(label);
    $("sampleOut").appendChild(document.createTextNode(l.output));
  }
  $("codeHeading").textContent = "The whole program after this step";
  $("lessonTitle").textContent = l.title;
  document.title = `${l.title} · Golang Roguelike Tutorial`;

  const prose = $("lessonProse");
  prose.innerHTML = l.prose;
  Prism.highlightAllUnder(prose);
  // diff blocks are highlighted line by line so the +/- backgrounds survive
  prose.querySelectorAll(".diff-line code").forEach((c) => {
    c.innerHTML = Prism.highlight(c.textContent, Prism.languages.go, "go");
  });
  enhanceSnippets(prose);
  // "try it first": hidden solutions, load-previous-step, always-show preference
  prose.querySelectorAll("details.reveal").forEach((d) => { if (state.reveal) d.open = true; });
  prose.querySelectorAll(".try-always-box").forEach((b) => {
    b.checked = state.reveal;
    b.addEventListener("change", () => {
      state.reveal = b.checked;
      localStorage.setItem("rogueCuteReveal", state.reveal ? "1" : "");
      prose.querySelectorAll("details.reveal").forEach((d) => { d.open = state.reveal; });
    });
  });
  prose.querySelectorAll(".try-load").forEach((b) => {
    if (i === 0) { b.remove(); return; }
    b.addEventListener("click", () => loadPreviousStep(i));
  });

  $("runCmd").textContent = `go run ./${l.base}`;
  // open the first new/changed file, or main.go
  state.file = Math.max(0, l.files.findIndex((f) => f.status !== "same"));
  buildTabs(l);
  showFile(l, state.file);
  $("outputBox").classList.add("hidden");
  $("runHint").textContent = "";
  updateDoneBtn();
  loadExercise(l);

  $("prevBtn").disabled = i === 0;
  $("nextBtn").disabled = i === LESSONS.length - 1;

  buildNav($("search").value);
  if (push) history.replaceState(null, "", "#" + l.slug);
  $("content").scrollTop = 0;
  window.scrollTo({ top: 0, behavior: "smooth" });
}

function buildTabs(l) {
  const tabs = $("fileTabs");
  tabs.innerHTML = "";
  l.files.forEach((f, j) => {
    const b = document.createElement("button");
    b.className = "filetab" + (j === state.file ? " active" : "") + (state.drafts[draftKey(l, f)] !== undefined ? " edited" : "");
    b.textContent = f.name;
    if (f.status !== "same") {
      const badge = document.createElement("i");
      badge.className = "badge " + f.status;
      badge.textContent = f.status === "new" ? "new" : "changed";
      b.appendChild(badge);
    }
    b.addEventListener("click", () => { state.file = j; buildTabs(l); showFile(l, j); });
    tabs.appendChild(b);
  });
}

function showFile(l, j) {
  setCode($("codeArea"), codeFor(l, l.files[j]));
}

function updateDoneBtn() {
  const l = LESSONS[state.idx];
  const btn = $("doneBtn");
  const done = state.done.has(l.slug);
  btn.classList.toggle("is-done", done);
  btn.innerHTML = done ? "mark as not done" : "mark done";
}

$("doneBtn").addEventListener("click", () => {
  const l = LESSONS[state.idx];
  if (state.done.has(l.slug)) {
    state.done.delete(l.slug);
    squeak("step marked as not done.");
  } else {
    state.done.add(l.slug);
    confetti();
    squeak(state.done.size === LESSONS.length
      ? "kobold says: all steps done. You built a roguelike."
      : "kobold says: step done. Next?");
  }
  save(); updateProgress(); updateDoneBtn(); buildNav($("search").value);
});

/* ---------- code editing & building ---------- */
$("codeArea").addEventListener("input", () => {
  const l = LESSONS[state.idx];
  const f = l.files[state.file];
  const v = $("codeArea").value;
  if (v === f.code) delete state.drafts[draftKey(l, f)];
  else state.drafts[draftKey(l, f)] = v;
  save();
  buildTabs(l);
});
$("resetCodeBtn").addEventListener("click", () => {
  const l = LESSONS[state.idx];
  const f = l.files[state.file];
  delete state.drafts[draftKey(l, f)];
  setCode($("codeArea"), f.code);
  save(); buildTabs(l);
  squeak("restored the original file.");
});
$("copyRunBtn").addEventListener("click", () => {
  navigator.clipboard?.writeText($("runCmd").textContent);
  squeak("copied.");
});

function showOutput(box, ok, text, okLabel, errLabel) {
  box.classList.remove("hidden", "err", "pass");
  box.classList.add(ok ? "pass" : "err");
  box.textContent = "";
  const label = document.createElement("span");
  label.className = "out-label";
  label.textContent = ok ? okLabel : errLabel;
  box.appendChild(label);
  box.appendChild(document.createTextNode(text || "(no output)"));
}

async function buildCode() {
  const l = LESSONS[state.idx];
  const btn = $("runBtn");
  const box = $("outputBox");
  const isRun = l.kind === "stdout";
  btn.disabled = true;
  btn.textContent = isRun ? "running…" : "building…";
  $("runHint").textContent = isRun ? "compiling and running with a real Go toolchain…" : "go vet with a real toolchain (every file, with your edits)…";
  box.classList.remove("hidden", "err", "pass");
  box.textContent = "compiling…";
  try {
    const res = await fetch(isRun ? "/api/run" : "/api/build", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ base: l.base, files: editedFiles(l) }),
    });
    const data = await res.json();
    if (isRun) {
      let out = data.output || "";
      if (data.buildErr) out = data.buildErr;
      if (data.stderr) out += (out ? "\n" : "") + data.stderr;
      const ok = !data.buildErr && !data.stderr;
      showOutput(box, ok, out, "output 🐾", "uh-oh! 🙀");
      squeak(ok ? "it runs." : "it failed; see the output.");
    } else {
      showOutput(box, data.ok, data.output, "it builds", "build error");
      squeak(data.ok ? "it compiles. press play." : "build failed; see the output.");
    }
  } catch (e) {
    box.classList.add("err");
    box.textContent = "no runner available 😿 (start the server: docker compose up, or go run ./server)";
  } finally {
    btn.disabled = false;
    btn.textContent = isRun ? "run it" : "just build";
    $("runHint").textContent = "";
  }
}
$("runBtn").addEventListener("click", buildCode);
$("runBtn2").addEventListener("click", buildCode);

/* ---------- play here: a real PTY streamed into an xterm.js terminal ---------- */
let term = null, sock = null;

function editedFiles(l) {
  const files = {};
  l.files.forEach((f) => { const d = state.drafts[draftKey(l, f)]; if (d !== undefined) files[f.name] = d; });
  return files;
}

function stopGame() {
  if (sock) { sock.onclose = null; sock.close(); sock = null; }
}

function playHere() {
  const l = LESSONS[state.idx];
  stopGame();
  $("outputBox").classList.add("hidden");
  $("termWrap").classList.remove("hidden");
  if (!term) {
    term = new Terminal({
      cols: 80, rows: 50, cursorBlink: true, fontSize: 14,
      fontFamily: '"Fira Code", "JetBrains Mono", Consolas, monospace',
      theme: { background: "#000000", foreground: "#e0e0e0", cursor: "#ff8fc7" },
    });
    term.open($("term"));
    term.onData((d) => { if (sock && sock.readyState === 1) sock.send(d); });
  }
  term.reset();
  term.focus();
  const proto = location.protocol === "https:" ? "wss:" : "ws:";
  sock = new WebSocket(`${proto}//${location.host}/api/pty?base=${encodeURIComponent(l.base)}`);
  sock.binaryType = "arraybuffer";
  sock.onopen = () => sock.send(JSON.stringify({ files: editedFiles(l) }));
  sock.onmessage = (e) => {
    if (typeof e.data === "string") term.write("\r\n\x1b[33m" + e.data.replace(/\n/g, "\r\n") + "\x1b[0m\r\n");
    else term.write(new Uint8Array(e.data));
  };
  sock.onclose = () => { sock = null; };
  sock.onerror = () => term.write("\r\n\x1b[31mno game server 😿 (start it: go run ./server)\x1b[0m\r\n");
  squeak("arrows or hjkl move, Esc quits.");
  $("termWrap").scrollIntoView({ behavior: "smooth", block: "nearest" });
}
$("playBtn").addEventListener("click", playHere);
$("termClose").addEventListener("click", () => { stopGame(); $("termWrap").classList.add("hidden"); });

/* ---------- play in a terminal window on this machine ---------- */
async function openTerminal() {
  const l = LESSONS[state.idx];
  const btn = $("openBtn");
  const box = $("outputBox");
  btn.disabled = true;
  $("runHint").textContent = "building, then opening your terminal emulator…";
  box.classList.remove("hidden", "err", "pass");
  box.textContent = "building…";
  try {
    const res = await fetch("/api/open", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ base: l.base, files: editedFiles(l) }),
    });
    const data = await res.json();
    showOutput(box, data.ok, data.output, "launched", "error");
    squeak(data.ok ? "a terminal window should have opened." : "see the message below.");
  } catch (e) {
    box.classList.add("err");
    box.textContent = "no game server 😿 (start it on this machine: go run ./server)";
  } finally {
    btn.disabled = false;
    $("runHint").textContent = "";
  }
}
$("openBtn").addEventListener("click", openTerminal);

/* ---------- nav buttons ---------- */
$("prevBtn").addEventListener("click", () => openLesson(state.idx - 1));
$("nextBtn").addEventListener("click", () => {
  const l = LESSONS[state.idx];
  if (!state.done.has(l.slug)) {
    state.done.add(l.slug);
    save(); updateProgress();
  }
  openLesson(state.idx + 1);
});

$("startBtn").addEventListener("click", () => openLesson(0));
$("search").addEventListener("input", (e) => buildNav(e.target.value));
$("resetBtn").addEventListener("click", () => {
  state.done.clear();
  state.exDone.clear();
  Object.keys(state.drafts).forEach((k) => delete state.drafts[k]);
  Object.keys(state.exDrafts).forEach((k) => delete state.exDrafts[k]);
  save(); updateProgress(); buildNav($("search").value);
  if (state.idx >= 0) { const l = LESSONS[state.idx]; buildTabs(l); showFile(l, state.file); loadExercise(l); }
  squeak("progress reset.");
});

document.addEventListener("keydown", (e) => {
  if (e.target.tagName === "TEXTAREA" || e.target.tagName === "INPUT") return;
  if ((e.ctrlKey || e.metaKey) && e.key === "Enter") { buildCode(); return; }
  if (state.idx < 0) return;
  if (e.key === "ArrowRight" && state.idx < LESSONS.length - 1) openLesson(state.idx + 1);
  if (e.key === "ArrowLeft" && state.idx > 0) openLesson(state.idx - 1);
});

/* ---------- copy buttons + resizable code blocks ---------- */
function copyText(text, what) {
  const done = () => squeak(`${what} copied! 📋`);
  if (navigator.clipboard?.writeText) navigator.clipboard.writeText(text).then(done, () => fallbackCopy(text, done));
  else fallbackCopy(text, done);
}
function fallbackCopy(text, done) {
  const t = document.createElement("textarea");
  t.value = text; t.style.position = "fixed"; t.style.opacity = "0";
  document.body.appendChild(t); t.select();
  try { document.execCommand("copy"); done(); } catch (e) { squeak("couldn't copy 😿 select it by hand~"); }
  t.remove();
}
// what a block "is" when copied: a diff copies the resulting code (kept + added
// lines, without the removed ones), everything else copies its text.
function snippetText(snip) {
  if (snip.classList.contains("diff")) {
    return [...snip.querySelectorAll(".diff-line:not(.del) code")].map((c) => c.textContent).join("\n") + "\n";
  }
  const code = snip.querySelector("pre code") || snip.querySelector("pre");
  return code.textContent;
}
function makeResizable(body, bar) {
  bar.addEventListener("pointerdown", (e) => {
    e.preventDefault();
    const y0 = e.clientY, h0 = body.offsetHeight;
    const move = (ev) => { body.style.maxHeight = "none"; body.style.height = Math.max(80, h0 + ev.clientY - y0) + "px"; };
    const up = () => { window.removeEventListener("pointermove", move); window.removeEventListener("pointerup", up); };
    window.addEventListener("pointermove", move);
    window.addEventListener("pointerup", up);
  });
}
function enhanceSnippets(root) {
  root.querySelectorAll(".snippet").forEach((snip) => {
    if (snip.dataset.enhanced) return;
    snip.dataset.enhanced = "1";
    let cap = snip.querySelector(".snippet-file");
    if (!cap) {
      cap = document.createElement("span");
      cap.className = "snippet-file";
      cap.textContent = snip.classList.contains("language-sh") || snip.querySelector("code.language-sh") ? "shell" : "";
      snip.prepend(cap);
    }
    const btn = document.createElement("button");
    btn.className = "ghost-btn snippet-copy";
    btn.textContent = "copy";
    btn.title = snip.classList.contains("diff") ? "copy the resulting code (added + unchanged lines)" : "copy this code";
    btn.addEventListener("click", () => copyText(snippetText(snip), cap.textContent.trim() || "code"));
    cap.appendChild(btn);
    if (snip.classList.contains("diff") && snip.querySelector(".diff-line.far")) {
      const tog = document.createElement("button");
      tog.className = "ghost-btn snippet-toggle";
      const label = () => { tog.textContent = snip.classList.contains("full") ? "changes only" : "whole file"; };
      label();
      tog.addEventListener("click", () => { snip.classList.toggle("full"); label(); });
      cap.appendChild(tog);
    }
    // wrap the <pre> in a scrollable, resizable body with a drag bar
    const pre = snip.querySelector("pre");
    const body = document.createElement("div");
    body.className = "snippet-body";
    pre.replaceWith(body);
    body.appendChild(pre);
    const bar = document.createElement("div");
    bar.className = "resize-bar snippet-resize";
    bar.title = "drag to resize";
    snip.appendChild(bar);
    makeResizable(body, bar);
    if (body.scrollHeight <= body.clientHeight + 2) bar.classList.add("idle");
  });
}
$("copyCodeBtn").addEventListener("click", () => {
  const l = LESSONS[state.idx];
  copyText($("codeArea").value, l.files[state.file].name);
});

/* ---------- syntax highlight overlay ---------- */
function setCode(t, v) { t.value = v; highlight(t); }
function highlight(t) {
  const hl = t.previousElementSibling.firstElementChild;
  const src = t.value.endsWith("\n") ? t.value + " " : t.value;
  hl.innerHTML = Prism.highlight(src, Prism.languages.go, "go")
    .split("\n").map((l) => '<span class="ln"></span>' + l).join("\n");
  hl.parentElement.style.height = t.offsetHeight + "px";
}
function wireEditor(t) {
  t.addEventListener("input", () => highlight(t));
  t.addEventListener("keydown", editorKeys);
  t.addEventListener("scroll", () => { t.previousElementSibling.scrollTop = t.scrollTop; });
  new ResizeObserver(() => highlight(t)).observe(t);
  t.nextElementSibling.addEventListener("pointerdown", (e) => {
    e.preventDefault();
    const y0 = e.clientY, h0 = t.offsetHeight;
    const move = (ev) => { t.style.height = Math.max(120, h0 + ev.clientY - y0) + "px"; };
    const up = () => { window.removeEventListener("pointermove", move); window.removeEventListener("pointerup", up); };
    window.addEventListener("pointermove", move);
    window.addEventListener("pointerup", up);
  });
}
wireEditor($("codeArea"));

/* ---------- editor niceties: tab + auto-indent ---------- */
function editorKeys(e) {
  if ((e.ctrlKey || e.metaKey) && e.key === "Enter") return;
  const t = e.target, s = t.selectionStart, v = t.value;
  if (e.key === "Tab") {
    e.preventDefault();
    t.setRangeText("\t", s, t.selectionEnd, "end");
  } else if (e.key === "Enter") {
    e.preventDefault();
    const line = v.slice(v.lastIndexOf("\n", s - 1) + 1, s);
    const indent = line.match(/^[\t ]*/)[0] + (/[{(\[]\s*$/.test(line) ? "\t" : "");
    t.setRangeText("\n" + indent, s, t.selectionEnd, "end");
  } else if (e.key === "}" && /^[\t ]+$/.test(v.slice(v.lastIndexOf("\n", s - 1) + 1, s))) {
    e.preventDefault();
    t.setRangeText("}", s - 1, t.selectionEnd, "end");
  } else return;
  t.dispatchEvent(new Event("input"));
}

/* ---------- exercises ---------- */
const exKeys = (slug) => Object.keys(EXERCISES).filter((k) => k === slug || k.startsWith(slug + "-") && /^\d+$/.test(k.slice(slug.length + 1)));

function loadExercise(l) {
  const wrap = $("exercises");
  wrap.innerHTML = "";
  const keys = exKeys(l.slug);
  keys.forEach((key, i) => {
    const ex = EXERCISES[key];
    const card = $("exerciseTpl").content.firstElementChild.cloneNode(true);
    card.dataset.key = key;
    if (keys.length > 1) {
      const n = document.createElement("span");
      n.className = "ex-n"; n.textContent = `${i + 1}/${keys.length}`;
      card.querySelector(".exercise-title").prepend(n);
    }
    card.querySelector(".exercise-prompt").innerHTML = ex.prompt
      .replace(/&/g, "&amp;").replace(/</g, "&lt;")
      .replace(/`([^`]+)`/g, "<code>$1</code>");
    card.querySelector(".ex-file").textContent = ex.file;
    const t = card.querySelector(".ex-code");
    wireEditor(t);
    setCode(t, state.exDrafts[key] ?? ex.starter);
    t.addEventListener("input", () => {
      const v = t.value;
      if (v === ex.starter) delete state.exDrafts[key]; else state.exDrafts[key] = v;
      save();
    });
    card.querySelector(".ex-copy").addEventListener("click", () => copyText(t.value, ex.file));
    card.querySelector(".ex-reset").addEventListener("click", () => {
      delete state.exDrafts[key];
      setCode(t, ex.starter);
      save();
      squeak("restored the starter code.");
    });
    card.querySelector(".ex-check").addEventListener("click", () => checkExercise(card, key));
    updateExerciseBadge(card, key);
    wrap.appendChild(card);
  });
}

function updateExerciseBadge(card, key) {
  const passed = state.exDone.has(key);
  const title = card.querySelector(".exercise-title");
  title.classList.toggle("passed", passed);
  title.querySelector(".exercise-sub").textContent = passed ? "solved 🌟" : "solve it yourself";
}

async function checkExercise(card, key) {
  const btn = card.querySelector(".ex-check");
  const box = card.querySelector(".output");
  const hint = card.querySelector(".hint");
  btn.disabled = true;
  btn.textContent = "checking…";
  hint.textContent = "compiling your file into the part & running the hidden tests…";
  box.classList.remove("hidden", "err", "pass");
  box.textContent = "checking…";
  try {
    const res = await fetch("/api/check", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ slug: key, code: card.querySelector(".ex-code").value }),
    });
    const data = await res.json();
    const passed = !!data.pass;
    showOutput(box, passed, data.output || "", "pass", "fail");
    if (passed) {
      if (!state.exDone.has(key)) {
        state.exDone.add(key);
        save();
        confetti();
      }
      squeak("exercise solved.");
    } else {
      squeak("not yet; read the test output.");
    }
    updateExerciseBadge(card, key);
  } catch (e) {
    box.classList.add("err");
    box.textContent = "no checker available 😿 (start the server: docker compose up, or go run ./server)";
  } finally {
    btn.disabled = false;
    btn.textContent = "check my answer";
    hint.textContent = "";
  }
}

/* ---------- sidebar toggle ---------- */
document.body.classList.toggle("nav-collapsed", localStorage.getItem("rogueCuteNav") === "collapsed");
$("navToggle").addEventListener("click", () => {
  const collapsed = document.body.classList.toggle("nav-collapsed");
  localStorage.setItem("rogueCuteNav", collapsed ? "collapsed" : "");
});

/* ---------- themes ---------- */
$("themePick").value = document.documentElement.dataset.theme;
$("themePick").addEventListener("change", (e) => {
  document.documentElement.dataset.theme = e.target.value;
  localStorage.setItem("rogueCuteTheme", e.target.value);
  squeak(`theme: ${e.target.selectedOptions[0].textContent.trim()}`);
});

/* ---------- boot ---------- */
(async () => {
  try {
    const r = await fetch("exercises.json");
    EXERCISES = await r.json();
  } catch (e) {
    console.error("no exercises.json", e);
  }
  buildNav();
  updateProgress();
  const start = location.hash.slice(1);
  const i = LESSONS.findIndex((l) => l.slug === start);
  if (i >= 0) openLesson(i, false);
  setTimeout(() => squeak(`kobold says: ${state.done.size}/${LESSONS.length} steps done, ${state.exDone.size} exercises solved.`), 800);
})();
