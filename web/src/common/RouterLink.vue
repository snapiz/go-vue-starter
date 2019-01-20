<template>
  <component :is="tag" @click="go" v-bind="currentAttrs">
    <slot></slot>
  </component>
</template>

<script>
import router from "@/router";

function isLeftClickEvent(event) {
  return event.button === 0;
}

function isModifiedEvent(event) {
  return !!(event.metaKey || event.altKey || event.ctrlKey || event.shiftKey);
}

export default {
  name: "router-link",
  data() {
    let currentAttrs = {
      ...this.$attrs
    };

    if (this.tag === "a") {
      currentAttrs.href = this.to;
    }

    return { currentAttrs };
  },
  props: {
    tag: {
      type: String,
      default: "a"
    },
    to: String
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
      router.context.history.push(this.to);
    }
  }
};
</script>
