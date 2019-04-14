<template>
  <div id="login">
    <a :href="oAuth2.google" target="_blank" rel="noopener noreferrer">Google</a>
    <a :href="oAuth2.facebook" target="_blank" rel="noopener noreferrer">Facebook</a>
  </div>
</template>

<style>
#login {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

#login a {
  padding: 0.5rem 1rem;
  background-color: rgb(95, 158, 160);
  color: white;
  border-radius: 3px;
  margin-bottom: 1em;
  min-width: 150px;
}

#login a:hover {
  background-color: rgb(86, 141, 143);
}
</style>

<script>
export default {
  data() {
    return {
      ...this.$route.data
    };
  }
};
</script>
