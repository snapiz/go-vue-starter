<template>
  <div>
    <a :href="oAuth2.google" target="_blank" rel="noopener noreferrer"
      >Google</a
    >
    <a :href="oAuth2.facebook" target="_blank" rel="noopener noreferrer"
      >Facebook</a
    >
  </div>
</template>

<script>
export default {
  data() {
    return {
      ...this.$route.data
    };
  }
};
</script>
