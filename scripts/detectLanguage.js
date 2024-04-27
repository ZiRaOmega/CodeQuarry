function checkHighlight() {
  // Ensure both hljs and Prism are defined
  if (typeof hljs === "undefined") {
    console.error("Highlight.js not found!");
    return;
  }

  document.querySelectorAll("pre code").forEach((block) => {
    // Apply Highlight.js
    hljs.highlightElement(block);
  });
}

// Call checkHighlight when the DOM is fully loaded and after any AJAX content load
document.addEventListener("DOMContentLoaded", checkHighlight);
