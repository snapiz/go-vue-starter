<template>
  <Layout v-bind:me="me">
    <div class="home">
      <h1>This is an home page</h1>
    </div>
  </Layout>
</template>

<script>
import Layout from "@/common/Layout.vue";

export default {
  name: "home",
  components: {
    Layout
  }
};
</script>
