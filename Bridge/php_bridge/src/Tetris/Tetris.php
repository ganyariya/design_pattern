<?php

declare(strict_types=1);

namespace Ganariya\PhpBridge\Tetris;

use Ganariya\PhpBridge\Random\RandomInterface;

class Tetris
{
    /**
     * @var string[]
     */
    private static array $MINOS  = ["LONG", "L", "R", "SHORT", "..."];

    private RandomInterface $random;
    private int $playTurns;

    public function __construct(
        int $playTurns,
        RandomInterface $random
    ) {
        $this->playTurns = $playTurns;
        $this->random = $random;
    }

    public function getMino(): string
    {
        return $this->random->getRandomMino(self::$MINOS);
    }

    public function play(): array
    {
        $history = [];
        for ($i = 0; $i < $this->playTurns; $i++) {
            $history[] = $this->getMino();
        }
        return $history;
    }

    public function locateMino(): void
    {
        /**
         * ミノを入力に従って配置する
         * デザパタのお勉強なので実装とかはしない　適当スタブ
         */
    }
}
