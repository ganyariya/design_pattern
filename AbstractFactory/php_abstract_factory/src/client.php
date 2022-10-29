<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory;

require __DIR__ . '/../vendor/autoload.php';

use Ganariya\PhpAbstractFactory\AbstractGUI\GUIFactory;
use Ganariya\PhpAbstractFactory\IosGUI\IosFactory;
use Ganariya\PhpAbstractFactory\MacGUI\MacFactory;

/**
 * 本来は動的に取得する
 */
$os = "Mac";

/** @var GUIFactory $guiFactory */
$guiFactory = null;
switch ($os) {
    case "Mac":
        $guiFactory = new MacFactory();
        break;
    case "Ios":
        $guiFactory = new IosFactory();
        break;
}
if (is_null($guiFactory)) {
    print("Null");
}

assert($guiFactory->createButton()->draw() === "MacButton");
