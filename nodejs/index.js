const fs = require("fs")
const Docxmerge = require("docxmerge") // npm i docxmerge
const out = fs.createWriteStream("./tmp/hello-world.pdf")
const docxmerge = new Docxmerge(
  "API_KEY",
  "https://api.docxmerge.com",
)
docxmerge
  .mergeAndTransformTemplate("example-hello-world", {
    logo: "https://docxmerge.com/assets/android-chrome-512x512.png",
    name: "James Bond",
  })
  .then(stream => stream.pipe(out))
  .catch(err => console.error(err))
