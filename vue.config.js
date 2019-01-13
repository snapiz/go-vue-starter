module.exports = {
  configureWebpack: {
    resolve: {
      alias: {
        "@": __dirname + "/web"
      }
    },
    entry: {
      app: "./web/main.js"
    }
  },
  devServer: {
    port: 9000,
    proxy: {
      "^/auth|graphql": {
        target: "http://localhost:3000"
      }
    }
  }
};
