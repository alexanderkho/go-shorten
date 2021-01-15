# TODO 🦆

## Client

### Chores
- Get rid of annoying linting errors on every save
- Get Vue intellisense + syntax highlighting configured

### Features
- Add "Recently created urls" list, use local storage to save
- Add 404 page
- Cleaner UI

## Server
- Handle 404 (send to client 404 page)
- Clean up error handling
- Check incoming url for scheme, default to HTTP if not provided
- Split `main.go` up into a few files. 

## Shared
- Setup config for different environmentes (dev, prod)
- Setup a build pipeline