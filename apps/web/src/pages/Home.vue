<template>
  <div>
    <h1>{{ $t("hello") }} {{ me ? me.displayName : $t("anonymous") }}</h1>
  </div>
</template>

<script>
export default {
  i18nOptions: { namespaces: ["home", "common"] },
  data() {
    return {
      me: this.$route.data.me
    };
  }
};
</script>
