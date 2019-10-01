using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Docxmerge;
namespace dotnet
{
  class Program
  {
    static void Main(string[] args)
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
