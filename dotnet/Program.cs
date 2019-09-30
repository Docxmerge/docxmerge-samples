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
      var docxmerge = new Docxmerge.Docxmerge("API_KEY", "https://api.docxmerge.com", "default");
      var output = docxmerge.MergeAndTransform("example-hello-world", new Dictionary<string, object>{
          {"name", "James Bond"},
          {"logo", "https://docxmerge.com/assets/android-chrome-512x512.png"}
      }).Result;
      using (var ms = new MemoryStream())
      {
        output.CopyTo(ms);
        var outputFile = "../tmp/hello_world_dotnet.pdf";
        File.WriteAllBytes(outputFile, ms.ToArray());
        System.Console.WriteLine($"Check {outputFile}");
      }
    }
  }
}
