using System.Collections.Generic;
using System.IO;

namespace DocxmergeExample
{
  class Program
  {
    static void Main(string[] args)
    {
      var docxmerge = new Docxmerge.Docxmerge("API_KEY");
      var output = docxmerge.RenderTemplate(
          "hello-world",
          new Dictionary<string, object>{
              {"name", "James Bond"},
              {"logo", "https://docxmerge.com/assets/android-chrome-512x512.png"}
          },
          "PDF",
          "latest" // version, can be ommited, just for reference
      ).Result;

      using (var ms = new MemoryStream())
      {
        output.CopyTo(ms);
        var outputFile = "./hello_world_template_dotnet.pdf";
        File.WriteAllBytes(outputFile, ms.ToArray());
        System.Console.WriteLine($"Check {outputFile}");
      }
    }
    static void Main3(string[] args)
    {
      var docxmerge = new Docxmerge.Docxmerge("API_KEY");
      var input = File.OpenRead("../fixtures/hello_world.docx");
      var output = docxmerge.RenderFile(input, new Dictionary<string, object>{
          {"name", "James Bond"},
          {"logo", "https://docxmerge.com/assets/android-chrome-512x512.png"}
      }, "PDF").Result;
      using (var ms = new MemoryStream())
      {
        output.CopyTo(ms);
        var outputFile = "./hello_world_file_dotnet.pdf";
        File.WriteAllBytes(outputFile, ms.ToArray());
        System.Console.WriteLine($"Check {outputFile}");
      }
    }
    static void Main2(string[] args)
    {
      var docxmerge = new Docxmerge.Docxmerge("API_KEY");
      var url = "https://api.docxmerge.com/api/v1/File/GetContenido?id=cdb9842d-5e38-4149-a06b-e1079a208fc3";
      var output = docxmerge.RenderUrl(url, new Dictionary<string, object>{
          {"name", "James Bond"},
          {"logo", "https://docxmerge.com/assets/android-chrome-512x512.png"}
      }, "PDF").Result;
      using (var ms = new MemoryStream())
      {
        output.CopyTo(ms);
        var outputFile = "./hello_world_dotnet.pdf";
        File.WriteAllBytes(outputFile, ms.ToArray());
        System.Console.WriteLine($"Check {outputFile}");
      }
    }
  }
}
