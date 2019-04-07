<template>
  <a @click="go" v-bind:href="href">
    <slot></slot>
  </a>
</template>

<script>
function isLeftClickEvent(event) {
  return event.button === 0;
}

function isModifiedEvent(event) {
  return !!(event.metaKey || event.altKey || event.ctrlKey || event.shiftKey);
}

export default {
  name: "router-link",
  data() {
    return {
      href: this.to
    };
  },
  props: {
    to: String
  },
  mounted() {
    this.$el.removeAttribute("to");
  },
  methods: {
    go(event) {
      if (isModifiedEvent(event) || !isLeftClickEvent(event)) {
        return;
      }

      if (event.defaultPrevented === true) {
        return;
      }

      event.preventDefault();
      this.$router.context.history.push(this.href);
    }
  }
};
</script>
