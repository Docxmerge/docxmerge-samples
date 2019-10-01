require "docxmerge"

docxmerge = Docxmerge::Docxmerge.new("API_KEY", "default", "https://api.docxmerge.com")
response = docxmerge.render_template(
    "hello-world",
    {
        :logo => "https://docxmerge.com/assets/android-chrome-512x512.png",
        :name => "James Bond",
    },
    "PDF",
    "latest" # version, can be ommited, just for reference
)

file_output = "./hello_world_template_ruby.pdf"

f = File.open(file_output, "wb")
f << response
f.close
puts "Check #{file_output}"