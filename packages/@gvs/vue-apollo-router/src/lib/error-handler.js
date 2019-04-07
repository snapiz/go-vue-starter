export default function errorHandler(error, { Vue, errorPage }) {
  Vue.prototype.$routeError = error;

  return {
    title: "500 Error (Internal Server Error)",
    component: errorPage
  };
}
