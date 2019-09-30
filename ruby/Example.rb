require "docxmerge"

docxmerge = Docxmerge::Docxmerge.new("API_KEY", "default", "https://api.docxmerge.com")
response = docxmerge.merge_and_transform_template("example-hello-world", {
    :logo => "https://docxmerge.com/assets/android-chrome-512x512.png",
    :name => "James Bond",
})

file_output = "../tmp/hello_world_ruby.pdf"

f = File.open(file_output, "wb")
f << response
f.close
puts "Check #{file_output}"