<?php
require_once('vendor/autoload.php');

use Docxmerge\Docxmerge;

$docxmerge = new Docxmerge("API_KEY");

$fp = fopen("./hello_world_php.pdf", "w");

$docxmerge->renderFile($fp, "../fixtures/hello_world.docx", array(
    "name" => "James bond",
    "logo" => "https://docxmerge.com/assets/android-chrome-512x512.png"
), "PDF");
