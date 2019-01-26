import Vue from "vue";
import Vuelidate from "vuelidate";
import SuiVue from "semantic-ui-vue";
import i18next from "i18next";
import LngDetector from "i18next-browser-languagedetector";
import i18nextFetchBackend from "i18next-fetch-backend";
import VueI18Next from "@panter/vue-i18next";

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

Vue.use(VueI18Next);
Vue.use(Vuelidate);
Vue.use(SuiVue);

async function start() {
  await i18next
    .use(LngDetector)
    .use(i18nextFetchBackend)
    .init({
      whitelist: ["en", "fr"],
      load: "languageOnly",
      fallbackLng: "en",
      backend: {
        loadPath: "locales/{{lng}}.json"
      }
    });

  Vue.filter("t", (value, args) => (value && i18next.t(value, args)) || "");

  const i18n = new VueI18Next(i18next);

  new Vue({
    i18n,
    render: h => h(App)
  }).$mount("#app");
}

start();
