// step 1, install docxmerge
// npm install docxmerge

// Step 2, import libraries
const Docxmerge = require("docxmerge") // npm i docxmerge
const fs = require("fs")

// Step 3, execute
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
  .then(stream => stream.pipe(fs.createWriteStream("./hello_world_nodejs.pdf")))
  .catch(err => {
    console.error(err)
  })
