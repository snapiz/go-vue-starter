module.exports = {
  devServer: {
    port: 9000,
    proxy: {
      "^/auth|graphql": {
        target: "http://localhost:3000"
      }
    }
  }
};
