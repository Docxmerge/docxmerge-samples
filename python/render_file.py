# pip install docxmerge_sdk 
# https://pypi.org/project/docxmerge-sdk/

from docxmerge_sdk import Docxmerge
docxmerge = Docxmerge("API_KEY", "default", "API_URL")

pdf = docxmerge.render_file(
    open("../fixtures/hello_world.docx"),
    {
        "logo": "https://docxmerge.com/assets/android-chrome-512x512.png",
        "name": "James Bond"
    },
    "PDF"
)

open("./example-file-hello-world.pdf", "wb").write(pdf)