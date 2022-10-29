<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory\MacGUI;

use Ganariya\PhpAbstractFactory\AbstractGUI\Checkbox;

final class MacCheckbox implements Checkbox
{
    public function draw(): string
    {
        return "MacCheckbox";
    }
}
