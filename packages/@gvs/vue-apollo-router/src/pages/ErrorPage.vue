<template>
  <div>
    <h1>Error</h1>
    <p>{{ error }}</p>
  </div>
</template>

<script>
export default {
  name: "ErrorPage",
  data() {
    return {
      error: this.$routeError
    };
  }
};
</script>

