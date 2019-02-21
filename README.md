## Ceph Exporter

This is an unofficial fork of the original exporter that could be found here: https://github.com/digitalocean/ceph_exporter

This projects adds the capability to enrich exposed metrics by lookup to a redis db which in turn could hold
metrics about usage of images or others

# Usage
Usage is basically the same as original, please have a look on above mentioned site.

To make us of redis lookup function set follwoing envs:
- REDIS_ENDPOINT: "localhost:6379"
- REDIS_DB: "0"