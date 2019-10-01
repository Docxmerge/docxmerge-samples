require "docxmerge"

docxmerge = Docxmerge::Docxmerge.new("API_KEY", "default", "https://api.docxmerge.com")

response = docxmerge.render_file(
    File.open("../fixtures/hello_world.docx", "r"),
    {
        :logo => "https://docxmerge.com/assets/android-chrome-512x512.png",
        :name => "James Bond",
    },
    "PDF"
)

file_output = "./hello_world_ruby_file.pdf"

f = File.open(file_output, "wb")
f << response
f.close
puts "Check #{file_output}"