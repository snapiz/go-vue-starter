import Vue from "vue";
import Vuelidate from "vuelidate";
import VueApollo from "vue-apollo";

import App from "./App.vue";
import apollo from "./apollo";
import router from "./router";
import store from "./store";
import VueRouteData from "./vue-route-data";

window.addEventListener("message", e => {
  if (e.data === "login_success") {
    router.push(router.currentRoute.query.redirect || "/");
  }
});

Vue.config.productionTip = false;

Vue.use(Vuelidate);
Vue.use(VueApollo);
Vue.use(VueRouteData);

const apolloProvider = new VueApollo({
  defaultClient: apollo,
  defaultOptions: {
    $query: {
      fetchPolicy: "network-only"
    }
  }
});

new Vue({
  apolloProvider,
  router,
  store,
  render: h => h(App)
}).$mount("#app");
