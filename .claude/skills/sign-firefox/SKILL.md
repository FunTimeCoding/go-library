---
name: sign-firefox
description: Sign the Firefox Browser Bridge extension via Mozilla AMO
---

Sign the Firefox extension in `js/firefox_bridge/` for self-distribution.

1. Bump the version in `js/firefox_bridge/manifest.json`.
2. Run the signing command:

```bash
podman run --rm \
  -e FIREFOX_ISSUER -e FIREFOX_SECRET \
  -v ./js/firefox_bridge:/ext \
  -v ./tmp/firefox:/out \
  node:lts-alpine sh -c "
    npm install -g web-ext 2>/dev/null &&
    cd /ext &&
    web-ext sign \
      --channel=unlisted \
      --api-key=\$FIREFOX_ISSUER \
      --api-secret=\$FIREFOX_SECRET \
      --artifacts-dir=/out
  "
```

3. The output `.xpi` will have a hex-prefixed filename — this is normal. Rename it to `browser-bridge-<version>.xpi`.
4. Report the path to the signed `.xpi` so the user can install it via `about:addons` > gear > "Install Add-on From File".
