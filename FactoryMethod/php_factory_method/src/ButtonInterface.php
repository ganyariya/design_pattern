<?php

namespace Ganariya\PhpFactoryMethod;

interface ButtonInterface
{
    public function onClick(): void;
    public function render(): string;
}
