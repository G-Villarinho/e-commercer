import { cn } from "@/lib/utils";
import { Link, useLocation, useParams } from "react-router-dom";

interface route {
  to: string;
  label: string;
  active: boolean;
}

export function MainNav({ className }: React.HtmlHTMLAttributes<HTMLElement>) {
  const { storeId } = useParams();
  const pathname = useLocation().pathname;

  const routes: route[] = [
    {
      to: `/${storeId}`,
      label: "Resumo",
      active: pathname === `/${storeId}`,
    },
    {
      to: `/${storeId}/billboards`,
      label: "Painés",
      active: pathname === `/${storeId}/billboards`,
    },
    {
      to: `/${storeId}/categories`,
      label: "Categorias",
      active: pathname === `/${storeId}/categories`,
    },
    {
      to: `/${storeId}/sizes`,
      label: "Tamanhos",
      active: pathname === `/${storeId}/sizes`,
    },
    {
      to: `/${storeId}/colors`,
      label: "Cores",
      active: pathname === `/${storeId}/colors`,
    },
    {
      to: `/${storeId}/products`,
      label: "Produtos",
      active: pathname === `/${storeId}/products`,
    },
    {
      to: `/${storeId}/settings`,
      label: "Configurações",
      active: pathname === `/${storeId}/settings`,
    },
  ];

  return (
    <nav className={cn("flex items-center space-x-4 lg:space-x-6", className)}>
      {routes.map((route) => (
        <Link
          key={route.to}
          to={route.to}
          className={cn(
            "text-sm font-medium transition-colors hover:text-primary",
            route.active
              ? "text-black dark:text-white"
              : "text-muted-foreground"
          )}
        >
          {route.label}
        </Link>
      ))}
    </nav>
  );
}
