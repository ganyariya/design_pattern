<?php

declare(strict_types=1);

namespace Ganariya\PhpBridge\Random;

final class LongOnlyRandom implements RandomInterface
{
    public function getRandomMino(array $minos): string
    {
        return "Long";
    }
}
