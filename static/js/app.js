// Typing effect for the role line
(function () {
  const el = document.getElementById("typed");
  if (!el) return;
  const text = el.dataset.text || "";
  if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) { el.textContent = text; return; }
  let i = 0;
  (function tick() {
    el.textContent = text.slice(0, i++);
    if (i <= text.length) setTimeout(tick, 70);
  })();
})();

// Voice-over: plays /static/audio/intro.mp3 if it exists, otherwise reads the intro aloud
(function () {
  const btn = document.getElementById("voice-btn");
  const audio = document.getElementById("intro-audio");
  const text = (document.getElementById("intro-text") || {}).textContent || "";
  if (!btn) return;
  let mode = null; // "file" | "speech"

  function setPlaying(on) {
    btn.setAttribute("aria-pressed", String(on));
    btn.textContent = on ? "■ Stop intro" : "▶ Hear my intro";
  }

  function speak() {
    if (!("speechSynthesis" in window)) { btn.disabled = true; btn.textContent = "Voice not supported"; return; }
    mode = "speech";
    const u = new SpeechSynthesisUtterance(text.trim());
    u.rate = 1;
    u.onend = u.onerror = () => setPlaying(false);
    speechSynthesis.cancel();
    speechSynthesis.speak(u);
    setPlaying(true);
  }

  function stop() {
    audio.pause(); audio.currentTime = 0;
    if ("speechSynthesis" in window) speechSynthesis.cancel();
    setPlaying(false);
  }

  btn.addEventListener("click", () => {
    if (btn.getAttribute("aria-pressed") === "true") return stop();
    mode = "file";
    audio.onended = () => setPlaying(false);
    audio.onerror = speak; // no recording uploaded yet: fall back to speech
    audio.play().then(() => setPlaying(true)).catch(speak);
  });
})();

// Demo modal (live site in an iframe, or a demo video)
(function () {
  const modal = document.getElementById("demo-modal");
  if (!modal || !modal.showModal) return;
  const body = document.getElementById("modal-body");
  const title = document.getElementById("modal-title");
  const openLink = document.getElementById("modal-open");

  function open(name, url, node) {
    title.textContent = name;
    openLink.href = url;
    body.replaceChildren(node);
    modal.showModal();
  }

  document.querySelectorAll(".demo-btn").forEach((b) =>
    b.addEventListener("click", () => {
      const f = document.createElement("iframe");
      f.src = b.dataset.url;
      f.title = b.dataset.name + " live demo";
      f.loading = "lazy";
      f.referrerPolicy = "no-referrer";
      open(b.dataset.name + " (live)", b.dataset.url, f);
    })
  );

  document.querySelectorAll(".video-btn").forEach((b) =>
    b.addEventListener("click", () => {
      const v = document.createElement("video");
      v.src = b.dataset.video;
      v.controls = true;
      v.autoplay = true;
      open(b.dataset.name + " (video demo)", b.dataset.video, v);
    })
  );

  function close() { modal.close(); body.replaceChildren(); }
  document.getElementById("modal-close").addEventListener("click", close);
  modal.addEventListener("click", (e) => { if (e.target === modal) close(); });
  modal.addEventListener("close", () => body.replaceChildren());
})();
(function () {
  const modal = document.getElementById("demo-modal");
  if (!modal || !modal.showModal) return;
  const body = document.getElementById("modal-body");
  const title = document.getElementById("modal-title");
  const openLink = document.getElementById("modal-open");
  document.querySelectorAll(".thumb").forEach((b) =>
    b.addEventListener("click", () => {
      const img = document.createElement("img");
      img.src = b.dataset.src;
      img.alt = b.dataset.caption;
      title.textContent = b.dataset.caption;
      openLink.href = b.dataset.src;
      body.replaceChildren(img);
      modal.showModal();
    })
  );
})();