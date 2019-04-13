<template>
  <div>
    <img :src="avatar" alt="avatar" />
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
