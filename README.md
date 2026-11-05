> [!WARNING]
> **templUI is now shadcn-templ.** This quickstart is for templUI v1 and is no longer maintained. Use [github.com/axadrn/shadcn-templ](https://github.com/axadrn/shadcn-templ) instead.
>
> This repository is archived and will be deleted in May 2027.

# templUI Quickstart

This is a ready-to-run quickstart project for templUI.

## Run

```sh
git clone https://github.com/templui/templui-quickstart.git myapp
rm -rf myapp/.git
cd myapp
cp .env.example .env
go mod tidy
task dev
```

Open `http://localhost:7331` for templ live preview or `http://localhost:8090` for the app.

If you want your own module path later, run:

```sh
go mod edit -module your/module/path
```

## Need more?

If you outgrow the starter, check out these next steps:

- [goilerplate](https://goilerplate.com/) for premium UI blocks and a production-ready Go SaaS foundation

## Docker

```sh
docker build -t templui-quickstart .
docker run --rm -p 8088:8090 templui-quickstart
```

Then open `http://localhost:8088`.
