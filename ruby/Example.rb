require "docxmerge"

docxmerge = Docxmerge::Docxmerge.new("API_KEY", "default", "https://api.docxmerge.com")
url = "https://api.docxmerge.com/api/v1/File/GetContenido?id=cdb9842d-5e38-4149-a06b-e1079a208fc3"
response = docxmerge.render_url(url, {
    :logo => "https://docxmerge.com/assets/android-chrome-512x512.png",
    :name => "James Bond",
}, "PDF")

file_output = "./hello_world_ruby.pdf"

f = File.open(file_output, "wb")
f << response
f.close
puts "Check #{file_output}"