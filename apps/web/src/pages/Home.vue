<template>
  <div>
    <h1>Hello {{ me ? me.displayName : "Anonymous" }}</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      me: this.$routeData.me
    };
  }
};
</script>
