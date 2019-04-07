<template>
  <component :is="routeComponent"/>
</template>

<script>
export default {
  name: "RouterView",
  data() {
    return {
      routeComponent: null
    };
  },
  created() {
    this.render();

    this.unlisten = this.$router.context.history.listen(() => {
      this.routeComponent && this.render();
    });

    this.$root.$on("vue-apollo-router:render", this.reRender);
  },
  destroyed() {
    this.unlisten && this.unlisten();
    this.$root.$off("vue-apollo-router:render", this.reRender);
  },
  methods: {
    reRender() {
      this.currentPathname = null;
      this.render()
    },
    render() {
      const {
        context: { history, Vue }
      } = this.$router;

      const { pathname } = history.location;

      if (this.currentPathname === pathname) {
        return;
      }

      this.currentPathname = pathname;

      this.$router.resolve(pathname).then(route => {
        const { title, component, redirect, data } = route;

        if (redirect) {
          history.push(redirect);

          return;
        }

        window.document.title = title;
        Vue.prototype.$route = route;
        Vue.prototype.$routeData = data;

        this.routeComponent = component;
      });
    }
  }
};
</script>

