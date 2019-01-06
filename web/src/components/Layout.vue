<template>
  <div id="layout">
    <div id="nav">
      <router-link to="/">Home</router-link>|
      <router-link to="/about">About</router-link>|
      <router-link v-if="!me" to="/login">Login</router-link>
      <router-link v-if="me" to="/login">Logout</router-link>
      <span v-if="me">({{me.displayName || 'John doe'}})</span>
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
