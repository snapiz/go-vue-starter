<template>
  <div id="app">
    <ul>
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
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#app ul {
  list-style: none;
  display: flex;
}

#app li {
  margin-right: 2em;
}
</style>
