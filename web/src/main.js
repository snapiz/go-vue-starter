import Vue from "vue";
import Vuelidate from "vuelidate";
import SuiVue from "semantic-ui-vue";

import "semantic-ui-css/semantic.min.css";

import router from "./router";
import App from "./App.vue";
import { getUrlParameter } from "./utils";

window.addEventListener("message", e => {
  if (e.data.key === "login_success") {
    router.context.history.push(getUrlParameter("redirect") || "/");
  }
});

Vue.config.productionTip = false;

Vue.use(Vuelidate);
Vue.use(SuiVue);

new Vue({
  render: h => h(App)
}).$mount("#app");
