export default function(Vue) {
  Vue.mixin({
    data(vm) {
      return vm.$route
        ? {
            ...vm.$route.params.$data
          }
        : {};
    }
  });
}
