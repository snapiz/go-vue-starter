<template>
  <Page>
    <div class="error">
      <h1>{{ "pages_error_header" | t }}</h1>
    </div>
  </Page>
</template>

<script>
import Page from "@/common/Page.vue";

export default {
  name: "error",
  components: {
    Page
  }
};
</script>
