
```bash
cargo install --git https://github.com/rust-lang/mdBook.git mdbook
```

#### Vercel

```bash
# build commands
# "v0.4.16" fails with no gcc
curl -L https://github.com/rust-lang/mdBook/releases/download/v0.4.15/mdbook-v0.4.15-x86_64-unknown-linux-gnu.tar.gz | tar xvz && ./mdbook build

# output directory
# book
```

#### Netlify

```bash
# build command
# "v0.4.16" fails with no gcc
curl -L https://github.com/rust-lang/mdBook/releases/download/v0.4.15/mdbook-v0.4.15-x86_64-unknown-linux-gnu.tar.gz | tar xvz && ./mdbook build

# publish directory
# book

# set "Domain management"
```
