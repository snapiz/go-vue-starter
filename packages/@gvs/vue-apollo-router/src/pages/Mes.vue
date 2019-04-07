<template>
  <div>
    <h1>{{ email }}</h1>
    <h1>{{ role }}</h1>
  </div>
</template>

<script>
export default {
  name: "Mes",
  data() {
    const [main1, main2] = this.$routeData;

    return {
      email: main1.me.email,
      role: main2.me.role
    };
  }
};
</script>
