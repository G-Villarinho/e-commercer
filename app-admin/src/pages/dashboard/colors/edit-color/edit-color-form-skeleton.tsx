import { Skeleton } from "@/components/ui/skeleton";
import { Button } from "@/components/ui/button";

export function EditColorFormSkeleton() {
  return (
    <div className="space-y-8 w-full">
      <div className="grid grid-cols-3 gap-8">
        <div className="space-y-1">
          <Skeleton className="w-24 h-5" />
          <Skeleton className="h-10 w-full" />
        </div>

        <div className="space-y-1">
          <Skeleton className="w-24 h-5" />
          <Skeleton className="h-10 w-full" />
        </div>

        <div className="flex items-center justify-center">
          <Skeleton className="border p-4 rounded-full w-8 h-8" />
        </div>
      </div>

      <Button disabled className="w-40">
        <Skeleton className="h-10 w-full" />
      </Button>
    </div>
  );
}
