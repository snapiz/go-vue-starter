module.exports = {
  devServer: {
    port: 9000,
    proxy: {
      "^/(main|auth)": {
        target: "http://localhost:3000"
      }
    }
  }
};
