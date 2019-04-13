<template>
  <div>
    <h1>Admin</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      avatar: this.$routeData.me.avatar
    };
  }
};
</script>
