import Vue from "vue";
import ApolloClient from "apollo-boost";
import gql from "graphql-tag";
import App from "./App.vue";
import VueApolloRouter from "./lib";
import router from "./router";

Vue.config.productionTip = false;

const apolloClientA = new ApolloClient({
  uri: "/main/api"
});

const apolloClientB = new ApolloClient({
  uri: "/main/api"
});

apolloClientA.query({
  query: gql`
  {
    me {
      id
      displayName
      avatar
      role
    }
  }
`
}).then(x => {
  console.log(x)
})

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
