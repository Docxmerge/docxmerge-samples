<?php
require_once('vendor/autoload.php');

use Docxmerge\Docxmerge;

$apiKey = "API_KEY";
$baseUrl = "https://api.docxmerge.com";
$docxmerge = new Docxmerge($apiKey, $baseUrl);
$fp = fopen("../tmp/hello_world_php.pdf", "w");

$docxmerge->mergeAndTransformTemplate($fp, "example-hello-world", array(
    "name" => "David Viejo",
    "logo" => "https://docxmerge.com/assets/android-chrome-512x512.png"
));
