const fs = require("fs")
const Docxmerge = require("docxmerge") // npm i docxmerge
const out = fs.createWriteStream("./hello_world_nodejs.pdf")
const docxmerge = new Docxmerge("API_KEY", "default")
docxmerge
  .renderUrl(
    "https://api.docxmerge.com/api/v1/File/GetContenido?id=cdb9842d-5e38-4149-a06b-e1079a208fc3&download=true",
    {
      logo: "https://docxmerge.com/assets/android-chrome-512x512.png",
      name: "James Bond",
    },
    "PDF",
  )
  .then(stream => stream.pipe(out))
  .catch(err => {
    err.pipe(fs.createWriteStream("./out.json"))
    console.error(err)
  })
