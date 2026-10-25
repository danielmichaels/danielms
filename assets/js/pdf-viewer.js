document.querySelectorAll(".pdf-viewer").forEach((figure) => {
  const track = figure.querySelector(".pdf-viewer-track");
  const prev = figure.querySelector('[data-action="prev"]');
  const next = figure.querySelector('[data-action="next"]');
  const pageNum = figure.querySelector("[data-page-num]");
  const last = track.children.length - 1;
  let current = 0;

  const go = (index) =>
    track.scrollTo({ left: index * track.clientWidth, behavior: "smooth" });

  prev.addEventListener("click", () => go(current - 1));
  next.addEventListener("click", () => go(current + 1));
  track.addEventListener("scroll", () => {
    const index = Math.round(track.scrollLeft / track.clientWidth);
    if (index === current) return;
    current = index;
    pageNum.textContent = current + 1;
    prev.disabled = current === 0;
    next.disabled = current === last;
  });
});
