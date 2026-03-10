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

## Datastar Repro Page

This quickstart contains a minimal repro route for templUI script behavior with Datastar-style SSE patching:

- Open `http://localhost:8090/repro`
- Click `Open Repro Dialog`
- Click `Patch Fragment` several times
- Click `Log Script Counts` and inspect browser console
- Try opening the selectbox and datepicker

The fragment includes `selectbox`, `datepicker`, and `collapsible` inside dialog content patched over SSE.

If you want your own module path later, run:

```sh
go mod edit -module your/module/path
```

## Docker

```sh
docker build -t templui-quickstart .
docker run --rm -p 8088:8090 templui-quickstart
```

Then open `http://localhost:8088`.
