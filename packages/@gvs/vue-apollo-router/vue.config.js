module.exports = {
  devServer: {
    port: 9000,
    proxy: {
      "^/main": {
        target: "http://localhost:3000"
      }
    }
  }
};
