package config

const defaultYAML string = `
service:
  name: omo.msa.tag
  address: :9607
  ttl: 15
  interval: 10
logger:
  level: trace
  dir: /var/log/msa/
database:
  mongodb:
    address: 127.0.0.1:27017
    timeout: 10
    user: root
    password: mongodb@OMO
    db: msa_tag
publisher:
- /signup
- /signin
- /signout
- /reset/password
- /profile/query
- /profile/update
`
