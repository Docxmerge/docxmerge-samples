import com.docxmerge.*;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.nio.file.Files;
import java.util.HashMap;
import java.util.Map;

public class Example {
    public static void main(String[] args) throws IOException {
        HelloWorldRequest data = new HelloWorldRequest();
        data.setLogo("https://docxmerge.com/assets/android-chrome-512x512.png");
        data.setName("James Bond");
        Docxmerge docxmerge = new Docxmerge("API_KEY", "https://api.docxmerge.com");
        byte[] bytes = docxmerge.mergeAndTransformTemplate("example-hello-world", data);
        final String outputFile = "../tmp/hello_world_java.pdf";
        File file = new File(outputFile);
        System.out.println("Check " + file.getCanonicalPath());
        Files.write(file.toPath(), bytes);

    }
}
