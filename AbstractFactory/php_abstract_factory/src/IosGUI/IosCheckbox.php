<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory\IosGUI;

use Ganariya\PhpAbstractFactory\AbstractGUI\Checkbox;

final class IosCheckbox implements Checkbox
{
    public function draw(): string
    {
        return "IosCheckbox";
    }
}
