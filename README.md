# Chongo-app

Chongo App is a modular webapp that anyone can selfhost for their community. Current modules(Quiz,Chat).

*Note that you will need to have [NPM](https://nodejs.org),[GO](https://go.dev/dl/),[MongoDB](https://docs.mongodb.com/manual/installation/) installed.*

## Get started
Install the dependencies...

```bash
npm install
```

```bash
go mod download
```

If you're using [Visual Studio Code](https://code.visualstudio.com/) we recommend installing the official extension [Svelte for VS Code](https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode) and the official extension [Go](https://open-vsx.org/vscode/item?itemName=golang.Go). If you are using other editors you may need to install a plugin in order to get syntax highlighting and intellisense.

## Development mode

The following command starts the go server on port 8888 and 4443 and every modification you make in the svelte files are live on it.

```bash
npm run dev
```

## Building and running in production mode

To create an optimised version of the app:

```bash
npm run build
```

You can run the newly built app with `npm run start`.