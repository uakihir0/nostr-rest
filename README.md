[>> 日本語](./docs/md/README.ja.md)

# Nostr REST

This project aims to treat Nostr as a microblog and create a server that provides a REST API. It is still a work-in-progress project, and we are looking for individuals to collaborate on its implementation.

## Purpose

When trying to use Nostr on mobile devices, the following issues may arise:

- The need to connect to multiple relays and communicate large amounts of data.
- Lack of notification mechanisms, making it difficult to be aware of reactions.
- Limited choices for client applications.

To address these issues, we aim to resolve them by building a REST API server as a gateway instead of directly accessing relays.

## Mastodon-Compatible API

Implementing a REST API that adheres to the Mastodon API specification is the primary goal of this project. While there are differences in data structures between Mastodon and Nostr, we believe that providing Mastodon API compatibility by accommodating these differences will help achieve our goal.

### Compatibility Status

**[Specification](https://www.uakihir0.com/nostr-rest/mastodon.html)**

- Accounts
    - [x] /v1/accounts/verify_credentials
    - [x] /v1/accounts/{uid}
    - [x] /v1/accounts/{uid}/statuses

## Simple API

Since the Mastodon-compatible API may have performance-related issues, we are also considering creating a simplified API and allowing clients to create their own.

### Compatibility Status

**[Specification](https://uakihir0.github.io/nostr-rest/)**

## Usage Example

A sample server reflecting the current implementation status is deployed, so you can refer to it to see what kind of data can be obtained.

```shell
# Retrieve Nostr posts by the author in Mastodon API format
curl --request GET \
  --url https://nostr-rest-ervoyfxxqq-an.a.run.app/api/v1/accounts/776ea4437354381f14a720be3c476937dce7257ed1073e54a192dbc99f3b7ecc/statuses \
  --header 'Authorization: Bearer npub1wah2gsmn2sup7998yzlrc3mfxlwwwft76yrnu49pjtdun8em0mxq6appzu' \
  --header 'Content-Type: application/json'
```

## Author

- Nostr: [uakihir0](https://iris.to/profile/npub1wah2gsmn2sup7998yzlrc3mfxlwwwft76yrnu49pjtdun8em0mxq6appzu)
- X: [uakihir0](https://x.com/uakihir0)