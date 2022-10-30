<?php

declare(strict_types=1);

namespace Ganariya\PhpBridge;

use Ganariya\PhpBridge\Random\LongOnlyRandom;
use Ganariya\PhpBridge\Tetris\Tetris;

require __DIR__ . '/../vendor/autoload.php';

$playTurns = 5;

$tetris = new Tetris($playTurns, new LongOnlyRandom());
$results = $tetris->play();

assert(count($results) === $playTurns);
foreach ($results as $mino) {
    assert($mino === "Long");
}
