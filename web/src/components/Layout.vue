<template>
  <div id="layout">
    <div id="nav">
      <span id="nav-left">
        <router-link to="/">Home</router-link>
        <router-link to="/about">About</router-link>
      </span>
      <router-link v-if="me" to="/me">Me ({{me.displayName || 'John doe'}})</router-link>
      <router-link v-if="!me" to="/login">Login</router-link>
      <router-link v-if="me" to="/login">Logout</router-link>
    </div>
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: "Layout",
  methods: {
    async logout(e) {
      e.preventDefault();
    }
  }
};
</script>
