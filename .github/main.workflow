steps:
- uses: actions/checkout@master
- uses: actions/setup-python@v1
  with:
    python-version: '3.x'
    architecture: 'x64' # (x64 or x86)
- run: pip install certbot-dns-cloudflare ndg_httpsclient
