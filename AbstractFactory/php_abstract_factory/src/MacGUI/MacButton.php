<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory\MacGUI;

use Ganariya\PhpAbstractFactory\AbstractGUI\Button;

final class MacButton implements Button
{
    public function draw(): string
    {
        return "MacButton";
    }
}
