import Vue from "vue";
import ApolloClient from "apollo-boost";
import VueApolloRouter from "vue-apollo-router";
import i18next from "i18next";
import VueI18Next from "@panter/vue-i18next";

import App from "./App.vue";
import router from "./router";
import i18nextConfig from "./i18next";
import i18nextLazy from "./middlewares/i18next-lazy";

Vue.config.productionTip = false;

Vue.use(VueI18Next);

i18next.use(i18nextLazy).init(i18nextConfig);

const i18n = new VueI18Next(i18next);

Vue.use(VueApolloRouter, {
  router,
  apollo: {
    main: new ApolloClient({
      uri: "/api/main"
    })
  }
});

new Vue({
  i18n,
  render: h => h(App)
}).$mount("#app");
