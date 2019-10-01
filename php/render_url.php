<?php
require_once('vendor/autoload.php');

use Docxmerge\Docxmerge;

$apiKey = "API_KEY";
$docxmerge = new Docxmerge($apiKey);
$fp = fopen("./hello_world_php.pdf", "w");
$url = "https://api.docxmerge.com/api/v1/File/GetContenido?id=cdb9842d-5e38-4149-a06b-e1079a208fc3";
$docxmerge->renderUrl($fp, $url, array(
    "name" => "James bond",
    "logo" => "https://docxmerge.com/assets/android-chrome-512x512.png"
), "PDF");
