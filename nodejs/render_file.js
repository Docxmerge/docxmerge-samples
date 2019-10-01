// step 1, install docxmerge
// npm install docxmerge

// Step 2, import libraries
const Docxmerge = require("docxmerge") // npm i docxmerge
const fs = require("fs")

// Step 3, execute
const docxmerge = new Docxmerge("API_KEY", "default")
docxmerge
  .renderFile(
    fs.createReadStream("./hello_world.docx"),
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
