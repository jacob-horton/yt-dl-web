const form = document.getElementById("download-form");
const button = document.getElementById("download-btn");

form.addEventListener("submit", handleSubmit);
async function handleSubmit(e) {
  e.preventDefault();

  button.disabled = true;
  button.innerHTML = "Downloading...";

  const formData = Object.fromEntries(new FormData(e.target));
  await fetch("/download", {
    method: "POST",
    body: formData.url,
  })
    .then((response) => response.blob())
    .then((blob) => {
      const url = URL.createObjectURL(blob);

      const a = document.createElement("a");
      a.href = url;
      a.setAttribute("target", "_blank");
      a.download = "audio.m4a";
      a.click();

      URL.revokeObjectURL(url);
    })
    .catch(() => alert("failed to download"));

  button.disabled = false;
  button.innerHTML = "Download";
}
