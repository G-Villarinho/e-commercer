import { createBrowserRouter } from "react-router-dom";

// Layouts
import { AuthLayout } from "@/pages/_layouts/auth";
import { DashboardLayout } from "@/pages/_layouts/dashboard";
import { RootLayout } from "@/pages/_layouts/root";

// Auth pages
import { Login } from "@/pages/auth/login/login";
import { AccountNotFound } from "@/pages/auth/account-not-found/account-not-found";
import { Register } from "@/pages/auth/register/register";
import { VerifyCode } from "@/pages/auth/verify-code/verify-code";

// Dashboard pages
import { Overview } from "@/pages/dashboard/overview/overview";
import { Settings } from "@/pages/dashboard/settings/settings";
import { Billboards } from "@/pages/dashboard/billboards/billboards";
import { CreateBillboard } from "@/pages/dashboard/billboards/create-billboard/create-billboard";
import { EditBillboard } from "@/pages/dashboard/billboards/edit-billboard/edit-billboard";
import { Categories } from "@/pages/dashboard/categories/categories";
import { CreateCategory } from "@/pages/dashboard/categories/create-category/create-category";
import { EditCategory } from "@/pages/dashboard/categories/edit-category/edit-category";
import { Sizes } from "@/pages/dashboard/sizes/sizes";
import { CreateSize } from "@/pages/dashboard/sizes/create-size/create-size";
import { EditSize } from "@/pages/dashboard/sizes/edit-size/edit-size";
import { Colors } from "@/pages/dashboard/colors/colors";
import { CreateColor } from "@/pages/dashboard/colors/create-color/create-color";
import { EditColor } from "@/pages/dashboard/colors/edit-color/edit-color";
import { Products } from "@/pages/dashboard/products/products";
import { CreateProduct } from "@/pages/dashboard/products/create-product/create-product";

// Root page
import { Root } from "@/pages/root";

// Error pages
import { NotFound } from "@/pages/404";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    errorElement: <NotFound />,
    children: [
      {
        path: "/",
        element: <Root />,
      },
    ],
  },

  {
    path: "/:storeId",
    element: <DashboardLayout />,
    errorElement: <NotFound />,
    children: [
      {
        path: "",
        element: <Overview />,
      },
      {
        path: "settings",
        element: <Settings />,
      },
      {
        path: "billboards",
        children: [
          {
            path: "",
            element: <Billboards />,
          },
          {
            path: "new",
            element: <CreateBillboard />,
          },
          {
            path: ":billboardId",
            element: <EditBillboard />,
          },
        ],
      },
      {
        path: "categories",
        children: [
          {
            path: "",
            element: <Categories />,
          },
          {
            path: "new",
            element: <CreateCategory />,
          },
          {
            path: ":categoryId",
            element: <EditCategory />,
          },
        ],
      },
      {
        path: "sizes",
        children: [
          {
            path: "",
            element: <Sizes />,
          },
          {
            path: "new",
            element: <CreateSize />,
          },
          {
            path: ":sizeId",
            element: <EditSize />,
          },
        ],
      },
      {
        path: "colors",
        children: [
          {
            path: "",
            element: <Colors />,
          },
          {
            path: "new",
            element: <CreateColor />,
          },
          {
            path: ":colorId",
            element: <EditColor />,
          },
        ],
      },
      {
        path: "products",
        children: [
          {
            path: "",
            element: <Products />,
          },
          {
            path: "new",
            element: <CreateProduct />,
          },
          {
            path: ":colorId",
            element: <EditColor />,
          },
        ],
      },
    ],
  },

  {
    path: "/",
    element: <AuthLayout />,
    errorElement: <NotFound />,
    children: [
      {
        path: "/login",
        element: <Login />,
      },
      {
        path: "/account-not-found",
        element: <AccountNotFound />,
      },
      {
        path: "/register",
        element: <Register />,
      },
      {
        path: "/verify-code",
        element: <VerifyCode />,
      },
    ],
  },
]);
