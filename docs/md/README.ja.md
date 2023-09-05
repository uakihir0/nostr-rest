# Nostr REST

このプロジェクトは Nostr をマイクロブログとして扱い、
REST API を提供するサーバーを作成するプロジェクトです。
まだ作成途中のプロジェクトで、一緒に実装してくれる方を募集しています。

## 目的

Nostr をモバイルで利用しようと思った時に以下のような問題が発生します。

- 多くのリレーに接続し大量のデータを通信する必要がある
- 通知の仕組みがなく、リアクションに対して気付きにくい
- クライアントアプリケーションの選択肢が少ない

これらの問題を解決するために、リレーを直接参照するのではなく、
ゲートウェイとして REST API サーバーを構築することによって解決を目指します。

## Mastodon 互換 API

Mastodon API 仕様に準じる REST API を実装することを、
本プロジェクトにおいてはの第一目標にしています。
Mastodon と Nostr で扱われているデータ構造には差異がありますが、
その差異を吸収して Mastodon API として提供することで、目的が達成できると考えています。

### 対応状況

**[仕様書](https://www.uakihir0.com/nostr-rest/mastodon.html)**

- Accounts
  - [x] /v1/accounts/verify_credentials
  - [x] /v1/accounts/{uid}
  - [x] /v1/accounts/{uid}/statuses
- Timelines
  - [x] /v1/timelines/public

## 簡易 API

Mastodon 互換 API は処理が重い問題を抱える可能性が高いので、
簡易的な API を作成して、クライアントを独自に作成してもらうという方向も考えています。

### 対応状況

**[仕様書](https://uakihir0.github.io/nostr-rest/)**

## 実行例

現在の実装状況を反映したサンプルサーバーがデプロイされているので、
どのようなデータが得られるか参考にしてみてください。

```shell
# 作者の Nostr の投稿を Mastodon API 形式で取得
curl --request GET \
  --url https://nostr-rest-ervoyfxxqq-an.a.run.app/api/v1/accounts/776ea4437354381f14a720be3c476937dce7257ed1073e54a192dbc99f3b7ecc/statuses \
  --header 'Authorization: Bearer npub1wah2gsmn2sup7998yzlrc3mfxlwwwft76yrnu49pjtdun8em0mxq6appzu' \
  --header 'Content-Type: application/json'
```

## 作者

- Nostr: [uakihir0](https://iris.to/profile/npub1wah2gsmn2sup7998yzlrc3mfxlwwwft76yrnu49pjtdun8em0mxq6appzu)  
- X: [uakihir0](https://x.com/uakihir0)
