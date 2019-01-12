import Vue from "vue";
import Vuelidate from "vuelidate";

import App from "./App.vue";
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
Vue.use(VueRouteData);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
