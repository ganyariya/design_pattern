<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory\IosGUI;

use Ganariya\PhpAbstractFactory\AbstractGUI\Button;
use Ganariya\PhpAbstractFactory\AbstractGUI\Checkbox;
use Ganariya\PhpAbstractFactory\AbstractGUI\GUIFactory;

final class IosFactory implements GUIFactory
{
    public function createButton(): Button
    {
        return new IosButton();
    }

    public function createCheckbox(): Checkbox
    {
        return new IosCheckbox();
    }
}
