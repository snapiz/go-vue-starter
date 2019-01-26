<template>
  <component :is="currentComponent" v-bind="currentProps"></component>
</template>

<script>
import router from "@/router";

export default {
  name: "AppRender",
  data: () => ({
    currentComponent: null,
    currentProps: null
  }),
  created() {
    this.render();

    const {
      context: { history }
    } = router;

    this.unlisten = history.listen(() => {
      this.render();
    });
  },
  destroyed() {
    this.unlisten && this.unlisten();
  },
  methods: {
    render() {
      const { pathname } = window.location;

      router.resolve(pathname).then(route => {
        if (!route) {
          throw `Unable to resolve ${pathname}`;
        }

        const {
          title,
          component,
          redirect,
          ctx: { history }
        } = route;

        if (redirect) {
          history.push(redirect);

          return;
        }

        if (title) {
          window.document.title =
            this.$i18n.i18next.t(title, route.data) || title;
        }

        const promise = component.beforeRouteEnter
          ? component.beforeRouteEnter(route)
          : Promise.resolve("");

        return promise.then(() => {
          this.currentProps = route.data;
          this.currentComponent = {
            ...component,
            props: {
              ...(component.props || {}),
              me: Object
            }
          };
        });
      });
    }
  }
};
</script>
