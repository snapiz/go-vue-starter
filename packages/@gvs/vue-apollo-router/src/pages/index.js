import gql from "graphql-tag";

export default [
    {
      path: "",
      title: "Vue router - Home",
      apollo: false,
      component: () => import(/* webpackChunkName: 'home' */ "./Home")
    },
    {
      path: "/me",
      title: "Vue router - Me",
      component: () => import(/* webpackChunkName: 'me' */ "./Me")
    },
    {
      path: "/admin",
      title: "Vue router - Admin",
      acl: "ADMIN",
      component: () => import(/* webpackChunkName: 'admin' */ "./Admin")
    },
    {
      path: "/staff",
      title: "Vue router - Staff",
      acl: ["ADMIN", "STAFF"],
      component: () => import(/* webpackChunkName: 'staff' */ "./Staff")
    },
    {
      path: "/mes",
      title: "Vue router - Mes",
      apollo: [
        {
          query: gql`
            {
              me {
                email
              }
            }
          `
        },
        {
          $client: 'b',
          query: gql`
            {
              me {
                role
              }
            }
          `
        }
      ],
      component: () => import(/* webpackChunkName: 'mes' */ "./Mes")
    },
    {
      path: "/avatar",
      title: "Vue router - Avatar",
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
    },
  ];
  