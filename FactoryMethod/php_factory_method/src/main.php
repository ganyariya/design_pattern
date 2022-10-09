<?php

namespace Ganariya\PhpFactoryMethod;

require __DIR__ . '/../vendor/autoload.php';


class Main
{
    private Dialog $dialog;
    public function __construct(Dialog $dialog)
    {
        $this->dialog = $dialog;
    }
    public function getDialog(): Dialog
    {
        return $this->dialog;
    }
}

$main = new Main(new MacDialog(0));
$dialog = $main->getDialog();

$button = $dialog->createButton();
$button->onClick();

$dialog->registerButton($button);
$result = $dialog->render();

assert($result == "mac: 1");
