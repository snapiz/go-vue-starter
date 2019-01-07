<h1>
  Go vue Starter Kit for Heroku.
</h1>


<p><strong>View</strong> <a href="https://go-vue-starter.netlify.com">online demo</a> (<a href="https://go-vue-starter.netlify.com/graphql">API</a>), Inspired by <a href="https://github.com/kriasoft/react-firebase-starter">React firebase starter</a>.

### Tech Stack

- [Vue CLI][vsc] (★ 18k) for development and test infrastructure (see [user guide][vscdocs])
- [golang.org/x/oauth2][oauth2] (★ 2k) for authentication configured with stateless JWT tokens for sessions
- [graphql-go/graphql][gqljs] (★ 4.1k) for declarative data fetching and efficient client stage management
- [PostgreSQL][psql] database pre-configured with a query builder and migrations using [Sqlboiler][sboiler] (★ 1.7k) and [Goose][goose]
- [Heroku][heroku] & [Netlify][netlify] for app architecture - Cloud SQL, CDN hosting ([docs][netlifydocs])

Also, you need to be familiar with [HTML][html], [CSS][css], [JavaScript][js] ([ES2015][es2015]) and [Vue](https://vuejs.org/v2/guide/).

### Prerequisites

- [Go][golang] v1.11 or newer
- [Node.js][nodejs] v8.11 or higher + [Yarn][yarn] v1.6 or higher &nbsp; (_HINT: On Mac install
  them via [Brew][brew]_)
- [VS Code][vc] editor (preferred) + [Project Snippets][vcsnippets], [EditorConfig][vceditconfig],
  [ESLint][vceslint], [Prettier][vcprettier], and [Babel JavaScript][vcjs] plug-ins
- [PostgreSQL][postgres] v9.6 or newer, only if you're planning to use a local db for development

### Getting Started

Just clone the repo, update environment variables in `.env` and/or `.env.local` file, and start
hacking:

```bash
$ go get bitbucket.org/liamstask/goose/cmd/goose
$ get -u -t github.com/volatiletech/sqlboiler
$ go get github.com/snapiz/go-vue-starter
$ cd ~/go/src/github.com/snapiz/go-vue-starter
$ go run server.go

# open new terminal
$ cd web
$ yarn install                     # Installs dependencies; creates PostgreSQL database
$ yarn serve                       # Compile the app and opens it in a browser with "live reload"
```

Then open [http://localhost:9000/](http://localhost:9000/) to see your app.<br>

### How to Migrate Database Schema

```bash
$ goose create AddSomeColumns sql  # Create a new database migration file
$ goose up                         # Migrate database to the latest version
$ goose down                       # Rollback the latest migration
$ sqlboiler psql                   # Generate models from db
```

### How to Test

```bash
$ go test ./..                     # Run unit tests for server
$ cd web                           
$ yarn test                        # Run unit tests. Or, `yarn test -- --watch`
```

### How to Deploy

1.  Create a new heroku project and postgres database.
2.  Configure heroku environement variables by running `heroku config:set APP_ENV=production` and for all variables in .env.
3.  Deploy your application by running `git push heroku master`.
4.  Migrate db schema by running `heroku run goose -- -env=production up` file.
5.  Update `/web/public/_redirects` with your own domain.
5.  Build your static files by running `yarn build` in web folder.
6.  Finally, drag and drop dist to netlify deploy.

### How to Update

If you keep the original Git history after cloning this repo, you can always fetch and merge
the recent updates back into your project by running:

```bash
git remote add go-vue-starter https://github.com/snapiz/go-vue-starter.git
git checkout master
git fetch go-vue-starter
git merge go-vue-starter/master
govendor sync
cd web
yarn install
yarn relay
```

_NOTE: Try to merge as soon as the new changes land on the master branch in the upstream repository,
otherwise your project may differ too much from the base/upstream repo._

### License

Copyright © 2019 snapiz. This source code is licensed under the MIT license found in
the [LICENSE](https://github.com/snapiz/go-vue-starter/LICENSE) file.

---

[vsc]: https://github.com/vuejs/vue-cli
[golang]: https://github.com/golang/go
[govendor]: https://github.com/kardianos/govendor
[vscdocs]: https://cli.vuejs.org/guide/
[psql]: https://www.postgresql.org/
[brew]: https://brew.sh/
[sboiler]: https://github.com/volatiletech/sqlboiler
[goose]: https://bitbucket.org/liamstask/goose
[gqljs]: http://graphql.org/graphql-js/
[oauth2]: https://github.com/golang/oauth2
[heroku]: https://github.com/golang/oauth2
[netlify]: https://www.netlify.com/
[netlifydocs]: https://www.netlify.com/docs/
[html]: https://developer.mozilla.org/en-US/docs/Web/HTML
[css]: https://developer.mozilla.org/en-US/docs/Web/CSS
[js]: https://developer.mozilla.org/en-US/docs/Web/JavaScript
[es2015]: http://babeljs.io/learn-es2015/
[nodejs]: https://nodejs.org/
[yarn]: https://yarnpkg.com/
[vc]: https://code.visualstudio.com/
[vcsnippets]: https://marketplace.visualstudio.com/items?itemName=rebornix.project-snippets
[vceditconfig]: https://marketplace.visualstudio.com/items?itemName=EditorConfig.EditorConfig
[vceslint]: https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint
[vcprettier]: https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode
[vcjs]: https://marketplace.visualstudio.com/items?itemName=mgmcdermott.vscode-language-babel
[watchman]: https://github.com/facebook/watchman
[postgres]: https://www.postgresql.org/