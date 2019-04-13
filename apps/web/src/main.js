import Vue from "vue";
import ApolloClient from "apollo-boost";
import VueApolloRouter from "vue-apollo-router";
import App from "./App.vue";
import router from "./router";

Vue.config.productionTip = false;

Vue.use(VueApolloRouter, {
  router,
  apollo: {
    main: new ApolloClient({
      uri: "/main/api"
    })
  }
});

new Vue({
  render: h => h(App)
}).$mount("#app");
