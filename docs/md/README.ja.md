**WIP**

# Nostr REST Gateway Server

**[API 仕様書](https://uakihir0.github.io/nostr-rest/)**

**[Mastodon API 仕様書](https://www.uakihir0.com/nostr-rest/mastodon.html)**

このアプリケーションは Go 言語で書かれた、Nostr の REST Wrapper サーバーです。

このサーバーは、Nostr のリレーサーバーにアクセスし、Mastodon API と互換のある REST API を提供します。 Nostr のプロトコルでは、複数のリレーサーバーと WebSocket で通信を行いサービスを実現します。一方で、プロトコルの特性上、通信量が膨れ上がってしまう問題が発生します。特にモバイルでは、通信量やバッテリーの制限があるため、この問題は非常にクリティカルです。

また、Nostr のリレーサーバーには通知の機能ないため、他のユーザーが自分の投稿に対して、返信などのリアクションを行った際に気付きにくいという問題もあります。

これらの問題を解決するために、リレーサーバーの前段に REST API のゲートウェイを置くことで解決を行います。

本プロジェクトは実装途中です。


## 作者

Nostr: [uakihir0](https://iris.to/profile/npub1wah2gsmn2sup7998yzlrc3mfxlwwwft76yrnu49pjtdun8em0mxq6appzu)  
X(Twitter): [uakihir0](https://x.com/uakihir0)
