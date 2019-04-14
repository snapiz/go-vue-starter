<template>
  <div>
    <img :src="avatar" alt="avatar" />
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
