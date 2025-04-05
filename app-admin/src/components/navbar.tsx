import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { MainNav } from "@/components/main-nav";
import { StoreSwitcher } from "./store-switcher";
import { StoreSwitcherSkeleton } from "./store-switcher-skeleton";
import { getUserStores } from "@/api/get-user-stores";
import { useQuery } from "@tanstack/react-query";

export function Navbar() {
  const { data: stores, isLoading } = useQuery({
    queryKey: ["stores"],
    queryFn: getUserStores,
    refetchOnWindowFocus: false,
  });

  return (
    <div className="border-b">
      <div className="flex h-16 items-center px-4">
        {isLoading ? (
          <StoreSwitcherSkeleton />
        ) : (
          <StoreSwitcher items={stores || []} />
        )}
        <MainNav className="mx-6" />
        <div className="ml-auto flex items-center space-x-4">
          <Avatar>
            <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
        </div>
      </div>
    </div>
  );
}
