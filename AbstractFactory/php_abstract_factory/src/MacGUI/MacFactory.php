<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory\MacGUI;

use Ganariya\PhpAbstractFactory\AbstractGUI\Button;
use Ganariya\PhpAbstractFactory\AbstractGUI\Checkbox;
use Ganariya\PhpAbstractFactory\AbstractGUI\GUIFactory;

final class MacFactory implements GUIFactory
{
    public function createButton(): Button
    {
        return new MacButton();
    }

    public function createCheckbox(): Checkbox
    {
        return new MacCheckbox();
    }
}
