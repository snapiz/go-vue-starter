<template>
  <Layout v-bind:me="me">
    <div class="home">
      <h1>{{ "pages_home_header" | t }}</h1>
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
