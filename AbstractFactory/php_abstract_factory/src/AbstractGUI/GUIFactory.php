<?php

declare(strict_types=1);

namespace Ganariya\PhpAbstractFactory\AbstractGUI;

/**
 * Windows or Mac or IOS
 * ゴブリン or バード
 * Twig or Blade or Smarty
 * 
 * 上記のように「サブクラスやサブシステム」を大きく変更する要件がある & かつ共通項がある場合に
 * 「サブクラスやサブシステム」自体を生成する Abstract Factory を複数用意しようとするもの
 */
interface GUIFactory
{
    public function createButton(): Button;
    public function createCheckbox(): Checkbox;
}
