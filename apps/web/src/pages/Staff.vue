<template>
  <div>
    <h1>Staff</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      avatar: this.$route.data.me.avatar
    };
  }
};
</script>
