package config

const defaultYAML string = `
service:
  name: xtc.api.ogm.tag
  address: :9607
  ttl: 15
  interval: 10
logger:
  level: trace
  dir: /var/log/ogm/
database:
  mongodb:
    address: localhost:27017
    timeout: 10
    user: root
    password: mongodb@OMO
    db: ogm_tag
publisher:
- /signup
`
