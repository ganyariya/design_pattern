<?php

declare(strict_types=1);

namespace Ganariya\PhpBridge\Random;

interface RandomInterface
{
    /**
     * @param string[] $minos
     */
    public function getRandomMino(array $minos): string;
}
