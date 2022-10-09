<?php

namespace Ganariya\PhpFactoryMethod;

final class MacButton implements ButtonInterface
{
    private int $count;
    public function __construct(int $count)
    {
        $this->count = $count;
    }

    public function onClick(): void
    {
        $this->count++;
    }

    public function render(): string
    {
        return "mac: $this->count";
    }
}
