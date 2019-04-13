<template>
  <div>
    <h1>Error</h1>
    <p>{{ error }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      error: this.$routeError.stack
    };
  }
};
</script>
