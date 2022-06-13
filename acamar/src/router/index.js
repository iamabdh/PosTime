import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/home.page";
import RegisterPage from "@/views/register.page";
import LoginPage from "@/views/login.page";
import ProfilePage from "@/views/profile.page";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: {
      title: "Home",
    },
  },
  {
    path: "/register",
    name: "RegisterPage",
    component: RegisterPage,
      meta: {
          title: "Register",
      },

  },
  {
    path: "/login",
    name: "LoginPage",
    component: LoginPage,
      meta: {
          title: "Login",
      },
  },
  {
    path: "/profile",
    name: "ProfilePage",
    component: ProfilePage,
      meta: {
          title: "Profile",
      },
  },
];


const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});
router.beforeEach((to, from, next) => {
    document.title = to.meta.title;
    next();
});
export default router;
