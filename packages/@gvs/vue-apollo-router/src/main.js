import Vue from "vue";
import ApolloClient from "apollo-boost";
import App from "./App.vue";
import VueApolloRouter from "./lib";
import router from "./router";

Vue.config.productionTip = false;

const apolloClientA = new ApolloClient({
  uri: "/main/api",
  credentials: "same-origin"
});

const apolloClientB = new ApolloClient({
  uri: "/main/api",
  credentials: "same-origin"
});

Vue.use(VueApolloRouter, {
  apollo: {
    a: apolloClientA,
    b: apolloClientB
  },
  router
});

new Vue({
  render: h => h(App)
}).$mount("#app");
