# pip install docxmerge_sdk 
# https://pypi.org/project/docxmerge-sdk/

from docxmerge_sdk import Docxmerge
docxmerge = Docxmerge("API_KEY", "default")

pdf = docxmerge.render_template(
    "hello-world",
    {
        "logo": "https://docxmerge.com/assets/android-chrome-512x512.png",
        "name": "James Bond"
    },
    "PDF",
    "latest"
)

open("./example-hello-world-template.pdf", "wb").write(pdf)
