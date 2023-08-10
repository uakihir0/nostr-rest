# Test

## TestAccounts

* AppTest
  * npub16chc5cm8e8dy54j2n2knp50as6h6vyxzy7lqrd6kwa8a4vz7487q7pwctm
    * d62f8a6367c9da4a564a9aad30d1fd86afa610c227be01b756774fdab05ea9fc (Hex)
  * nsec1x8rwwhqserc2te6e9cj6uc972y8v5krdyq2muahjnreux3k4pukq3llk9z
    * 31c6e75c10c8f0a5e7592e25ae60be510eca586d2015be76f298f3c346d50f2c (Hex)

## NIP
`curl -s https://api.nostr.watch/v1/nip/40 | jq`

## Play Ground

[https://snowcait.github.io/nostr-playground/]()

### Get Replay
```json
[
  "REQ",
  "48040",
  {
    "kinds": [1], 
    "#p": ["776ea4437354381f14a720be3c476937dce7257ed1073e54a192dbc99f3b7ecc"],
    "limit": 100
  }
]
```

### Get Following
```json
[
  "REQ",
  "299",
  {
    "kinds": [3],
    "authors": ["776ea4437354381f14a720be3c476937dce7257ed1073e54a192dbc99f3b7ecc"],
    "limit": 1
  }
]
```

### Get Reactions

```json
[
  "REQ",
  "299",
  {
    "kinds": [7],
    "#e": ["279ed364b7f141d40606eccc13326368e06c05cb126a9db82bde6acffefb791c"],
    "limit": 100
  }
]

```