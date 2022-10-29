<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory\AbstractGUI;

/**
 * Interface で定義すると name, size などの要素が与えられない
 * 
 * AbstractFactory パターンは「共通項」のみを扱う
 * よって、共通項がある場合は abstract class にして外部から与えられるようにしたほうが良い（はず）
 * 
 */
interface Checkbox
{
    public function draw(): string;
}
