<template>
  <div id="app">
    <ul class="nav">
      <li><router-link to="/">Home</router-link></li>
      <li><router-link to="/login">Login</router-link></li>
      <li><router-link to="/avatar">Avatar</router-link></li>
      <li><router-link to="/admin">Admin</router-link></li>
      <li><router-link to="/staff">Staff</router-link></li>
      <li><router-link to="/erfefre">Not found</router-link></li>
    </ul>
    <RouterView />
  </div>
</template>

<script>
import { RouterView } from "vue-apollo-router";

export default {
  name: "app",
  components: {
    RouterView
  }
};
</script>

<style>
body {
  background-color: #fafafa;
  margin: 0px;
}

html,
body {
  height: 100%;
  width: 100%;
}

#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}

#app ul.nav {
  list-style: none;
  display: flex;
  background-color: white;
  margin: 0px;
  position: fixed;
  width: 100%;
  top: 0;
}

#app ul.nav li {
  padding: 1em;
}

#app a {
  text-decoration: none;
  color: inherit;
}

#app li:hover a {
  color: rgba(44, 62, 80, 0.5);
}

#app h1 {
  margin-top: 2em;
}
</style>
