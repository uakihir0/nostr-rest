[> Japanese](./docs/README.ja.md)

**WIP**

# Nostr REST Gateway Server

This application is a Nostr REST Wrapper server written in the Go language.

The Nostr protocol uses WebSockets to communicate with multiple relay servers to provide services. However, due to the nature of the protocol, the volume of communication can become overwhelming. This problem is especially critical in mobile applications, where there are limits to the amount of traffic and battery power.

In addition, Nostr's relay server does not have a notification function, making it difficult to notice when other users respond to your posts.

To solve these problems, we will place a REST API gateway in front of the relay server.

This project is in the process of implementation.

## Author

Nostr: [uakihir0](https://iris.to/profile/npub1wah2gsmn2sup7998yzlrc3mfxlwwwft76yrnu49pjtdun8em0mxq6appzu)  
Twitter: [uakihir0](https://twitter.com/uakihir0)