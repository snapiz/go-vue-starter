<template>
  <div>
    <h1>Login</h1>
    <button @click="login">Login</button>
  </div>
</template>

<script>
export default {
  methods: {
    login() {
      window.open("/auth/google", "", "width=452,height=633");
    }
  }
};
</script>
