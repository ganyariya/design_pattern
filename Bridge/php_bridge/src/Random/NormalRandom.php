<?php

declare(strict_types=1);

namespace Ganariya\PhpBridge\Random;

final class NormalRandom implements RandomInterface
{
    public function getRandomMino(array $minos): string
    {
        $idx = array_rand($minos);
        return $minos[$idx[0]];
    }
}
