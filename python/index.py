from docxmerge_sdk import Docxmerge # pip install docxmerge_sdk
docxmerge = Docxmerge("API_KEY","https://api.docxmerge.com")
pdf = docxmerge.merge_and_transform_template("example-hello-world",{
    "logo": "https://docxmerge.com/assets/android-chrome-512x512.png",
    "name": "James Bond"
})
open("./tmp/example-hello-world.pdf", "wb").write(pdf)