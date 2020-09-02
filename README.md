# StartKit

See [omo-msa-startkit](https://github.com/xtech-cloud/omo-msa-startkit)

# Protoc

See [omo-msp-tag](https://github.com/xtech-cloud/omo-msp-tag)

# Docker

See [omo-docker-tag](https://github.com/xtech-cloud/omo-docker-tag)

# 消息订阅

- 地址
  omo.msa.tag.notification

- 消息
  | Action | Head | Body|
  |:--|:--|:--|
  |/signup||uuid|
  |/signin|accessToken|uuid|
  |/signout|accessToken||
  |/reset/password|accessToken||
  |/profile/update|accessToken||
  |/profile/query|accessToken||
