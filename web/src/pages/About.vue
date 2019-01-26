<template>
  <Layout v-bind:me="me">
    <div class="about">
      <h1>{{ "pages_about_header" | t }}</h1>
    </div>
  </Layout>
</template>

<script>
import Layout from "@/common/Layout.vue";

export default {
  name: "about",
  components: {
    Layout
  }
};
</script>
