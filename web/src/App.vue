<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>
body {
  margin: 0;
}

#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 15px;
  display: flex;
}

#nav-left {
  flex-grow: 1;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
  margin: 0 5px;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
