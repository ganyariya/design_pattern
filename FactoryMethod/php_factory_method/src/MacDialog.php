<?php

namespace Ganariya\PhpFactoryMethod;

final class MacDialog extends Dialog
{
    private int $count;
    public function __construct(int $count)
    {
        $this->count = $count;
    }

    public function createButton(): ButtonInterface
    {
        return new MacButton($this->count);
    }
}
