import gql from "graphql-tag";

export default [
  {
    path: "",
    title: "Go vue starter - Home",
    acl: false,
    component: () => import(/* webpackChunkName: 'home' */ "./Home")
  },
  {
    path: "/login",
    title: "Go vue starter - Login",
    apollo: false,
    component: () => import(/* webpackChunkName: 'login' */ "./Login")
  },
  {
    path: "/admin",
    title: "Go vue starter - Admin",
    acl: "ADMIN",
    component: () => import(/* webpackChunkName: 'admin' */ "./Admin")
  },
  {
    path: "/staff",
    title: "Go vue starter - Staff",
    acl: ["ADMIN", "STAFF"],
    component: () => import(/* webpackChunkName: 'staff' */ "./Staff")
  },
  {
    path: "/avatar",
    title: "Go vue starter - Avatar",
    apollo: {
      query: gql`
        {
          me {
            avatar
          }
        }
      `
    },
    component: () => import(/* webpackChunkName: 'avatar' */ "./Avatar")
  }
];
