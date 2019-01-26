export default [
  {
    path: "",
    title: "Go vue starter - Home",
    acl: false,
    component: () => import(/* webpackChunkName: 'home' */ "./Home")
  },
  {
    path: "/about",
    title: "Go vue starter - About",
    component: () => import(/* webpackChunkName: 'home' */ "./About")
  }
];
