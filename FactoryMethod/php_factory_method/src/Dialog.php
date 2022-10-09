<?php

namespace Ganariya\PhpFactoryMethod;

abstract class Dialog
{
    /** @var ButtonInterface[] $buttons */
    public array $buttons = [];

    public function registerButton(ButtonInterface $button)
    {
        $this->buttons[] = $button;
    }
    public function render(): string
    {
        $ret = "";
        foreach ($this->buttons as $button) {
            $ret .= $button->render();
        }
        return $ret;
    }
    abstract public function createButton(): ButtonInterface;
}
