<template>
  <div>
    <h1>Home</h1>
    <button @click="test">Test</button>
  </div>
</template>

<script>
export default {
  name: 'Home',
  methods: {
    test() {
      this.$render()
    }
  }
}
</script>

