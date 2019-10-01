from docxmerge_sdk import Docxmerge # pip install docxmerge_sdk
docxmerge = Docxmerge("API_KEY")

url =  "https://api.docxmerge.com/api/v1/File/GetContenido?id=cdb9842d-5e38-4149-a06b-e1079a208fc3"

pdf = docxmerge.render_url(url,{
    "logo": "https://docxmerge.com/assets/android-chrome-512x512.png",
    "name": "James Bond"
}, "PDF")

open("./example-hello-world.pdf", "wb").write(pdf)
