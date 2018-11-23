# golang-design-patterns
GOF Design Patterns written in Golang.


## 2. Adapterパターン

adaptという単語は日本語で「適合させる」という意味で、adapterとは「適合させるもの」という意味になります。Adapterパターンは、インタフェースに互換性の無いクラス同士を組み合わせることを目的としたパターンです。

例えば、これまで利用していたメソッドと同じ機能を、よりすぐれた形で提供するメソッドを持つクラスの存在を知ったとします。しかし、このすぐれたメソッドは、これまで利用していたメソッドとは異なるインタフェースを持つため、乗り換えるとなると多大な変更を余儀なくされる場合があります。こんなとき、この2つのメソッドのインタフェースの違いを吸収してやる Adapter を準備することで、少ない変更で新しいメソッドに乗り換えることができるのです。

> - [Adapter Golang Example - Inheritance](./adapter/adapter_inheritance.go)
> - [Adapter Golang Example - Delegation](./adapter/adapter_delegation.go)
> - [Adapterパターン｜techscore](http://www.techscore.com/tech/DesignPattern/Adapter/Adapter1.html/)


## 15. Facadeパターン

Facadeパターンは、既存のクラスを複数組み合わせて使う手順を、「窓口」となるクラスを作ってシンプルに利用できるようにするパターンです。

> - [Facade Golang Example](./facade/facade.go)
> - [Facadeパターン｜techscore](http://www.techscore.com/tech/DesignPattern/Facade.html/)
