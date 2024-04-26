document.addEventListener("DOMContentLoaded", function () {
  // Select all code blocks
  document.querySelectorAll("pre code").forEach((block) => {
    // Highlight.js detects the language
    hljs.highlightElement(block);
    // Extract the detected language
    let detectedLanguage = block.classList.value.match(/language-(\w+)/)[1];

    // Now set the Prism class for that language
    block.className = "";
    block.classList.add("language-" + detectedLanguage);
    // Use Prism to highlight
    Prism.highlightElement(block);
  });
});
