# About

ganyariya の Design Pattern についてのサンプル実装まとめです。
Go もしくは PHP で実装しています。
 参考

- [Java言語で学ぶデザインパターン入門](https://www.amazon.co.jp/Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E7%B5%90%E5%9F%8E-%E6%B5%A9/dp/4797316462/ref=sr_1_4?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&crid=21WG87PIJZQX2&keywords=Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80&qid=1665281104&qu=eyJxc2MiOiIyLjMyIiwicXNhIjoiMS42MiIsInFzcCI6IjEuNDQifQ%3D%3D&s=books&sprefix=java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80%2Cstripbooks%2C178&sr=1-4)
  - 持っていた初版版をそのまま参考にしています（今は 3 版まで出ていますが、旧バージョンを利用しています）。
  - 説明のみ参考にし、独自実装に置き換えています。
- [REFACTORING GURU](https://refactoring.guru/ja/design-patterns/catalog)
- [デザインパターン習得編](http://marupeke296.com/DP_main.html)
  - ゲームを題材としたデザインパターンの説明

# Pattern

1. [Iterator](./Iterator/)
2. [Adapter](./Adapter/)
3. [Template Method](./TemplateMethod/)
4. [Factory Method](./FactoryMethod/)
5. [Singleton](./Singleton/)
6. [Prototype](./Prototype/)
7. [Builder](./Builder/)
8. [Abstract Factory](./AbstractFactory/)
9. [Bridge](./Bridge/)
10. [Strategy](./Strategy/)
11. [Composite](./Composite/)
12. [Decorator](./Decorator/)
13. [Visitor](./Visitor/)
14. [ChainOfResponsibility](./ChainOfResponsibility/)
15. [Facade](./Facade/)
16. [Mediator](./Mediator/)
17. [Observer](./Observer/)
18. [Memento](./Memento/)
19. [State](./State/)

# Category

### 生成パターン

インスタンス・オブジェクトの生成に関するパターン。インスタンスの生成のロジックのみ分離する。

- [Factory Method](./FactoryMethod)
- [Abstract Factory](./AbstractFactory/)
- [Builder](./Builder/)
- [Prototype](./Prototype/)
- [Singleton](./Singleton/)

### 構造パターン

プログラムの構造に関するパターン。大枠の構造を規定する。

- [Adapter](./Adapter/)
- [Bridge](./Bridge/)
- [Composite](./Composite/)
- [Decorator](./Decorator/)
- [Facade](./Facade/)
- [Flyweight]()
- [Proxy]()

### 振る舞いパターン

アルゴリズムやオブジェクトの振る舞いに関して、責任の分離を行う。

- [Chain of Responsibility](./ChainOfResponsibility/)
- [Command]()
- [Iterator](./Iterator/)
- [Mediator](./Mediator/)
- [Memento](./Memento/)
- [Observer](./Observer/)
- [State](./State/)
- [Strategy](./Strategy/)
- [TemplateMethod](./TemplateMethod/)
- [Visitor](./Visitor/)

