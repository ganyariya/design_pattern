<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory\IosGUI;

use Ganariya\PhpAbstractFactory\AbstractGUI\Button;

final class IosButton implements Button
{
    public function draw(): string
    {
        return "IosButton";
    }
}
