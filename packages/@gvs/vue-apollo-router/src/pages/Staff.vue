<template>
  <div>
    <h1>Staff</h1>
  </div>
</template>

<script>
export default {
  name: 'Staff',
  data() {
    return {
      avatar: this.$routeData.me.avatar
    }
  }
}
</script>
