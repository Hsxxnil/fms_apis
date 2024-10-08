package auth

const _CASBIN_RULES = `
[
  {
    "ptype": "p",
    "v0": "admin",
    "v1": "/fms/*",
    "v2": "GET"
  },
  {
    "ptype": "p",
    "v0": "admin",
    "v1": "/fms/*",
    "v2": "POST"
  },
  {
    "ptype": "p",
    "v0": "admin",
    "v1": "/fms/*",
    "v2": "PATCH"
  },
  {
    "ptype": "p",
    "v0": "admin",
    "v1": "/fms/*",
    "v2": "DELETE"
  }
]
`
